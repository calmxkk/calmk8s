package pod

import (
	"calmk8s/internal/model/input/k8sin"
	"calmk8s/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "calmk8s/api/k8s/pod/v1"
)

func (c *ControllerV1) CreatePod(ctx context.Context, req *v1.CreatePodReq) (res *v1.CreatePodRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *ControllerV1) GetPod(ctx context.Context, req *v1.GetPodReq) (res *v1.GetPodRes, err error) {
	res = &v1.GetPodRes{}
	res.Data, err = service.Pod().GetPodByName(ctx, &req.GetPodInp)
	if err != nil {
		g.Log().Debugf(ctx, "get pod error, %v", err)
		return nil, err
	}
	return res, err
}
func (c *ControllerV1) ListPod(ctx context.Context, req *v1.ListPodReq) (res *v1.ListPodRes, err error) {
	res = &v1.ListPodRes{
		Data: make([]*k8sin.GetPodRes, 0),
	}

	data, err := service.Pod().List(ctx, &req.GetPodInp)
	if err != nil {
		g.Log().Debugf(ctx, "list all pod error, %v", err)
	}
	res.Data = append(res.Data, data...)
	return res, err
}
