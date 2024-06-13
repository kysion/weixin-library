package weixin

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
)

/*
	微信支付SDK的通用方法
*/

// NewPayClient 初始化微信支付客户端对象
func NewPayClient(ctx context.Context, mchId, privateKey, mchCertificateSerialNumber, mchAPIv3Key string) (client *core.Client, err error) {
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, _ := LoadPrivateKey(privateKey)

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchId, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}

	client, err = core.NewClient(ctx, opts...)
	if err != nil {
		//log.Fatal("new pay client error")
		return nil, err
	}

	return client, nil
}

// LoadPrivateKey 加载商户私钥
func LoadPrivateKey(privateKey string) (res *rsa.PrivateKey, err error) {
	var mchPrivateKey *rsa.PrivateKey

	if gfile.IsFile(privateKey) {
		mchPrivateKey, err = utils.LoadPrivateKeyWithPath(privateKey)
	} else {
		//mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/data/kysion-files/weixin/weixin-pay/1642565036_cert/apiclient_key.pem")
		mchPrivateKey, err = utils.LoadPrivateKey(privateKey)
	}

	if err != nil {
		log.Fatal("load merchant private key error")

		return nil, err
	}

	return mchPrivateKey, nil
}

// LoadWeXinCert 加载平台证书
func LoadWeXinCert(cert string) (res *x509.Certificate, err error) {
	var certificate *x509.Certificate

	if gfile.IsFile(cert) {
		certificate, err = utils.LoadCertificateWithPath(cert)
	} else {
		certificate, err = utils.LoadCertificate(cert)
	}

	if err != nil {
		log.Fatal("load certificate  error")

		return nil, err
	}

	return certificate, nil

}

// SignSHA256WithRSA 通过私钥对字符串以 SHA256WithRSA 算法生成签名信息
func SignSHA256WithRSA(source string, privateKey *rsa.PrivateKey) (signature string, err error) {
	if privateKey == nil {
		return "", fmt.Errorf("private key should not be nil")
	}
	//var hash crypto.Hash
	//h := hash.New()
	//
	h := sha256.New()
	_, err = h.Write([]byte(source))
	if err != nil {
		return "", nil
	}
	hashed := h.Sum(nil)
	signatureByte, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signatureByte), nil
}

// NewNotifyHandler 使用微信支付平台证书（验签）和商户 APIv3 密钥（解密）初始化 notify.Handler
func NewNotifyHandler(ctx context.Context, spMerchant *weixin_model.PayMerchant) *notify.Handler {
	// 1. 使用 `RegisterDownloaderWithPrivateKey` 注册下载器
	pri, _ := LoadPrivateKey(spMerchant.PayPrivateKeyPem)
	mchId := gconv.String(spMerchant.Mchid)

	err := downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, pri, spMerchant.CertSerialNumber, mchId, spMerchant.ApiV3Key)
	if err != nil {
		return nil
	}
	// 2. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(mchId)

	// 3. 使用证书访问器初始化 `notify.Handler`
	handler := notify.NewNotifyHandler(spMerchant.ApiV3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))

	return handler
}
