// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Cluster is the golang structure for table cluster.
type Cluster struct {
	Id        int64       `json:"id"        ` // ä¸»é”®
	Name      string      `json:"name"      ` // é›†ç¾¤åç§°
	Configstr string      `json:"configstr" ` // é…ç½®æ–‡ä»¶
	CreatedBy *gtime.Time `json:"createdBy" ` // åˆ›å»ºç”¨æˆ·
	CreatedAt *gtime.Time `json:"createdAt" ` // åˆ›å»ºæ—¶é—´
	UpdatedAt *gtime.Time `json:"updatedAt" ` // ä¿®æ”¹æ—¶é—´
}
