package task

import (
	"calmk8s/internal/library/procutil"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type sSystemTask struct {
	timeDuration time.Duration
}

func NewUser() *sSystemTask {
	return &sSystemTask{}
}

func (d *sSystemTask) Run(ctx context.Context) {
	d.timeDuration = g.Cfg().MustGet(ctx, "task.duration", "").Duration()
	go d.GetMemory(ctx)
	go d.GetSystemInfo(ctx)
}

func (d *sSystemTask) GetMemory(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			g.Log().Info(ctx, "grouting getMemory exit")
			return
		default:
			res, err := procutil.GetMemoryInfo(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "%v", err.Error())
				return
			}

			fmt.Printf("%s memory: total %v, used %v\n", time.Now().Format("2006-01-02 15:04:05"), procutil.FmtByte(int64(res.Total)), procutil.FmtByte(int64(res.Used)))
			time.Sleep(d.timeDuration)
		}
	}
}

func (d *sSystemTask) GetSystemInfo(ctx context.Context) {
	for {
		select {
		default:
			res, err := procutil.GetSystemInfo(ctx)
			if err != nil {
				return
			}

			fmt.Printf("%s get system Info %v\n", time.Now().Format("2006-01-02 15:04:05"), res)
			time.Sleep(d.timeDuration)
		case <-ctx.Done():
			g.Log().Info(ctx, "grouting getMemory exit")
			return
		}
	}
}
