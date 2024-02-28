package merchant

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	dao "github.com/kysion/weixin-library/weixin_model/weixin_dao"
)

// GetJwtToken 根据平台userId换取JwtToken
func GetJwtToken(ctx context.Context, sysUserId int64) (*sys_model.TokenInfo, error) {
	// 根据平台userId去第三方平台用户关系表查询该用户是否存在
	//info, err := service.PlatformUser().GetPlatformUserByUserId(ctx, platFormUserId)
	//info, err := service.PlatformUser().GetPlatformUserByUserIdAndType(ctx, platFormUserId, userType)

	//if err != nil {
	//	return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据平台用户id获取接口Token失败", dao.PlatformUser.Table())
	//}

	// 根据我们系统的用户id获取用户数据，然后生成Token
	sysUser, err := sys_service.SysUser().GetSysUserById(ctx, sysUserId)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "用户不存在，未注册请先注册", sys_dao.SysUser.Table())
	}

	// 存在返回JwtToken
	if sysUser != nil {
		jwtToken, err := sys_service.Jwt().GenerateToken(ctx, sysUser)
		if err != nil {
			return nil, err
		}

		return jwtToken, nil
	}

	// 不存在返回 "",err
	return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "换取接口JwtToken失败", dao.PlatformUser.Table())
}
