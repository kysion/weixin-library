package app_manager

type appManager struct {
	AuditState auditState
}

var AppManager = appManager{
	AuditState: AuditState,
}
