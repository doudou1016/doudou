package gormplus

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DBService is a database engine object.
type DBService struct {
	Default *gorm.DB            // the default database engine
	List    map[string]*gorm.DB // database engine list
}

var dbService = func() (serv *DBService) {
	serv = &DBService{
		List: map[string]*gorm.DB{},
	}

	var errs []string
	defer func() {
		if len(errs) > 0 {
			panic("[gorm] " + strings.Join(errs, "\n"))
		}
		if serv.Default == nil {
			log.Panicf("[gorm] the `default` database engine must be configured and enabled")
		}
	}()

	err := loadDBConfig()
	if err != nil {
		log.Panicf("[gorm]", err.Error())
		return
	}

	for _, conf := range dbConfigs {
		if !conf.Enable {
			continue
		}
		engine, err := gorm.Open(conf.Driver, conf.Connstring)
		if err != nil {
			log.Panicf("[gorm]", err.Error())
			errs = append(errs, err.Error())
			continue
		}
		engine.LogMode(conf.ShowSql)
		engine.DB().SetMaxOpenConns(conf.MaxOpenConns)
		engine.DB().SetMaxIdleConns(conf.MaxIdleConns)
		engine.SingularTable(true)
		engine.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
		engine.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
		engine.Callback().Delete().Replace("gorm:delete", deleteCallback)
		serv.List[conf.Name] = engine
		if DEFAULTDB_NAME == conf.Name {
			serv.Default = engine
		}
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return conf.TableFix + defaultTableName
		}
	}
	return
}()

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedAt", time.Now().Unix())
	}
}

// deleteCallback will set `DeletedOn` where deleting
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedAt")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
