// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Containers is the golang structure for table containers.
type Containers struct {
	Id          int64       `json:"id"          ` // 主键
	Timestamp   *gtime.Time `json:"timestamp"   ` // 时间戳
	ContainerId string      `json:"containerId" ` // 容器ID
	Name        string      `json:"name"        ` // 名称
	Image       string      `json:"image"       ` // 镜像
	Ip          string      `json:"ip"          ` // IP地址
	State       string      `json:"state"       ` // 状态
	Uptime      string      `json:"uptime"      ` // 运行时间
	CpuPercent  float64     `json:"cpuPercent"  ` // CPU使用率
	MemPercent  float64     `json:"memPercent"  ` // 内存使用率
	MemUsage    float64     `json:"memUsage"    ` // 内存使用量
	MemLimit    float64     `json:"memLimit"    ` // 内存限制
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 修改时间
}
