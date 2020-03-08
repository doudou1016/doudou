package admin

import (
	"github.com/gin-gonic/gin"
)

type DashboardApi struct {
	AdminApi
}

func NewDashboardApi() *DashboardApi {
	return &DashboardApi{}
}
func (api *DashboardApi) Router(router *gin.RouterGroup) {
	router.POST("/dashboard/index", api.Index)
}

// Login 用户登录
// @Summary 用户登录
// @Success 200 schema.HTTPStatus "{status:OK}"
// @Router POST /api/admin/dashboard/index
func (api *DashboardApi) Index(c *gin.Context) {
	api.ok(c, "")
}
