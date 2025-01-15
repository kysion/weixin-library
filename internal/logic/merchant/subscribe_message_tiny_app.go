package merchant

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
	"regexp"
	"strings"
)

/*
小程序订阅消息：
	- 微信接口文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-message-management/subscribe-message/getMessageTemplateList.html
*/

// 小程序订阅消息
type sSubscribeMessage struct{}

func init() {
	weixin_service.RegisterSubscribeMessage(NewSubscribeMessage())
}

func NewSubscribeMessage() weixin_service.ISubscribeMessage {
	return &sSubscribeMessage{}
}

/*
Bug：应用授权了权限。但是第三方token调用接口的时候，报错：{"errcode":48001,"errmsg":"api unauthorized rid: 6687c0aa-3fcb5c05-794bfa90"}
*/

// SendMessage 发送订阅消息
func (s *sSubscribeMessage) SendMessage(ctx context.Context, appId string, info *weixin_model.SendMessage) (*weixin_model.SendMessageRes, error) {
	//	POST https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=ACCESS_TOKEN
	token := ""
	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}
	token = merchantApp.AppAuthToken

	//if merchantApp.ThirdAppId != "" {
	//	thirdConfig, _ := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, merchantApp.ThirdAppId)
	//	if thirdConfig.AppAuthToken != "" {
	//		token = thirdConfig.AppAuthToken
	//	}
	//}

	url := "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=" + token

	g.Dump(info)
	reqData, _ := gjson.Encode(info)
	g.Dump(reqData)

	result := g.Client().PostContent(ctx, url, reqData)

	res := weixin_model.SendMessageRes{}
	_ = gjson.DecodeTo(result, &res)
	//Bug:{"errcode":48001,"errmsg":"api unauthorized rid: 66852ea5-4a805617-41e722d4"} 使用第三方的token，会报此错误。
	//Bug: {"errcode":43101,"errmsg":"user refuse to accept the msg rid: 66852f4d-05bdc04c-3984eb83"}  用户未授权订阅。或一次性订阅已使用完。
	//Bug：{"errcode":47003,"errmsg":"argument invalid! data.thing4.value invalid rid: 6685312c-738a8045-626e019c"}
	return &res, err
}

// GetCategory 获取小程序账号的类目
func (s *sSubscribeMessage) GetCategory(ctx context.Context, appId string) (*weixin_model.GetCategoryRes, error) {
	//GET https://api.weixin.qq.com/wxaapi/newtmpl/getcategory?access_token=ACCESS_TOKEN
	token := ""
	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}
	token = merchantApp.AppAuthToken

	url := "https://api.weixin.qq.com/wxaapi/newtmpl/getcategory?access_token=" + token

	result := g.Client().GetContent(ctx, url)

	res := weixin_model.GetCategoryRes{}
	_ = gjson.DecodeTo(result, &res)

	return &res, err
}

// GetMyTemplateList 获取个人模板列表
func (s *sSubscribeMessage) GetMyTemplateList(ctx context.Context, appId string) (*weixin_model.GetMyTemplateListRes, error) {
	//GET https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate?access_token=ACCESS_TOKEN
	token := ""
	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}
	token = merchantApp.AppAuthToken

	//if merchantApp.ThirdAppId != "" {
	//	thirdConfig, _ := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, merchantApp.ThirdAppId)
	//	if thirdConfig.AppAuthToken != "" {
	//		token = thirdConfig.AppAuthToken
	//	}
	//}

	url := "https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate?access_token=" + token

	// BUG:{"errcode":48001,"errmsg":"api unauthorized rid: 6687c0aa-3fcb5c05-794bfa90"}  使用第三方的token，会报此错误。但是明明将应用授权给第三方应用了啊
	result := g.Client().GetContent(ctx, url)
	fmt.Println("获取模板列表：")
	fmt.Println(result)

	res := weixin_model.GetMyTemplateListRes{}
	_ = gjson.DecodeTo(result, &res)

	// 持久化到SQL数据库
	if res.Errcode == 0 && res.Errmsg == "ok" {
		newCtx := context.Background()
		go func(ctx context.Context) {
			s.saveMyTemplateList(ctx, appId, &res)
		}(newCtx)
	}

	return &res, err

	/*
			     模板列表：
					{
		  "data": [
		    {
		      "priTmplId": "SnMMnZ5m-5jMkcZ82v09hJN76CDhVarCTglONTFGZxA",
		      "title": "活动即将开始提醒",
		      "content": "活动名称:{{thing1.DATA}}\n活动简介:{{thing12.DATA}}\n温馨提示:{{thing5.DATA}}\n",
		      "example": "活动名称:66天街主题活动\n活动简介:本活动主要用于……\n温馨提示:请您提前到达活动地点参与活动\n",
		      "type": 2,
		      "keywordEnumValueList": []
		    },
		    {
		      "priTmplId": "YvYH8JcXzFfegoVV8uOQx3KaaykQDuyt24j3QcZwmEU",
		      "title": "优惠券使用通知",
		      "content": "优惠券类型:{{phrase1.DATA}}\n使用时间:{{date2.DATA}}\n备注:{{thing3.DATA}}\n",
		      "example": "优惠券类型:优惠券\n使用时间:2019-11-03 11:12:13\n备注:优惠券已使用，请及时发放奖励\n",
		      "type": 2,
		      "keywordEnumValueList": []
		    },
		    {
		      "priTmplId": "amFgAO1n3QFZfv63uDgnP9iPDTGkIDs5ltIGscsiiV8",
		      "title": "优惠券使用提醒",
		      "content": "优惠券名称:{{thing1.DATA}}\n温馨提示:{{thing4.DATA}}\n过期时间:{{time3.DATA}}\n",
		      "example": "优惠券名称:优惠券名称\n温馨提示:已领取外卖大牌优惠券，请尽快使用\n过期时间:2021年11月30日\n",
		      "type": 2,
		      "keywordEnumValueList": []
		    },
		    {
		      "priTmplId": "EM9Qxptod2MVNQ24gDZanPRndrAjz_KgA-l_yZ2AdVc",
		      "title": "核销成功通知",
		      "content": "优惠券名称:{{thing6.DATA}}\n券编号:{{character_string12.DATA}}\n商户名称:{{thing13.DATA}}\n核销时间:{{time10.DATA}}\n",
		      "example": "优惠券名称:满100减10\n券编号:0938099813098134\n商户名称:某某门店\n核销时间:2020年2月12日  11:00\n",
		      "type": 2,
		      "keywordEnumValueList": []
		    },
		    {
		      "priTmplId": "4L5yV0qP1x0tppLaPIXjH6SC5a7esgKI3MuVHIMi7YA",
		      "title": "优惠券过期提醒",
		      "content": "优惠券名称:{{thing1.DATA}}\n优惠券类型:{{thing2.DATA}}\n到期日:{{thing3.DATA}}\n温馨提示:{{thing4.DATA}}\n",
		      "example": "优惠券名称:10元现金抵扣券\n优惠券类型:现金抵扣券\n到期日:3天后到期\n温馨提示:您领取的优惠券即将到期，请尽快使用！\n",
		      "type": 2,
		      "keywordEnumValueList": []
		    },
		    {
		      "priTmplId": "SqbQjeS-hI25N782IjYq3UnY4rSWdnsDIZCR7XsKSdc",
		      "title": "活动即将结束提醒",
		      "content": "活动名称:{{thing1.DATA}}\n结束时间:{{time2.DATA}}\n温馨提示:{{thing3.DATA}}\n选手排名:{{thing4.DATA}}\n",
		      "example": "活动名称:最美设计大赛\n结束时间:2020-09-01 12:12:12\n温馨提示:活动最后一天啦，快来参与吧！\n选手排名:第1名\n",
		      "type": 2,
		      "keywordEnumValueList": []
		    },
		    {
		      "priTmplId": "O_UBqu_UqfOa32tyqXu1RMDE9BW2WgWtmIK0HIS505Y",
		      "title": "活动开始通知",
		      "content": "活动名称:{{thing1.DATA}}\n活动时间:{{time3.DATA}}\n活动地址:{{thing4.DATA}}\n备注:{{thing6.DATA}}\n",
		      "example": "活动名称:花卉活动\n活动时间:2021年7月12日\n活动地址:xx市xx区\n备注:您参加的活动还有30分钟就要开始了\n",
		      "type": 2,
		      "keywordEnumValueList": []
		    },
		    {
		      "priTmplId": "fe-BqUsXRXQpZx0HgIEe3-5ua_MoOASuzrx-_63-Sw4",
		      "title": "活动状态通知",
		      "content": "活动名称:{{thing1.DATA}}\n温馨提醒:{{thing3.DATA}}\n活动说明:{{thing4.DATA}}\n活动进度:{{thing2.DATA}}\n",
		      "example": "活动名称:小程序试用7天\n温馨提醒:恭喜您，已抽到奖品\n活动说明:活动店铺地址在朝阳街德庄火锅举行活动\n活动进度:还差1人即可开奖\n",
		      "type": 2,
		      "keywordEnumValueList": []
		    },
		    {
		      "priTmplId": "Sd2HV2B9Dq5nIo2lT0LhKEERy35y0cElemoBYYfqH8Y",
		      "title": "活动提醒",
		      "content": "活动描述:{{thing1.DATA}}\n活动日期:{{time2.DATA}}\n活动说明:{{thing3.DATA}}\n",
		      "example": "活动描述:新年开工大吉，好礼翻不停\n活动日期:2020年2月1日\n活动说明:关注小程序，更多精彩内容等着你~\n",
		      "type": 2,
		      "keywordEnumValueList": []
		    },
		    {
		      "priTmplId": "pPoqYm68T5ZGGkqi7QpFVheS9cOGFfFB2bvSVzp0n5c",
		      "title": "活动奖励通知",
		      "content": "奖励内容:{{thing4.DATA}}\n备注:{{thing5.DATA}}\n活动名称:{{thing6.DATA}}\n使用说明:{{thing7.DATA}}\n",
		      "example": "奖励内容:5积分\n备注:奖励可在小程序“行贾有术”本店详情查看\n活动名称:签到满x次领取大奖\n使用说明:小程序商城交易满10元可使用\n",
		      "type": 2,
		      "keywordEnumValueList": []
		    },
		    {
		      "priTmplId": "LMeD6AptC7fu1y2rmNjuk1lLA9zGxKOLMtQaGtVowZM",
		      "title": "优惠券到账通知",
		      "content": "优惠券名称:{{thing1.DATA}}\n优惠券类型:{{short_thing2.DATA}}\n温馨提示:{{thing6.DATA}}\n过期时间:{{time5.DATA}}\n起始日期:{{time4.DATA}}\n",
		      "example": "优惠券名称:送你一张3折优惠券\n优惠券类型:折扣券\n温馨提示:优惠券已入账 快去使用吧\n过期时间:2022年1月1日 12:00\n起始日期:2021-09-01\n",
		      "type": 2,
		      "keywordEnumValueList": []
		    },
		    {
		      "priTmplId": "bJfhU_sSVzr6go1oLcX_ueCf8MKoKl15pms9LmqyYqw",
		      "title": "新任务发布通知",
		      "content": "任务名称:{{thing1.DATA}}\n奖励积分:{{number3.DATA}}\n温馨提示:{{thing4.DATA}}\n任务编号:{{number5.DATA}}\n产品名称:{{thing18.DATA}}\n",
		      "example": "任务名称:需要专业法律服务\n奖励积分:200\n温馨提示:3天内再次揭盖，限时加赠100积分赢更多好礼\n任务编号:202012312\n产品名称:某某产品\n",
		      "type": 2,
		      "keywordEnumValueList": []
		    }
		  ],
		  "errmsg": "ok",
		  "errcode": 0
		}
	*/
}

func (s *sSubscribeMessage) saveMyTemplateList(ctx context.Context, appId string, res *weixin_model.GetMyTemplateListRes) (err error) {
	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return err
	}

	// 插入模板列表到数据库
	for _, item := range res.Data {
		//item.Example = string(g.Md5(item.Example))

		// 是否存在模板
		template, _ := weixin_service.SubscribeMessageTemplate().GetSubscribeMessageTemplateByTemplateId(ctx, item.PriTmplId)
		if template == nil { // 创建

			// 1、找到模版的关键词
			keyWords := make([]string, 0)
			words, _ := s.GetPubTemplateKeyWords(ctx, appId, item.PriTmplId)
			for _, datum := range words.Data {
				keyWords = append(keyWords, datum.Name)
			}

			// 2、找到模版的服务分类 TODO 没有相关API。。。

			// 3、处理模板内容Json，去掉多余文字，需要符合消息发送时候的格式
			contentDataJson := s.getContentJson(item.Content)

			keywordEnumValueListStr, _ := gjson.EncodeString(item.KeywordEnumValueList)
			itemJson, _ := gjson.EncodeString(item)

			tempInfo := &weixin_model.WeixinSubscribeMessageTemplate{
				TemplateId:               item.PriTmplId,
				Type:                     item.Type,
				Title:                    item.Title,
				KeyWords:                 strings.Join(keyWords, ","),
				ServerCategory:           "", // 服务分类
				ServerCategoryId:         0,
				SceneDesc:                "", // 场景描述: 后设置
				SceneType:                0,  // 场景类型: 后设置
				MessageType:              0,  // 消息类型: 后设置
				Content:                  item.Content,
				ContentExample:           item.Example,
				ContentDataJson:          contentDataJson,
				KeyWordEnumValueListJson: keywordEnumValueListStr,
				MerchantAppId:            appId,
				MerchantAppType:          merchantApp.AppType,
				ThirdAppId:               merchantApp.ThirdAppId,
				UserId:                   merchantApp.UserId,
				UnionMainId:              merchantApp.UnionMainId,
				ExtJson:                  itemJson,
			}
			_, err = weixin_service.SubscribeMessageTemplate().CreateSubscribeMessageTemplate(ctx, tempInfo)
			if err != nil {
				g.Log().Error(ctx, err)
			}

		} else { // 更新
			// 1、找到模版的关键词

			// 2、找到模版的服务分类 TODO 没有相关API。。。

			// 3、处理模板内容Json，去掉多余文字，需要符合消息发送时候的格式
			contentDataJson := s.getContentJson(item.Content)

			keywordEnumValueListStr, _ := gjson.EncodeString(item.KeywordEnumValueList)
			itemJson, _ := gjson.EncodeString(item)

			tempInfo := &weixin_model.UpdateWeixinSubscribeMessageTemplate{
				//KeyWords:                 nil,
				//ServerCategory:           nil,
				//ServerCategoryId:         nil,
				Content:                  &item.Content,
				ContentExample:           &item.Example,
				ContentDataJson:          &contentDataJson,
				KeyWordEnumValueListJson: &keywordEnumValueListStr,
				//SceneDesc:                nil,
				//SceneType:                nil,
				//MessageType:              nil,
				//UserId:      nil,
				//UnionMainId: nil,
				ExtJson: &itemJson,
			}

			_, err = weixin_service.SubscribeMessageTemplate().UpdateSubscribeMessageTemplate(ctx, template.Id, tempInfo)
			if err != nil {
				g.Log().Error(ctx, err)
			}
		}

	}

	return err
}

func (s *sSubscribeMessage) getContentJson(templateStr string) string {
	//	templateStr := `优惠券名称:{{thing1.DATA}}\n温馨提示:{{thing4.DATA}}\n过期时间:{{time3.DATA}}`

	// 使用正则表达式匹配模板中的字段，如 {{thing1.DATA}}
	re := regexp.MustCompile(`{{(.*?)\.DATA}}`)
	matches := re.FindAllStringSubmatch(templateStr, -1)

	// 提取匹配到dataJson
	dataValue := map[string]weixin_model.DataValue{}
	for _, match := range matches {
		dataValue[match[1]] = weixin_model.DataValue{
			Value: "",
		}
	}
	contentDataJson, _ := gjson.EncodeString(dataValue)
	return contentDataJson
}

// DeleteTemplate 删除模板
func (s *sSubscribeMessage) DeleteTemplate(ctx context.Context, appId string, info *weixin_model.DeleteTemplate) (*weixin_model.DeleteTemplateRes, error) {
	// POST https://api.weixin.qq.com/wxaapi/newtmpl/deltemplate?access_token=ACCESS_TOKEN

	token := ""
	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}
	token = merchantApp.AppAuthToken

	//if merchantApp.ThirdAppId != "" {
	//	thirdConfig, _ := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, merchantApp.ThirdAppId)
	//	if thirdConfig.AppAuthToken != "" {
	//		token = thirdConfig.AppAuthToken
	//	}
	//}

	url := "https://api.weixin.qq.com/wxaapi/newtmpl/deltemplate?access_token=" + token

	reqData, _ := gjson.Encode(info)

	result := g.Client().PostContent(ctx, url, reqData)

	res := weixin_model.DeleteTemplateRes{}
	_ = gjson.DecodeTo(result, &res)

	// 持久化到SQL数据库 (数据同步)
	if res.Errcode == 0 && res.Errmsg == "ok" {
		newCtx := context.Background()
		go func(ctx context.Context) {
			_, err = weixin_service.SubscribeMessageTemplate().DeleteSubscribeMessageTemplate(ctx, appId, info.PriTmplId)
		}(newCtx)
	}

	return &res, err
}

// GetPubTemplateKeyWords 获取模板的关键词列表
func (s *sSubscribeMessage) GetPubTemplateKeyWords(ctx context.Context, appId string, templateId string) (*weixin_model.GetPubTemplateKeyWordsRes, error) {
	// GET https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatekeywords?access_token=ACCESS_TOKEN

	token := ""
	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}
	token = merchantApp.AppAuthToken

	// https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatekeywords?access_token=ACCESS_TOKEN&tid=99
	url := "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatekeywords?access_token=" + token + "&tid=" + templateId

	result := g.Client().GetContent(ctx, url)

	res := weixin_model.GetPubTemplateKeyWordsRes{}
	_ = gjson.DecodeTo(result, &res)

	return &res, err
}

// GetPubTemplateTitleList 获取指定类目下的公共模板列表
func (s *sSubscribeMessage) GetPubTemplateTitleList(ctx context.Context, appId string) (*weixin_model.GetPubTemplateTitleListRes, error) {
	// GET https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatetitles?access_token=ACCESS_TOKEN
	token := ""
	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}
	token = merchantApp.AppAuthToken

	url := "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatetitles?access_token=" + token

	result := g.Client().GetContent(ctx, url)

	res := weixin_model.GetPubTemplateTitleListRes{}
	_ = gjson.DecodeTo(result, &res)

	return &res, err
}
