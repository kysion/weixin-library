package weixin

import (
	"context"
	"encoding/xml"
	"github.com/go-pay/gopay"
	wechat3 "github.com/go-pay/gopay/wechat/v3"
	"github.com/go-pay/xlog"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
	"github.com/kysion/weixin-library/weixin_utility/weixin_encrypt"
)

// NewClient 初始化微信V3客户端对象
func NewClient(ctx context.Context, mchId, serialNo, aPIv3Key, privateKeyContent string) (*wechat3.ClientV3, error) {
	// 初始化V3版本客户端 (商户ID,商户证书的证书序列号, APIv3Key，商户平台获取,私钥 apiClient_key.pem)
	client, err := wechat3.NewClientV3(mchId, serialNo, aPIv3Key, privateKeyContent)
	if err != nil {
		xlog.Error(err)
		return nil, err
	}

	// 启用自动同步返回验签，并定时更新微信平台API证书
	err = client.AutoVerifySign()
	if err != nil {
		xlog.Error(err)
		return nil, err
	}

	client.DebugSwitch = gopay.DebugOff

	return client, nil
}

// DecryptEvent 解密事件推送
func DecryptEvent(ctx context.Context, eventInfo weixin_model.EventEncryptMsgReq, msgInfo weixin_model.MessageEncryptReq) *weixin_model.EventMessageBody {
	var msgEncryptKey string
	var token string
	// 第三方待开发模式
	config, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, eventInfo.AppId)
	if config != nil && err == nil {
		msgEncryptKey = config.MsgEncryptKey
		token = config.MsgVerifyToken
	}

	// 自开发模式
	if config == nil || config.Id == 0 {
		merchantConfig, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, eventInfo.AppId)
		if merchantConfig != nil && err == nil {
			msgEncryptKey = merchantConfig.MsgEncryptKey
			token = merchantConfig.MsgVerifyToken
		}
	}

	if token != "" && msgEncryptKey != "" {
		// 创建解密对象
		instance := weixin_encrypt.NewWechatMsgCrypt(token, msgEncryptKey, eventInfo.AppId)
		// 微信消息推送事件解密
		decryptData := instance.WechatEventDecrypt(eventInfo, msgInfo.MsgSignature, msgInfo.TimeStamp, msgInfo.Nonce)

		//fmt.Println("解密后的密文：", decryptData)
		// 消息事件内容结构体
		data := weixin_model.EventMessageBody{}

		gconv.Struct(decryptData, &data)

		return &data
	}

	return nil
}

// DecryptMessage 解密消息通知
func DecryptMessage(ctx context.Context, eventInfo weixin_model.EventEncryptMsgReq, msgInfo weixin_model.MessageEncryptReq) *weixin_model.MessageBodyDecrypt {
	var msgEncryptKey string
	var token string
	config, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, eventInfo.AppId)
	if config != nil && err == nil {
		msgEncryptKey = config.MsgEncryptKey
		token = config.MsgVerifyToken
	}

	// TODO 代码暂时比较丑陋，后续优化
	if config == nil || config.Id == 0 {
		merchantConfig, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, eventInfo.AppId)
		if merchantConfig != nil && err == nil {
			msgEncryptKey = merchantConfig.MsgEncryptKey
			token = merchantConfig.MsgVerifyToken
		}
	}

	if token != "" && msgEncryptKey != "" {
		// 创建解密对象
		instance := weixin_encrypt.NewWechatMsgCrypt(token, msgEncryptKey, eventInfo.AppId)
		// 微信消息通知事件解密
		//decryptData := instance.WechatMessageDecrypt(eventInfo, msgInfo.MsgSignature, msgInfo.TimeStamp, msgInfo.Nonce)
		decryptData := instance.WechatMessageDecrypt(weixin_encrypt.MessageEncryptRequest{
			XMLName:      xml.Name{},
			Encrypt:      eventInfo.Encrypt,
			MsgSignature: msgInfo.MsgSignature,
			TimeStamp:    msgInfo.TimeStamp,
			Nonce:        msgInfo.Nonce,
		})

		//fmt.Println("解密后的密文：", decryptData)
		// 消息通知内容结构体
		data := weixin_model.MessageBodyDecrypt{}

		gconv.Struct(decryptData, &data)

		return &data
	}

	return nil
}
