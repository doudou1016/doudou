package schema

// LoginParam 登录参数
type LoginParam struct {
	UserName string `json:"username" binding:"required" swaggo:"true,用户名"`
	Password string `json:"password" binding:"required" swaggo:"true,密码(md5加密)"`
}

// User 用户对象
type UserParam struct {
	UserId   string `json:"user_id" swaggo:"false,记录ID"`
	UserName string `json:"username" binding:"required" swaggo:"true,用户名"`
	Nickname string `json:"nickname" binding:"required" swaggo:"true,真实姓名"`
	Password string `json:"password" swaggo:"false,密码"`
	Phone    string `json:"phone" swaggo:"false,手机号"`
	Email    string `json:"email" swaggo:"true,邮箱"`
	Status   int64  `json:"status" swaggo:"true,用户状态(1:启用 2:停用)"`
}
