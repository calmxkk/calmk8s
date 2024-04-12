package v1

import (
	"calmk8s/internal/model/input/k8sin"

	"github.com/gogf/gf/v2/frame/g"
)

type ListClusterReq struct {
	g.Meta `path:"/cluster/list" method:"get" tags:"clister" summary:"获得集群列表"`
}

type ListClusterRes struct {
	*k8sin.Cluster
}

type CreateClusterReq struct {
	g.Meta `path:"/cluster/create" method:"post" tags:"clister" summary:"获得集群列表"`
	k8sin.Cluster
}

type CreateClusterRes struct {
	*k8sin.Cluster
}
