package controller

import (
	"firstHouse/logic"
	"firstHouse/model/request"

	"github.com/gin-gonic/gin"
)

type RecordController struct{}

// Add 添加记录
func (m *RecordController) Add(c *gin.Context) {
	req := new(request.RecordAddReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Record.Add(c, req)
	})
}

//
//// Update 更新记录
//func (m *UserController) Update(c *gin.Context) {
//	req := new(request.UserUpdateReq)
//	Run(c, req, func() (interface{}, interface{}) {
//		return logic.User.Update(c, req)
//	})
//}

// List 解析记录列表
func (m *RecordController) List(c *gin.Context) {
	req := new(request.RecordListReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Record.List(c, req)
	})
}

// Delete 删除记录
func (m RecordController) Delete(c *gin.Context) {
	req := new(request.RecordDeleteReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Record.Delete(c, req)
	})
}

//
//// ChangePwd 更新密码
//func (m UserController) ChangePwd(c *gin.Context) {
//	req := new(request.UserChangePwdReq)
//	Run(c, req, func() (interface{}, interface{}) {
//		return logic.User.ChangePwd(c, req)
//	})
//}
//
//// ChangeUserStatus 更改用户状态
//func (m UserController) ChangeUserStatus(c *gin.Context) {
//	req := new(request.UserChangeUserStatusReq)
//	Run(c, req, func() (interface{}, interface{}) {
//		return logic.User.ChangeUserStatus(c, req)
//	})
//}
//
//// GetUserInfo 获取当前登录用户信息
//func (uc UserController) GetUserInfo(c *gin.Context) {
//	req := new(request.UserGetUserInfoReq)
//	Run(c, req, func() (interface{}, interface{}) {
//		return logic.User.GetUserInfo(c, req)
//	})
//}
