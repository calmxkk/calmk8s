// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Memories is the golang structure of table ck8s_memories for DAO operations like Where/Data.
type Memories struct {
	g.Meta     `orm:"table:ck8s_memories, do:true"`
	Id         interface{} // 主键
	Timestamp  *gtime.Time // 时间戳
	MemPercent interface{} // 内存使用百分比
	MemTotal   interface{} // 内存总量
	MemUsed    interface{} // 内存已使用量
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 修改时间
}
