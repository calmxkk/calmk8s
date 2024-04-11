package main

import (
	_ "calmk8s/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"calmk8s/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
