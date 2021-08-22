package main

import (
	"github.com/asim/go-micro/v3"
	pb "github.com/leeprince/protobuf/grpc/gin_micro"
	"micro_server/services"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/8/13 下午10:04
 * @Desc:
 */

func main() {
	// 创建服务
	service := micro.NewService(
		micro.Name("micro.user.server"),
		micro.Address(":8081"),
		// micro
	)
	// 初始化服务
	service.Init()
	
	// 注册服务
	pb.RegisterUserServiceHandler(service.Server(), new(services.UserService))
	
	service.Run()
}
