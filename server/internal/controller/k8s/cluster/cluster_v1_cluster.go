package cluster

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"calmk8s/api/k8s/cluster/v1"
)

func (c *ControllerV1) ListCluster(ctx context.Context, req *v1.ListClusterReq) (res *v1.ListClusterRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerV1) CreateCluster(ctx context.Context, req *v1.CreateClusterReq) (res *v1.CreateClusterRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
