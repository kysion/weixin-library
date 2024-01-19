package merchant

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/weixin-library/internal/logic/internal/weixin"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
	"github.com/kysion/weixin-library/weixin_utility"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	payments_jspai "github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"log"
)

/*
	微信支付 - 直连模式
*/

// JsapiCreateOrderByDirect JsApi 支付下单 - 直连模式
func (s *sWeiXinPay) JsapiCreateOrderByDirect(ctx context.Context, info *weixin_model.TradeOrder, openId string) (tradeNo string, err error) {
	sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------JSAPI 创建支付订单，预下单 ------- ", "WeiXin-Pay")

	appId := weixin_utility.GetAppIdFormContext(ctx) // 绑定的AppId

	// 获取商家应用
	subApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)

	// 通过AppId拿到特约商户商户号
	//subMerchant, err := weixin_service.PaySubMerchant().GetPaySubMerchantByAppId(ctx, appId)
	//if err != nil {
	//	return "", sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的特约商户", "WeiXin-Pay")
	//}

	// 通过SpMchId拿到微信支付服务商商户号
	//spMerchant, err := weixin_service.PayMerchant().GetPayMerchantByMchid(ctx, subMerchant.SpMchid)
	spMerchant, err := weixin_service.PayMerchant().GetPayMerchantByAppId(ctx, appId)
	if err != nil {
		return "", sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的商户号", "WeiXin-Pay")
	}

	//subMchId := subMerchant.SubMchid // 特约商家商户号
	//spMchId := spMerchant.Mchid      // 服务商商户号
	mchid := gconv.String(spMerchant.Mchid)
	payClient, _ := weixin.NewPayClient(ctx, gconv.String(mchid), spMerchant.PayPrivateKeyPem, spMerchant.CertSerialNumber, spMerchant.ApiV3Key)

	// 微信支付服务商商户号  --> 特约商户XXX ...

	svc := payments_jspai.JsapiApiService{Client: payClient}

	req := payments_jspai.PrepayRequest{
		//SpAppid:     core.String(spMerchant.AppId),
		//SpMchid:     core.String(gconv.String(spMchId)),
		//SubAppid:    core.String(subMerchant.SubAppid),
		//SubMchid:    core.String(gconv.String(subMchId)),

		Appid:         core.String(appId),
		Mchid:         core.String(mchid),
		Description:   core.String(info.Order.ProductName),
		OutTradeNo:    core.String(gconv.String(info.Order.Id)), // out_trade_no = order.Id
		TimeExpire:    nil,                                      // 交易结束时间。订单失效时间
		Attach:        nil,                                      // 附加信息
		NotifyUrl:     core.String(subApp.NotifyUrl),            // 支付成功后的异步通知地址
		GoodsTag:      nil,                                      // 订单优惠标记
		LimitPay:      nil,                                      // 指定支付方式
		SupportFapiao: nil,                                      // 电子发票入口开放标识：传入true时，支付成功消息和支付详情页将出现开票入口。需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效。
		Amount: &payments_jspai.Amount{ // 金额
			Total:    core.Int64(gconv.Int64(info.Amount)),
			Currency: core.String("CNY"),
		},
		Payer: &payments_jspai.Payer{ // 支付者，二选一
			//SpOpenid: core.String(openId),// 用户在服务商AppId下面的唯一标识
			//SubOpenid: core.String(openId), // 用户在子商户AppId下面的唯一标识
			Openid: core.String(openId),
		},
		Detail: nil, // 优惠功能
		SceneInfo: &payments_jspai.SceneInfo{ // 支付场景信息
			PayerClientIp: core.String(g.RequestFromCtx(ctx).Request.RemoteAddr), // 用户终端IP
			//DeviceId:      ,	// 商户端设备号
			StoreInfo: nil,
		},
		SettleInfo: &payments_jspai.SettleInfo{ // TODO 是否指定分账
			//ProfitSharing: core.Bool(true),
			ProfitSharing: core.Bool(false),
		},
	}

	log.Println("微信JASAPI支付下单数据：", req)

	// 这里是预下单
	resp, _, err := svc.Prepay(ctx, req) // wx18150015642076d683b4336866f9370000

	if err == nil {
		log.Println("支付订单prepay_id: ", resp)
	} else {
		log.Println(err)
		return "", sys_service.SysLogs().ErrorSimple(ctx, err, "支付订单下单失败！", "WeiXin-Pay")
	}

	return *resp.PrepayId, nil
}
