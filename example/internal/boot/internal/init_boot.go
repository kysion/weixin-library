package internal

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_consts"
)

func init() {
	InitGlobal()
	// InitFileData()

}

// InitGlobal 初始化公共对象
func InitGlobal() {
	//weixin_consts.Global.AppId = g.Cfg().MustGet(context.Background(), "service.appId").String()
	//weixin_consts.Global.AppSecret = g.Cfg().MustGet(context.Background(), "service.appSecret").String()
	//weixin_consts.Global.Token = g.Cfg().MustGet(context.Background(), "service.token").String()
	//weixin_consts.Global.DecryptKey = g.Cfg().MustGet(context.Background(), "service.decryptKey").String()
	weixin_consts.Global.PayCertP12Path = g.Cfg().MustGet(context.Background(), "service.payCertP12").String()
	weixin_consts.Global.PayPublicKeyPemPath = g.Cfg().MustGet(context.Background(), "service.payPublicKeyPem").String()
	weixin_consts.Global.PayPrivateKeyPemPath = g.Cfg().MustGet(context.Background(), "service.payPrivateKeyPem").String()

	// 交易Hook失效时间
	weixin_consts.Global.TradeHookExpireAt = g.Cfg().MustGet(context.Background(), "service.tradeHookExpireAt").Int64()
}

//// InitGlobal 初始化公共对象
//func InitGlobal() {
//	weixin_consts.Global.PayCertP12Path = g.Cfg().MustGet(context.Background(), "service.payCertP12").String()
//	weixin_consts.Global.PayPublicKeyPemPath = g.Cfg().MustGet(context.Background(), "service.payPublicKeyPem").String()
//	weixin_consts.Global.PayPrivateKeyPemPath = g.Cfg().MustGet(context.Background(), "service.payPrivateKeyPem").String()
//
//}
//
//func InitFileData() {
//	// 加载证书文件
//	privateData, _ := file.GetFile(alipay_consts.Global.PriPath)
//	publicCertData, _ := file.GetFile(alipay_consts.Global.PublicCrtPath)
//	appCertPublicKeyData, _ := file.GetFile(alipay_consts.Global.AppCertPublicKeyPath)
//	alipayRootCertData, _ := file.GetFile(alipay_consts.Global.AlipayRootCertPath)
//	alipayCertPublicKeyData, _ := file.GetFile(alipay_consts.Global.AlipayCertPublicKeyPath)
//	info := alipay_model.UpdateThirdKeyCertReq{
//		AppId:                   "2021003179681073",
//		PrivateKey:              string(privateData),
//		PublicKey:               string(publicCertData),
//		PublicKeyCert:           string(alipayCertPublicKeyData),
//		AppPublicCertKey:        string(appCertPublicKeyData),
//		AlipayRootCertPublicKey: string(alipayRootCertData),
//	}
//	fmt.Println(info)
//
//	_, err := alipay_service.ThirdAppConfig().UpdateThirdKeyCert(context.Background(), &info)
//	if err != nil {
//		fmt.Println("证书文件存储失败啦~")
//	}
//}
