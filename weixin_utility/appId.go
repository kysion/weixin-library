package weixin_utility

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

func GetAppIdFormContext(ctx context.Context) string {
	// wx6pgm80gb0ejfj    编码后，32进制
	// wx6cc2c80416074df3 原始AppID，16进制
	pathAppId := g.RequestFromCtx(ctx).Get("appId").String() // wx6pgm80gb0ejfj

	appId := ""

	appIdLen := len(pathAppId)
	subAppId := gstr.SubStr(pathAppId, 2, appIdLen) // 6pgm80gb0ejfj

	if len(subAppId) < 16 {
		id := Base32ToHex(subAppId)
		appId = gconv.String(id) // 6cc2c80416074df3
	} else {
		appId = subAppId
	}

	return "wx" + appId // wx6cc2c80416074df3
}

// WeiXinAppIdEncode 微信 - appId的32进制编码
func WeiXinAppIdEncode(weiXinAppId string) (appIdEncode string) {
	appLen := len(weiXinAppId)                      // wxcaf4b7b8d6620f00
	subAppId := gstr.SubStr(weiXinAppId, 2, appLen) // caf4b7b8d6620f00

	// 十六进制转32进制
	appIdBase32Encode := HexToBase32(subAppId) // clt5nn3b643o0
	return "wx" + appIdBase32Encode            // wxclt5nn3b643o0
}

// WeiXinAppIdDecode 微信 - appId的32进制解码
func WeiXinAppIdDecode(appIdEncode string) (weiXinAppId string) {
	appIdLen := len(appIdEncode)                      // wxclt5nn3b643o0
	subAppId := gstr.SubStr(appIdEncode, 2, appIdLen) // clt5nn3b643o0

	id := Base32ToHex(subAppId)
	appId := gconv.String(id) // caf4b7b8d6620f00

	return "wx" + appId
}

func exampleTest() {
	appID := "wx60be13f2f91586f9"
	fmt.Println()
	encode := WeiXinAppIdEncode(appID)
	fmt.Println(encode)

	decode := WeiXinAppIdDecode(encode)
	fmt.Println(decode)
}
