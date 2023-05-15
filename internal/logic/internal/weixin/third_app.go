package weixin

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
)

// GetAuthorizerList 拉取已授权的帐号信息|列表
func GetAuthorizerList(ctx context.Context, info *weixin_model.GetAuthorizerList) (*weixin_model.GetAuthorizerListRes, error) {
	app, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, info.ComponentAppid)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "服务商应用不存在", "WeiXin服务商服务")
	}

	queryAuthUrl := "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_list?access_token=" + app.AppAuthToken
	req := weixin_model.GetAuthorizerList{
		ComponentAppid: info.ComponentAppid,
		Offset:         info.Offset,
		Count:          info.Count,
	}

	if info.Count == 0 {
		req.Count = 20
	}
	//if info.Offset == 0 {
	//	req.Offset = 1
	//}

	reqJson, _ := gjson.Encode(req)

	res := g.Client().PostContent(ctx, queryAuthUrl, reqJson)

	//	{"errcode":41002,"errmsg":"appid missing rid: 645a2314-3a072a6c-18abe153"}

	resData := weixin_model.GetAuthorizerListRes{}
	gjson.DecodeTo(res, &resData)

	fmt.Println("已授权的帐号信息列表数量：", resData.TotalCount)

	return &resData, nil
}

// GetOpenAccount 获取应用绑定的开放平台账号
func GetOpenAccount(ctx context.Context, appId, authorizerAccessToken string) (*weixin_model.GetOpenAccountRes, error) {
	//app, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, info.ComponentAppid)
	//if err != nil {
	//	return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "服务商应用不存在", "WeiXin服务商服务")
	//}

	queryAuthUrl := "https://api.weixin.qq.com/cgi-bin/open/get?access_token=" + authorizerAccessToken
	req := g.Map{}
	if appId != "" {
		req["appid"] = appId
	}

	//if info.Offset == 0 {
	//	req.Offset = 1
	//}

	reqJson, _ := gjson.Encode(req)

	res := g.Client().PostContent(ctx, queryAuthUrl, reqJson)

	resData := weixin_model.GetOpenAccountRes{}
	gjson.DecodeTo(res, &resData)

	fmt.Println("绑定的开放平台账号AppId为：", resData.OpenAppid)

	return &resData, nil
}
