package weixin_hook

import (
	"context"
	entity "github.com/kysion/weixin-library/weixin_model/weixin_entity"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
)

type PlatFormUserHookFunc func(ctx context.Context, info entity.PlatformUser) int64

type PlatFormUserHookInfo struct {
	Key   weixin_enum.ConsumerAction
	Value PlatFormUserHookFunc
}
