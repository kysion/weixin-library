package merchant

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/weixin-library/internal/logic/internal/weixin"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
	"github.com/kysion/weixin-library/weixin_utility"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/certificates"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/jsapi"
	"log"
)

/*
	微信支付 - 服务商代调用模式
*/

// 微信支付（模式：第三方模式 + 服务商代调用）

// DownloadCertificates 测试SDK ，下载微信支付平台证书
func (s *sWeiXinPay) DownloadCertificates(ctx context.Context, appID ...string) (*certificates.DownloadCertificatesResponse, error) {
	appId := ""
	if appID[0] == "" && !gstr.HasPrefix(appID[0], "wx") {
		appId = weixin_utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId
	}
	appId = appID[0]

	// 通过AppId拿到特约商户商户号
	subMerchant, err := weixin_service.PaySubMerchant().GetPaySubMerchantByAppId(ctx, appId)
	//if err != nil {
	//	return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的特约商户", "WeiXin-Pay")
	//}

	spMerchant, err := weixin_service.PayMerchant().GetPayMerchantByAppId(ctx, appId)

	//spMerchant, err := weixin_service.PayMerchant().GetPayMerchantByMchid(ctx, subMerchant.SpMchid)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的商户号", "WeiXin-Pay")
	}

	mchID := ""
	if subMerchant != nil {
		mchID = gconv.String(subMerchant.SpMchid) // 商户号
	}
	if spMerchant != nil {
		mchID = gconv.String(spMerchant.Mchid) // 商户号
	}

	//mchID string = "1642565036" // 商户号
	//mchCertificateSerialNumber string = "298D4028EC0F48748DF237A226DB4D5281EFE86E" // 商户证书序列号
	//mchAPIv3Key                string = "655957AD45E5FE85F1BF3B9E0D82B96D"         // 商户APIv3密钥

	client, _ := weixin.NewPayClient(ctx, mchID, spMerchant.PayPrivateKeyPem, spMerchant.CertSerialNumber, spMerchant.ApiV3Key)
	//client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
	}

	// 发送请求，以下载微信支付平台证书为例
	// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
	svc := certificates.CertificatesApiService{Client: client}
	resp, result, err := svc.DownloadCertificates(ctx)
	log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)

	if err == nil {
		log.Println(resp)
	} else {
		log.Println(err)
		return &certificates.DownloadCertificatesResponse{}, sys_service.SysLogs().ErrorSimple(ctx, err, "支付订单下单失败！", "WeiXin-Pay")
	}

	return resp, nil
}

// JsapiCreateOrder JsApi 支付下单 - 服务商待调用
func (s *sWeiXinPay) JsapiCreateOrder(ctx context.Context, info *weixin_model.TradeOrder, openId string) (tradeNo string, err error) {
	sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------JSAPI 创建支付订单，预下单 ------- ", "WeiXin-Pay")

	appId := weixin_utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId

	// 获取商家应用
	subApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)

	// 通过AppId拿到特约商户商户号
	subMerchant, err := weixin_service.PaySubMerchant().GetPaySubMerchantByAppId(ctx, appId)
	if err != nil {
		return "", sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的特约商户", "WeiXin-Pay")
	}

	// 通过SpMchId拿到微信支付服务商商户号
	spMerchant, err := weixin_service.PayMerchant().GetPayMerchantByMchid(ctx, subMerchant.SpMchid)
	if err != nil {
		return "", sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的商户号", "WeiXin-Pay")
	}

	subMchId := subMerchant.SubMchid // 特约商家商户号
	spMchId := spMerchant.Mchid      // 服务商商户号

	payClient, _ := weixin.NewPayClient(ctx, gconv.String(spMchId), spMerchant.PayPrivateKeyPem, spMerchant.CertSerialNumber, spMerchant.ApiV3Key)

	// 微信支付服务商商户号  --> 特约商户XXX ...

	svc := jsapi.JsapiApiService{Client: payClient}

	req := jsapi.PrepayRequest{
		SpAppid:     core.String(spMerchant.AppId),
		SpMchid:     core.String(gconv.String(spMchId)),
		SubAppid:    core.String(subMerchant.SubAppid),
		SubMchid:    core.String(gconv.String(subMchId)),
		Description: core.String(info.Order.ProductName),
		OutTradeNo:  core.String(gconv.String(info.Order.Id)), // out_trade_no = order.Id
		NotifyUrl:   core.String(subApp.NotifyUrl),            // 支付成功后的异步通知地址

		Amount: &jsapi.Amount{ // 金额
			Total:    core.Int64(gconv.Int64(info.Amount)),
			Currency: core.String("CNY"),
		},

		Payer: &jsapi.Payer{ // 支付者，二选一
			//SpOpenid: core.String(openId),// 用户在服务商AppId下面的唯一标识
			SubOpenid: core.String(openId), // 用户在子商户AppId下面的唯一标识
		},

		SceneInfo: &jsapi.SceneInfo{ // 支付场景信息
			PayerClientIp: core.String(g.RequestFromCtx(ctx).Request.RemoteAddr), // 用户终端IP
			//DeviceId:      ,	// 商户端设备号
			StoreInfo: nil,
		},

		SettleInfo: &jsapi.SettleInfo{ // TODO 是否指定分账
			ProfitSharing: core.Bool(true),
			//ProfitSharing: core.Bool(false),
		},

		//TimeExpire:  core.Time(gtime.Now().Add(time.Minute * 5)), // 订单失效失效时间
		//Attach:        core.String(""),// 附加数据
		//GoodsTag:      core.String(), // 订单优惠标记
		//LimitPay:      core.String(), // 指定支付方式
		//SupportFapiao: core.String(), // 开票相关
		//Detail:     &jsapi.Detail{}, // 优惠功能
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

// TODO 补充下单 （Native、JSAPI、APP等不同场景生成交易串调起支付。）

// QueryOrderByIdMchID 查询订单 （1.根据tradeNo 2.根据mchId）
func (s *sWeiXinPay) QueryOrderByIdMchID(ctx context.Context, transactionId string, appID ...string) (*weixin_model.TradeOrderRes, error) {
	appId := ""
	if appID[0] == "" && !gstr.HasPrefix(appID[0], "wx") {
		appId = weixin_utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId
	}
	appId = appID[0]
	//
	//appId := weixin_utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId

	// 通过AppId拿到特约商户商户号
	subMerchant, err := weixin_service.PaySubMerchant().GetPaySubMerchantByAppId(ctx, appId)
	if err != nil {
		return &weixin_model.TradeOrderRes{}, sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的特约商户", "WeiXin-Pay")
	}

	// 通过SpMchId拿到微信支付服务商商户号
	spMerchant, err := weixin_service.PayMerchant().GetPayMerchantByMchid(ctx, subMerchant.SpMchid)
	if err != nil {
		return &weixin_model.TradeOrderRes{}, sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的商户号", "WeiXin-Pay")
	}

	payClient, _ := weixin.NewPayClient(ctx, gconv.String(spMerchant.Mchid), spMerchant.PayPrivateKeyPem, spMerchant.CertSerialNumber, spMerchant.ApiV3Key)

	svc := jsapi.JsapiApiService{Client: payClient}

	resp, result, err := svc.QueryOrderById(ctx,
		jsapi.QueryOrderByIdRequest{
			//TransactionId: core.String("4200000985202103031441826014"),
			TransactionId: core.String(transactionId),
			SpMchid:       core.String(gconv.String(spMerchant.Mchid)),
			SubMchid:      core.String(gconv.String(subMerchant.SubMchid)),
		},
	)

	if err == nil {
		log.Println(resp)
	} else {
		log.Println(err)
	}

	ret := weixin_model.TradeOrderRes{
		Transaction: resp,
		APIResult:   result,
	}

	return &ret, err
}

// QueryOrderByIdOutTradeNo 根据支付编号查询订单
func (s *sWeiXinPay) QueryOrderByIdOutTradeNo(ctx context.Context, outTradeNo string, appID ...string) (*weixin_model.TradeOrderRes, error) {
	appId := ""
	if appID[0] == "" && !gstr.HasPrefix(appID[0], "wx") {
		appId = weixin_utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId
	}
	appId = appID[0]

	// 通过AppId拿到特约商户商户号
	subMerchant, err := weixin_service.PaySubMerchant().GetPaySubMerchantByAppId(ctx, appId)
	if err != nil {
		return &weixin_model.TradeOrderRes{}, sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的特约商户", "WeiXin-Pay")
	}

	// 通过SpMchId拿到微信支付服务商商户号
	spMerchant, err := weixin_service.PayMerchant().GetPayMerchantByMchid(ctx, subMerchant.SpMchid)
	if err != nil {
		return &weixin_model.TradeOrderRes{}, sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的商户号", "WeiXin-Pay")
	}

	payClient, _ := weixin.NewPayClient(ctx, gconv.String(spMerchant.Mchid), spMerchant.PayPrivateKeyPem, spMerchant.CertSerialNumber, spMerchant.ApiV3Key)

	svc := jsapi.JsapiApiService{Client: payClient}

	resp, result, err := svc.QueryOrderByOutTradeNo(ctx,
		jsapi.QueryOrderByOutTradeNoRequest{
			//TransactionId: core.String("4200000985202103031441826014"),
			OutTradeNo: core.String(outTradeNo),
			SpMchid:    core.String(gconv.String(spMerchant.Mchid)),
			SubMchid:   core.String(gconv.String(subMerchant.SubMchid)),
		},
	)

	if err == nil {
		log.Println(resp)
	} else {
		log.Println("查询订单失败：outTRradeNo：", outTradeNo, err)
		return &weixin_model.TradeOrderRes{}, sys_service.SysLogs().ErrorSimple(ctx, err, "查询支付订单失败！", "WeiXin-Pay")
	}

	ret := weixin_model.TradeOrderRes{
		Transaction: resp,
		APIResult:   result,
	}

	return &ret, err
}

// 接收支付结果通知接口 NotifyUrl

// CloseOrder 关闭订单接口
func (s *sWeiXinPay) CloseOrder(ctx context.Context, outTradeNo string, appID ...string) (bool, error) {
	//appId := weixin_utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId
	appId := ""
	if appID[0] == "" && !gstr.HasPrefix(appID[0], "wx") {
		appId = weixin_utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId
	}
	appId = appID[0]

	// 通过AppId拿到特约商户商户号
	subMerchant, err := weixin_service.PaySubMerchant().GetPaySubMerchantByAppId(ctx, appId)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的特约商户", "WeiXin-Pay")
	}

	// 通过SpMchId拿到微信支付服务商商户号
	spMerchant, err := weixin_service.PayMerchant().GetPayMerchantByMchid(ctx, subMerchant.SpMchid)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的商户号", "WeiXin-Pay")
	}

	payClient, _ := weixin.NewPayClient(ctx, gconv.String(spMerchant.Mchid), spMerchant.PayPrivateKeyPem, spMerchant.CertSerialNumber, spMerchant.ApiV3Key)

	svc := jsapi.JsapiApiService{Client: payClient}
	result, err := svc.CloseOrder(ctx,
		jsapi.CloseOrderRequest{
			OutTradeNo: core.String(outTradeNo),
			SpMchid:    core.String(gconv.String(subMerchant.SpMchid)),
			SubMchid:   core.String(gconv.String(subMerchant.SubMchid)),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("关闭订单微信订单失败：call CloseOrder err:%s", err)
		return false, err
	} else {
		// 处理返回结果
		log.Printf("关闭订单微信订单：status=%d", result.Response.StatusCode)
	}

	return result.Response.StatusCode == 200, nil
}

// DownloadAccountBill 账单下载接口
func (s *sWeiXinPay) DownloadAccountBill(ctx context.Context, mchId string) {

}
