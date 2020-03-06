package gormplus

import "time"

type Model struct {
	ID        int        `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at;type:datetime;" json:"created_at"` // 创建时间
	UpdatedAt time.Time  `gorm:"column:updated_at;type:datetime;" json:"updated_at"` // 更新时间
	DeletedAt *time.Time `gorm:"column:updated_at;type:datetime;" json:"deleted_at" sql:"index"`
}

type ModelID struct {
	ID int `gorm:"primary_key;auto_increment" json:"id"`
}
