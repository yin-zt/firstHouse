package routes

import (
	"firstHouse/controller"
	"github.com/gin-gonic/gin"
)

// 注册用户路由
func InitRecordRoutes(r *gin.RouterGroup) gin.IRoutes {
	record := r.Group("/record")
	// 开启jwt认证中间件
	//user.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	//user.Use(middleware.CasbinMiddleware())
	{
		record.POST("/add", controller.Record.Add)       // 增加解析
		record.POST("/delete", controller.Record.Delete) // 删除解析
		record.POST("/modify", controller.Record.Add)    // 修改解析(暂不提供)
		record.GET("/list", controller.Record.List)      // 查询解析
		//record.POST("/count", controller.Record.Delete)   // 统计解析
	}
	return r
}
