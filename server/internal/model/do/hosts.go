// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Hosts is the golang structure of table ck8s_hosts for DAO operations like Where/Data.
type Hosts struct {
	g.Meta          `orm:"table:ck8s_hosts, do:true"`
	Id              interface{} // 主键
	Timestamp       *gtime.Time // 时间戳
	Uptime          interface{} // 运行时间
	Hostname        interface{} // 主机名
	Os              interface{} // 操作系统
	Platform        interface{} // 平台
	PlatformVersion interface{} // 平台版本
	KernelVersion   interface{} // 内核版本
	KernelArch      interface{} // 内核架构
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 修改时间
}
