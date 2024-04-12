// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id                 int64       `json:"id"                 ` // ä¸»é”®
	Role               int64       `json:"role"               ` // æƒé™ç­‰çº§
	RealName           string      `json:"realName"           ` // çœŸå®žå§“å
	Username           string      `json:"username"           ` // å¸å·
	PasswordHash       string      `json:"passwordHash"       ` // å¯†ç 
	Salt               string      `json:"salt"               ` // å¯†ç ç›
	PasswordResetToken string      `json:"passwordResetToken" ` // å¯†ç é‡ç½®ä»¤ç‰Œ
	Avatar             string      `json:"avatar"             ` // å¤´åƒ
	Sex                int         `json:"sex"                ` // æ€§åˆ«
	Qq                 string      `json:"qq"                 ` // qq
	Email              string      `json:"email"              ` // é‚®ç®±
	Mobile             string      `json:"mobile"             ` // æ‰‹æœºå·ç 
	Birthday           *gtime.Time `json:"birthday"           ` // ç”Ÿæ—¥
	Address            string      `json:"address"            ` // è”ç³»åœ°å€
	LastActiveAt       *gtime.Time `json:"lastActiveAt"       ` // æœ€åŽæ´»è·ƒæ—¶é—´
	Remark             string      `json:"remark"             ` // å¤‡æ³¨
	Status             int         `json:"status"             ` // çŠ¶æ€
	CreatedAt          *gtime.Time `json:"createdAt"          ` // åˆ›å»ºæ—¶é—´
	UpdatedAt          *gtime.Time `json:"updatedAt"          ` // ä¿®æ”¹æ—¶é—´
}
