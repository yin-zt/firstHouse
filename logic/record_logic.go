package logic

import (
	"firstHouse/model"
	"firstHouse/model/request"
	"firstHouse/model/response"
	"firstHouse/public/tools"
	"firstHouse/service/isql"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RecordLogic struct{}

// Add 添加数据
func (l RecordLogic) Add(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.RecordAddReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	//if isql.Record.Exist(tools.H{"username": r.Username}) {
	//	return nil, tools.NewValidatorError(fmt.Errorf("用户名已存在,请勿重复添加"))
	//}
	//if isql.Record.Exist(tools.H{"mobile": r.Mobile}) {
	//	return nil, tools.NewValidatorError(fmt.Errorf("手机号已存在,请勿重复添加"))
	//}
	//if isql.Record.Exist(tools.H{"job_number": r.JobNumber}) {
	//	return nil, tools.NewValidatorError(fmt.Errorf("工号已存在,请勿重复添加"))
	//}
	//if isql.Record.Exist(tools.H{"mail": r.Mail}) {
	//	return nil, tools.NewValidatorError(fmt.Errorf("邮箱已存在,请勿重复添加"))
	//}

	// 密码通过RSA解密
	// 密码不为空就解密
	//if r.Password != "" {
	//	decodeData, err := tools.RSADecrypt([]byte(r.Password), config.Conf.System.RSAPrivateBytes)
	//	if err != nil {
	//		return nil, tools.NewValidatorError(fmt.Errorf("密码解密失败"))
	//	}
	//	r.Password = string(decodeData)
	//	if len(r.Password) < 6 {
	//		return nil, tools.NewValidatorError(fmt.Errorf("密码长度至少为6位"))
	//	}
	//} else {
	//	r.Password = "123456"
	//}

	//// 当前登陆用户角色排序最小值（最高等级角色）以及当前登陆的用户
	//currentRoleSortMin, ctxUser, err := isql.User.GetCurrentUserMinRoleSort(c)
	//if err != nil {
	//	return nil, tools.NewValidatorError(fmt.Errorf("获取当前登陆用户角色排序最小值失败"))
	//}
	//
	//// 根据角色id获取角色
	//if r.RoleIds == nil || len(r.RoleIds) == 0 {
	//	r.RoleIds = []uint{2} // 默认添加为普通用户角色
	//}
	//
	//roles, err := isql.Role.GetRolesByIds(r.RoleIds)
	//if err != nil {
	//	return nil, tools.NewValidatorError(fmt.Errorf("根据角色ID获取角色信息失败"))
	//}
	//
	//var reqRoleSorts []int
	//for _, role := range roles {
	//	reqRoleSorts = append(reqRoleSorts, int(role.Sort))
	//}
	//// 前端传来用户角色排序最小值（最高等级角色）
	//reqRoleSortMin := uint(funk.MinInt(reqRoleSorts).(int))
	//
	//// 如果登录用户的角色ID为1，亦即为管理员，则直接放行，保障管理员拥有最大权限
	//if currentRoleSortMin != 1 {
	//	// 当前用户的角色排序最小值 需要小于 前端传来的角色排序最小值（用户不能创建比自己等级高的或者相同等级的用户）
	//	if currentRoleSortMin >= reqRoleSortMin {
	//		return nil, tools.NewValidatorError(fmt.Errorf("用户不能创建比自己等级高的或者相同等级的用户"))
	//	}
	//}
	record := model.Record{
		Fview:             r.Fview,
		Fzone:             r.Fzone,
		Fdomain:           r.Fdomain,
		Ftype:             r.Ftype,
		Ftarget:           r.Ftarget,
		Ftarget_slave:     r.Ftarget_slave,
		Ftarget_flag:      r.Ftarget_flag,
		Fonline_status:    r.Fonline_status,
		Fsystem_name:      r.Fsystem_name,
		Fbusiness_line:    r.Fbusiness_line,
		Fcharge_person:    r.Fcharge_person,
		Fcandidate_person: r.Fcandidate_person,
		Fdeveloper:        r.Fdeveloper,
		Fbackup_person:    r.Fbackup_person,
		Fprotocol_port:    r.Fprotocol_port,
		Fdetection_time:   r.Fdetection_time,
		Ftest_type:        r.Ftest_type,
		Flevel:            r.Flevel,
		Fmaintain:         r.Fmaintain,
		Fversion:          r.Fversion,
	}

	err := CommonAddRecord(&record)
	if err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("添加用户失败" + err.Error()))
	}
	return nil, nil
}

// List 数据列表
func (l RecordLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.RecordListReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	records, err := isql.Record.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取用户列表失败：" + err.Error()))
	}

	rets := make([]model.Record, 0)
	for _, record := range records {
		rets = append(rets, *record)
	}
	count, err := isql.Record.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取用户总数失败：" + err.Error()))
	}

	return response.UserListRsp{
		Total: int(count),
		Users: rets,
	}, nil
}

//// Update 更新数据
//func (l UserLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
//	r, ok := req.(*request.UserUpdateReq)
//	if !ok {
//		return nil, ReqAssertErr
//	}
//	_ = c
//
//	if !isql.User.Exist(tools.H{"id": r.ID}) {
//		return nil, tools.NewMySqlError(fmt.Errorf("该记录不存在"))
//	}
//
//	// 获取当前登陆用户
//	ctxUser, err := isql.User.GetCurrentLoginUser(c)
//	if err != nil {
//		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户失败"))
//	}
//
//	// 获取当前登陆用户角色ID集合
//	var currentRoleSorts []int
//	for _, role := range ctxUser.Roles {
//		currentRoleSorts = append(currentRoleSorts, int(role.Sort))
//	}
//
//	// 获取将要操作的用户角色ID集合
//	var reqRoleSorts []int
//	roles, err := isql.Role.GetRolesByIds(r.RoleIds)
//	if err != nil || len(roles) == 0 {
//		return nil, tools.NewValidatorError(fmt.Errorf("根据角色ID获取角色信息失败"))
//	}
//	for _, role := range roles {
//		reqRoleSorts = append(reqRoleSorts, int(role.Sort))
//	}
//
//	// 当前登陆用户角色排序最小值（最高等级角色）
//	currentRoleSortMin := funk.MinInt(currentRoleSorts).(int)
//	// 前端传来用户角色排序最小值（最高等级角色）
//	reqRoleSortMin := funk.MinInt(reqRoleSorts).(int)
//
//	// 如果登录用户的角色ID为1，亦即为管理员，则直接放行，保障管理员拥有最大权限
//	if currentRoleSortMin != 1 {
//		// 判断是更新自己还是更新别人,如果操作的ID与登陆用户的ID一致，则说明操作的是自己
//		if int(r.ID) == int(ctxUser.ID) {
//			// 不能更改自己的角色
//			reqDiff, currentDiff := funk.Difference(reqRoleSorts, currentRoleSorts)
//			if len(reqDiff.([]int)) > 0 || len(currentDiff.([]int)) > 0 {
//				return nil, tools.NewValidatorError(fmt.Errorf("不能更改自己的角色"))
//			}
//		}
//
//		// 如果是更新别人，操作者不能更新比自己角色等级高的或者相同等级的用户
//		minRoleSorts, err := isql.User.GetUserMinRoleSortsByIds([]uint{uint(r.ID)}) // 根据userIdID获取用户角色排序最小值
//		if err != nil || len(minRoleSorts) == 0 {
//			return nil, tools.NewValidatorError(fmt.Errorf("根据用户ID获取用户角色排序最小值失败"))
//		}
//		if currentRoleSortMin >= minRoleSorts[0] || currentRoleSortMin >= reqRoleSortMin {
//			return nil, tools.NewValidatorError(fmt.Errorf("用户不能更新比自己角色等级高的或者相同等级的用户"))
//		}
//	}
//
//	// 先获取用户信息
//	oldData := new(model.User)
//	err = isql.User.Find(tools.H{"id": r.ID}, oldData)
//	if err != nil {
//		return nil, tools.NewMySqlError(err)
//	}
//
//	// 拼装新的用户信息
//	user := model.User{
//		Model:         oldData.Model,
//		Username:      r.Username,
//		Nickname:      r.Nickname,
//		GivenName:     r.GivenName,
//		Mail:          r.Mail,
//		JobNumber:     r.JobNumber,
//		Mobile:        r.Mobile,
//		Avatar:        r.Avatar,
//		PostalAddress: r.PostalAddress,
//		Departments:   r.Departments,
//		Position:      r.Position,
//		Introduction:  r.Introduction,
//		Creator:       ctxUser.Username,
//		DepartmentId:  tools.SliceToString(r.DepartmentId, ","),
//		Source:        oldData.Source,
//		Roles:         roles,
//	}
//
//	if err = CommonUpdateUser(oldData, &user, r.DepartmentId); err != nil {
//		return nil, tools.NewOperationError(fmt.Errorf("更新用户失败" + err.Error()))
//	}
//
//	return nil, nil
//}

// Delete 删除数据
func (l RecordLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.RecordDeleteReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	for _, id := range r.RecordIds {
		filter := tools.H{"id": int(id)}
		if !isql.Record.Exist(filter) {
			return nil, tools.NewMySqlError(fmt.Errorf("有解析不存在"))
		}
	}

	// 再将用户从MySQL中删除
	err := isql.Record.Delete(r.RecordIds)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("在MySQL删除解析记录失败: " + err.Error()))
	}

	return nil, nil
}

//
//// ChangePwd 修改密码
//func (l UserLogic) ChangePwd(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
//	r, ok := req.(*request.UserChangePwdReq)
//	if !ok {
//		return nil, ReqAssertErr
//	}
//	_ = c
//	// 前端传来的密码是rsa加密的,先解密
//	// 密码通过RSA解密
//	decodeOldPassword, err := tools.RSADecrypt([]byte(r.OldPassword), config.Conf.System.RSAPrivateBytes)
//	if err != nil {
//		return nil, tools.NewValidatorError(fmt.Errorf("原密码解析失败"))
//	}
//	decodeNewPassword, err := tools.RSADecrypt([]byte(r.NewPassword), config.Conf.System.RSAPrivateBytes)
//	if err != nil {
//		return nil, tools.NewValidatorError(fmt.Errorf("新密码解析失败"))
//	}
//	r.OldPassword = string(decodeOldPassword)
//	r.NewPassword = string(decodeNewPassword)
//	// 获取当前用户
//	user, err := isql.User.GetCurrentLoginUser(c)
//	if err != nil {
//		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户失败"))
//	}
//	if tools.NewParPasswd(user.Password) != r.OldPassword {
//		return nil, tools.NewValidatorError(fmt.Errorf("原密码错误"))
//	}
//
//	// 更新密码
//	err = isql.User.ChangePwd(user.Username, tools.NewGenPasswd(r.NewPassword))
//	if err != nil {
//		return nil, tools.NewMySqlError(fmt.Errorf("在MySQL更新密码失败: " + err.Error()))
//	}
//
//	return nil, nil
//}
//
//// ChangeUserStatus 修改用户状态
//func (l UserLogic) ChangeUserStatus(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
//	r, ok := req.(*request.UserChangeUserStatusReq)
//	if !ok {
//		return nil, ReqAssertErr
//	}
//	_ = c
//	// 校验工作
//	filter := tools.H{"id": r.ID}
//	if !isql.User.Exist(filter) {
//		return nil, tools.NewValidatorError(fmt.Errorf("该用户不存在"))
//	}
//	user := new(model.User)
//	err := isql.User.Find(filter, user)
//	if err != nil {
//		return nil, tools.NewMySqlError(fmt.Errorf("在MySQL查询用户失败: " + err.Error()))
//	}
//
//	if r.Status == 1 && r.Status == user.Status {
//		return nil, tools.NewValidatorError(fmt.Errorf("用户已经是在职状态"))
//	}
//	if r.Status == 2 && r.Status == user.Status {
//		return nil, tools.NewValidatorError(fmt.Errorf("用户已经是离职状态"))
//	}
//
//	// 获取当前登录用户，只有管理员才能够将用户状态改变
//	// 获取当前登陆用户角色排序最小值（最高等级角色）以及当前用户
//	minSort, _, err := isql.User.GetCurrentUserMinRoleSort(c)
//	if err != nil {
//		return nil, tools.NewValidatorError(fmt.Errorf("获取当前登陆用户角色排序最小值失败"))
//	}
//
//	if int(minSort) != 1 {
//		return nil, tools.NewValidatorError(fmt.Errorf("只有管理员才能更改用户状态"))
//	}
//
//	err = isql.User.ChangeStatus(int(r.ID), int(r.Status))
//	if err != nil {
//		return nil, tools.NewMySqlError(fmt.Errorf("在MySQL更新用户状态失败: " + err.Error()))
//	}
//	return nil, nil
//}
//
//// GetUserInfo 获取用户信息
//func (l UserLogic) GetUserInfo(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
//	r, ok := req.(*request.UserGetUserInfoReq)
//	if !ok {
//		return nil, ReqAssertErr
//	}
//
//	_ = c
//	_ = r
//
//	user, err := isql.User.GetCurrentLoginUser(c)
//	if err != nil {
//		return nil, tools.NewMySqlError(fmt.Errorf("获取当前用户信息失败: " + err.Error()))
//	}
//	return user, nil
//}
