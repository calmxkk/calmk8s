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
	Id                 interface{} // ä¸»é”®
	Role               interface{} // æƒé™ç­‰çº§
	RealName           interface{} // çœŸå®žå§“å
	Username           interface{} // å¸å·
	PasswordHash       interface{} // å¯†ç 
	Salt               interface{} // å¯†ç ç›
	PasswordResetToken interface{} // å¯†ç é‡ç½®ä»¤ç‰Œ
	Avatar             interface{} // å¤´åƒ
	Sex                interface{} // æ€§åˆ«
	Qq                 interface{} // qq
	Email              interface{} // é‚®ç®±
	Mobile             interface{} // æ‰‹æœºå·ç 
	Birthday           *gtime.Time // ç”Ÿæ—¥
	Address            interface{} // è”ç³»åœ°å€
	LastActiveAt       *gtime.Time // æœ€åŽæ´»è·ƒæ—¶é—´
	Remark             interface{} // å¤‡æ³¨
	Status             interface{} // çŠ¶æ€
	CreatedAt          *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt          *gtime.Time // ä¿®æ”¹æ—¶é—´
}
