package weixin

import (
	"context"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
	wechat "github.com/go-pay/gopay/wechat"
	wechat3 "github.com/go-pay/gopay/wechat/v3"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/kys-weixin-library/utility"
	"log"
	"sort"
	"strings"

	"github.com/kysion/kys-weixin-library/service"
	"github.com/kysion/kys-weixin-library/utility/file"
)

type sWeiXin struct {
}

func init() {
	service.RegisterWeiXin(New())
}

func New() *sWeiXin {
	// 初始化文件内容
	privateData, _ = file.GetFile(priPath)
	publicCertData, _ = file.GetFile(publicCrtPath)
	appCertPublicKeyData, _ = file.GetFile(appCertPublicKeyPath)
	alipayRootCertData, _ = file.GetFile(alipayRootCertPath)
	alipayCertPublicKeyData, _ = file.GetFile(alipayCertPublicKeyPath)

	return &sWeiXin{}
}

// NewClient 初始化微信客户端
func NewClient(ctx context.Context, appId, mchId, apiKey string, isProd bool) *wechat.Client {

	// 初始化客户端 (应用ID,商户ID, API秘钥值, 是否是正式环境)
	client := wechat.NewClient(appId, mchId, apiKey, isProd)
	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugOn

	return client
}

// NewClient3 初始化微信V3客户端对象
func NewClient3(ctx context.Context, mchId, serialNo, aPIv3Key, privateKeyContent string) (*wechat3.ClientV3, error) {
	// 初始化V3版本客户端 (商户ID,商户证书的证书序列号, APIv3Key，商户平台获取,私钥 apiclient_key.pem)
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

func (s *sWeiXin) UserInfoAuth(ctx context.Context) (string, error) {

	return "", nil
}

// WeiXinServices 接收消息通知  B端消息
func (s *sWeiXin) WeiXinServices(ctx context.Context) (string, error) {

	fmt.Println("hello")

	return "success", nil
}

// WeiXinCallback 接收消息回调  C端消息
func (s *sWeiXin) WeiXinCallback(ctx context.Context) (string, error) {

	return "", nil

}

// WXCheckSignature 微信接入校验
func (s *sWeiXin) WXCheckSignature(ctx context.Context, signature, timestamp, nonce, echostr string) string {
	// 与填写的服务器配置中的Token一致
	const Token = "kysion.kuaimk"
	fmt.Println(signature + "、" + timestamp + "、" + nonce + "、" + echostr)
	arr := []string{timestamp, nonce, Token}
	// 字典序排序
	sort.Strings(arr)

	n := len(timestamp) + len(nonce) + len(Token)
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < len(arr); i++ {
		b.WriteString(arr[i])
	}

	sign := utility.Sha1(b.String())

	ok := utility.CheckSignature(sign, timestamp, nonce, Token)

	if !ok {
		log.Println("微信公众号接入校验失败!")
		return ""
	}

	log.Println("微信公众号接入校验成功!")

	g.RequestFromCtx(ctx).Response.Write(echostr)
	return echostr
}
