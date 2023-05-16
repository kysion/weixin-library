package merchant

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/kysion/weixin-library/api/weixin_v1/weixin_merchant_app_v1"
	"github.com/kysion/weixin-library/weixin_service"
	// 	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/jsapi"  服务商微信支付
	// 	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"  商户微信支付
)

var WeiXinPay = cWeiXinPay{}

type cWeiXinPay struct{}

func (c *cWeiXinPay) TestSDK(ctx context.Context, req *weixin_merchant_app_v1.TestSDKReq) (api_v1.BoolRes, error) {
	weixin_service.WeiXinPay().TestSDK(ctx)

	return true, nil
}

//// JsapiCreateOrder JsApi 支付下单
//func (c *cWeiXinPay) JsapiCreateOrder(ctx context.Context, req *weixin_merchant_app_v1.JsapiCreateOrderReq) (tradeNo api_v1.StringRes, err error) {
//
//}
