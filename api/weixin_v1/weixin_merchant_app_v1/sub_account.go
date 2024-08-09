package weixin_merchant_app_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
)

// 分账相关接口

type GetSubAccountMaxRatioReq struct {
	g.Meta `path:"/:appId/postSubAccountMaxRatio" method:"get" summary:"查询最大分账比例" tags:"WeiXin分账"`
}

type QuerySubAccountOrderReq struct {
	g.Meta `path:"/:appId/querySubAccountOrder" method:"post" summary:"查询分账结果" tags:"WeiXin分账"`
	weixin_model.QueryOrderRequest
}

type UnfreezeOrderReq struct {
	g.Meta `path:"/:appId/unfreezeOrder" method:"post" summary:"解冻剩余资金API" tags:"WeiXin分账"`
	weixin_model.UnfreezeOrderRequest
}

type SubAccountRequestReq struct {
	g.Meta `path:"/:appId/subAccountRequest" method:"post" summary:"请求分账" tags:"WeiXin分账"`
	weixin_model.SubAccountReq
}

type QueryOrderAmountReq struct {
	g.Meta `path:"/:appId/queryOrderAmount" method:"post" summary:"查询剩余待分金额API" tags:"WeiXin分账"`
	weixin_model.QueryOrderAmountRequest
}

type AddReceiverReq struct {
	g.Meta `path:"/:appId/addReceiver" method:"post" summary:"添加分账接收方" tags:"WeiXin分账"`
	weixin_model.AddReceiverRequest
}

type AddReceiverRequestList []weixin_model.AddReceiverRequest
type AddProfitSharingReceiversReq struct {
	g.Meta `path:"/:appId/addProfitSharingReceivers" method:"post" summary:"添加多个分账关系" tags:"WeiXin分账"`
	AddReceiverRequestList
}

type DeleteReceiverReq struct {
	g.Meta `path:"/:appId/deleteReceiver" method:"post" summary:"删除分账接收方" tags:"WeiXin分账"`
	weixin_model.DeleteReceiverRequest
}
