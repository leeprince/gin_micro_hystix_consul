package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	pb "github.com/leeprince/protobuf/grpc/gin_micro"
	"micro_server/services"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/8/13 下午10:04
 * @Desc:	micro 服务，并注册到 consul 中
 */

func main() {
	// 创建 consul 服务
	consulReg := initConsul()
	
	// 创建服务
	service := micro.NewService(
		micro.Name("micro.user.server"),
		micro.Address(":8082"),
		
		// 注册 micro 服务到 consul
		micro.Registry(consulReg),
	)
	// 初始化服务
	service.Init()
	
	// 注册服务
	pb.RegisterUserServiceHandler(service.Server(), new(services.UserService))
	
	service.Run()
}

// 创建 consul 服务
func initConsul() registry.Registry {
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.0.102:8500"),
	)
	return consulReg
}
