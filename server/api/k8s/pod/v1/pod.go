package v1

import (
	"calmk8s/internal/model/input/k8sin"

	"github.com/gogf/gf/v2/frame/g"
)

type CreatePodReq struct {
	g.Meta `path:"/cluster/pod/create" method:"post" tags:"pod" summary:"创建pod"`
	Mode   string `json:"mode" v:"in:0,1" dc:"0:yaml;1:config"`
	k8sin.Pod
}

type CreatePodRes struct {
}

type GetPodReq struct {
	g.Meta `path:"/cluster/pod/get" method:"post" tags:"pod" summary:"获得pod"`
	k8sin.GetPodInp
}

type GetPodRes struct {
	Data *k8sin.GetPodRes
}

type ListPodReq struct {
	g.Meta `path:"/cluster/pod/list" method:"post" tags:"pod" summary:"获得所有pod"`
	k8sin.GetPodInp
}

type ListPodRes struct {
	Data []*k8sin.GetPodRes `json:"data"`
}
