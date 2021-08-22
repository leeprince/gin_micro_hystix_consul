package hystrix

import (
	"context"
	"fmt"
	"gin_client/consts"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/v3/client"
	"github.com/go-redis/redis/v8"
	"github.com/leeprince/protobuf/grpc/gin_micro"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/8/21 下午1:21
 * @Desc:	micro wrap 中间件 + hystrix 熔断 + redis 降级
 */

type clientWrapper struct {
	client.Client
}

// 创建中间件对象
func NewClientWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
}

// 中间件的执行方法
func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Println("====>>> micro wrap 中间件 + hystrix 熔断 + redis 降级 ===>> start")
	
	// 目标：执行时间超过1秒就执行熔断
	// 1. 定义 hystrix
	configHy := hystrix.CommandConfig{
		Timeout:                1000, // 单位毫秒
		MaxConcurrentRequests:  0,
		RequestVolumeThreshold: 0,
		SleepWindow:            0,
		ErrorPercentThreshold:  0,
	}
	// 2. 配置 hystrix
	hystrix.ConfigureCommand(req.Service(), configHy)
	// 3. 使用 hystrix
	err := hystrix.Do(req.Service(), func() error {
		// 具体的业务代码
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(err error) error {
		// 返回的错误是: hystrix.Do 的 runFunc 返回的错误
		// 熔断后的处理
		// 降级方法特点，切记不建议执行复杂的业务；尽量简单，不要存在异常
		// 文件 ， 缓存 => 提前准备
		fmt.Println("降级处理>")
		userReq := req.Body().(*gin_micro.GetUsersReq)
		userId := userReq.UserId
		fmt.Println("降级处理>userId:", userId)
		err = demotion(userId, rsp)
		fmt.Println("降级处理后返回>", rsp)
		return err
	})
	
	fmt.Println("====>>> micro wrap 中间件 + hystrix 熔断 + redis 降级 ===>> end")
	return err
}


func demotion(userId int64, rsp interface{}) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	userKey := fmt.Sprintf(consts.USER_KEY, userId)
	val, err := rdb.Get(context.Background(), userKey).Bytes()
	fmt.Println("redis 返回：", string(val), err)
	if err != nil && err != redis.Nil {
		return err
	}
	
	data := make(map[string]interface{})
	if string(val) == "" {
		return nil
	}
	fmt.Println("demotion>data:", data)
	
	resp := rsp.(*gin_micro.GetUsersRsp)
	resp.Code = 0
	resp.Message = "micro hystrix redis demotion return successful."
	resp.Data = val
	
	return nil
}
