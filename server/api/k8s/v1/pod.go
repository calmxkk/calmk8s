package v1

import (
	"calmk8s/internal/model/input/k8sin"

	"github.com/gogf/gf/v2/frame/g"
)

type CreatePodReq struct {
	g.Meta `path:"/cluster/pod/create" method:"post" tags:"clister" summary:"获得集群列表"`
	k8sin.Pod
}

type CreatePodByYamlRes struct {
	*k8sin.Pod
}

type CreatePodRes struct {
	*k8sin.Pod
}

type GetPodReq struct {
	*k8sin.Pod
}
