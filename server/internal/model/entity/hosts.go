// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Hosts is the golang structure for table hosts.
type Hosts struct {
	Id              int64       `json:"id"              description:"主键"`
	Timestamp       *gtime.Time `json:"timestamp"       description:"时间戳"`
	Uptime          string      `json:"uptime"          description:"运行时间"`
	Hostname        string      `json:"hostname"        description:"主机名"`
	Os              string      `json:"os"              description:"操作系统"`
	Platform        string      `json:"platform"        description:"平台"`
	PlatformVersion string      `json:"platformVersion" description:"平台版本"`
	KernelVersion   string      `json:"kernelVersion"   description:"内核版本"`
	KernelArch      string      `json:"kernelArch"      description:"内核架构"`
	CreatedAt       *gtime.Time `json:"createdAt"       description:"创建时间"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       description:"修改时间"`
}
