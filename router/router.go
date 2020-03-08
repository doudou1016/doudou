package router

import (
	"doudou/controller"
	"doudou/controller/admin"

	ginplus "github.com/dllgo/go-gin"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.NoRoute(ginplus.NoRouteHandler())
	router.NoMethod(ginplus.NoMethodHandler())
	router.Use(ginplus.CorsMiddleware())
	router.Use(ginplus.LogMiddleware())
	//注册admin api
	adminRouter(router)
	//注册api
	normalRouter(router)
	//注册api
	authorizedRouter(router)

	return router
}
func normalRouter(r *gin.Engine) {
	//注册api
	v1 := r.Group("/api/v1")
	{
		controller.NewAuthApi().Router(v1)
	}
}

func authorizedRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.Use(ginplus.AuthMiddleware())
	{
		controller.NewUserApi().Router(v1)
	}
}

func adminRouter(r *gin.Engine) {
	adm := r.Group("/api/admin")
	adm.Use(ginplus.AuthMiddleware())
	{
		admin.NewDashboardApi().Router(adm)
	}
}
