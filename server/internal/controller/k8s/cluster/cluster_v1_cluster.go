package cluster

import (
	"calmk8s/api/k8s/cluster/v1"
	"calmk8s/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
)

func (c *ControllerV1) ListCluster(ctx context.Context, req *v1.ListClusterReq) (res *v1.ListClusterRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerV1) CreateCluster(ctx context.Context, req *v1.CreateClusterReq) (res *v1.CreateClusterRes, err error) {
	req.Authentication.ConfigFileContent = gfile.GetBytes(gfile.Join(gfile.Pwd(), "kubeconfig"))
	req.Authentication.Mode = "configfile"
	err = service.Cluster().Create(ctx, &req.Cluster)
	return nil, err
}
