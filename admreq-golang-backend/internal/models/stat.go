package models

type StatByDepartment struct {
	DepartmentName string
	TsrInWork      int32
	TsrFinished    int32
	TsrApplyed     int32
}

type StatByDepartmentReq struct {
	TargetDepartmentUUID string
	SourceDepartmentUUID string
}
