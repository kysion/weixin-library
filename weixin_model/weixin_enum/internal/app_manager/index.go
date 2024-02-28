package app_manager

type appManager struct {
	AuditState auditState
	AppType    appType
}

var AppManager = appManager{
	AuditState: AuditState,
	AppType:    AppType,
}
