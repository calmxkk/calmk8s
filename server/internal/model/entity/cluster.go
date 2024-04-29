// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Cluster is the golang structure for table cluster.
type Cluster struct {
	Id        int64       `json:"id"        ` // 主键
	Name      string      `json:"name"      ` // 集群名称
	ConfigStr string      `json:"configStr" ` // 配置文件
	CreatedBy *gtime.Time `json:"createdBy" ` // 创建用户
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 修改时间
}
