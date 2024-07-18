package main

import (
	"github.com/crisis111/jwt_gf/example/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
