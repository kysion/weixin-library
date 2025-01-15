// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package weixin_service

import (
	"context"

	"github.com/kysion/base-library/base_model"
	"github.com/kysion/weixin-library/weixin_model"
)

type (
	IPaySubMerchant interface {
		// GetPaySubMerchantById 根据id查找特约商户配置信息
		GetPaySubMerchantById(ctx context.Context, id int64) (*weixin_model.WeixinPaySubMerchant, error)
		// GetPaySubMerchantByAppId 根据AppId查找特约商户配置信息
		GetPaySubMerchantByAppId(ctx context.Context, appId string) (*weixin_model.WeixinPaySubMerchant, error)
		// GetPaySubMerchantByMchid 根据Mchid查找特约商户配置信息
		GetPaySubMerchantByMchid(ctx context.Context, id int) (*weixin_model.WeixinPaySubMerchant, error)
		// GetPaySubMerchantBySysUserId  根据用户id查询特约商户配置信息
		GetPaySubMerchantBySysUserId(ctx context.Context, sysUserId int64) (*weixin_model.WeixinPaySubMerchant, error)
		// QueryPaySubMerchant 查询列表
		QueryPaySubMerchant(ctx context.Context, params *base_model.SearchParams, isExport bool) (*weixin_model.WeixinPaySubMerchantList, error)
		// CreatePaySubMerchant  创建特约商户配置信息
		CreatePaySubMerchant(ctx context.Context, info *weixin_model.WeixinPaySubMerchant) (*weixin_model.WeixinPaySubMerchant, error)
		// UpdatePaySubMerchant 更新特约商户配置信息
		UpdatePaySubMerchant(ctx context.Context, id int64, info *weixin_model.UpdatePaySubMerchant) (bool, error)
		// SetAuthPath 设置特约商户授权目录
		SetAuthPath(ctx context.Context, info *weixin_model.SetSubMerchantAuthPath) (bool, error)
	}
	ISubscribeMessageTemplate interface {
		// GetSubscribeMessageTemplateByTemplateId 根据模板ID查找消息模板
		GetSubscribeMessageTemplateByTemplateId(ctx context.Context, templateId string) (*weixin_model.WeixinSubscribeMessageTemplateRes, error)
		// GetSubscribeMessageTemplateById 根据id查找消息模板信息
		GetSubscribeMessageTemplateById(ctx context.Context, id int64) (*weixin_model.WeixinSubscribeMessageTemplateRes, error)
		// CreateSubscribeMessageTemplate  创建消息模板信息
		CreateSubscribeMessageTemplate(ctx context.Context, info *weixin_model.WeixinSubscribeMessageTemplate) (*weixin_model.WeixinSubscribeMessageTemplateRes, error)
		// UpdateSubscribeMessageTemplate 更新消息模板信息
		UpdateSubscribeMessageTemplate(ctx context.Context, id int64, info *weixin_model.UpdateWeixinSubscribeMessageTemplate) (bool, error)
		// DeleteSubscribeMessageTemplate 删除模板
		DeleteSubscribeMessageTemplate(ctx context.Context, appId string, templateId string) (bool, error)
		// QuerySubscribeMessageTemplate 查询模板｜列表
		QuerySubscribeMessageTemplate(ctx context.Context, params *base_model.SearchParams, isExport bool) (*weixin_model.WeixinSubscribeMessageTemplateListRes, error)
		// GetSubscribeMessageTemplateByAppAndSceneTypeAndMessageType 查询模板
		GetSubscribeMessageTemplateByAppAndSceneTypeAndMessageType(ctx context.Context, appId string, appType int, sceneType, messageType int) (*weixin_model.WeixinSubscribeMessageTemplateRes, error)
	}
	IThirdAppConfig interface {
		// GetThirdAppConfigByAppId 根据AppId查找第三方应用配置信息
		GetThirdAppConfigByAppId(ctx context.Context, id string) (*weixin_model.WeixinThirdAppConfig, error)
		// GetThirdAppConfigById 根据id查找第三方应用配置信息
		GetThirdAppConfigById(ctx context.Context, id int64) (*weixin_model.WeixinThirdAppConfig, error)
		// CreateThirdAppConfig  创建第三方应用配置信息
		CreateThirdAppConfig(ctx context.Context, info *weixin_model.WeixinThirdAppConfig) (*weixin_model.WeixinThirdAppConfig, error)
		// UpdateThirdAppConfig 更新第三方应用配置信息
		UpdateThirdAppConfig(ctx context.Context, id int64, info *weixin_model.UpdateThirdAppConfig) (bool, error)
		// UpdateReleaseState 修改发布状态
		UpdateReleaseState(ctx context.Context, id int64, releaseState int) (bool, error)
		// UpdateState 修改应用状态
		UpdateState(ctx context.Context, id int64, state int) (bool, error)
		// UpdateAppAuthToken 更新Token  服务商应用授权token
		UpdateAppAuthToken(ctx context.Context, info *weixin_model.UpdateAppAuthToken) (bool, error)
		// UpdateAppConfig 修改服务商基础信息
		UpdateAppConfig(ctx context.Context, info *weixin_model.UpdateThirdAppConfigReq) (bool, error)
		// UpdateAppConfigHttps 修改服务商应用Https配置
		UpdateAppConfigHttps(ctx context.Context, info *weixin_model.UpdateThirdAppConfigHttpsReq) (bool, error)
	}
	IConsumer interface {
		// GetConsumerById 根据id查找消费者信息
		GetConsumerById(ctx context.Context, id int64) (*weixin_model.WeixinConsumerConfig, error)
		// GetConsumerBySysUserId  根据用户id查询消费者信息
		GetConsumerBySysUserId(ctx context.Context, sysUserId int64) (*weixin_model.WeixinConsumerConfig, error)
		// GetConsumerByOpenId  根据用户openId查询消费者信息
		GetConsumerByOpenId(ctx context.Context, openId string, unionId ...string) (*weixin_model.WeixinConsumerConfig, error)
		// GetConsumerByOpenIdAndAppId  根据用户openId和appId查询消费者信息
		GetConsumerByOpenIdAndAppId(ctx context.Context, openId string, appId string, unionId ...string) (*weixin_model.WeixinConsumerConfig, error)
		// QueryConsumerByUnionId  根据用户unionId查询消费者|列表
		QueryConsumerByUnionId(ctx context.Context, unionId string) (*weixin_model.WeixinConsumerConfigListRes, error)
		// CreateConsumer  创建消费者信息
		CreateConsumer(ctx context.Context, info *weixin_model.WeixinConsumerConfig) (*weixin_model.WeixinConsumerConfig, error)
		// UpdateConsumer 更新消费者信息
		UpdateConsumer(ctx context.Context, id int64, info *weixin_model.UpdateConsumerReq) (bool, error)
		// UpdateConsumerByUserId 更新消费者信息
		UpdateConsumerByUserId(ctx context.Context, userId int64, info *weixin_model.UpdateConsumerReq) (bool, error)
		// UpdateConsumerState 修改用户状态
		UpdateConsumerState(ctx context.Context, id int64, state int) (bool, error)
		// UpdateConsumerAuthState 修改用户授权状态
		UpdateConsumerAuthState(ctx context.Context, id int64, state int) (bool, error)
		// SetIsFollowPublic 设置用户是够关注公众号
		SetIsFollowPublic(ctx context.Context, openId string, appID string, isFollowPublic int) (bool, error)
		// UpdateConsumerToken 更新消费者token等数据信息
		UpdateConsumerToken(ctx context.Context, openId string, info *weixin_model.UpdateConsumerTokenReq) (bool, error)
	}
	IMerchantAppConfig interface {
		// GetMerchantAppConfigById 根据id查找商家应用配置信息
		GetMerchantAppConfigById(ctx context.Context, id int64) (*weixin_model.WeixinMerchantAppConfig, error)
		// GetMerchantAppConfigByAppId 根据AppId查找商家应用配置信息
		GetMerchantAppConfigByAppId(ctx context.Context, id string) (*weixin_model.WeixinMerchantAppConfig, error)
		// GetMerchantAppConfigBySysUserId  根据商家id查询商家应用配置信息
		GetMerchantAppConfigBySysUserId(ctx context.Context, sysUserId int64) (*weixin_model.WeixinMerchantAppConfig, error)
		// CreateMerchantAppConfig  创建商家应用配置信息
		CreateMerchantAppConfig(ctx context.Context, info *weixin_model.WeixinMerchantAppConfig) (*weixin_model.WeixinMerchantAppConfig, error)
		// UpdateMerchantAppConfig 更新商家应用配置信息
		UpdateMerchantAppConfig(ctx context.Context, id int64, info *weixin_model.UpdateMerchantAppConfig) (bool, error)
		// UpdateState 修改应用状态
		UpdateState(ctx context.Context, id int64, state int) (bool, error)
		// UpdateAppAuth 修改应用授权信息 (绑定&解绑第三方服务商)
		UpdateAppAuth(ctx context.Context, appId string, thirdAppId, isFullProxy int) (bool, error)
		// UpdateAppAuthToken 更新Token  商家应用授权token
		UpdateAppAuthToken(ctx context.Context, info *weixin_model.UpdateMerchantAppAuthToken) (bool, error)
		// UpdateAppConfig 修改商家基础信息
		UpdateAppConfig(ctx context.Context, info *weixin_model.UpdateMerchantAppConfigReq) (bool, error)
		// UpdateAppConfigHttps 修改商家应用Https配置
		UpdateAppConfigHttps(ctx context.Context, info *weixin_model.UpdateMerchantAppConfigHttpsReq) (bool, error)
		// GetPolicy 获取协议
		GetPolicy(ctx context.Context, appId string) (*weixin_model.GetPolicyRes, error)
	}
	IPayMerchant interface {
		// GetPayMerchantById 根据id查找商户号配置信息
		GetPayMerchantById(ctx context.Context, id int64) (*weixin_model.PayMerchant, error)
		// GetPayMerchantByAppId 根据AppId查找商户号配置信息
		GetPayMerchantByAppId(ctx context.Context, appId string) (*weixin_model.PayMerchant, error)
		// GetPayMerchantByMchid 根据Mchid查找商户号配置信息
		GetPayMerchantByMchid(ctx context.Context, id int) (*weixin_model.PayMerchant, error)
		// GetPayMerchantBySysUserId  根据商家id查询商户号配置信息
		GetPayMerchantBySysUserId(ctx context.Context, sysUserId int64) (*weixin_model.PayMerchant, error)
		// CreatePayMerchant  创建商户号配置信息
		CreatePayMerchant(ctx context.Context, info *weixin_model.PayMerchant) (*weixin_model.PayMerchant, error)
		// UpdatePayMerchant 更新商户号配置信息
		UpdatePayMerchant(ctx context.Context, id int64, info *weixin_model.UpdatePayMerchant) (bool, error)
		// SetCertAndKey  设置商户号证书及密钥文件
		SetCertAndKey(ctx context.Context, id int64, info *weixin_model.SetCertAndKey) (bool, error)
		// SetAuthPath 设置商户号授权目录
		SetAuthPath(ctx context.Context, info *weixin_model.SetAuthPath) (bool, error)
		// SetPayMerchantUnionId 设置商户号关联的AppId
		SetPayMerchantUnionId(ctx context.Context, info *weixin_model.SetPayMerchantUnionId) (bool, error)
		// SetBankcardAccount 设置商户号银行卡号
		SetBankcardAccount(ctx context.Context, info *weixin_model.SetBankcardAccount) (bool, error)
	}
)

var (
	localPayMerchant              IPayMerchant
	localPaySubMerchant           IPaySubMerchant
	localSubscribeMessageTemplate ISubscribeMessageTemplate
	localThirdAppConfig           IThirdAppConfig
	localConsumer                 IConsumer
	localMerchantAppConfig        IMerchantAppConfig
)

func Consumer() IConsumer {
	if localConsumer == nil {
		panic("implement not found for interface IConsumer, forgot register?")
	}
	return localConsumer
}

func RegisterConsumer(i IConsumer) {
	localConsumer = i
}

func MerchantAppConfig() IMerchantAppConfig {
	if localMerchantAppConfig == nil {
		panic("implement not found for interface IMerchantAppConfig, forgot register?")
	}
	return localMerchantAppConfig
}

func RegisterMerchantAppConfig(i IMerchantAppConfig) {
	localMerchantAppConfig = i
}

func PayMerchant() IPayMerchant {
	if localPayMerchant == nil {
		panic("implement not found for interface IPayMerchant, forgot register?")
	}
	return localPayMerchant
}

func RegisterPayMerchant(i IPayMerchant) {
	localPayMerchant = i
}

func PaySubMerchant() IPaySubMerchant {
	if localPaySubMerchant == nil {
		panic("implement not found for interface IPaySubMerchant, forgot register?")
	}
	return localPaySubMerchant
}

func RegisterPaySubMerchant(i IPaySubMerchant) {
	localPaySubMerchant = i
}

func SubscribeMessageTemplate() ISubscribeMessageTemplate {
	if localSubscribeMessageTemplate == nil {
		panic("implement not found for interface ISubscribeMessageTemplate, forgot register?")
	}
	return localSubscribeMessageTemplate
}

func RegisterSubscribeMessageTemplate(i ISubscribeMessageTemplate) {
	localSubscribeMessageTemplate = i
}

func ThirdAppConfig() IThirdAppConfig {
	if localThirdAppConfig == nil {
		panic("implement not found for interface IThirdAppConfig, forgot register?")
	}
	return localThirdAppConfig
}

func RegisterThirdAppConfig(i IThirdAppConfig) {
	localThirdAppConfig = i
}
