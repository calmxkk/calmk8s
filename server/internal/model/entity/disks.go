// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Disks is the golang structure for table disks.
type Disks struct {
	Id          int64       `json:"id"          ` // 主键
	Timestamp   *gtime.Time `json:"timestamp"   ` // 时间戳
	Device      string      `json:"device"      ` // 设备名称
	DiskPercent float64     `json:"diskPercent" ` // 磁盘使用百分比
	DiskTotal   float64     `json:"diskTotal"   ` // 磁盘总量
	DiskUsed    float64     `json:"diskUsed"    ` // 磁盘已使用量
	DiskRead    float64     `json:"diskRead"    ` // 磁盘读取量
	DiskWrite   float64     `json:"diskWrite"   ` // 磁盘写入量
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 修改时间
}
