package routes

import (
	"github.com/gin-gonic/gin"
	"gin_client/middleware"
)

// 约定自动注册路由的类型
type Router func(*gin.Engine)

// 自动注册路由的切片
var routers = []Router{}

// 注册路由
func RegisterRouter(router ...Router)  {
	routers = append(routers, router...)
}

// 初始化路由
func IniterRouter() *gin.Engine  {
	// 默认启动方式，包含 Logger、Recovery 中间件
	// r := gin.New()
	r := gin.Default()
	
	// 注册中间件
	middleware.InitMiddeware(r)
	
	// 初始化路由
	for _, router := range routers {
		router(r)
	}
	return r
}