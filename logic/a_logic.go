package logic

import (
	"fmt"

	"firstHouse/model"
	"firstHouse/public/tools"
	"firstHouse/service/isql"
)

var (
	ReqAssertErr = tools.NewRspError(tools.SystemErr, fmt.Errorf("请求异常"))

	//Api          = &ApiLogic{}
	Record = &RecordLogic{}
	//Group        = &GroupLogic{}
	//Role         = &RoleLogic{}
	//Menu         = &MenuLogic{}
	//OperationLog = &OperationLogLogic{}
	//Base         = &BaseLogic{}
)

//// CommonAddGroup 标准创建分组
//func CommonAddGroup(group *model.Group) error {
//	// 在数据库中创建组
//	err := isql.Group.Add(group)
//	if err != nil {
//		return err
//	}
//
//	// 默认创建分组之后，需要将admin添加到分组中
//	adminInfo := new(model.User)
//	err = isql.User.Find(tools.H{"id": 1}, adminInfo)
//	if err != nil {
//		return err
//	}
//
//	err = isql.Group.AddUserToGroup(group, []model.User{*adminInfo})
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

//// CommonUpdateGroup 标准更新分组
//func CommonUpdateGroup(oldGroup, newGroup *model.Group) error {
//	err := isql.Group.Update(newGroup)
//	if err != nil {
//		return err
//	}
//	return nil
//}

// CommonAddRecord 标准创建解析
func CommonAddRecord(record *model.Record) error {
	// 解析信息的预置处理

	// 先将用户添加到MySQL
	err := isql.Record.Add(record)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("向MySQL创建解析失败：" + err.Error()))
	}
	return err
}

//
//// CommonUpdateUser 标准更新用户
//func CommonUpdateUser(oldUser, newUser *model.User, groupId []uint) error {
//	err := isql.User.Update(newUser)
//	if err != nil {
//		return tools.NewMySqlError(fmt.Errorf("在MySQL更新用户失败：" + err.Error()))
//	}
//
//	//判断部门信息是否有变化有变化则更新相应的数据库
//	oldDeptIds := tools.StringToSlice(oldUser.DepartmentId, ",")
//	addDeptIds, removeDeptIds := tools.ArrUintCmp(oldDeptIds, groupId)
//
//	// 先处理添加的部门
//	addgroups, err := isql.Group.GetGroupByIds(addDeptIds)
//	if err != nil {
//		return tools.NewMySqlError(fmt.Errorf("根据部门ID获取部门信息失败" + err.Error()))
//	}
//	for _, group := range addgroups {
//		// 先将用户和部门信息维护到MySQL
//		err := isql.Group.AddUserToGroup(group, []model.User{*newUser})
//		if err != nil {
//			return tools.NewMySqlError(fmt.Errorf("向MySQL添加用户到分组关系失败：" + err.Error()))
//		}
//	}
//
//	// 再处理删除的部门
//	removegroups, err := isql.Group.GetGroupByIds(removeDeptIds)
//	if err != nil {
//		return tools.NewMySqlError(fmt.Errorf("根据部门ID获取部门信息失败" + err.Error()))
//	}
//	for _, group := range removegroups {
//		err := isql.Group.RemoveUserFromGroup(group, []model.User{*newUser})
//		if err != nil {
//			return tools.NewMySqlError(fmt.Errorf("在MySQL将用户从分组移除失败：" + err.Error()))
//		}
//	}
//	return nil
//}
