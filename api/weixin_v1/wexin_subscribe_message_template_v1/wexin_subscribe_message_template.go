package wexin_subscribe_message_template_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

type QueryMessageTemplateReq struct {
	g.Meta `path:"/queryMessageTemplate" method:"post" summary:"获取订阅消息模板｜列表" tags:"WeiXin订阅消息模板"`

	AppId string `json:"appId" dc:"应用ID" v:"required#应用ID不能为空"`

	base_model.SearchParams
	IsExport bool `json:"isExport" dc:"是否导出" `
}

type GetMessageTemplateByIdReq struct {
	g.Meta `path:"/getMessageTemplateById" method:"post" summary:"获取订阅消息模板｜信息" tags:"WeiXin订阅消息模板"`

	Id int64 `json:"id" dc:"模板记录ID" v:"required#模板记录ID不能为空"`
}

//
//type GetMessageTemplateByTemplateIdReq struct {
//	g.Meta `path:"/:appId/getMessageTemplateByTemplateId" method:"post" summary:"获取订阅消息模板｜信息" tags:"WeiXin订阅消息模板管理"`
//}
