package app_manager

import "github.com/kysion/base-library/utility/enum"

// AuditStateEnum  审核状态：0审核成功、1审核被拒绝、2审核中、3已撤回、4审核延后、

type AuditStateEnum enum.IEnumCode[int]

type auditState struct {
	Success   AuditStateEnum
	Refuse    AuditStateEnum
	InReview  AuditStateEnum
	Withdraw  AuditStateEnum
	Postponed AuditStateEnum
}

var AuditState = auditState{
	Success:   enum.New[AuditStateEnum](0, "0审核成功"),
	Refuse:    enum.New[AuditStateEnum](1, "1审核被拒绝"),
	InReview:  enum.New[AuditStateEnum](2, "2审核中"),
	Withdraw:  enum.New[AuditStateEnum](3, "3已撤回"),
	Postponed: enum.New[AuditStateEnum](4, "4审核延后"),
}

func (e auditState) New(code int, description string) AuditStateEnum {
	if (code&AuditState.Success.Code()) == AuditState.Success.Code() ||
		(code&AuditState.Refuse.Code()) == AuditState.Refuse.Code() ||
		(code&AuditState.InReview.Code()) == AuditState.InReview.Code() ||
		(code&AuditState.Withdraw.Code()) == AuditState.Withdraw.Code() ||
		(code&AuditState.Postponed.Code()) == AuditState.Postponed.Code() {
		return enum.New[AuditStateEnum](code, description)
	}
	panic("consumerAuditState: error")
}
