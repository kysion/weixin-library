package weixin

import (
	"context"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
	wechat "github.com/go-pay/gopay/wechat"
	wechat3 "github.com/go-pay/gopay/wechat/v3"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/kys-weixin-library/utility"
	"github.com/kysion/kys-weixin-library/utility/weixin_encrypt"
	"github.com/kysion/kys-weixin-library/weixin_consts"
	"github.com/kysion/kys-weixin-library/weixin_model"
	service "github.com/kysion/kys-weixin-library/weixin_service"
	"log"
	"sort"
	"strings"
	"time"
)

type sWeiXin struct {
}

func init() {
	service.RegisterWeiXin(New())
}

func New() *sWeiXin {
	// 初始化文件内容

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
func (s *sWeiXin) WeiXinServices(ctx context.Context, eventInfo *weixin_model.EventEncryptMsgReq, msgInfo *weixin_model.MessageEncryptReq) (string, error) {
	fmt.Println("hello~ 我是消息接收")

	// 1.验签
	ok := utility.VerifyByteDanceServer(weixin_consts.Global.Token, msgInfo.TimeStamp, msgInfo.Nonce, msgInfo.Encrypt, msgInfo.MsgSignature)
	if !ok {
		fmt.Println("验签失败")
		return "success", nil
	}

	// 2.解密
	data := Decrypt(*eventInfo, *msgInfo)
	fmt.Println("解密后的内容：", data)

	//{"errcode":41018,"errmsg":"missing component_appid rid: 64021f9c-2ffaddeb-572ebf0a"}

	// 通过Hook处理不同类型的请求。。。。。。。。 先写

	// 3.获取token令牌 (第三方平台接口的调用凭据 access_token)
	tokenUrl := "https://api.weixin.qq.com/cgi-bin/component/api_component_token?component_appid=" + weixin_consts.Global.AppId +
		"&component_appsecret=" + weixin_consts.Global.AppSecret +
		"&component_verify_ticket=" + data.ComponentVerifyTicket
	componentAccessToken := g.Client().PostContent(ctx, tokenUrl)

	fmt.Println(tokenUrl)

	componentAccessTokenRes := weixin_model.ComponentAccessTokenRes{}
	gjson.DecodeTo(componentAccessToken, &componentAccessTokenRes)
	fmt.Println("令牌：", componentAccessTokenRes)

	// 缓存componentAccessToken 第三方接口调用凭据
	if componentAccessTokenRes.ComponentAccessToken != "" {
		gcache.Set(ctx, weixin_consts.Global.AppId+"_component_access_token", componentAccessTokenRes.ComponentAccessToken, time.Duration(componentAccessTokenRes.ExpiresIn))
	}

	/*
		AppId = {string} "wx534d1a08aa84c529"
		CreateTime = {int} 1677855987
		InfoType = {string} "component_verify_ticket"
		ComponentVerifyTicket = {string} "ticket@@@Bb3RjaKczF7YiV-mdama4Qzmo6x5H72QsZSWsCSfs1fs0XiWoMF5UY7Yix_-24W9RdKXn-yHHHOKyLwD8t79FA"
		AuthorizerAppid = {string} ""
		AuthorizationCode = {string} ""
		AuthorizationCodeExpiredTime = {string} ""
		PreAuthCode = {string} ""
	*/

	return "success", nil
}

func Decrypt(eventInfo weixin_model.EventEncryptMsgReq, msgInfo weixin_model.MessageEncryptReq) *weixin_model.EventMessageBody {
	// 创建解密对象
	instance := weixin_encrypt.NewWechatMsgCrypt(weixin_consts.Global.Token, weixin_consts.Global.DecryptKey, weixin_consts.Global.AppId)
	// 微信消息推送事件解密
	decryptData := instance.WechatEventDecrypt(eventInfo, msgInfo.MsgSignature, msgInfo.TimeStamp, msgInfo.Nonce)

	// 消息事件内容结构体
	data := weixin_model.EventMessageBody{}

	gconv.Struct(decryptData, &data)

	return &data
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
