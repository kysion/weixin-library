package weixin_merchant_app_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
)

type TestSDKReq struct {
	g.Meta `path:"/testSDK" method:"get" summary:"测试SDK" tags:"WeiXin支付"`
}

type JsapiCreateOrderReq struct {
	g.Meta `path:"/:appId/jsapiCreateOrder" method:"get" summary:"JsApi下单" tags:"WeiXin支付"`
	weixin_model.TradeOrder
}
