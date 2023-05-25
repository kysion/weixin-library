package weixin_utility

import (
	"context"
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
