package controller

import (
	"doudou/schema"
	"errors"
	"fmt"

	user "gitee.com/go90/user-srv/proto/user"
	ginplus "github.com/dllgo/go-gin"
	grpcplus "github.com/dllgo/go-grpc"
	jwtplus "github.com/dllgo/go-jwt"
	"github.com/gin-gonic/gin"
)

func NewUserSvr1Client() user.UserService {
	return user.NewUserService("com.lcb123.srv.user", grpcplus.NewClient())
}

type AuthApi struct {
	BaseApi
	userClient user.UserService
}

func NewAuthApi() *AuthApi {
	return &AuthApi{
		userClient: NewUserSvr1Client(),
	}
}
func (api *AuthApi) Router(router *gin.RouterGroup) {
	router.POST("/login", api.Login)
	router.POST("/logout", api.Logout)
	router.POST("/register", api.Register)
}

// Login 用户登录
// @Summary 用户登录
// @Success 200 schema.HTTPStatus "{status:OK}"
// @Router POST /api/v1/login
func (api *AuthApi) Login(c *gin.Context) {
	var item schema.LoginParam
	if err := ginplus.ParseJSON(c, &item); err != nil {
		api.faild(c, errors.New("参数异常"))
		return
	}
	rsp, err := api.userClient.Verify(c, &user.VerifyRequest{
		UserName: item.UserName,
		Password: item.Password,
	})
	if err != nil {
		api.faild(c, errors.New("用户名或密码错误"))
		return
	}
	token, err := jwtplus.GenToken(&jwtplus.Userdata{UserId: rsp.UserId})
	if err != nil {
		api.faild(c, errors.New("token异常"))
		return
	}
	api.ok(c, token)
}

// Logout 用户登出
// @Summary 用户登出
// @Success 200 schema.HTTPStatus "{status:OK}"
// @Router POST /api/v1/logout

func (api *AuthApi) Logout(c *gin.Context) {
	// 检查用户是否处于登录状态，如果是则执行销毁
	token := ginplus.GetToken(c)
	fmt.Println(token)
	if len(token) > 0 {
		err := jwtplus.DestroyToken(token)
		if err != nil {
			api.faild(c, errors.New("退出异常"))
			return
		}
	}
	api.ok(c, ginplus.HTTPStatus{Status: "退出成功"})
}

// Register 用户注册
// @Summary 用户注册
// @Success 200 schema.HTTPStatus "{status:OK}"
// @Router POST /api/v1/register

func (api *AuthApi) Register(c *gin.Context) {
	var suser schema.UserParam
	if err := ginplus.ParseJSON(c, &suser); err != nil {
		api.faild(c, errors.New("参数异常"))
		return
	}
	_, err := api.userClient.Create(c, &user.UserSchema{
		UserName: suser.UserName,
		NickName: suser.Nickname,
		Password: suser.Password,
		Phone:    suser.Phone,
		Email:    suser.Email,
		Status:   suser.Status,
	})
	if err != nil {
		api.faild(c, errors.New("注册失败"))
		return
	}
	api.ok(c, ginplus.HTTPStatus{Status: "注册成功"})
}
