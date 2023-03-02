package main

import (
	_ "github.com/SupenBysz/gf-admin-community"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/kysion/kys-weixin-library/example/internal/boot"
	_ "github.com/kysion/kys-weixin-library/internal/logic"
)

func main() {
	boot.Main.Run(gctx.New())
}
