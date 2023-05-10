package weixin_controller

import (
	"context"
	"github.com/kysion/weixin-library/internal/logic/internal/weixin"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
)

var ThirdService = sThirdService{}

type sThirdService struct{}

func init() {
	weixin_service.RegisterThirdService(NewThirdService())
}

func NewThirdService() *sThirdService {

	result := &sThirdService{}

	//result.injectHook()
	return result
}

// GetAuthorizerList 拉取已授权的账号列表
func (s *sThirdService) GetAuthorizerList(ctx context.Context, info *weixin_model.GetAuthorizerList) (*weixin_model.GetAuthorizerListRes, error) {
	return weixin.GetAuthorizerList(ctx, info)
}
