package v1

import (
	"calmk8s/internal/model/input/k8sin"
	"github.com/gogf/gf/v2/frame/g"
)

type ListReq struct {
	g.Meta `path:"/ck8s/deployment/{namespace}" method:"get" tags:"deployment" summary:"获得所有deployment"`
}

type ListRes struct {
	Data []*k8sin.GetPodRes `json:"data"`
}

type GetReq struct {
	g.Meta `path:"/ck8s/deployment/{namespace}/{name}" method:"get" tags:"deployment" summary:"获得指定deployment"`
}

type GetRes struct {
	Data []*k8sin.GetPodRes `json:"data"`
}
