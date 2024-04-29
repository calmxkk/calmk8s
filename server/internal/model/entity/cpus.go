// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Cpus is the golang structure for table cpus.
type Cpus struct {
	Id         int64       `json:"id"         description:"主键"`
	Timestamp  *gtime.Time `json:"timestamp"  description:"时间戳"`
	CpuPercent float64     `json:"cpuPercent" description:"CPU使用百分比"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"修改时间"`
}
