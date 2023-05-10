package boot

import (
	"context"
	_ "github.com/SupenBysz/gf-admin-community"
	"github.com/SupenBysz/gf-admin-community/sys_controller"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	_ "github.com/kysion/weixin-library/example/internal/boot/internal"
	"github.com/kysion/weixin-library/weixin_controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			s.Group("/weixin", func(group *ghttp.RouterGroup) {
				// 注册中间件
				group.Middleware(
					sys_service.Middleware().CTX,
					sys_service.Middleware().ResponseHandler,
				)

				// 不需要鉴权，但是需要登录的路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					// 注册中间件
					group.Middleware(
						sys_service.Middleware().Auth,
					)
					// 文件上传
					group.Group("/common/sys_file", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysFile) })
				})

				// 匿名路由绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					// 鉴权：登录，注册，找回密码等
					group.Group("/auth", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.Auth) })
					// 图型验证码、短信验证码、地区
					group.Group("/common", func(group *ghttp.RouterGroup) {
						group.Bind(
							// 图型验证码
							sys_controller.Captcha,
							// 短信验证码
							sys_controller.SysSms,
							// 地区
							sys_controller.SysArea,
						)
					})
				})

				// 微信网关
				group.Bind(
					weixin_controller.WeiXin.WeiXinServices,
					weixin_controller.WeiXin.WeiXinCallback,
					weixin_controller.WeiXin.CheckSignature,
				)

				group.Group("/third_app", func(group *ghttp.RouterGroup) {
					// 服务商应用配置
					group.Bind(weixin_controller.WeiXinThirdAppConfig)

					// 服务商服务 （WeiXin）
					group.Bind(weixin_controller.ThirdService)
				})

				group.Group("/merchant_app", func(group *ghttp.RouterGroup) {
					group.Bind(weixin_controller.WeiXinMerchantAppConfig)
				})

				// 引入用户进入授权页
				// https://weixin.jditco.com/weixin/gateway.call
				//group.GET("/:appId/gateway.auth", func(r *ghttp.Request) {
				//	// 通过appId将具体第三方应用配置信息从数据库获取出来
				//
				//	appId := g.RequestFromCtx(r.Context()).Get("appId").Int64()
				//	app, _ := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, appId)
				//
				//	// 4.获取与授权码
				//	proAuthCodeReq := weixin_model.ProAuthCodeReq{
				//		ComponentAppid: weixin_consts.Global.AppId,
				//		// ComponentAccessToken: token,  // 不能写json结构体里面，一半数据写在上面url上，一半数据写在json结构体
				//	}
				//	encode, _ := gjson.Encode(proAuthCodeReq)
				//	proAuthCodeUrl := "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=" + token
				//	fmt.Println(string(encode))
				//
				//	proAuthCode := g.Client().PostContent(ctx, proAuthCodeUrl, encode)
				//	proAuthCodeRes := weixin_model.ProAuthCodeRes{}
				//	gjson.DecodeTo(proAuthCode, &proAuthCodeRes)
				//	/*
				//		{
				//			"pre_auth_code": "preauthcode@@@pxvu7JW0hDQqNf38HcEXF6ejB4pnzVnA_GXlqqb1XcSmS3GjEhy-TfJOIqjAODk3MmmTZpNHi7Brgc_ugz0RCg",
				//			"expires_in": 1800
				//		}
				//	*/
				//
				//	// 5.引导用户进入授权页面
				//	redirect_url := gurl.Encode("https://weixin.jditco.com/weixin/$APPID$/gateway.callback")
				//	//authUrl := "https://mp.weixin.qq.com/cgi-bin/componentloginpage?" +
				//	authUrl := "https://mp.weixin.qq.com/safe/bindcomponent?" +
				//		"component_appid=" + weixin_consts.Global.AppId +
				//		"&no_scan=1&auth_type=3&pre_auth_code=" + proAuthCodeRes.PreAuthCode +
				//		"&redirect_url=" + redirect_url
				//	fmt.Println("授权全链接：\n", authUrl)
				//
				//	r.Response.Header().Set("referer", "https://douyin.jditco.com/weixin/gateway.services")
				//
				//	r.Response.RedirectTo(authUrl)
				//	//r.Response.Header().Set("Content-Type", "text/html; charset=UTF-8")
				//	//r.Response.WriteTplContent(`<html lang="zh"><head><meta charset="utf-8"></head><body>测试页面：<a href="{{.url}}">{{.label}}</a></body></html>`, g.Map{
				//	//	"url":   authUrl,
				//	//	"label": "授权",
				//	//})
				//})

			})

			s.Run()
			return nil
		},
	}
)
