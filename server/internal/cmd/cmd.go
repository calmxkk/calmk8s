package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"calmk8s/internal/router"
	"calmk8s/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)

				router.Admin(ctx, group)
				router.K8s(ctx, group)
			})

			waitCleanChan := make(chan struct{}, 1)
			signalListen(ctx, signalHandlerForOverall)

			serverWg.Add(1)

			systemTaskCtx, systemTaskcancel := context.WithCancel(ctx)
			service.SystemTask().Run(systemTaskCtx)

			go func() {
				defer serverWg.Done()
				<-serverCloseSignal
				g.Log().Info(ctx, "stop system task grouting")
				systemTaskcancel()
				_ = s.Shutdown()
				g.Log().Info(ctx, "stop server")
				waitCleanChan <- struct{}{}
			}()

			s.Run()
			<-waitCleanChan
			g.Log().Info(ctx, "clean all server, web exit")
			return nil
		},
	}
)
