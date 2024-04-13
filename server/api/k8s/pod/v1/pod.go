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
	k8sin.Pod
}

type GetPodReq struct {
	g.Meta `path:"/cluster/pod/get" method:"post" tags:"pod" summary:"获得pod"`
	k8sin.GetPodInp
}

type GetPodRes struct {
	*k8sin.Pod
}

type ListPodReq struct {
	g.Meta `path:"/cluster/pod/list" method:"post" tags:"pod" summary:"获得所有pod"`
	k8sin.GetPodInp
}

type ListPodRes struct {
}
