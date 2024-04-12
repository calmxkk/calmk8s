// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package cluster

import (
	"context"

	"calmk8s/api/k8s/cluster/v1"
)

type IClusterV1 interface {
	ListCluster(ctx context.Context, req *v1.ListClusterReq) (res *v1.ListClusterRes, err error)
	CreateCluster(ctx context.Context, req *v1.CreateClusterReq) (res *v1.CreateClusterRes, err error)
}
