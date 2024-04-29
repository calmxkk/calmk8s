package cmd

import (
	"calmk8s/utility/simple"
	"context"
	"os"
	"sync"

	"github.com/gogf/gf/v2/os/gproc"
)

var (
	serverCloseSignal = make(chan struct{}, 1)
	serverWg          = sync.WaitGroup{}
	once              sync.Once
)

// signalHandlerForOverall 关闭信号处理
func signalHandlerForOverall(sig os.Signal) {
	serverCloseSignal <- struct{}{}
}

// signalListen 信号监听
func signalListen(ctx context.Context, handler ...gproc.SigHandler) {
	simple.SafeGo(ctx, func(ctx context.Context) {
		gproc.AddSigHandlerShutdown(handler...)
		gproc.Listen()
	})
}
