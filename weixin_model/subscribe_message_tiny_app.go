package weixin_model

// MessageCommonRes 消息公共返回值
type MessageCommonRes struct {
	Errcode int    `json:"errcode" dc:"错误码"`
	Errmsg  string `json:"errmsg" dc:"错误信息，无错误是ok"`
}

// SendMessage 发送订阅消息。
type SendMessage struct {
	Touser           string      `json:"touser" dc:"接收者（用户）的 openid"`
	TemplateId       string      `json:"template_id" dc:"所需下发的订阅模板id"`
	Page             string      `json:"page" dc:"点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转"`
	MiniprogramState string      `json:"miniprogram_state" dc:"跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版"`
	Lang             string      `json:"lang" dc:"进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN"`
	Data             interface{} `json:"data" dc:"模板内容"`
}

// SendMessageRes 发送订阅消息Res
type SendMessageRes struct {
	MessageCommonRes
}

// DeleteTemplate 删除模板。
type DeleteTemplate struct {
	PriTmplId string `json:"priTmplId" dc:"模板ID"`
}

// DeleteTemplateRes 删除模板Res
type DeleteTemplateRes struct {
	MessageCommonRes
}

// GetCategoryRes 获取小程序类目Res
type GetCategoryRes struct {
	MessageCommonRes
	Data []struct {
		Id   int    `json:"id" dc:"类目ID"`
		Name string `json:"name"  dc:"类目Name"`
	} `json:"data" dc:"类目列表"`
}

// GetMyTemplateListRes 获取我的模版列表Res
type GetMyTemplateListRes struct {
	MessageCommonRes
	Data []struct {
		PriTmplId            string `json:"priTmplId" dc:"添加至账号下的模板 id，发送小程序订阅消息时所需"`
		Title                string `json:"title" dc:"模版标题"`
		Content              string `json:"content" dc:"模版内容"`
		Example              string `json:"example" dc:"模板内容示例"`
		Type                 int    `json:"type" dc:"模版类型，2 为一次性订阅，3 为长期订阅"`
		KeywordEnumValueList []struct {
			EnumValueList []string `json:"enumValueList" dc:"枚举参数的 key"`
			KeywordCode   string   `json:"keywordCode" dc:"枚举参数值范围列表"`
		} `json:"keywordEnumValueList,omitempty" dc:"枚举参数值范围"`
	} `json:"data" dc:"模板列表"`
}

// GetPubTemplateKeyWordsRes 获取模板关键词列表Res
type GetPubTemplateKeyWordsRes struct {
	MessageCommonRes

	Data []struct {
		Kid     int    `json:"kid" dc:"关键词 id，选用模板时需要"`
		Name    string `json:"name" dc:"关键词内容"`
		Example string `json:"example" dc:"关键词内容对应的示例"`
		Rule    string `json:"rule" dc:"参数类型"`
	} `json:"data" dc:"模板关键词列表"`
}

// GetPubTemplateTitleListRes 获取指定类目下的公共模板列表
type GetPubTemplateTitleListRes struct {
	MessageCommonRes

	Count int `json:"count" dc:"模版标题列表总数"`
	Data  []struct {
		Tid        int    `json:"tid" dc:"模版标题 id"`
		Title      string `json:"title" dc:"模版标题"`
		Type       int    `json:"type" dc:"	模版类型，2 为一次性订阅，3 为长期订阅"`
		CategoryId string `json:"categoryId" dc:"模版所属类目 id"`
	} `json:"data" dc:"模板标题列表"`
}
