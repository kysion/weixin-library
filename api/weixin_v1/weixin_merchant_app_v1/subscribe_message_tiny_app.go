package weixin_merchant_app_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
)

/*
小程序消息管理：
消息分类：
- 订阅消息
- 统一服务消息 【已下架】
- 客服消息
- 位置消息
*/

// API: 用户对消息模板 进行订阅授权 （前端实现）

// API：对用户发送订阅消息

type SendMessageReq struct {
	g.Meta `path:"/:appId/sendMessage" method:"post" summary:"发送订阅消息" tags:"WeiXin消息/小程序订阅消息"`

	weixin_model.SendMessage
}

// API：获取小程序账号的类目

type GetCategoryReq struct {
	g.Meta `path:"/:appId/getCategory" method:"get" summary:"获取小程序账号的类目" tags:"WeiXin消息/小程序订阅消息"`
}

// API：获取个人模板列表

type GetMyTemplateListReq struct {
	g.Meta `path:"/:appId/getMyTemplateList" method:"get" summary:"获取个人模板列表" tags:"WeiXin消息/小程序订阅消息"`
}

// API：删除模板

type DeleteTemplateReq struct {
	g.Meta `path:"/:appId/deleteTemplate" method:"post" summary:"删除模板" tags:"WeiXin消息/小程序订阅消息"`
	
	TemplateId string `json:"templateId" dc:"模板ID" v:"required#模板ID不能为空"`
}

// API：获取模板的关键词列表

type GetPubTemplateKeyWordsReq struct {
	g.Meta `path:"/:appId/getPubTemplateKeyWords" method:"get" summary:"获取模板的关键词列表" tags:"WeiXin消息/小程序订阅消息"`

	TemplateId string `json:"templateId" dc:"模板ID" v:"required#模板ID不能为空"`
}

// API：获取指定类目下的公共模板列表

type GetPubTemplateTitleListReq struct {
	g.Meta `path:"/:appId/getPubTemplateTitleList" method:"get" summary:"获取指定类目下的公共模板列表" tags:"WeiXin消息/小程序订阅消息"`
}
