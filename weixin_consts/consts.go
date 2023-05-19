package weixin_consts

type global struct {
	PayCertP12Path       string // 包含私钥的证书文件
	PayPublicKeyPemPath  string // 公钥
	PayPrivateKeyPemPath string // 私钥

	TradeHookExpireAt int64 // 交易Hook过期时间
}

var (
	Global = global{}
)

//type global struct {
//	AppId      string
//	AppSecret  string
//	Token      string
//	DecryptKey string
//}
//
//var (
//	Global = global{
//		AppId:      "",
//		AppSecret:  "",
//		Token:      "",
//		DecryptKey: "",
//	}
//)

const (
	API_COMPONENT_TOKEN = "https://api.weixin.qq.com/cgi-bin/component/api_component_token"

	API_QUERY_AUTH = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token="
)
