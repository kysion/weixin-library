package weixin_model

// 小程序代码提交审核------------------------------------------------------------------------------------------------------------------------
//
//
//{
//"item_list": [{
//	"address":"index",
//	"tag":"学习 生活",
//	"first_class": "文娱",
//	"second_class": "资讯",
//	"first_id":1,
//	"second_id":2,
//	"title": "首页"
//}],
//"feedback_info": "blablabla",
//"feedback_stuff": "xx|yy|zz",
//"preview_info" : {
//	"video_id_list": ["xxxx"],
//	"pic_id_list": ["xxxx", "yyyy", "zzzz" ]
//},
//"version_desc":"blablabla",
//"ugc_declare": {
//	"scene": [1,2],
//	"method": [1],
//	"has_audit_team": 1,
//	"audit_desc": "blablabla"
//}
//}
//

type SubmitAppVersionAuditReq struct {
	ItemList []ItemList `json:"item_list" dc:"审核项列表（选填，至多填写 5 项）；类目是必填的，且要填写已经在小程序配置好的类目。"`

	PreviewInfo `json:"preview_info" dc:"预览信息（小程序页面截图和操作录屏）"`

	FeedbackInfo     string `json:"feedback_info" dc:"【可选】反馈内容，至多 200 字"`
	FeedbackStuff    string `json:"feedback_stuff" dc:"【可选】用 | 分割的 media_id 列表，至多 5 张图片, 可以通过新增临时素材接口上传而得到"`
	VersionDesc      string `json:"version_desc" dc:"小程序版本说明和功能解释"`
	PrivacyApiNotUse bool   `json:"privacy_api_not_use" dc:"【可选】用于声明是否不使用“代码中检测出但是未配置的隐私相关接口"`
	OrderPath        string `json:"order_path" dc:"【可选】订单中心path"`

	UgcDeclare `json:"ugc_declare" dc:"用户生成内容场景（UGC）信息安全声明"`
}

type ItemList struct {
	Address     string `json:"address" dc:"否	小程序的页面，可通过获取小程序的页面列表getCodePage 接口获取"`
	Tag         string `json:"tag" dc:"否	小程序的标签，用空格分隔，标签至多10 个，标签长度至多 20"`
	FirstClass  string `json:"first_class" dc:"是	一级类目名称，可通过getAllCategoryName 接口获取"`
	SecondClass string `json:"second_class" dc:"是	二级类目名称，可通过getAllCategoryName 接口获取"`
	ThirdClass  string `json:"third_class" dc:"否	三级类目名称，可通过getAllCategoryName 接口获取"`
	Title       string `json:"title" dc:"否	小程序页面的标题,标题长度至多 32"`
	FirstId     int    `json:"first_id" dc:"是	一级类目id，可通过getAllCategoryName 接口获取"`
	SecondId    int    `json:"second_id" dc:"是	二级类目id，可通过getAllCategoryName 接口获取"`
	ThirdId     int    `json:"third_id" dc:"否	三级类目id，可通过getAllCategoryName 接口获取"`
}

type PreviewInfo struct {
	VideoIdList []string `json:"video_id_list" dc:"录屏mediaid列表，可以通过提审素材上传接口获得"`
	PicIdList   []string `json:"pic_id_list" dc:"截屏mediaid列表，可以通过提审素材上传接口获得"`
}

type UgcDeclare struct {
	Scene  []int `json:"scene" dc:"否	UGC场景 0,不涉及用户生成内容, 1.用户资料,2.图片,3.视频,4.文本,5音频, 可多选,当scene填0时无需填写下列字段"`
	Method []int `json:"method" dc:"否	内容安全机制 1.使用平台建议的内容安全API,2.使用其他的内容审核产品,3.通过人工审核把关,4.未做内容审核把关"`
	//OtherSceneDesc string `json:"other_scene_desc" dc:"否	当scene选其他时的说明,不超时256字"`
	HasAuditTeam int    `json:"has_audit_team" dc:"否	是否有审核团队, 0.无,1.有,默认0"`
	AuditDesc    string `json:"audit_desc" dc:"否	说明当前对UGC内容的审核机制,不超过256字"`
}

type ErrorCommon struct {
	Errcode int    `json:"errcode" dc:"返回码"`
	Errmsg  string `json:"errmsg" dc:"错误信息"`
}
type AppVersionAuditRes struct {
	ErrorCommon
	Auditid int `json:"auditid" dc:"审核编号"`
}

// 小程序版本回退------------------------------------------------------------------------------------------------------------------------

type CancelAppVersionReq struct {
	AccessToken string `json:"access_token" dc:"是	接口调用凭证，该参数为 URL 参数，非 Body 参数。使用authorizer_access_token"`
	Action      string `json:"action" dc:"否	只能填get_history_version。表示获取可回退的小程序版本。该参数为 URL 参数，非 Body 参数。"`
	AppVersion  string `json:"app_version" dc:"否	默认是回滚到上一个版本；也可回滚到指定的小程序版本，可通过get_history_version获取app_version。该参数为 URL 参数，非 Body 参数。"`
}

type CancelAppVersionRes struct {
	ErrorCommon
	VersionList VersionList `json:"version_list" dc:"模板信息列表。当action=get_history_version，才会返回。"`
}

type VersionList struct {
	AppVersion  int    `json:"app_version" dc:"小程序版本"`
	UserVersion string `json:"user_version" dc:"模板版本号，开发者自定义字段"`
	UserDesc    string `json:"user_desc" dc:"模板描述，开发者自定义字段"`
	CommitTime  int    `json:"commit_time" dc:"更新时间，时间戳"`
}

// 撤销小程序版本审核------------------------------------------------------------------------------------------------------------------------

type CancelAppVersionAuditRes struct {
	ErrorCommon
}

// 小程序已经上传代码页面的列表------------------------------------------------------------------------------------------------------------------------

type QueryAppVersionListRes struct {
	ErrorCommon
	PageList []string `json:"page_list" dc:"page_list 页面配置列表"`
}

// 小程序版本详情------------------------------------------------------------------------------------------------------------------------

type QueryAppVersionDetailRes struct {
	ErrorCommon
	ExpInfo     ExpInfo     `json:"exp_info" dc:"体验版信息"`
	ReleaseInfo ReleaseInfo `json:"release_info" dc:"线上版信息"`
}

type ExpInfo struct {
	ExpTime    int    `json:"exp_time" dc:"提交体验版的时间"`
	ExpVersion string `json:"exp_version" dc:"体验版版本信息"`
	ExpDesc    string `json:"exp_desc" dc:"体验版版本描述"`
}

type ReleaseInfo struct {
	ReleaseTime    int    `json:"release_time" dc:"发布线上版的时间"`
	ReleaseVersion string `json:"release_version" dc:"线上版版本信息"`
	ReleaseDesc    string `json:"release_desc" dc:"线上版本描述"`
}

// GetAppLatestVersionAuditRes 查询最新一次审核单状态 ------------------------------------------------------------------------------------------------------------------------

type GetAppLatestVersionAuditRes struct {
	ErrorCommon

	Auditid         int    `json:"auditid" dc:"最新的审核id"`
	Status          int    `json:"status" dc:"审核状态"`
	Reason          string `json:"reason" dc:"当审核被拒绝时，返回的拒绝原因"`
	ScreenShot      string `json:"ScreenShot" dc:"当审核被拒绝时，会返回审核失败的小程序截图示例。用 竖线I 分隔的 media_id 的列表，可通过获取永久素材接口拉取截图内容"`
	UserVersion     string `json:"user_version" dc:"审核版本"`
	UserDesc        string `json:"user_desc" dc:"版本描述"`
	SubmitAuditTime int    `json:"submit_audit_time" dc:"时间戳，提交审核的时间"`
}

// 小程序审核素材上传接口 ------------------------------------------------------------------------------------------------------------------------

type UploadAppMediaToAuditRes struct {
	ErrorCommon
	Type    string `json:"type"dc:"类型"`
	Mediaid string `json:"mediaid" dc:"媒体id"`
}

// 小程序上传代码接口------------------------------------------------------------------------------------------------------------------------

type CommitAppAuditCodeReq struct {
	//AccessToken string `json:"access_token" dc:"是，接口调用凭证，该参数为 URL 参数，非 Body 参数。使用authorizer_access_token"`
	TemplateId  int    `json:"template_id" dc:"是	代码库中的代码模板 ID，可通过getTemplateList接口获取代码模板template_id。注意，如果该模板id为标准模板库的模板id"` // 则ext_json可支持的参数为：{"extAppid":" ", "ext": {}, "window": {}}
	ExtJson     string `json:"ext_json" dc:"是	为了方便第三方平台的开发者引入 extAppid 的开发调试工作，引入ext.json配置文件概念，该参数则是用于控制ext.json配置文件的内容。关于该参数的补充说明请查看下方的"`
	UserVersion string `json:"user_version" dc:"是	代码版本号，开发者可自定义（长度不要超过 64 个字符）"`
	UserDesc    string `json:"user_desc" dc:"是	代码描述，开发者可自定义"`
}

type CommitAppAuditCodeRes struct {
	ErrorCommon
}
