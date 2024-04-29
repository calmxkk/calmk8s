// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Cluster is the golang structure for table cluster.
type Cluster struct {
	Id        int64       `json:"id"        description:"主键"`
	Name      string      `json:"name"      description:"集群名称"`
	ConfigStr string      `json:"configStr" description:"配置文件"`
	CreatedBy *gtime.Time `json:"createdBy" description:"创建用户"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改时间"`
}
