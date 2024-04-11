// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"server/api/user/v1"
)

type IUserV1 interface {
	Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error)
	LoginCaptcha(ctx context.Context, req *v1.LoginCaptchaReq) (res *v1.LoginCaptchaRes, err error)
	AccountLogin(ctx context.Context, req *v1.AccountLoginReq) (res *v1.AccountLoginRes, err error)
	LoginLogout(ctx context.Context, req *v1.LoginLogoutReq) (res *v1.LoginLogoutRes, err error)
}
