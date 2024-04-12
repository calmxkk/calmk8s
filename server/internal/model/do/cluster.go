// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Cluster is the golang structure of table ck8s_cluster for DAO operations like Where/Data.
type Cluster struct {
	g.Meta    `orm:"table:ck8s_cluster, do:true"`
	Id        interface{} // ä¸»é”®
	Name      interface{} // é›†ç¾¤åç§°
	Configstr interface{} // é…ç½®æ–‡ä»¶
	CreatedBy *gtime.Time // åˆ›å»ºç”¨æˆ·
	CreatedAt *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt *gtime.Time // ä¿®æ”¹æ—¶é—´
}
