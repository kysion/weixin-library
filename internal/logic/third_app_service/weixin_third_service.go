package third_app_service

import (
	"context"
	"github.com/kysion/weixin-library/internal/logic/internal/weixin"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
)

/*
	微信第三方平台应用服务：
	1.拉取已授权的账号列表
	2.获取应用绑定的开放平台账号
*/

var ThirdService = sThirdService{}

type sThirdService struct{}

//
//func init() {
//	weixin_service.RegisterThirdService(NewThirdService())
//}

func NewThirdService() weixin_service.IThirdService {

	result := &sThirdService{}

	//result.injectHook()
	return result
}

// GetAuthorizerList 拉取已授权的账号列表
func (s *sThirdService) GetAuthorizerList(ctx context.Context, info *weixin_model.GetAuthorizerList) (*weixin_model.GetAuthorizerListRes, error) {
	return weixin.GetAuthorizerList(ctx, info)
}

// GetOpenAccount 获取应用绑定的开放平台账号 open_app_id
func (s *sThirdService) GetOpenAccount(ctx context.Context, appId, authorizerAccessToken string) (*weixin_model.GetOpenAccountRes, error) {
	return weixin.GetOpenAccount(ctx, appId, authorizerAccessToken)
}
