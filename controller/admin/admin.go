package admin

import (
	ginplus "github.com/dllgo/go-gin"
	"github.com/gin-gonic/gin"
)

type AdminApi struct {
}

func (api *AdminApi) ok(c *gin.Context, v interface{}) {
	ginplus.ResSuccess(c, v)
}
func (api *AdminApi) faild(c *gin.Context, err error) {
	ginplus.ResError(c, err)
}
