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

type StatByEmployee struct {
	EmployeeName string
	TsrInWork    int32
	TsrFinished  int32
	TsrApplyed   int32
}

type StatByEmployeeReq struct {
	EmplotyeeUUID string
}

type FullStat struct {
	ByDepartment []*StatByDepartment
	ByEmployee   []*StatByEmployee
}
