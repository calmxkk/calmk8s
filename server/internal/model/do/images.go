// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Images is the golang structure of table ck8s_images for DAO operations like Where/Data.
type Images struct {
	g.Meta    `orm:"table:ck8s_images, do:true"`
	Id        interface{} // 主键
	Timestamp *gtime.Time // 时间戳
	ImageId   interface{} // 镜像ID
	Name      interface{} // 名称
	Tag       interface{} // 标签
	Created   interface{} // 创建时间
	Size      interface{} // 大小
	Number    interface{} // 数量
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}
