package weixin_model

import "time"

type SubMerchantAuditStateRes struct {
	BusinessCode      string        `json:"business_code" dc:"业务申请编号"`
	ApplymentId       int64         `json:"applyment_id" dc:"微信支付申请单号" `
	SubMchid          string        `json:"sub_mchid" dc:"特约商户号"`
	SignUrl           string        `json:"sign_url" dc:"超级管理员签约链接"`
	ApplymentState    string        `json:"applyment_state" dc:"申请单状态"`
	ApplymentStateMsg string        `json:"applyment_state_msg" dc:"申请状态描述"`
	AuditDetail       []AuditDetail `json:"audit_detail" dc:"audit_detail驳回原因详情"`
}

/*
申请单状态：
	1、APPLYMENT_STATE_EDITTING（编辑中）：提交申请发生错误导致，请尝试重新提交。
	2、APPLYMENT_STATE_AUDITING（审核中）：申请单正在审核中，超级管理员用微信打开“签约链接”，完成绑定微信号后，申请单进度将通过微信公众号通知超级管理员，引导完成后续步骤。
	3、APPLYMENT_STATE_REJECTED（已驳回）：请按照驳回原因修改申请资料，超级管理员用微信打开“签约链接”，完成绑定微信号，后续申请单进度将通过微信公众号通知超级管理员。
	4、APPLYMENT_STATE_TO_BE_CONFIRMED（待账户验证）：请超级管理员使用微信打开返回的“签约链接”，根据页面指引完成账户验证。
	5、APPLYMENT_STATE_TO_BE_SIGNED（待签约）：请超级管理员使用微信打开返回的“签约链接”，根据页面指引完成签约。
	6、APPLYMENT_STATE_SIGNING（开通权限中）：系统开通相关权限中，请耐心等待。
	7、APPLYMENT_STATE_FINISHED（已完成）：商户入驻申请已完成。
	8、APPLYMENT_STATE_CANCELED（已作废）：申请单已被撤销。
*/

type AuditDetail struct {
	Field        string `json:"field" dc:"字段名"`
	FieldName    string `json:"field_name" dc:"字段名称"`
	RejectReason string `json:"reject_reason" dc:"驳回原因"`
}

type UpdateSettlementReq struct {
	ModifyMode  string `json:"modify_mode" dc:"修改模式,可选，无论是否传入，均按照受理模式执行"`
	AccountType string `json:"account_type" dc:"根绝特约商户主体类型，可选择的账户类型如下：ACCOUNT_TYPE_BUSINESS: 对公银行账户 ，ACCOUNT_TYPE_PRIVATE: 经营者个人银行卡"`
	/*
		账户类型如下：
			1、小微主体：经营者个人银行卡
			2、个体工商户主体：经营者个人银行卡/ 对公银行账户
			3、企业主体：对公银行账户
			4、党政、机关及事业单位主体：对公银行账户
			5、其他组织主体：对公银行账户
		可选取值：
			ACCOUNT_TYPE_BUSINESS: 对公银行账户
			ACCOUNT_TYPE_PRIVATE: 经营者个人银行卡
	*/
	AccountBank     string `json:"account_bank" dc:"开户银行名称。"`
	BankAddressCode string `json:"bank_address_code" dc:"开户银行省市编码"`
	BankName        string `json:"bank_name" dc:"开户银行全称（含支行）"`
	BankBranchId    string `json:"bank_branch_id" dc:"开户银行联行号"`
	AccountNumber   string `json:"account_number" dc:"银行账号，该字段需进行敏感信息加密处理"`
	AccountName     string `json:"account_name" dc:"开户名称，该字段需进行敏感信息加密处理"`
}

/*
UpdateSettlementReq 示例值：
{
	"modify_mode" : "MODIFY_MODE_ASYNC",
	"account_type" : "ACCOUNT_TYPE_BUSINESS",
	"account_bank" : "工商银行",
	"bank_address_code" : "110000",
	"bank_name" : "中国工商银行股份有限公司北京市分行营业部",
	"bank_branch_id" : "402713354941",
	"account_number" : "d+xT+MQCvrLHUVDWv/8MR/dB7TkXM2YYZlokmXzFsWs35NXUot7C0NcxIrUF5FnxqCJHkNgKtxa6RxEYyba1+VBRLnqKG2fSy/Y5qDN08Ej9zHCwJjq52Wg1VG8MRugli9YMI1fI83KGBxhuXyemgS/hqFKsfYGiOkJqjTUpgY5VqjtL2N4l4z11T0ECB/aSyVXUysOFGLVfSrUxMPZy6jWWYGvT1+4P633f+R+ki1gT4WF/2KxZOYmli385ZgVhcR30mr4/G3HBcxi13zp7FnEeOsLlvBmI1PHN4C7Rsu3WL8sPndjXTd75kPkyjqnoMRrEEaYQE8ZRGYoeorwC+w==",
	"account_name" : "VyOMa+SncfM4lLha65dsxZ/xYW1zqBVVp6/W5mNkolESJU9fqgMt0lxjtuiWdhR+qUjnC2dTfuJuCOZs/Qi6kmicogGFjDC9ZxzFpdjR7AidWDuCIId5WRnRN8lGUcVyxctZZ4WcxxL2ADq57h7dZoFxNgyRYR4Y6q37LpYDccmYO5SiCkUP3rMX1CrTwKJysVhHij62HiU/P/yScImgdKrc+/MBWb1O6TT2RgwU3U6IwSZRWx4QH4EmYBLAQTdcEyUz2wuDmPA4nMSeXJVyzKl/WB+QYBh4Yj+BLT0HkA2IbTRyGX1U2wvv3N/w59Xq0pWYSXMHlmxhle2Cqj/7Cw=="
}
*/

type UpdateSettlementRes struct {
	ApplicationNo string `json:"application_no" dc:"修改结算账户申请返回的申请单号"`
}

type SettlementRes struct {
	AccountType   string `json:"account_type" dc:"根绝特约商户主体类型，可选择的账户类型如下：ACCOUNT_TYPE_BUSINESS: 对公银行账户 ，ACCOUNT_TYPE_PRIVATE: 经营者个人银行卡"`
	AccountBank   string `json:"account_bank" dc:"开户银行名称。"`
	BankName      string `json:"bank_name" dc:"开户银行全称（含支行）"`
	BankBranchId  string `json:"bank_branch_id" dc:"开户银行联行号开户银行全称（含支行）"`
	AccountNumber string `json:"account_number" dc:"银行账号"`
	VerifyResult  string `json:"verify_result" dc:"【验证结果】 返回特约商户的结算账户-验证结果"`
	// 验证结果:
	// 		VERIFY_SUCCESS: 验证成功，该账户可正常发起提现。
	// 		VERIFY_FAIL: 验证失败，该账户无法发起提现，请检查修改。
	// 		VERIFYING: 验证中，商户可发起提现尝试。
	VerifyFailReason string    `json:"verify_fail_reason" dc:"【验证失败原因】"`
	VerifyFinishTime time.Time `json:"verify_finish_time" dc:"【审核结果更新时间】"`
}
