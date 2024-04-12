package pod

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "calmk8s/api/k8s/pod/v1"
)

func (c *ControllerV1) CreatePod(ctx context.Context, req *v1.CreatePodReq) (res *v1.CreatePodRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *ControllerV1) GetPod(ctx context.Context, req *v1.GetPodReq) (res *v1.GetPodRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerV1) ListPod(ctx context.Context, req *v1.ListPodReq) (res *v1.ListPodRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
