// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Hosts is the golang structure for table hosts.
type Hosts struct {
	Id              int64       `json:"id"              ` // 主键
	Timestamp       *gtime.Time `json:"timestamp"       ` // 时间戳
	Uptime          string      `json:"uptime"          ` // 运行时间
	Hostname        string      `json:"hostname"        ` // 主机名
	Os              string      `json:"os"              ` // 操作系统
	Platform        string      `json:"platform"        ` // 平台
	PlatformVersion string      `json:"platformVersion" ` // 平台版本
	KernelVersion   string      `json:"kernelVersion"   ` // 内核版本
	KernelArch      string      `json:"kernelArch"      ` // 内核架构
	CreatedAt       *gtime.Time `json:"createdAt"       ` // 创建时间
	UpdatedAt       *gtime.Time `json:"updatedAt"       ` // 修改时间
}
