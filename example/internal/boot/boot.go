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
	"github.com/kysion/weixin-library/weixin_controller/merchant"
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
							//sys_controller.SysSms,
							// 地区
							sys_controller.SysArea,
						)
					})
				})

				// 微信网关
				group.Bind(
					weixin_controller.WeiXin.WeiXinServices,     // 消息接收
					weixin_controller.WeiXin.WeiXinCallback,     // 网关回调 Get
					weixin_controller.WeiXin.WeiXinCallbackPost, // 网关回调 Post
					weixin_controller.WeiXin.NotifyServices,     // 支付异步通知
					weixin_controller.WeiXin.CheckSignature,     // 微信接入校验，设置服务器配置需要验证

					// 异步通知
					//alipay_controller.MerchantNotify.NotifyServices,

					// 刷新授权应用的Token
					merchant.MerchantService.RefreshToken,

					// 商家授权
					merchant.MerchantService.AppAuthReq,

					// 商家授权回调地址
					merchant.MerchantService.AppAuthRes,

					// 用户授权（公众号）
					merchant.MerchantService.UserAuth,

					// 用户授权（小程序） 会额外包装统一的 /appId/login 接口

					// 用户授权回调地址（公众号）
					merchant.MerchantService.UserAuthRes,

					// 用户登陆（小程序）
					merchant.MerchantService.UserLogin,

					// 获取用户信息
					merchant.UserInfo.GetUserInfo,
				)

				group.Group("/pay", func(group *ghttp.RouterGroup) {
					// 微信支付
					group.Bind(merchant.WeiXinPay)
				})

				// 分账
				group.Group("/sub_account", func(group *ghttp.RouterGroup) {
					group.Bind(merchant.SubAccount)
				})

				// 商户进件
				group.Group("/sub_merchant", func(group *ghttp.RouterGroup) {
					group.Bind(merchant.SubMerchant)
				})

				// 小程序开发管理
				group.Group("/app_version_manager", func(group *ghttp.RouterGroup) {
					group.Bind(merchant.AppVersionManager)
				})

				// 消息
				group.Group("/message", func(group *ghttp.RouterGroup) {
					// 消息 【小程序】
					group.Group("/tiny_app", func(group *ghttp.RouterGroup) {
						// 订阅消息 【小程序】
						group.Bind(merchant.SubscribeMessage)
					})

					group.Group("/template", func(group *ghttp.RouterGroup) {
						// 订阅消息模板管理 【小程序 & 公众号】
						group.Bind(weixin_controller.SubscribeMessageTemplate)
					})
				})

				// 微信支付
				group.Group("/weixin_pay", func(group *ghttp.RouterGroup) {
					// 微信支付商户号
					group.Bind(weixin_controller.WeiXinPayMerchant)

					// 微信支付特约商户
					group.Bind(weixin_controller.WeiXinPaySubMerchant)
				})

				group.Group("/third_app", func(group *ghttp.RouterGroup) {
					// 服务商应用配置
					group.Bind(weixin_controller.WeiXinThirdAppConfig)

					// 服务商服务 （WeiXin）
					group.Bind(weixin_controller.ThirdService)
				})

				group.Group("/merchant_app", func(group *ghttp.RouterGroup) {
					group.Bind(weixin_controller.WeiXinMerchantAppConfig)
				})

			})

			s.Run()
			return nil
		},
	}
)
