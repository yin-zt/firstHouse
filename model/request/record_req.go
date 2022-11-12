package request

// UserAddReq 创建资源结构体
type RecordAddReq struct {
	Fview             string `json:"view" validate:"required,min=2,max=20"`
	Fzone             string `json:"zone" validate:"required,min=2,max=30"`
	Fdomain           string `json:"domain" validate:"required,min=2,max=50"`
	Ftype             string `json:"type" validate:"required,min=1,max=10"`
	Ftarget           string `json:"target" validate:"required,min=2,max=50"`
	Ftarget_slave     string `json:"target_slave" validate:"min=0,max=50"`
	Ftarget_flag      string `json:"target_flag" validate:"required,min=2,max=20"`
	Fonline_status    string `json:"status" validate:"required,min=1,max=10"`
	Fsystem_name      string `json:"system" validate:"min=0,max=255"`
	Fbusiness_line    string `json:"line" validate:"required,min=0,max=100"`
	Fcharge_person    string `json:"charge" validate:"required,min=0,max=100"`
	Fcandidate_person string `json:"candidate" validate:"min=0,max=300"`
	Fdeveloper        string `json:"developer" validate:"min=0,max=300"`
	Fbackup_person    string `json:"backupPerson" validate:"min=0,max=300"`
	Fprotocol_port    string `json:"port" validate:"required,min=2,max=30"`
	Fdetection_time   string `json:"detect" validate:"min=0,max=100"`
	Ftest_type        string `json:"testtype" validate:"required,min=0,max=100"`
	Flevel            string `json:"level" validate:"required,min=0,max=100"`
	Fmaintain         string `json:"maintain" validate:"required,min=0,max=100"`
	Fversion          string `json:"version" validate:"min=0,max=100"`
}

// DingUserAddReq 钉钉用户创建资源结构体
type DingUserAddReq struct {
	Username      string `json:"username" validate:"required,min=2,max=20"`
	Password      string `json:"password"`
	Nickname      string `json:"nickname" validate:"required,min=0,max=20"`
	GivenName     string `json:"givenName" validate:"min=0,max=20"`
	Mail          string `json:"mail" validate:"required,min=0,max=100"`
	JobNumber     string `json:"jobNumber" validate:"required,min=0,max=20"`
	PostalAddress string `json:"postalAddress" validate:"min=0,max=255"`
	Departments   string `json:"departments" validate:"min=0,max=255"`
	Position      string `json:"position" validate:"min=0,max=255"`
	Mobile        string `json:"mobile" validate:"required,checkMobile"`
	Avatar        string `json:"avatar"`
	Introduction  string `json:"introduction" validate:"min=0,max=255"`
	Status        uint   `json:"status" validate:"oneof=1 2"`
	DepartmentId  []uint `json:"departmentId" validate:"required"`
	Source        string `json:"source" validate:"min=0,max=20"`
	RoleIds       []uint `json:"roleIds" validate:"required"`
	SourceUserId  string `json:"sourceUserId"`  // 第三方用户id
	SourceUnionId string `json:"sourceUnionId"` // 第三方唯一unionId
}

// WeComUserAddReq 企业微信用户创建资源结构体
type WeComUserAddReq struct {
	Username      string `json:"username" validate:"required,min=2,max=20"`
	Password      string `json:"password"`
	Nickname      string `json:"nickname" validate:"required,min=0,max=20"`
	GivenName     string `json:"givenName" validate:"min=0,max=20"`
	Mail          string `json:"mail" validate:"required,min=0,max=100"`
	JobNumber     string `json:"jobNumber" validate:"required,min=0,max=20"`
	PostalAddress string `json:"postalAddress" validate:"min=0,max=255"`
	Departments   string `json:"departments" validate:"min=0,max=255"`
	Position      string `json:"position" validate:"min=0,max=255"`
	Mobile        string `json:"mobile" validate:"required,checkMobile"`
	Avatar        string `json:"avatar"`
	Introduction  string `json:"introduction" validate:"min=0,max=255"`
	Status        uint   `json:"status" validate:"oneof=1 2"`
	DepartmentId  []uint `json:"departmentId" validate:"required"`
	Source        string `json:"source" validate:"min=0,max=20"`
	RoleIds       []uint `json:"roleIds" validate:"required"`
	SourceUserId  string `json:"sourceUserId"`  // 第三方用户id
	SourceUnionId string `json:"sourceUnionId"` // 第三方唯一unionId
}

// UserUpdateReq 更新资源结构体
type UserUpdateReq struct {
	ID            uint   `json:"id" validate:"required"`
	Username      string `json:"username" validate:"required,min=2,max=20"`
	Nickname      string `json:"nickname" validate:"min=0,max=20"`
	GivenName     string `json:"givenName" validate:"min=0,max=20"`
	Mail          string `json:"mail" validate:"min=0,max=100"`
	JobNumber     string `json:"jobNumber" validate:"min=0,max=20"`
	PostalAddress string `json:"postalAddress" validate:"min=0,max=255"`
	Departments   string `json:"departments" validate:"min=0,max=255"`
	Position      string `json:"position" validate:"min=0,max=255"`
	Mobile        string `json:"mobile" validate:"checkMobile"`
	Avatar        string `json:"avatar"`
	Introduction  string `json:"introduction" validate:"min=0,max=255"`
	DepartmentId  []uint `json:"departmentId" validate:"required"`
	Source        string `json:"source" validate:"min=0,max=20"`
	RoleIds       []uint `json:"roleIds" validate:"required"`
}

// RecordDeleteReq 批量删除资源结构体
type RecordDeleteReq struct {
	RecordIds []uint `json:"recordIds" validate:"required"`
}

// UserChangePwdReq 修改密码结构体
type UserChangePwdReq struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}

// UserChangeUserStatusReq 修改用户状态结构体
type UserChangeUserStatusReq struct {
	ID     uint `json:"id" validate:"required"`
	Status uint `json:"status" validate:"oneof=1 2"`
}

// UserGetUserInfoReq 获取用户信息结构体
type UserGetUserInfoReq struct {
}

// SyncDingUserReq 同步钉钉用户信息
type SyncDingUserReq struct {
}

// SyncWeComUserReq 同步企业微信用户信息
type SyncWeComUserReq struct {
}

// SyncFeiShuUserReq 同步飞书用户信息
type SyncFeiShuUserReq struct {
}

// SyncOpenLdapUserReq 同步ldap用户信息
type SyncOpenLdapUserReq struct {
}

// RecordListReq 获取解析记录列表结构体
type RecordListReq struct {
	Fview             string `json:"view" form:"view"`
	Fzone             string `json:"zone" form:"zone"`
	Fdomain           string `json:"domain" form:"domain"`
	Ftype             string `json:"type" form:"type"`
	Ftarget           string `json:"target" form:"target"`
	Ftarget_slave     string `json:"target_slave" form:"target_slave"`
	Ftarget_flag      string `json:"target_flag" form:"flag"`
	Fonline_status    string `json:"status" form:"status"`
	Fsystem_name      string `json:"system" form:"system"`
	Fbusiness_line    string `json:"line" form:"line"`
	Fcharge_person    string `json:"charge" form:"charge"`
	Fcandidate_person string `json:"candidate" form:"candidate"`
	Fdeveloper        string `json:"developer" form:"developer"`
	Fbackup_person    string `json:"backupPerson" form:"backupPerson"`
	Fprotocol_port    string `json:"port" form:"port"`
	Fdetection_time   string `json:"detect" form:"detect"`
	Ftest_type        string `json:"testtype" form:"testtype"`
	Flevel            string `json:"level" form:"level"`
	Fmaintain         string `json:"maintain" form:"maintain"`
	Fversion          string `json:"version" form:"version"`
	PageNum           int    `json:"pageNum" form:"pageNum"`
	PageSize          int    `json:"pageSize" form:"pageSize"`
}

// RegisterAndLoginReq 用户登录结构体
type RegisterAndLoginReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
