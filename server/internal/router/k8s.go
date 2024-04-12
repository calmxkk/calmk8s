package router

import (
	"calmk8s/internal/controller/k8s/cluster"
	"calmk8s/internal/controller/k8s/pod"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

func K8s(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/k8s", func(group *ghttp.RouterGroup) {

		group.Middleware()
		group.Bind(
			pod.NewV1(),
			cluster.NewV1(),
		)
	})
}
