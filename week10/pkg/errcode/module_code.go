package errcode

var (
	ErrorUploadFileFail       = NewError(20030001, "上传文件失败")
	ErrorParentBalance        = NewError(20040001, "获取家长用户帐户余额失败")
	ErrorStudentBalanceRecord = NewError(20040002, "获取学生帐户pcpu记录失败")
	ErrorGetUserLineID        = NewError(20040003, "获取用户所属业务线失败")
)
