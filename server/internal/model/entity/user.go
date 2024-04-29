// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id                 int64       `json:"id"                 description:"主键"`
	Role               int64       `json:"role"               description:"权限等级"`
	RealName           string      `json:"realName"           description:"真实姓名"`
	Username           string      `json:"username"           description:"帐号"`
	PasswordHash       string      `json:"passwordHash"       description:"密码"`
	Salt               string      `json:"salt"               description:"密码盐"`
	PasswordResetToken string      `json:"passwordResetToken" description:"密码重置令牌"`
	Avatar             string      `json:"avatar"             description:"头像"`
	Sex                int         `json:"sex"                description:"性别"`
	Qq                 string      `json:"qq"                 description:"qq"`
	Email              string      `json:"email"              description:"邮箱"`
	Mobile             string      `json:"mobile"             description:"手机号码"`
	Birthday           *gtime.Time `json:"birthday"           description:"生日"`
	Address            string      `json:"address"            description:"联系地址"`
	LastActiveAt       *gtime.Time `json:"lastActiveAt"       description:"最后活跃时间"`
	Remark             string      `json:"remark"             description:"备注"`
	Status             int         `json:"status"             description:"状态"`
	CreatedAt          *gtime.Time `json:"createdAt"          description:"创建时间"`
	UpdatedAt          *gtime.Time `json:"updatedAt"          description:"修改时间"`
}
