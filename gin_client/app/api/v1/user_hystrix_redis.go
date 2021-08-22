package v1

import (
	"fmt"
	"gin_client/rpc"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/leeprince/protobuf/grpc/gin_micro"
	"github.com/spf13/cast"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/8/21 下午12:10
 * @Desc:	hystrix实现熔断 + redis实现降级
 */

func GetUserHystrixRedis(c *gin.Context) {
	fmt.Println("GetUserHystrixRedis - controller")
	userId := cast.ToInt64(c.Query("userId"))
	
	var (
		userList *gin_micro.GetUsersRsp
		err      error
	)
	data := make(map[string]interface{})
	
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
	hystrix.ConfigureCommand("userHy", configHy)
	// 3. 使用 hystrix
	err = hystrix.Do("userHy", func() error {
		userList, err = rpc.GetUsers(userId)
		return err
	}, func(err error) error {
		// 返回的错误是: hystrix.Do 的 runFunc 返回的错误
		// 熔断后的处理
		// 降级方法特点，切记不建议执行复杂的业务；尽量简单，不要存在异常
		// 文件 ， 缓存 => 提前准备
		fmt.Println("降级处理>")
		data, err = demotion(c, userId)
		fmt.Println("降级处理后返回>", data)
		return err
	})
	
	// err 有两种来源：1. rpc 调用出错；2. 熔断返回报错
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "err: " + err.Error(),
			"data":    nil,
		})
		return
	}
	// 降级返回
	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "demetion successful",
			"data":    data,
		})
		return
	}
	message := userList.GetMessage()
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": message,
		"data":    data,
	})
}
