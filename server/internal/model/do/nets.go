// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Nets is the golang structure of table ck8s_nets for DAO operations like Where/Data.
type Nets struct {
	g.Meta    `orm:"table:ck8s_nets, do:true"`
	Id        interface{} // 主键
	Timestamp *gtime.Time // 时间戳
	Ethernet  interface{} // 以太网接口
	NetRecv   interface{} // 网络接收量
	NetSend   interface{} // 网络发送量
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}
