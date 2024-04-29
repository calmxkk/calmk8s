// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Containers is the golang structure for table containers.
type Containers struct {
	Id          int64       `json:"id"          description:"主键"`
	Timestamp   *gtime.Time `json:"timestamp"   description:"时间戳"`
	ContainerId string      `json:"containerId" description:"容器ID"`
	Name        string      `json:"name"        description:"名称"`
	Image       string      `json:"image"       description:"镜像"`
	Ip          string      `json:"ip"          description:"IP地址"`
	State       string      `json:"state"       description:"状态"`
	Uptime      string      `json:"uptime"      description:"运行时间"`
	CpuPercent  float64     `json:"cpuPercent"  description:"CPU使用率"`
	MemPercent  float64     `json:"memPercent"  description:"内存使用率"`
	MemUsage    float64     `json:"memUsage"    description:"内存使用量"`
	MemLimit    float64     `json:"memLimit"    description:"内存限制"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"修改时间"`
}
