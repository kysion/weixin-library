package weixin

import (
	"context"
	"crypto/rsa"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
)

/*
	微信支付SDK的通用方法
*/

// NewPayClient 初始化微信支付客户端对象
func NewPayClient(ctx context.Context, privateKey, mchId, mchCertificateSerialNumber, mchAPIv3Key string) (client *core.Client, err error) {
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, _ := loadPrivateKey(privateKey)

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchId, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}

	client, err = core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatal("new pay client error")
		return nil, err
	}

	return client, nil
}

// 加载商户私钥
func loadPrivateKey(privateKey string) (res *rsa.PrivateKey, err error) {
	var mchPrivateKey *rsa.PrivateKey

	if gfile.IsFile(privateKey) {
		mchPrivateKey, err = utils.LoadPrivateKey(privateKey)
	} else {
		//mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/data/kysion-files/weixin/weixin-pay/1642565036_cert/apiclient_key.pem")
		mchPrivateKey, err = utils.LoadPrivateKeyWithPath(privateKey)
	}

	if err != nil {
		log.Fatal("load merchant private key error")

		return nil, err
	}

	return mchPrivateKey, nil
}

// 加载平台证书
