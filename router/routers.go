package router

import (
	Controllers "dqh-test/app/controllers/api"
	"dqh-test/libs/system"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	sysBaseRouter(e)

	businessRouter(e)
}

//系统路由
func sysBaseRouter(e *gin.Engine) {
	e.GET("/sysinfo", system.SysInfo)
	e.GET("/ping", system.Ping)
}

//业务路由
func businessRouter(e *gin.Engine) {
	//API
	api := e.Group("/api")
	api.POST("/member/login", Controllers.Login)
	api.POST("/member/regist", Controllers.Regist)
}
