// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Containers is the golang structure of table ck8s_containers for DAO operations like Where/Data.
type Containers struct {
	g.Meta      `orm:"table:ck8s_containers, do:true"`
	Id          interface{} // 主键
	Timestamp   *gtime.Time // 时间戳
	ContainerId interface{} // 容器ID
	Name        interface{} // 名称
	Image       interface{} // 镜像
	Ip          interface{} // IP地址
	State       interface{} // 状态
	Uptime      interface{} // 运行时间
	CpuPercent  interface{} // CPU使用率
	MemPercent  interface{} // 内存使用率
	MemUsage    interface{} // 内存使用量
	MemLimit    interface{} // 内存限制
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
}
