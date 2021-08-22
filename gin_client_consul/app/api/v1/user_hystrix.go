package v1

import (
	"encoding/json"
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
 * @Date:   2021/8/13 下午11:41
 * @Desc:	熔断
 */

func GetUserHystrix(c *gin.Context)  {
	fmt.Println("GetUserHystrix - controller")
	userId := cast.ToInt64(c.Query("userId"))
	
	var (
		userList *gin_micro.GetUsersRsp
		err error
	)
	
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
		// 熔断后的处理
		fmt.Println("降级处理>")
		return err
	})
	
	// err 有两种来源：1. rpc 调用出错；2. 熔断返回报错
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"message":err.Error(),
				"data":nil,
			})
		return
	}
	fmt.Println("userList:", userList)
	fmt.Println("userList.code:", userList.Code)
	
	data := make(map[string]interface{})
	json.Unmarshal(userList.Data, &data)
	
	fmt.Println("data:", data)
	fmt.Println("data.Username:", data["Username"])

	message := userList.GetMessage()
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message":message,
		"data": data,
	})
}