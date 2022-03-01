package request

type SysUserModel struct {
	Name       string `json:"name" validate:"required,min=3,max=32" name:"用户名"`
	Password   string `json:"password" validate:"required,min=8,max=32" name:"密码"`
	RememberMe bool   `json:"rememberMe" name:"记住密码"`
}
