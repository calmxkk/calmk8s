package v1

import (
	"calmk8s/internal/model/input"

	"github.com/gogf/gf/v2/frame/g"
)

// RegisterReq 提交账号注册
type RegisterReq struct {
	g.Meta `path:"/site/register" method:"post" tags:"后台基础" summary:"账号注册"`
	input.RegisterInp
}

type RegisterRes struct {
	*input.LoginModel
}

// LoginCaptchaReq 获取登录验证码
type LoginCaptchaReq struct {
	g.Meta `path:"/site/captcha" method:"get" tags:"后台基础" summary:"获取登录验证码"`
}

type LoginCaptchaRes struct {
	Cid    string `json:"cid" dc:"验证码ID"`
	Base64 string `json:"base64" dc:"验证码"`
}

// AccountLoginReq 登录
type AccountLoginReq struct {
	g.Meta `path:"/site/accountLogin" method:"post" tags:"后台基础" summary:"账号登录"`
	input.AccountLoginInp
}

type AccountLoginRes struct {
	*input.LoginModel
}

// LoginLogoutReq 注销登录
type LoginLogoutReq struct {
	g.Meta `path:"/site/logout" method:"post" tags:"后台基础" summary:"注销登录"`
}

type LoginLogoutRes struct{}
