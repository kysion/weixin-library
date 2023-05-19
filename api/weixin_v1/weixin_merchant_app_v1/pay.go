package weixin_merchant_app_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
)

type DownloadCertificatesReq struct {
	g.Meta `path:"/:appId/downloadCertificates" method:"get" summary:"获取平台证书" tags:"WeiXin支付"`
}

type PayTradeCreateReq struct {
	g.Meta `path:"/:appId/payTradeCreate" method:"post" summary:"支付下单" tags:"WeiXin支付"`
	OpenId string `json:"open_id" dc:"微信用户OpenId"`
	weixin_model.TradeOrder
}
