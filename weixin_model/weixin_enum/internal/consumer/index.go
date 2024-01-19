package consumer

type consumer struct {
	SexEnum        sex
	ActionEnum     action
	Category       category
	IsFollowPublic isFollowPublic
	AuthState      authState
}

var Consumer = consumer{
	SexEnum:        SexType,
	ActionEnum:     ActionType,
	Category:       CategoryType,
	IsFollowPublic: IsFollowPublic,
	AuthState:      AuthState,
}
