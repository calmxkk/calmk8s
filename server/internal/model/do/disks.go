// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Disks is the golang structure of table ck8s_disks for DAO operations like Where/Data.
type Disks struct {
	g.Meta      `orm:"table:ck8s_disks, do:true"`
	Id          interface{} // 主键
	Timestamp   *gtime.Time // 时间戳
	Device      interface{} // 设备名称
	DiskPercent interface{} // 磁盘使用百分比
	DiskTotal   interface{} // 磁盘总量
	DiskUsed    interface{} // 磁盘已使用量
	DiskRead    interface{} // 磁盘读取量
	DiskWrite   interface{} // 磁盘写入量
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
}
