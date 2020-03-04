package main

import (
	"doudou/pkg/auth"
	"doudou/pkg/errors"
	"doudou/pkg/ginplus"

	"github.com/gin-gonic/gin"
)

func main() {
	router := InitWeb()
	router.GET("/login", login) // hello函数处理"/hello"请求
	router.GET("/hello", hello) // hello函数处理"/hello"请求
	// 指定地址和端口号
	router.Run(":9090")
}
func InitWeb() *gin.Engine {
	gin.SetMode("debug")
	app := gin.Default()
	app.Use(ginplus.AuthMiddleware(ginplus.AllowMethodAndPathPrefixSkipper(
		ginplus.JoinRouter("GET", "/login"),
	)))
	return app
}
func login(context *gin.Context) {
	println(">>>> login start <<<<")
	var AuthS = auth.DefaultAuthServer()
	token, err := AuthS.GenerateToken(context, "123456")
	if err != nil {
		println(">>>> login start 3<<<<")
		errors.WithStack(err)
		ginplus.ResError(context, err)
		return
	}
	AuthS.DestroyToken(context, token)
	ginplus.ResOK(context, token)
}
func hello(context *gin.Context) {
	println(">>>> hello function start <<<<")

	ginplus.ResOK(context, nil)
	ginplus.ResList(context, nil, 0)
	ginplus.ResError(context, nil)
}
