package main

import (
	"doudou/pkg/errors"
	"fmt"

	ginplus "github.com/dllgo/go-gin"
	jwtplus "github.com/dllgo/go-jwt"
	"github.com/gin-gonic/gin"
)

func main() {
	mconf := ginplus.Config{Address: ":9090", ReadTimeout: 30, WriteTimeout: 30}
	httpserver, err := ginplus.NewServerHttp(mconf)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpserver.Router = InitRouter()
	err = httpserver.Listen()
	defer httpserver.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
func InitRouter() *gin.Engine {
	router := gin.New()
	router.NoRoute(ginplus.NoRouteHandler())
	router.NoMethod(ginplus.NoMethodHandler())
	router.Use(ginplus.CorsMiddleware())
	router.Use(ginplus.LogMiddleware())
	router.Use(ginplus.AuthMiddleware(ginplus.AllowMethodAndPathPrefixSkipper(
		ginplus.JoinRouter("GET", "/login"),
	)))
	//注册api
	router.GET("/login", login)
	router.GET("/hello", hello)
	return router
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
}
func hello(context *gin.Context) {
	println(">>>> hello function start <<<<")

	ginplus.ResOK(context, nil)
	ginplus.ResList(context, nil, 0)
}
