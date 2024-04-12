// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package pod

import (
	"context"

	"calmk8s/api/k8s/pod/v1"
)

type IPodV1 interface {
	CreatePod(ctx context.Context, req *v1.CreatePodReq) (res *v1.CreatePodRes, err error)
	CreatePodByYaml(ctx context.Context, req *v1.CreatePodByYamlReq) (res *v1.CreatePodByYamlRes, err error)
	GetPod(ctx context.Context, req *v1.GetPodReq) (res *v1.GetPodRes, err error)
	ListPod(ctx context.Context, req *v1.ListPodReq) (res *v1.ListPodRes, err error)
}
