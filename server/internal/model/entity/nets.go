// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Nets is the golang structure for table nets.
type Nets struct {
	Id        int64       `json:"id"        description:"主键"`
	Timestamp *gtime.Time `json:"timestamp" description:"时间戳"`
	Ethernet  string      `json:"ethernet"  description:"以太网接口"`
	NetRecv   float64     `json:"netRecv"   description:"网络接收量"`
	NetSend   float64     `json:"netSend"   description:"网络发送量"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改时间"`
}
