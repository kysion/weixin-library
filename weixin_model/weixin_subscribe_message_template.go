package weixin_model

import (
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/weixin-library/weixin_model/weixin_entity"
)

type WeixinSubscribeMessageTemplate struct {
	//Id                       int64       `json:"id"                       description:"ID"`
	TemplateId               string `json:"templateId"               description:"模板ID"`
	Type                     int    `json:"type"                     description:"模板类型：2一次性订阅、3长期订阅"`
	Title                    string `json:"title"                    description:"模板标题"`
	KeyWords                 string `json:"keyWords"                 description:"模板主题词/关键词"`
	ServerCategory           string `json:"serverCategory"           description:"模板服务类目"`
	ServerCategoryId         int    `json:"serverCategoryId"         description:"模板服务类目ID"`
	Content                  string `json:"content"                  description:"模板内容"`
	ContentExample           string `json:"contentExample"           description:"模板内容示例"`
	ContentDataJson          string `json:"contentDataJson"          description:"模板内容Json"`
	KeyWordEnumValueListJson string `json:"keyWordEnumValueListJson" description:"模板枚举参数值范围列表"`
	SceneDesc                string `json:"sceneDesc"                description:"场景描述"`
	SceneType                int    `json:"sceneType"                description:"场景类型【业务层自定义】：1活动即将开始提醒、2活动开始提醒、3活动即将结束提醒、4活动结束提醒、5活动获奖提醒、6券即将生效提醒、7券的生效提醒、8券的失效提醒、9券即将失效提醒、10券核销提醒、8192系统通知、"`
	MessageType              int    `json:"messageType"              description:"消息类型【业务层自定义】：1系统消息、2活动消息、4免啦券消息"`
	MerchantAppId            string `json:"merchantAppId"            description:"商家应用APPID"`
	MerchantAppType          int    `json:"merchantAppType"          description:"商家应用类型：1公众号、2小程序、4网站应用H5、8移动应用、16视频小店"`
	ThirdAppId               string `json:"thirdAppId"               description:"第三方平台应用APPID"`
	UserId                   int64  `json:"userId"                   description:"应用所属账号"`
	UnionMainId              int64  `json:"unionMainId"              description:"关联主体id"`
	ExtJson                  string `json:"extJson"                  description:"拓展字段Json"`
	//CreatedAt                *gtime.Time `json:"createdAt"                description:""`
	//UpdatedAt                *gtime.Time `json:"updatedAt"                description:""`
}

type UpdateWeixinSubscribeMessageTemplate struct {
	//Id                       int64       `json:"id"                       description:"ID"`
	//TemplateId               *string `json:"templateId"               description:"模板ID"`
	//Type                     *int    `json:"type"                     description:"模板类型：2一次性订阅、3长期订阅"`
	//Title                    *string `json:"title"                    description:"模板标题"`
	KeyWords                 *string `json:"keyWords"                 description:"模板主题词/关键词"`
	ServerCategory           *string `json:"serverCategory"           description:"模板服务类目"`
	ServerCategoryId         *int    `json:"serverCategoryId"         description:"模板服务类目ID"`
	Content                  *string `json:"content"                  description:"模板内容"`
	ContentExample           *string `json:"contentExample"           description:"模板内容示例"`
	ContentDataJson          *string `json:"contentDataJson"          description:"模板内容Json"`
	KeyWordEnumValueListJson *string `json:"keyWordEnumValueListJson" description:"模板枚举参数值范围列表"`
	//SceneDesc                *string `json:"sceneDesc"                description:"场景描述"`
	//SceneType                *int    `json:"sceneType"                description:"场景类型【业务层自定义】：1活动即将开始提醒、2活动开始提醒、3活动即将结束提醒、4活动结束提醒、5活动获奖提醒、6券即将生效提醒、7券的生效提醒、8券的失效提醒、9券即将失效提醒、10券核销提醒、8192系统通知、"`
	//MessageType              *int    `json:"messageType"              description:"消息类型【业务层自定义】：1系统消息、2活动消息、4免啦券消息"`
	//UserId                   *int64  `json:"userId"                   description:"应用所属账号"`
	//UnionMainId              *int64  `json:"unionMainId"              description:"关联主体id"`
	ExtJson *string `json:"extJson"                  description:"拓展字段Json"`
	//CreatedAt                *gtime.Time `json:"createdAt"                description:""`
	//UpdatedAt                *gtime.Time `json:"updatedAt"                description:""`
}

type WeixinSubscribeMessageTemplateRes weixin_entity.WeixinSubscribeMessageTemplate

type WeixinSubscribeMessageTemplateListRes base_model.CollectRes[WeixinSubscribeMessageTemplateRes]
