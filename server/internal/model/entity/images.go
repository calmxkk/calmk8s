// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Images is the golang structure for table images.
type Images struct {
	Id        int64       `json:"id"        description:"主键"`
	Timestamp *gtime.Time `json:"timestamp" description:"时间戳"`
	ImageId   string      `json:"imageId"   description:"镜像ID"`
	Name      string      `json:"name"      description:"名称"`
	Tag       string      `json:"tag"       description:"标签"`
	Created   string      `json:"created"   description:"创建时间"`
	Size      string      `json:"size"      description:"大小"`
	Number    int         `json:"number"    description:"数量"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改时间"`
}
