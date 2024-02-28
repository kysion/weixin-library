// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package weixin_service

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/weixin-library/weixin_model"
	hook "github.com/kysion/weixin-library/weixin_model/weixin_hook"
	"github.com/wechatpay-apiv3/wechatpay-go/services/certificates"
	"github.com/wechatpay-apiv3/wechatpay-go/services/profitsharing"
)

type (
	IAppVersion interface {
		// SubmitAppVersionAudit 提交应用版本审核
		SubmitAppVersionAudit(ctx context.Context, appId string, info *weixin_model.SubmitAppVersionAuditReq) (*weixin_model.AppVersionAuditRes, error)
		// CancelAppVersionAudit 撤销应用版本审核
		CancelAppVersionAudit(ctx context.Context, appId string) (*weixin_model.CancelAppVersionAuditRes, error)
		// CancelAppVersion 退回开发版本
		CancelAppVersion(ctx context.Context, appId string, info *weixin_model.CancelAppVersionReq) (*weixin_model.CancelAppVersionRes, error)
		// QueryAppVersionList 查询小程序版本列表,获取已上传的代码页面列表
		QueryAppVersionList(ctx context.Context, appId string) (*weixin_model.QueryAppVersionListRes, error)
		// GetAppVersionDetail 查询小程序版本详情
		GetAppVersionDetail(ctx context.Context, appId string) (*weixin_model.QueryAppVersionDetailRes, error)
		// GetAppLatestVersionAudit 最新一次提审单的审核状态
		GetAppLatestVersionAudit(ctx context.Context, appId string) (*weixin_model.GetAppLatestVersionAuditRes, error)
		// GetAllCategory 获取应用所有类目
		GetAllCategory(ctx context.Context, appId string) (*weixin_model.AppCategoryInfoRes, error)
		// GetAccountVBasicInfo 获取小程序基本信息
		GetAccountVBasicInfo(ctx context.Context, appId string) (*weixin_model.AccountVBasicInfoRes, error)
		// UploadAppMediaToAudit 应用提审素材上传接口
		UploadAppMediaToAudit(ctx context.Context, appId string, mediaPath string) (*weixin_model.UploadAppMediaToAuditRes, error)
		// CommitAppAuditCode 上传代码并生成体验版
		CommitAppAuditCode(ctx context.Context, appId string, info *weixin_model.CommitAppAuditCodeReq) (*weixin_model.CommitAppAuditCodeRes, error)
		// GetQrcode 获取小程序体验版二维码
		GetQrcode(ctx context.Context, appId string) (*weixin_model.ErrorCommonRes, error)
		// ReleaseApp 发布已通过审核的小程序
		ReleaseApp(ctx context.Context, appId string) (*weixin_model.ErrorCommonRes, error)
	}
	IMerchantNotify interface {
		// InstallNotifyHook 订阅异步通知Hook
		InstallNotifyHook(hookKey hook.NotifyKey, hookFunc hook.NotifyHookFunc)
		// InstallTradeHook 订阅支付Hook
		InstallTradeHook(hookKey hook.TradeHookKey, hookFunc hook.TradeHookFunc)
		// NotifyServices 异步通知地址  用于接收支付宝推送给商户的支付/退款成功的消息。
		NotifyServices(ctx context.Context) (string, error)
	}
	IWeiXinPay interface {
		// PayTradeCreate  1、创建交易订单   （AppId的H5是没有的，需要写死，小程序有的 ）
		PayTradeCreate(ctx context.Context, info *weixin_model.TradeOrder, openId string) (*weixin_model.PayParamsRes, error)
		// JsapiCreateOrderByDirect JsApi 支付下单 - 直连模式
		JsapiCreateOrderByDirect(ctx context.Context, info *weixin_model.TradeOrder, openId string) (tradeNo string, err error)
		// DownloadCertificates 测试SDK ，下载微信支付平台证书
		DownloadCertificates(ctx context.Context, appID ...string) (*certificates.DownloadCertificatesResponse, error)
		// JsapiCreateOrder JsApi 支付下单 - 服务商待调用
		JsapiCreateOrder(ctx context.Context, info *weixin_model.TradeOrder, openId string) (tradeNo string, err error)
		// QueryOrderByIdMchID 查询订单 （1.根据tradeNo 2.根据mchId）
		QueryOrderByIdMchID(ctx context.Context, transactionId string, appID ...string) (*weixin_model.TradeOrderRes, error)
		// QueryOrderByIdOutTradeNo 根据支付编号查询订单
		QueryOrderByIdOutTradeNo(ctx context.Context, outTradeNo string, appID ...string) (*weixin_model.TradeOrderRes, error)
		// CloseOrder 关闭订单接口
		CloseOrder(ctx context.Context, outTradeNo string, appID ...string) (bool, error)
		// DownloadAccountBill 账单下载接口
		DownloadAccountBill(ctx context.Context, mchId string)
	}
	ISubAccount interface {
		// GetSubAccountMaxRatio 查询最大分账比例
		GetSubAccountMaxRatio(ctx context.Context, appId string) (*weixin_model.QueryMerchantRatioRes, error)
		// QuerySubAccountOrder 查询分账结果
		QuerySubAccountOrder(ctx context.Context, appId string, info *weixin_model.QueryOrderRequest) (*profitsharing.OrdersEntity, error)
		// UnfreezeOrder 解冻剩余资金API
		UnfreezeOrder(ctx context.Context, appId string, info *weixin_model.UnfreezeOrderRequest) (*profitsharing.OrdersEntity, error)
		// SubAccountRequest 请求分账
		SubAccountRequest(ctx context.Context, appId string, info *weixin_model.SubAccountReq) (*profitsharing.OrdersEntity, error)
		// QueryOrderAmount 查询剩余待分金额API
		QueryOrderAmount(ctx context.Context, appId string, info *weixin_model.QueryOrderAmountRequest) (*profitsharing.QueryOrderAmountResponse, error)
		// AddReceiver 添加分账接收方（相当于绑定分账关系）
		AddReceiver(ctx context.Context, appId string, info *weixin_model.AddReceiverRequest) (*profitsharing.AddReceiverResponse, error)
		// AddProfitSharingReceivers 添加多个分账关系
		AddProfitSharingReceivers(ctx context.Context, appId string, info []weixin_model.AddReceiverRequest) (*profitsharing.AddReceiverResponse, error)
		// DeleteReceiver 删除分账接收方（相当于分账关系解绑）
		DeleteReceiver(ctx context.Context, appId string, info *weixin_model.DeleteReceiverRequest) (*profitsharing.DeleteReceiverResponse, error)
	}
	ISubMerchant interface {
		// GetAuditStateByBusinessCode 根据业务申请编号查询申请状态
		GetAuditStateByBusinessCode(ctx context.Context, spMchId, businessCode string) (*weixin_model.SubMerchantAuditStateRes, error)
		// GetAuditStateByApplymentId 根据申请单号查询申请状态
		GetAuditStateByApplymentId(ctx context.Context, spMchId, applymentId string) (*weixin_model.SubMerchantAuditStateRes, error)
		// GetSettlement 查询结算账号
		GetSettlement(ctx context.Context, subMchId string) (*weixin_model.SettlementRes, error)
		// UpdateSettlement 修改结算账号,成功会返回application_no，作为查询申请状态的唯一标识
		UpdateSettlement(ctx context.Context, subMchId string, info *weixin_model.UpdateSettlementReq) (string, error)
		// GetSettlementAuditState 查询结算账户修改审核状态
		GetSettlementAuditState(ctx context.Context, subMchId, applicationNo string) (*weixin_model.SettlementRes, error)
	}
	IUserAuth interface {
		InstallConsumerHook(infoType hook.ConsumerKey, hookFunc hook.ConsumerHookFunc)
		GetHook() base_hook.BaseHook[hook.ConsumerKey, hook.ConsumerHookFunc]
		// UserAuthCallback 处理网页授权回调请求 （公众号登录）
		UserAuthCallback(ctx context.Context, info g.Map) (int64, error)
		// GetMiniAppUserInfo 获取小程序用户唯一标识，用于检查是否注册,如果已经注册，返会openId
		GetMiniAppUserInfo(ctx context.Context, authCode string, appId string, getDetail bool) (*weixin_model.UserInfoRes, error)
		// UserLogin 获取微信用户openId和sessionKey会话key 进行login  （小程序登录）
		UserLogin(ctx context.Context, info g.Map) (string, error)
		// GetMinoUserAccessToken 获取小程序用户access_token TODO
		GetMinoUserAccessToken(ctx context.Context)
		// GetTinyAppUserInfo 小程序获取用户数据
		GetTinyAppUserInfo(ctx context.Context, sessionKey, encryptedData, iv, appId string, openId string) (*weixin_model.UserInfoRes, error)
	}
	IUserEvent interface {
		// UserEvent 用户相关事件
		UserEvent(ctx context.Context, info g.Map) bool
		// Subscribe 用户关注公众号
		Subscribe(ctx context.Context, appId string, info *weixin_model.MessageBodyDecrypt) (bool, error)
		// UnSubscribe 用户取消关注公众号
		UnSubscribe(ctx context.Context, appId string, info *weixin_model.MessageBodyDecrypt) (bool, error)
		// UserAuthorizationRevoke 用户撤回事件
		UserAuthorizationRevoke(ctx context.Context, appId string, info *weixin_model.MessageBodyDecrypt) (bool, error)
	}
	IAppAuth interface {
		// RefreshToken 刷新Token
		RefreshToken(ctx context.Context, merchantAppId, thirdAppId, refreshToken string) (bool, error)
		// AppAuth 应用授权具体服务
		AppAuth(ctx context.Context, info g.Map) bool
		// Authorized 授权成功
		Authorized(ctx context.Context, info g.Map) bool
		// UpdateAuthorized 授权更新
		UpdateAuthorized(ctx context.Context, info g.Map) bool
		// Unauthorized 授权取消
		Unauthorized(ctx context.Context, info g.Map) bool
	}
)

var (
	localUserAuth       IUserAuth
	localUserEvent      IUserEvent
	localAppAuth        IAppAuth
	localAppVersion     IAppVersion
	localMerchantNotify IMerchantNotify
	localWeiXinPay      IWeiXinPay
	localSubAccount     ISubAccount
	localSubMerchant    ISubMerchant
)

func UserAuth() IUserAuth {
	if localUserAuth == nil {
		panic("implement not found for interface IUserAuth, forgot register?")
	}
	return localUserAuth
}

func RegisterUserAuth(i IUserAuth) {
	localUserAuth = i
}

func UserEvent() IUserEvent {
	if localUserEvent == nil {
		panic("implement not found for interface IUserEvent, forgot register?")
	}
	return localUserEvent
}

func RegisterUserEvent(i IUserEvent) {
	localUserEvent = i
}

func AppAuth() IAppAuth {
	if localAppAuth == nil {
		panic("implement not found for interface IAppAuth, forgot register?")
	}
	return localAppAuth
}

func RegisterAppAuth(i IAppAuth) {
	localAppAuth = i
}

func AppVersion() IAppVersion {
	if localAppVersion == nil {
		panic("implement not found for interface IAppVersion, forgot register?")
	}
	return localAppVersion
}

func RegisterAppVersion(i IAppVersion) {
	localAppVersion = i
}

func MerchantNotify() IMerchantNotify {
	if localMerchantNotify == nil {
		panic("implement not found for interface IMerchantNotify, forgot register?")
	}
	return localMerchantNotify
}

func RegisterMerchantNotify(i IMerchantNotify) {
	localMerchantNotify = i
}

func WeiXinPay() IWeiXinPay {
	if localWeiXinPay == nil {
		panic("implement not found for interface IWeiXinPay, forgot register?")
	}
	return localWeiXinPay
}

func RegisterWeiXinPay(i IWeiXinPay) {
	localWeiXinPay = i
}

func SubAccount() ISubAccount {
	if localSubAccount == nil {
		panic("implement not found for interface ISubAccount, forgot register?")
	}
	return localSubAccount
}

func RegisterSubAccount(i ISubAccount) {
	localSubAccount = i
}

func SubMerchant() ISubMerchant {
	if localSubMerchant == nil {
		panic("implement not found for interface ISubMerchant, forgot register?")
	}
	return localSubMerchant
}

func RegisterSubMerchant(i ISubMerchant) {
	localSubMerchant = i
}
