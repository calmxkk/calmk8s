// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Images is the golang structure for table images.
type Images struct {
	Id        int64       `json:"id"        ` // 主键
	Timestamp *gtime.Time `json:"timestamp" ` // 时间戳
	ImageId   string      `json:"imageId"   ` // 镜像ID
	Name      string      `json:"name"      ` // 名称
	Tag       string      `json:"tag"       ` // 标签
	Created   string      `json:"created"   ` // 创建时间
	Size      string      `json:"size"      ` // 大小
	Number    int         `json:"number"    ` // 数量
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 修改时间
}
