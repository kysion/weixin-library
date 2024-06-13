package merchant

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/weixin-library/internal/logic/internal/weixin"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher/decryptors"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher/encryptors"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"io/ioutil"
	"log"
	"net/http"
)

type sSubMerchant struct {
}

func init() {
	//weixin_service.RegisterSubMerchant(NewSubMerchant())
}

func NewSubMerchant() *sSubMerchant {

	result := &sSubMerchant{}

	return result
}

func newSubClient(ctx context.Context, spMchId string, no ...string) (*core.Client, error) {
	spMerchant, err := weixin_service.PayMerchant().GetPayMerchantByMchid(ctx, gconv.Int(spMchId))
	if err != nil {
		return nil, err
	}
	//spMerchant, err := weixin_service.PayMerchant().GetPayMerchantByMchid(ctx, gconv.Int(subMerchant.SpMchid))
	//if err != nil {
	//	return nil, err
	//}

	// 使用敏感信息加解密器cipher.Cipher，根据 API 契约自动处理敏感信息：
	privateKey, err := weixin.LoadPrivateKey(spMerchant.PayPrivateKeyPem)

	// 一次性设置 签名/验签/敏感字段加解密，并注册 平台证书下载器，自动定时获取最新的平台证书
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(spMchId, spMerchant.CertSerialNumber, privateKey, spMerchant.ApiV3Key),
		option.WithWechatPayCipher(
			encryptors.NewWechatPayEncryptor(downloader.MgrInstance().GetCertificateVisitor(spMchId)),
			decryptors.NewWechatPayDecryptor(privateKey),
		),
	}

	client, err := core.NewClient(ctx, opts...)

	if err != nil {
		//log.Fatal("new pay client error")
		return nil, err
	}

	return client, nil
}

// 提交申请单 （人工）

// 查询申请单状态

// 修改结算账号

// 查询结算账号

// 查询结算账户修改审核状态

// GetAuditStateByBusinessCode 根据业务申请编号查询申请状态
func (s *sSubMerchant) GetAuditStateByBusinessCode(ctx context.Context, spMchId, businessCode string) (*weixin_model.SubMerchantAuditStateRes, error) {
	//client := core.Client{}

	client, _ := newSubClient(ctx, spMchId)

	result, err := client.Get(ctx, consts.WechatPayAPIServer+"/v3/applyment4sub/applyment/business_code/"+businessCode)

	res := weixin_model.SubMerchantAuditStateRes{}
	body, err := ioutil.ReadAll(result.Response.Body)

	err = gjson.DecodeTo(body, &res)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "特约商户进件申请状态查询失败", "WeiXin-Sub-Merchant")
	}

	return &res, nil
}

// GetAuditStateByApplymentId 根据申请单号查询申请状态
func (s *sSubMerchant) GetAuditStateByApplymentId(ctx context.Context, spMchId, applymentId string) (*weixin_model.SubMerchantAuditStateRes, error) {
	//client := core.Client{}

	client, _ := newSubClient(ctx, spMchId)

	result, err := client.Get(ctx, consts.WechatPayAPIServer+"/v3/applyment4sub/applyment/applyment_id/"+applymentId)

	res := weixin_model.SubMerchantAuditStateRes{}
	body, err := ioutil.ReadAll(result.Response.Body)

	err = gjson.DecodeTo(body, &res)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "特约商户进件申请状态查询失败", "WeiXin-Sub-Merchant")
	}

	return &res, nil
}

// GetSettlement 查询结算账号
func (s *sSubMerchant) GetSettlement(ctx context.Context, subMchId string) (*weixin_model.SettlementRes, error) {
	url := fmt.Sprintf("/v3/apply4sub/sub_merchants/%s/settlement", subMchId)

	subMerchant, err := weixin_service.PaySubMerchant().GetPaySubMerchantByMchid(ctx, gconv.Int(subMchId))
	if err != nil {
		return nil, err
	}

	client, _ := newSubClient(ctx, gconv.String(subMerchant.SpMchid))

	result, err := client.Get(ctx, consts.WechatPayAPIServer+url)

	resp := weixin_model.SettlementRes{}
	body, err := ioutil.ReadAll(result.Response.Body)
	gjson.DecodeTo(body, &resp)

	return &resp, err
}

// UpdateSettlement 修改结算账号,成功会返回application_no，作为查询申请状态的唯一标识
func (s *sSubMerchant) UpdateSettlement(ctx context.Context, subMchId string, info *weixin_model.UpdateSettlementReq) (string, error) {
	subMerchant, err := weixin_service.PaySubMerchant().GetPaySubMerchantByMchid(ctx, gconv.Int(subMchId))
	if err != nil {
		return "", err
	}

	//spMerchant, err := weixin_service.PayMerchant().GetPayMerchantByMchid(ctx, gconv.Int(subMerchant.SpMchid))
	//if err != nil {
	//	return "", err
	//}

	// 微信平台证书
	certificates, err := weixin_service.WeiXinPay().DownloadCertificates(ctx, subMerchant.SubAppid)

	client, _ := newSubClient(ctx, gconv.String(subMerchant.SpMchid), *certificates.Data[0].SerialNo)

	// 1.构建签名串
	/*
		HTTP请求方法\n
		URL\n
		请求时间戳\n
		请求随机串\n
		请求报文主体\n
	*/

	url := fmt.Sprintf("/v3/apply4sub/sub_merchants/%s/modify-settlement", subMchId)
	//bodyContentJson, err := gjson.Encode(info) //请求报文

	//var content = "GET" + "\n" +
	//	url + "\n" +
	//	strconv.FormatInt(time.Now().Unix(), 10) + "\n" +
	//	gconv.String(idgen.NextId()) + "\n" +
	//	gconv.String(bodyContentJson) + "\n"
	//
	//var authorization, _ = weixin.SignSHA256WithRSA(content, privateKey) // qDLKva8l1HPQ0GDjQA9cHMqIg8cI4JWv0/toKBoA+8dSgIKKySQniAv8AKapAj3DHX1Td6xS9Tgm2LPUewdP4KkZ6aYOdbtiDLaoCiuLNud4S0mTsek7Re9oOaA5OCIqsz2E5AYOWJkGxebrIOhWAWChKiT/+JKZXWdBozuYIN0tqtirfK3xuhaPszlx0sJwD0V7Gn2tYK9VVVVYfpFNdXZeQaehdpDVfj5xkVXaH8yQwweoljoy1qWC+UFmZ+/8TIu5w3OslMnbrWIlMOckJdfnv5bXyvkChzETfO4R46eiOdkXi1dP6759S9FZn7JVFglu22aJdTVk3g7e8BmtHA==

	// 2.发起请求，client里面包含生成签名，加密敏感数据逻辑
	//result, err := client.Post(ctx, consts.WechatPayAPIServer+url, info)  // TODO 报错没有设置请求头Wechatpay-Serial，

	result, err := client.Request(ctx, "POST", consts.WechatPayAPIServer+url, http.Header{ // TODO 请确认待处理的消息是否为加密后的密文
		"Wechatpay-Serial": []string{*certificates.Data[0].SerialNo},
	}, nil, info, "application/json")

	// 3.解析响应结果
	if err != nil {
		// 处理错误
		log.Printf("call AddReceiver err:%s", err)
		return "", err
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, result)
	}

	resp := new(weixin_model.UpdateSettlementRes)

	// 处理成功响应结果方法1
	//body, err := ioutil.ReadAll(result.Response.Body)
	//gjson.DecodeTo(body, &resp)

	// 处理成功响应结果方法2
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return "", err
	}

	applicationNo := resp.ApplicationNo

	return applicationNo, err
}

// GetSettlementAuditState 查询结算账户修改审核状态
func (s *sSubMerchant) GetSettlementAuditState(ctx context.Context, subMchId, applicationNo string) (*weixin_model.SettlementRes, error) {
	subMerchant, err := weixin_service.PaySubMerchant().GetPaySubMerchantByMchid(ctx, gconv.Int(subMchId))
	if err != nil {
		return nil, err
	}

	client, _ := newSubClient(ctx, gconv.String(subMerchant.SpMchid))

	url := fmt.Sprintf("/v3/apply4sub/sub_merchants/%s/application/%s", subMchId, applicationNo)

	result, err := client.Get(ctx, url)

	// 3.解析响应结果
	if err != nil {
		// 处理错误
		log.Printf("call AddReceiver err:%s", err)
		return nil, err
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, result)
	}

	resp := weixin_model.SettlementRes{}
	body, err := ioutil.ReadAll(result.Response.Body)
	gjson.DecodeTo(body, &resp)

	return &resp, err
}

// 	publicKey, err := utils.LoadPublicKey(spMerchant.PayPublicKeyPem)
//	fmt.Println(publicKey)
//	//accountNumber, err := utils.EncryptOAEPWithPublicKey(info.AccountNumber, publicKey)
