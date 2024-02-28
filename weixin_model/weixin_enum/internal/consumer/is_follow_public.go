package consumer

import "github.com/kysion/base-library/utility/enum"

// IsFollowPublicEnum 是否关注公众号：1关注、2取消关注
type IsFollowPublicEnum enum.IEnumCode[int]

type isFollowPublic struct {
	Subscribe   IsFollowPublicEnum
	UnSubscribe IsFollowPublicEnum
}

var IsFollowPublic = isFollowPublic{
	Subscribe:   enum.New[IsFollowPublicEnum](1, "关注"),
	UnSubscribe: enum.New[IsFollowPublicEnum](2, "取消关注"),
}

func (e isFollowPublic) New(code int, description string) IsFollowPublicEnum {
	if (code & IsFollowPublic.Subscribe.Code()) == IsFollowPublic.Subscribe.Code() {
		return e.Subscribe
	}
	if (code & IsFollowPublic.UnSubscribe.Code()) == IsFollowPublic.UnSubscribe.Code() {
		return e.UnSubscribe
	}
	return enum.New[IsFollowPublicEnum](code, description)
}
