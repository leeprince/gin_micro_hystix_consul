package api

import (
	v1 "gin_client/app/api/v1"
	"gin_client/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	routes.RegisterRouter(Routes)
}

func Routes(g *gin.Engine) {
	// 健康检查
	g.GET("/ping", Ping)
	
	hello := g.Group("v1")
	{
		hello.GET("/getUser", v1.GetUser)
		hello.GET("/getUser/hystrix", v1.GetUserHystrix)
		hello.GET("/getUser/redis", v1.GetUserRedis)
		hello.GET("/getUser/hystrix/redis", v1.GetUserHystrixRedis)
		hello.GET("/getUser/wrap", v1.GetUserWrap)
		hello.GET("/getUser/wrap/hystrix/redis", v1.GetUserWrapHystrixRedis)
	}
}
