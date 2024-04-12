package v1

import (
	"calmk8s/internal/model/input/k8sin"

	"github.com/gogf/gf/v2/frame/g"
)

type CreatePodReq struct {
	g.Meta `path:"/cluster/pod/create" method:"post" tags:"clister" summary:"创建pod"`
	Mode   string `json:"mode" v:"in:0,1" dc:"0:yaml;1:config"`
	k8sin.Pod
}

type CreatePodRes struct {
	k8sin.Pod
}

type GetPodReq struct {
	g.Meta `path:"/cluster/pod/get" method:"post" tags:"clister" summary:"创建pod"`
	*k8sin.Pod
}

type GetPodRes struct {
	*k8sin.Pod
}

type ListPodReq struct {
	g.Meta `path:"/cluster/pod/list" method:"post" tags:"clister" summary:"创建pod"`
}

type ListPodRes struct {
	*k8sin.Pod
}
