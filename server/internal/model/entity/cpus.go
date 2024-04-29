// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Cpus is the golang structure for table cpus.
type Cpus struct {
	Id         int64       `json:"id"         ` // 主键
	Timestamp  *gtime.Time `json:"timestamp"  ` // 时间戳
	CpuPercent float64     `json:"cpuPercent" ` // CPU使用百分比
	CreatedAt  *gtime.Time `json:"createdAt"  ` // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  ` // 修改时间
}
