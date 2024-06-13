package gateway

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/weixin-library/weixin_model"
	dao "github.com/kysion/weixin-library/weixin_model/weixin_dao"
	do "github.com/kysion/weixin-library/weixin_model/weixin_do"
	entity "github.com/kysion/weixin-library/weixin_model/weixin_entity"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

// 微信消费者配置表
type sConsumer struct {
	redisCache *gcache.Cache
	Duration   time.Duration
}

func NewConsumerConfig() *sConsumer {
	return &sConsumer{
		redisCache: gcache.New(),
	}
}

// GetConsumerById 根据id查找消费者信息
func (s *sConsumer) GetConsumerById(ctx context.Context, id int64) (*weixin_model.WeixinConsumerConfig, error) {
	result, err := daoctl.GetByIdWithError[weixin_model.WeixinConsumerConfig](dao.WeixinConsumerConfig.Ctx(ctx), id)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据ID查找消费者失败", dao.WeixinConsumerConfig.Table())
	}

	result = s.masker(ctx, result)

	return result, err
}

// GetConsumerBySysUserId  根据用户id查询消费者信息
func (s *sConsumer) GetConsumerBySysUserId(ctx context.Context, sysUserId int64) (*weixin_model.WeixinConsumerConfig, error) {
	result := weixin_model.WeixinConsumerConfig{}
	model := dao.WeixinConsumerConfig.Ctx(ctx)

	err := model.Where(dao.WeixinConsumerConfig.Columns().SysUserId, sysUserId).Scan(&result)

	if err != nil {
		return nil, err
	}

	s.masker(ctx, &result)

	return &result, nil
}

// GetConsumerByOpenId  根据用户openId查询消费者信息
func (s *sConsumer) GetConsumerByOpenId(ctx context.Context, openId string, unionId ...string) (*weixin_model.WeixinConsumerConfig, error) {
	result := weixin_model.WeixinConsumerConfig{}
	model := dao.WeixinConsumerConfig.Ctx(ctx)

	if len(unionId) > 0 {
		model = model.Where(dao.WeixinConsumerConfig.Columns().UnionId, unionId[0])
	}

	err := model.Where(dao.WeixinConsumerConfig.Columns().OpenId, openId).Scan(&result)

	if err != nil {
		return nil, err
	}

	s.masker(ctx, &result)

	return &result, nil
}

// GetConsumerByOpenIdAndAppId  根据用户openId和appId查询消费者信息
func (s *sConsumer) GetConsumerByOpenIdAndAppId(ctx context.Context, openId string, appId string, unionId ...string) (*weixin_model.WeixinConsumerConfig, error) {
	result := weixin_model.WeixinConsumerConfig{}
	model := dao.WeixinConsumerConfig.Ctx(ctx)

	if len(unionId) > 0 {
		model = model.Where(dao.WeixinConsumerConfig.Columns().UnionId, unionId[0])
	}

	err := model.Where(do.WeixinConsumerConfig{
		OpenId: openId,
		AppId:  appId,
	}).Scan(&result)

	if err != nil {
		return nil, err
	}

	s.masker(ctx, &result)

	return &result, nil
}

// QueryConsumerByUnionId  根据用户unionId查询消费者|列表
func (s *sConsumer) QueryConsumerByUnionId(ctx context.Context, unionId string) (*weixin_model.WeixinConsumerConfigListRes, error) {
	daoModel := dao.WeixinConsumerConfig.Ctx(ctx).Where(dao.WeixinConsumerConfig.Columns().UnionId, unionId)

	result, err := daoctl.Query[weixin_model.WeixinConsumerConfig](daoModel, nil, true) // 不分页
	if err != nil {
		return &weixin_model.WeixinConsumerConfigListRes{}, err
	}

	records := make([]weixin_model.WeixinConsumerConfig, 0)
	for _, record := range result.Records {
		masker := s.masker(ctx, &record)

		records = append(records, *masker)
	}

	result.Records = records

	return (*weixin_model.WeixinConsumerConfigListRes)(result), nil
}

// CreateConsumer  创建消费者信息
func (s *sConsumer) CreateConsumer(ctx context.Context, info *weixin_model.WeixinConsumerConfig) (*weixin_model.WeixinConsumerConfig, error) {
	data := do.WeixinConsumerConfig{}
	gconv.Struct(info, &data)

	if info.ExtJson == "" {
		data.ExtJson = nil
	}

	data.Id = idgen.NextId()
	data.UserState = 1 // 用户状态默认正常
	affected, err := daoctl.InsertWithError(
		dao.WeixinConsumerConfig.Ctx(ctx),
		data,
	)

	if affected == 0 || err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "消费者信息创建失败", dao.WeixinConsumerConfig.Table())
	}

	return s.GetConsumerById(ctx, gconv.Int64(data.Id))
}

// UpdateConsumer 更新消费者信息
func (s *sConsumer) UpdateConsumer(ctx context.Context, id int64, info *weixin_model.UpdateConsumerReq) (bool, error) {
	// 首先判断消费者信息是否存在
	consumerInfo, err := daoctl.GetByIdWithError[entity.WeixinConsumerConfig](dao.WeixinConsumerConfig.Ctx(ctx), id)
	if err != nil || consumerInfo == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "该消费者不存在", dao.WeixinConsumerConfig.Table())
	}
	data := do.WeixinConsumerConfig{}
	gconv.Struct(info, &data)

	model := dao.WeixinConsumerConfig.Ctx(ctx)
	affected, err := daoctl.UpdateWithError(model.Data(data).OmitEmptyData().Where(do.WeixinConsumerConfig{Id: id}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "消费者信息更新失败", dao.WeixinConsumerConfig.Table())
	}

	return affected > 0, nil
}

// UpdateConsumerByUserId 更新消费者信息
func (s *sConsumer) UpdateConsumerByUserId(ctx context.Context, userId int64, info *weixin_model.UpdateConsumerReq) (bool, error) {
	// 首先判断消费者信息是否存在
	//consumerInfo, err := daoctl.GetByIdWithError[entity.WeixinConsumerConfig](dao.WeixinConsumerConfig.Ctx(ctx), id)
	//if err != nil || consumerInfo == nil {
	//	return false, sys_service.SysLogs().ErrorSimple(ctx, err, "该消费者不存在", dao.WeixinConsumerConfig.Table())
	//}

	data := do.WeixinConsumerConfig{}
	gconv.Struct(info, &data)

	model := dao.WeixinConsumerConfig.Ctx(ctx)
	affected, err := daoctl.UpdateWithError(model.Data(data).OmitEmptyData().Where(do.WeixinConsumerConfig{SysUserId: userId}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "消费者信息更新失败", dao.WeixinConsumerConfig.Table())
	}

	return affected > 0, nil
}

// UpdateConsumerState 修改用户状态
func (s *sConsumer) UpdateConsumerState(ctx context.Context, id int64, state int) (bool, error) {
	affected, err := daoctl.UpdateWithError(dao.WeixinConsumerConfig.Ctx(ctx).Data(do.WeixinConsumerConfig{
		UserState: state,
	}).OmitNilData().Where(do.WeixinConsumerConfig{Id: id}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "消费者状态修改失败", dao.WeixinConsumerConfig.Table())
	}
	return affected > 0, err
}

// UpdateConsumerAuthState 修改用户授权状态
func (s *sConsumer) UpdateConsumerAuthState(ctx context.Context, id int64, state int) (bool, error) {
	affected, err := daoctl.UpdateWithError(dao.WeixinConsumerConfig.Ctx(ctx).Data(do.WeixinConsumerConfig{
		AuthState: state,
	}).OmitNilData().Where(do.WeixinConsumerConfig{Id: id}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "消费者授权状态修改失败", dao.WeixinConsumerConfig.Table())
	}
	return affected > 0, err
}

// SetIsFollowPublic 设置用户是够关注公众号
func (s *sConsumer) SetIsFollowPublic(ctx context.Context, openId string, appID string, isFollowPublic int) (bool, error) {
	affected, err := daoctl.UpdateWithError(dao.WeixinConsumerConfig.Ctx(ctx).Data(do.WeixinConsumerConfig{
		IsFollowPublic: isFollowPublic,
	}).OmitNilData().Where(do.WeixinConsumerConfig{OpenId: openId, AppId: appID}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "消费者是否关注公众号修改失败", dao.WeixinConsumerConfig.Table())
	}
	return affected > 0, err
}

// UpdateConsumerToken 更新消费者token等数据信息
func (s *sConsumer) UpdateConsumerToken(ctx context.Context, openId string, info *weixin_model.UpdateConsumerTokenReq) (bool, error) {
	// 首先判断消费者信息是否存在
	consumerInfo, err := daoctl.ScanWithError[entity.WeixinConsumerConfig](dao.WeixinConsumerConfig.Ctx(ctx).Where(do.WeixinConsumerConfig{OpenId: openId}))
	if err != nil || consumerInfo == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "该消费者不存在", dao.WeixinConsumerConfig.Table())
	}

	data := do.WeixinConsumerConfig{}
	gconv.Struct(info, &data)

	model := dao.WeixinConsumerConfig.Ctx(ctx)
	affected, err := daoctl.UpdateWithError(model.Data(data).OmitNilData().Where(do.WeixinConsumerConfig{OpenId: openId}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "消费者token信息更新失败", dao.WeixinConsumerConfig.Table())
	}

	return affected > 0, nil
}

func (s *sConsumer) masker(ctx context.Context, info *weixin_model.WeixinConsumerConfig) *weixin_model.WeixinConsumerConfig {
	if info == nil {
		return info
	}

	if (info.AuthState & weixin_enum.Consumer.AuthState.UnAuth.Code()) == weixin_enum.Consumer.AuthState.UnAuth.Code() {
		info.NickName = "微信用户"
		info.Avatar = "https://thirdwx.qlogo.cn/mmopen/vi_32/Q3auHgzwzM4ibNq0sk5FFv2hW3NXpfAlO936dcZQRYtwKF89coctv1I94F4qXMltNjGog9dv3yNwDZ3iaJMep3xw/132"
	}

	return info
}
