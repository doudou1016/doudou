package controller

import (
	"doudou/schema"
	"errors"

	user "gitee.com/go90/user-srv/proto/user"
	ginplus "github.com/dllgo/go-gin"
	grpcplus "github.com/dllgo/go-grpc"
	"github.com/gin-gonic/gin"
)

func NewUserSrvClient() user.UserService {
	return user.NewUserService("com.lcb123.srv.user", grpcplus.NewClient())
}

type UserApi struct {
	BaseApi
	userClient user.UserService
}

func NewUserApi() *UserApi {
	return &UserApi{
		userClient: NewUserSrvClient(),
	}
}
func (api *UserApi) Router(router *gin.RouterGroup) {
	router.POST("/users", api.UserList)
	router.GET("/currentuser", api.CurUser)
	router.GET("/user/:id", api.UserInfo)
}

// Logout 用户登录
// @Summary 用户登录
// @Success 200 schema.HTTPStatus "{status:OK}"
// @Router POST /api/v1/login
func (api *UserApi) UserList(c *gin.Context) {
	var params schema.HTTPPagination
	if err := ginplus.ParseJSON(c, &params); err != nil {
		ginplus.ResError(c, errors.New("参数异常"))
		return
	}
	var page = &user.PaginationParam{
		PageIndex: params.PageIndex,
		PageSize:  params.PageSize,
	}
	rsp, err := api.userClient.List(c, &user.QueryRequest{
		Page: page,
	})
	if err != nil {
		api.faild(c, errors.New("数据异常"))
		return
	}
	api.ok(c, rsp)
}

func (api *UserApi) UserInfo(c *gin.Context) {
	item, err := api.userClient.Get(c, &user.QueryRequest{
		UserId: c.Param("id"),
	})
	if err != nil {
		api.faild(c, errors.New("用户不存在"))
		return
	}
	api.ok(c, item)
}

func (api *UserApi) CurUser(c *gin.Context) {
	item, err := api.userClient.Get(c, &user.QueryRequest{
		UserId: ginplus.GetUserID(c),
	})
	if err != nil {
		api.faild(c, errors.New("重新登录"))
		return
	}
	api.ok(c, item)
}
