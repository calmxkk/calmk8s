package user

import (
	v1 "calmk8s/api/user/v1"
	"calmk8s/internal/library/captcha"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerV1) LoginCaptcha(ctx context.Context, req *v1.LoginCaptchaReq) (res *v1.LoginCaptchaRes, err error) {
	cid, base64 := captcha.Generate(ctx)
	res = &v1.LoginCaptchaRes{
		Cid:    cid,
		Base64: base64,
	}
	return
}
func (c *ControllerV1) AccountLogin(ctx context.Context, req *v1.AccountLoginReq) (res *v1.AccountLoginRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerV1) LoginLogout(ctx context.Context, req *v1.LoginLogoutReq) (res *v1.LoginLogoutRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
