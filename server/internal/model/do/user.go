// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table ck8s_user for DAO operations like Where/Data.
type User struct {
	g.Meta             `orm:"table:ck8s_user, do:true"`
	Id                 interface{} // 主键
	Role               interface{} // 权限等级
	RealName           interface{} // 真实姓名
	Username           interface{} // 帐号
	PasswordHash       interface{} // 密码
	Salt               interface{} // 密码盐
	PasswordResetToken interface{} // 密码重置令牌
	Avatar             interface{} // 头像
	Sex                interface{} // 性别
	Qq                 interface{} // qq
	Email              interface{} // 邮箱
	Mobile             interface{} // 手机号码
	Birthday           *gtime.Time // 生日
	Address            interface{} // 联系地址
	LastActiveAt       *gtime.Time // 最后活跃时间
	Remark             interface{} // 备注
	Status             interface{} // 状态
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 修改时间
}
