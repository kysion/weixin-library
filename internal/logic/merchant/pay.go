package merchant

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/pay-share-library/pay_model/pay_enum"
	"github.com/kysion/pay-share-library/pay_service"
	"github.com/kysion/weixin-library/internal/logic/internal/weixin"
	"github.com/kysion/weixin-library/utility"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/certificates"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/jsapi"
	"log"
	"strconv"
	"time"
)

// 微信支付（模式：第三方模式 + 服务商代调用）
type sWeiXinPay struct {
}

func init() {
	weixin_service.RegisterWeiXinPay(NewWeiXinPay())
}

func NewWeiXinPay() *sWeiXinPay {

	result := &sWeiXinPay{}

	//result.injectHook()
	return result
}

// PayTradeCreate  1、创建交易订单   （AppId的H5是没有的，需要写死，小程序有的 ）
func (s *sWeiXinPay) PayTradeCreate(ctx context.Context, info *weixin_model.TradeOrder, openId string) (*weixin_model.PayParamsRes, error) {
	sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------WeiXin创建交易订单 ------- ", "WeiXin-Pay")
	appId := utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId

	// 商家AppId解析，获取商家应用，创建微信支付客户端
	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}

	sysUser, err := sys_service.SysUser().GetSysUserById(ctx, merchantApp.SysUserId)
	if err != nil {
		return nil, err
	}

	info.Order.TradeSourceType = pay_enum.Order.TradeSourceType.Weixin.Code() // 交易源类型
	info.Order.UnionMainId = merchantApp.UnionMainId
	info.Order.UnionMainType = sysUser.Type

	// 支付前创建交易订单，支付后修改交易订单元数据
	orderInfo, err := pay_service.Order().CreateOrder(ctx, &info.Order) // CreatedOrder不能修改订单id
	if err != nil || orderInfo == nil {
		return nil, err
	}

	var prepayId string
	// 判断是小程序还是H5
	if merchantApp.AppType == 1 {
		//  公众号

	} else if merchantApp.AppType == 2 {
		// 小程序  JsApi支付产品
		prepayId, err = s.JsapiCreateOrder(ctx, &weixin_model.TradeOrder{
			ReturnUrl: info.ReturnUrl, // 支付成功后的返回地址
			Order:     info.Order,
		}, openId)

	} else if merchantApp.AppType == 4 {
		// APP

	}

	// 支付订单创建成功后，需要拼接好支付参数，然后返回给前端
	return s.makePayParams(ctx, gconv.String(orderInfo.Id), appId, prepayId)
}

// 生成支付所需参数
func (s *sWeiXinPay) makePayParams(ctx context.Context, orderId, appId, prepay_id string) (*weixin_model.PayParamsRes, error) {
	// 通过AppId拿到特约商户商户号
	subMerchant, err := weixin_service.PaySubMerchant().GetPaySubMerchantByAppId(ctx, appId)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的特约商户", "WeiXin-Pay")
	}
	// 通过AppId拿到特约商户商户号
	spMerchant, err := weixin_service.PayMerchant().GetPayMerchantByMchid(ctx, subMerchant.SpMchid)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的商户号", "WeiXin-Pay")
	}

	ret := &weixin_model.PayParamsRes{
		AppId:     appId,
		TimeStamp: strconv.FormatInt(time.Now().Unix(), 10),
		NonceStr:  utility.Md5Hash(orderId),
		Package:   "prepay_id=" + prepay_id,
		SignType:  "RSA",
		PaySign:   "",
	}

	// 1.构建签名串
	/*
		应用ID
		时间戳
		随机字符串
		订单详情扩展字符串
	*/
	var content = ret.AppId + "\n" +
		ret.TimeStamp + "\n" +
		ret.NonceStr + "\n" +
		ret.Package + "\n"

	// 2.计算签名值paySign = appId、timeStamp、nonceStr、package ==》 通过私钥进行SHA256 with RSA签名 ==》 对签名结果进行Base64编码得到签名值
	privateKey, err := weixin.LoadPrivateKey(spMerchant.PayPrivateKeyPem)

	var sign, _ = weixin.SignSHA256WithRSA(content, privateKey) // qDLKva8l1HPQ0GDjQA9cHMqIg8cI4JWv0/toKBoA+8dSgIKKySQniAv8AKapAj3DHX1Td6xS9Tgm2LPUewdP4KkZ6aYOdbtiDLaoCiuLNud4S0mTsek7Re9oOaA5OCIqsz2E5AYOWJkGxebrIOhWAWChKiT/+JKZXWdBozuYIN0tqtirfK3xuhaPszlx0sJwD0V7Gn2tYK9VVVVYfpFNdXZeQaehdpDVfj5xkVXaH8yQwweoljoy1qWC+UFmZ+/8TIu5w3OslMnbrWIlMOckJdfnv5bXyvkChzETfO4R46eiOdkXi1dP6759S9FZn7JVFglu22aJdTVk3g7e8BmtHA==
	ret.PaySign = sign

	// 3.将支付参数返回至前端

	return ret, err
	/*
		wx6cc2c80416074df3
		1684393307
		f9aa40a057cae16c37a2b97db23a86ed
		prepay_id=wx18150015642076d683b4336866f9370000
	*/
}

// DownloadCertificates 测试SDK ，下载微信支付平台证书
func (s *sWeiXinPay) DownloadCertificates(ctx context.Context) (*certificates.DownloadCertificatesResponse, error) {
	appId := utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId

	// 通过AppId拿到特约商户商户号
	subMerchant, err := weixin_service.PaySubMerchant().GetPaySubMerchantByAppId(ctx, appId)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的特约商户", "WeiXin-Pay")
	}

	spMerchant, err := weixin_service.PayMerchant().GetPayMerchantByMchid(ctx, subMerchant.SpMchid)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的商户号", "WeiXin-Pay")
	}

	var (
		mchID = gconv.String(subMerchant.SpMchid) // 商户号

		//mchID string = "1642565036" // 商户号
		//mchCertificateSerialNumber string = "298D4028EC0F48748DF237A226DB4D5281EFE86E" // 商户证书序列号
		//mchAPIv3Key                string = "655957AD45E5FE85F1BF3B9E0D82B96D"         // 商户APIv3密钥
	)

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

// JsapiCreateOrder JsApi 支付下单
func (s *sWeiXinPay) JsapiCreateOrder(ctx context.Context, info *weixin_model.TradeOrder, openId string) (tradeNo string, err error) {
	sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------JSAPI 创建支付订单，预下单 ------- ", "WeiXin-Pay")

	appId := utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId

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
		Description: core.String("筷满客充电站_"),
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

		SettleInfo: &jsapi.SettleInfo{ // 是否指定分账
			ProfitSharing: core.Bool(false),
			//ProfitSharing: core.Bool(true),
		},

		//TimeExpire:  core.Time(gtime.Now().Add(time.Minute * 5)), // 订单失效失效时间
		//Attach:        core.String(""),// 附加数据
		//GoodsTag:      core.String(), // 订单优惠标记
		//LimitPay:      core.String(), // 指定支付方式
		//SupportFapiao: core.String(), // 开票相关
		//Detail:     &jsapi.Detail{}, // 优惠功能
	}

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
		appId = utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId
	}
	appId = appID[0]
	//
	//appId := utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId

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
		appId = utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId
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
		log.Println(err)
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
	//appId := utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId
	appId := ""
	if appID[0] == "" && !gstr.HasPrefix(appID[0], "wx") {
		appId = utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId
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
			SpMchid:    core.String(gconv.String(subMerchant.SubMchid)),
			SubMchid:   core.String(gconv.String(subMerchant.SubMchid)),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CloseOrder err:%s", err)
		return false, err
	} else {
		// 处理返回结果
		log.Printf("status=%d", result.Response.StatusCode)
	}

	return result.Response.StatusCode == 200, nil
}

// DownloadAccountBill 账单下载接口
func (s *sWeiXinPay) DownloadAccountBill(ctx context.Context, mchId string) {

}

// 得到prepay_id，以及调起支付所需的参数和签名  ----- 直连模式
//resp, result, err := svc.PrepayWithRequestPayment(ctx,
//jsapi.PrepayRequest{
//	Appid:       core.String("wxd678efh567hg6787"),
//	Mchid:       core.String("1900009191"),
//	Description: core.String("Image形象店-深圳腾大-QQ公仔"),
//	OutTradeNo:  core.String("1217752501201407033233368018"),
//	Attach:      core.String("自定义数据说明"),
//	NotifyUrl:   core.String("https://www.weixin.qq.com/wxpay/pay.php"),
//	Amount: &jsapi.Amount{
//		Total: core.Int64(100),
//	},
//	Payer: &jsapi.Payer{
//		Openid: core.String("oUpF8uMuAJO_M2pxb1Q9zNjWeS6o"),
//	},
//},
