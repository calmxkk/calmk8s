// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Memories is the golang structure for table memories.
type Memories struct {
	Id         int64       `json:"id"         ` // 主键
	Timestamp  *gtime.Time `json:"timestamp"  ` // 时间戳
	MemPercent float64     `json:"memPercent" ` // 内存使用百分比
	MemTotal   float64     `json:"memTotal"   ` // 内存总量
	MemUsed    float64     `json:"memUsed"    ` // 内存已使用量
	CreatedAt  *gtime.Time `json:"createdAt"  ` // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  ` // 修改时间
}
