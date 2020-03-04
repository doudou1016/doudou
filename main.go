package main

import (
	"doudou/pkg/ginplus"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := InitWeb()
	router.GET("/hello", hello) // hello函数处理"/hello"请求
	// 指定地址和端口号
	router.Run(":9090")
}
func InitWeb() *gin.Engine {
	gin.SetMode("debug")
	app := gin.Default()
	app.Use(ginplus.NoMethodHandler())
	return app
}
func hello(context *gin.Context) {
	println(">>>> hello function start <<<<")

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"success": true,
	})
}
