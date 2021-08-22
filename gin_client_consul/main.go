package main

import (
	_ "gin_client/app"
	"gin_client/routes"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/web"
)

func main() {
	// 初始化路由
	gin := routes.IniterRouter()
	
	// 创建 consul 服务
	consulReg := initConsul()
	
	// 创建micro服务
	microUserServer := web.NewService(
		web.Name("micro.user.web.server"),
		web.Handler(gin),
		
		// 发现已注册到 consul 的 micro 服务
		web.Registry(consulReg),
	)
	// 初始化micro服务
	microUserServer.Init()
	
	gin.Run() // gin.Run(":8100") // 监听并在 0.0.0.0:8100 上启动服务
}

// 创建 consul 服务
func initConsul() registry.Registry {
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.0.102:8500"),
	)
	return consulReg
}