package weixin_consts

type global struct {
	AppId      string
	AppSecret  string
	Token      string
	DecryptKey string
}

var (
	Global = global{
		AppId:      "",
		AppSecret:  "",
		Token:      "",
		DecryptKey: "",
	}
)
