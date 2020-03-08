package controller

import (
	ginplus "github.com/dllgo/go-gin"
	"github.com/gin-gonic/gin"
)

type BaseApi struct {
}

func (api *BaseApi) ok(c *gin.Context, v interface{}) {
	ginplus.ResSuccess(c, v)
}
func (api *BaseApi) faild(c *gin.Context, err error) {
	ginplus.ResError(c, err)
}
