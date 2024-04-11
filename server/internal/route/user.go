package route

import (
	"calmk8s/internal/controller/user"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

func User(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/user", func(group *ghttp.RouterGroup) {

		group.Middleware()
		group.Bind(
			user.NewV1(),
		)
	})

}
