// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Nets is the golang structure for table nets.
type Nets struct {
	Id        int64       `json:"id"        ` // 主键
	Timestamp *gtime.Time `json:"timestamp" ` // 时间戳
	Ethernet  string      `json:"ethernet"  ` // 以太网接口
	NetRecv   float64     `json:"netRecv"   ` // 网络接收量
	NetSend   float64     `json:"netSend"   ` // 网络发送量
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 修改时间
}
