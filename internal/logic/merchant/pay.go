package merchant

import (
	"context"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/certificates"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
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

// TestSDK 测试SDK ，下载微信支付平台证书
func (s *sWeiXinPay) TestSDK(ctx context.Context) {
	var (
		mchID                      string = "1642565036"                               // 商户号
		mchCertificateSerialNumber string = "298D4028EC0F48748DF237A226DB4D5281EFE86E" // 商户证书序列号
		mchAPIv3Key                string = "655957AD45E5FE85F1BF3B9E0D82B96D"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/data/kysion-files/weixin/weixin-pay/1642565036_cert/apiclient_key.pem")
	if err != nil {
		log.Fatal("load merchant private key error")
	}

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
	}

	// 发送请求，以下载微信支付平台证书为例
	// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
	svc := certificates.CertificatesApiService{Client: client}
	resp, result, err := svc.DownloadCertificates(ctx)
	log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
}

// JsapiCreateOrder JsApi 支付下单
func (s *sWeiXinPay) JsapiCreateOrder(ctx context.Context, info *weixin_model.TradeOrder) (tradeNo string, err error) {
	//appId := utility.GetAppIdFormContext(ctx) // 特约商户绑定的App的Id
	//
	//mchId := info.MchId // 特约商家商户号
	//
	//// TODO 实现补充初始化客户端的通用方法 client
	//client, _ := weixin.NewPayClient(ctx)
	//
	//// 微信支付服务商商户号  --> 特约商户XXX ...
	//
	//svc := jsapi.JsapiApiService{Client: client}
	//// 得到prepay_id，以及调起支付所需的参数和签名
	////resp, result, err := svc.PrepayWithRequestPayment(ctx,
	//
	//// TODO 这里是预下单
	//resp, result, err := svc.Prepay(ctx,
	//	jsapi.PrepayRequest{
	//		Appid:       core.String("wxd678efh567hg6787"),
	//		Mchid:       core.String("1900009191"),
	//		Description: core.String("Image形象店-深圳腾大-QQ公仔"),
	//		OutTradeNo:  core.String("1217752501201407033233368018"),
	//		Attach:      core.String("自定义数据说明"),
	//		NotifyUrl:   core.String("https://www.weixin.qq.com/wxpay/pay.php"),
	//		Amount: &jsapi.Amount{
	//			Total: core.Int64(100),
	//		},
	//		Payer: &jsapi.Payer{
	//			Openid: core.String("oUpF8uMuAJO_M2pxb1Q9zNjWeS6o"),
	//		},
	//	},
	//)
	//
	//// TODO 补充下单 （Native、JSAPI、APP等不同场景生成交易串调起支付。）
	//
	//if err == nil {
	//	log.Println(resp)
	//} else {
	//	log.Println(err)
	//}

	// TODO 返回交易编号
	return "", nil
}

// QueryOrderByIdMchID 查询订单 （1.根据tradeNo 2.根据mchId）
func (s *sWeiXinPay) QueryOrderByIdMchID(ctx context.Context, mchId string) {

}

// QueryOrderByIdOutTradeNo 根据支付编号查询订单
func (s *sWeiXinPay) QueryOrderByIdOutTradeNo(ctx context.Context, outTradeNo string) {

}

// 接收支付结果通知接口 NotifyUrl

// CloseOrder 关闭订单接口
func (s *sWeiXinPay) CloseOrder(ctx context.Context, mchId string) {

}

// DownloadAccountBill 账单下载接口
func (s *sWeiXinPay) DownloadAccountBill(ctx context.Context, mchId string) {

}
