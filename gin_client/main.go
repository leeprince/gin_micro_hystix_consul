package main

import (
	_ "gin_client/app"
	"gin_client/routes"
	"github.com/asim/go-micro/v3/web"
)

func main() {
	gin := routes.IniterRouter()
	
	// 创建micro服务
	microUserServer := web.NewService(
		web.Name("micro.user.web.server"),
		web.Handler(gin),
	)
	// 初始化micro服务
	microUserServer.Init()
	
	gin.Run() // gin.Run(":8100") // 监听并在 0.0.0.0:8100 上启动服务
}
