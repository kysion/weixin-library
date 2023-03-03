package boot

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gcmd"
	_ "github.com/kysion/kys-weixin-library/example/internal/boot/internal"
	"github.com/kysion/kys-weixin-library/weixin_consts"
	"github.com/kysion/kys-weixin-library/weixin_controller"
	"github.com/kysion/kys-weixin-library/weixin_model"
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

				// 引入用户进入授权页
				// https://weixin.jditco.com/weixin/gateway.call
				group.GET("/gateway.call", func(r *ghttp.Request) {
					token := gcache.MustGet(ctx, weixin_consts.Global.AppId+"_component_access_token").String()
					if token == "" {
						// 实际从数据库找
						token = "ticket@@@Bb3RjaKczF7YiV-mdama4Qzmo6x5H72QsZSWsCSfs1fs0XiWoMF5UY7Yix_-24W9RdKXn-yHHHOKyLwD8t79FA"
					}
					// 4.获取与授权码
					proAuthCodeUrl := "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=" + token +
						"&component_appid=" + weixin_consts.Global.AppId

					proAuthCode := g.Client().PostContent(ctx, proAuthCodeUrl)
					proAuthCodeRes := weixin_model.ProAuthCodeRes{}
					gjson.DecodeTo(proAuthCode, &proAuthCodeRes)

					// 5.引导用户进入授权页面
					redirect_url := gurl.Encode("https://weixin.jditco.com/weixin/$APPID$/gateway.callback")
					authUrl := "https://mp.weixin.qq.com/cgi-bin/componentloginpage?" +
						"component_appid=" + weixin_consts.Global.AppId +
						"&pre_auth_code=" + proAuthCodeRes.PreAuthCode +
						"&redirect_url=" + redirect_url

					r.Response.Header().Set("referer", "https://douyin.jditco.com/weixin/gateway.services")

					//content := g.Client().GetContent(ctx, authUrl)
					//
					//fmt.Println(content)
					//r.Response.RedirectTo(authUrl)
					r.Response.Header().Set("Content-Type", "text/html; charset=UTF-8")
					r.Response.WriteTplContent(`<html lang="zh"><head><meta charset="utf-8"></head><body>测试页面：<a href="{{.url}}">{{.label}}</a></body></html>`, g.Map{
						"url":   authUrl,
						"label": "授权",
					})
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
