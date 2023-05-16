package weixin_model

import "github.com/kysion/pay-share-library/pay_model"

type TradeOrder struct {
	//  商户号
	MchId string `json:"mchId" dc:"商户号"`
	pay_model.Order
}
