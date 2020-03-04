package ginplus

import (
	"doudou/pkg/errors"
	"doudou/pkg/utils"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// HTTPError HTTP响应错误
type HTTPError struct {
	Error HTTPErrorItem `json:"error" swaggo:"true,错误项"`
}

// HTTPErrorItem HTTP响应错误项
type HTTPErrorItem struct {
	Code    int    `json:"code" swaggo:"true,错误码"`
	Message string `json:"message" swaggo:"true,错误信息"`
}

// HTTPStatus HTTP响应状态
type HTTPStatus struct {
	Status string `json:"status" swaggo:"true,状态(OK)"`
}

// HTTPList HTTP响应列表数据
type HTTPList struct {
	List       interface{}     `json:"list"`
	Pagination *HTTPPagination `json:"pagination,omitempty"`
}

// HTTPPagination HTTP分页数据
type HTTPPagination struct {
	Total    int `json:"total"`
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
}

// 定义上下文中的键
const (
	prefix = "doudou"
	// UserIDKey 存储上下文中的键(用户ID)
	UserIDKey = prefix + "/user_id"
	// TraceIDKey 存储上下文中的键(跟踪ID)
	TraceIDKey = prefix + "/trace_id"
	// ResBodyKey 存储上下文中的键(响应Body数据)
	ResBodyKey = prefix + "/res_body"
)

// ResList 响应列表数据
func ResList(c *gin.Context, v interface{}) {
	ResSuccess(c, HTTPList{List: v})
}

// ResOK 响应OK
func ResOK(c *gin.Context) {
	ResSuccess(c, HTTPStatus{Status: "OK"})
}

// ResSuccess 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	ResJSON(c, http.StatusOK, v)
}

// ResPage 响应分页数据
func ResPage(c *gin.Context, v interface{}) {

	ResSuccess(c, v)
}

// ParseJSON 解析请求JSON
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.ErrInvalidRequestParameter
	}
	return nil
}

// GetUserID 获取用户ID
func GetUserID(c *gin.Context) string {
	return c.GetString(UserIDKey)
}

// SetUserID 设定用户ID
func SetUserID(c *gin.Context, userID string) {
	c.Set(UserIDKey, userID)
}

// GetToken 获取用户令牌
func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

// GetBasicToken 获取basic认证信息
func GetBasicToken(c *gin.Context) (string, string, error) {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Basic "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	credential, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", "", err
	}
	userAndPassword := strings.Split(string(credential), ":")
	return userAndPassword[0], userAndPassword[1], nil
}

// ResJSON 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	buf, err := utils.JSONMarshal(v)
	if err != nil {
		panic(err)
	}
	c.Set(ResBodyKey, buf)
	c.Data(status, "application/json; charset=utf-8", buf)
	c.Abort()
}

// ResError 响应错误
func ResError(c *gin.Context, err error, status ...int) {
	statusCode := 500
	errItem := HTTPErrorItem{
		Code:    500,
		Message: "服务器发生错误",
	}

	if errCode, ok := errors.FromErrorCode(err); ok {
		errItem.Code = errCode.Code
		errItem.Message = errCode.Message
		statusCode = errCode.HTTPStatusCode
	}

	if len(status) > 0 {
		statusCode = status[0]
	}
	ResJSON(c, statusCode, HTTPError{Error: errItem})
}

// GetPageIndex 获取分页的页索引
func GetPageIndex(c *gin.Context) int {
	defaultVal := 1
	if v := c.Query("current"); v != "" {
		if iv := utils.S(v).DefaultInt(defaultVal); iv > 0 {
			return iv
		}
	}
	return defaultVal
}

// GetPageSize 获取分页的页大小(最大50)
func GetPageSize(c *gin.Context) int {
	defaultVal := 10
	if v := c.Query("pageSize"); v != "" {
		if iv := utils.S(v).DefaultInt(defaultVal); iv > 0 {
			if iv > 50 {
				iv = 50
			}
			return iv
		}
	}
	return defaultVal
}

// // GetPaginationParam 获取分页查询参数
// func GetPaginationParam(c *gin.Context) *user.PaginationParam {
// 	return &user.PaginationParam{
// 		PageIndex: int64(GetPageIndex(c)),
// 		PageSize:  int64(GetPageSize(c)),
// 	}
// }

// ResOpenshiftLoginError 响应openshfit登录错误
// func ResOpenshiftLoginError(c *gin.Context, err error) {
// 	ResJSON(c, http.StatusUnauthorized, schema.OpenshiftLoginError{Msg: err.Error()})
// }
