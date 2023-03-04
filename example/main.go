package main

import (
	_ "github.com/SupenBysz/gf-admin-community"
	_ "github.com/SupenBysz/gf-admin-company-modules"
	"github.com/gogf/gf/v2/os/gctx"
	_ "github.com/kysion/base-library/base_hook"
	"github.com/kysion/weixin-library/example/internal/boot"
	_ "github.com/kysion/weixin-library/internal/logic"
)

func main() {
	boot.Main.Run(gctx.New())
}
