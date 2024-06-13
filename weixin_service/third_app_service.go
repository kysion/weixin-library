// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package weixin_service

import (
	"context"

	"github.com/kysion/weixin-library/weixin_model"
)

type (
	IThirdService interface {
		// GetAuthorizerList 拉取已授权的账号列表
		GetAuthorizerList(ctx context.Context, info *weixin_model.GetAuthorizerList) (*weixin_model.GetAuthorizerListRes, error)
		// GetOpenAccount 获取应用绑定的开放平台账号 open_app_id
		GetOpenAccount(ctx context.Context, appId, authorizerAccessToken string) (*weixin_model.GetOpenAccountRes, error)
	}
)

var (
	localThirdService IThirdService
)

func ThirdService() IThirdService {
	if localThirdService == nil {
		panic("implement not found for interface IThirdService, forgot register?")
	}
	return localThirdService
}

func RegisterThirdService(i IThirdService) {
	localThirdService = i
}
