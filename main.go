package main

import (
	"doudou/pkg/errors"
	"doudou/pkg/ginplus"
	"doudou/pkg/jwtplus"
	"doudou/pkg/redisplus"

	"github.com/gin-gonic/gin"
)

func main() {
	InitDb()
	router := InitWeb()
	router.GET("/login", login) // hello函数处理"/hello"请求
	router.GET("/hello", hello) // hello函数处理"/hello"请求
	// 指定地址和端口号
	router.Run(":9090")
}
func InitWeb() *gin.Engine {
	gin.SetMode("debug")
	app := gin.Default()
	app.Use(ginplus.CorsMiddleware())
	app.Use(ginplus.LogMiddleware())
	app.Use(ginplus.AuthMiddleware(ginplus.AllowMethodAndPathPrefixSkipper(
		ginplus.JoinRouter("GET", "/login"),
	)))
	return app
}
func InitDb() {
	redisplus.NewRedisWithDefualtConfig()
}
func login(context *gin.Context) {
	println(">>>> login start <<<<")
	token, err := jwtplus.GenToken(&jwtplus.Userdata{UserId: "123456"})
	if err != nil {
		println(">>>> login start 3<<<<")
		errors.WithStack(err)
		ginplus.ResError(context, err)
		return
	}
	ginplus.ResOK(context, token)
	jwtplus.DestroyToken(token)
}
func hello(context *gin.Context) {
	println(">>>> hello function start <<<<")

	ginplus.ResOK(context, nil)
	ginplus.ResList(context, nil, 0)
	ginplus.ResError(context, nil)
}
