package weixin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/gopay"
	"github.com/kysion/gopay/pkg/xlog"
	wechat3 "github.com/kysion/gopay/wechat/v3"
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

func Decrypt(ctx context.Context, eventInfo weixin_model.EventEncryptMsgReq, msgInfo weixin_model.MessageEncryptReq) *weixin_model.EventMessageBody {
	config, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, eventInfo.AppId)

	token := config.MsgVerfiyToken

	if err == nil {
		// 创建解密对象
		instance := weixin_encrypt.NewWechatMsgCrypt(token, config.MsgEncryptKey, eventInfo.AppId)
		// 微信消息推送事件解密
		decryptData := instance.WechatEventDecrypt(eventInfo, msgInfo.MsgSignature, msgInfo.TimeStamp, msgInfo.Nonce)

		// 消息事件内容结构体
		data := weixin_model.EventMessageBody{}

		gconv.Struct(decryptData, &data)

		return &data
	}

	return nil
}
