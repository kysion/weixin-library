package boot

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/kysion/kys-weixin-library/weixin_controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/weixin", func(group *ghttp.RouterGroup) {
				group.Bind(
					weixin_controller.WeiXin.WeiXinServices,
					weixin_controller.WeiXin.WeiXinCallback,
					weixin_controller.WeiXin.CheckSignature,
				)
				// 直接通过回调获取用户信息

				group.GET("/gateway.call", func(r *ghttp.Request) {
					r.Response.RedirectTo("https://openauth.alipay.com/oauth2/publicAppAuthorize.htm?app_id=2021003179632101&scope=auth_user&redirect_uri=https%3A%2F%2Falipay.jditco.com%2Falipay%2Fgateway.callback")
				})

				//group.GET("/gateway.service", func(r *ghttp.Request) {
				//	r.Response.RedirectTo("https://weixin.jditco.com/weixin/gateway.service")
				//})

				group.POST("/api_start_push_ticket", func(r *ghttp.Request) {

				})

			})

			s.Run()
			return nil
		},
	}
)
