// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Disks is the golang structure for table disks.
type Disks struct {
	Id          int64       `json:"id"          description:"主键"`
	Timestamp   *gtime.Time `json:"timestamp"   description:"时间戳"`
	Device      string      `json:"device"      description:"设备名称"`
	DiskPercent float64     `json:"diskPercent" description:"磁盘使用百分比"`
	DiskTotal   float64     `json:"diskTotal"   description:"磁盘总量"`
	DiskUsed    float64     `json:"diskUsed"    description:"磁盘已使用量"`
	DiskRead    float64     `json:"diskRead"    description:"磁盘读取量"`
	DiskWrite   float64     `json:"diskWrite"   description:"磁盘写入量"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"修改时间"`
}
