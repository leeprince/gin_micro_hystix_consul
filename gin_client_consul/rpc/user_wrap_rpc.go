package rpc

import (
	"context"
	"fmt"
	"gin_client/wrapper/log"
	"github.com/asim/go-micro/v3"
	"github.com/leeprince/protobuf/grpc/gin_micro"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/8/15 下午12:21
 * @Desc:	经过 micro wrap 中间件，并调用 micro 服务端方法
 */

var (
	microUserClient gin_micro.UserService
)

// 注意：wrapper 中间件是 micro 的，不是 gin 的，只发生在调用 micro 服务时执行
func init() {
	microServer := micro.NewService(
		micro.Name("micro.user.wrap.server"),
		// micro.WrapClient(): 用于用一些中间件组件包装Client,包装器以相反的顺序应用，因此最后一个先执
		micro.WrapClient(
			// 测试 consul 建议把熔断暂时关闭，避免降级之后查看不到请求具体的是哪个 rpc 服务
			// hystrix.NewClientWrapper(),
			log.NewClientWrapper(),
		),
	)
	// 创建新的客户端
	microUserClient = gin_micro.NewUserService("micro.user.server", microServer.Client())
}

func GetUsersWrap(userId int64) (*gin_micro.GetUsersRsp, error) {
	fmt.Println("> 准备调用 micro 服务端方法")
	
	// 调用rpc方法
	rsp, err := microUserClient.GetUsers(context.Background(), &gin_micro.GetUsersReq{
		UserId: userId,
	})
	fmt.Println("micro 服务端返回结果：", rsp)
	if err != nil {
		return nil, err
	}
	return rsp, err
}
