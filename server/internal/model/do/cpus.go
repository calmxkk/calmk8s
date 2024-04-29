// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Cpus is the golang structure of table ck8s_cpus for DAO operations like Where/Data.
type Cpus struct {
	g.Meta     `orm:"table:ck8s_cpus, do:true"`
	Id         interface{} // 主键
	Timestamp  *gtime.Time // 时间戳
	CpuPercent interface{} // CPU使用百分比
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 修改时间
}
