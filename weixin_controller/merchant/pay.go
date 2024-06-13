package merchant

import (
	"context"
	"github.com/kysion/weixin-library/api/weixin_v1/weixin_merchant_app_v1"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
	"github.com/kysion/weixin-library/weixin_utility"
	// 	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/jsapi"  服务商微信支付
	// 	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"  商户微信支付
)

var WeiXinPay = cWeiXinPay{}

type cWeiXinPay struct{}

// DownloadCertificates 获取微信平台证书
func (c *cWeiXinPay) DownloadCertificates(ctx context.Context, _ *weixin_merchant_app_v1.DownloadCertificatesReq) (*weixin_model.CertificatesDownloadCertificatesRes, error) {
	appId := weixin_utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId

	ret, err := weixin_service.WeiXinPay().DownloadCertificates(ctx, appId)

	return (*weixin_model.CertificatesDownloadCertificatesRes)(ret), err
}

// PayTradeCreate 支付下单
func (c *cWeiXinPay) PayTradeCreate(ctx context.Context, req *weixin_merchant_app_v1.PayTradeCreateReq) (*weixin_model.PayParamsRes, error) {
	ret, err := weixin_service.WeiXinPay().PayTradeCreate(ctx, &req.TradeOrder, req.OpenId)

	return ret, err
}
