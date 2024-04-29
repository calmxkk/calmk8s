// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id                 int64       `json:"id"                 ` // 主键
	Role               int64       `json:"role"               ` // 权限等级
	RealName           string      `json:"realName"           ` // 真实姓名
	Username           string      `json:"username"           ` // 帐号
	PasswordHash       string      `json:"passwordHash"       ` // 密码
	Salt               string      `json:"salt"               ` // 密码盐
	PasswordResetToken string      `json:"passwordResetToken" ` // 密码重置令牌
	Avatar             string      `json:"avatar"             ` // 头像
	Sex                int         `json:"sex"                ` // 性别
	Qq                 string      `json:"qq"                 ` // qq
	Email              string      `json:"email"              ` // 邮箱
	Mobile             string      `json:"mobile"             ` // 手机号码
	Birthday           *gtime.Time `json:"birthday"           ` // 生日
	Address            string      `json:"address"            ` // 联系地址
	LastActiveAt       *gtime.Time `json:"lastActiveAt"       ` // 最后活跃时间
	Remark             string      `json:"remark"             ` // 备注
	Status             int         `json:"status"             ` // 状态
	CreatedAt          *gtime.Time `json:"createdAt"          ` // 创建时间
	UpdatedAt          *gtime.Time `json:"updatedAt"          ` // 修改时间
}
