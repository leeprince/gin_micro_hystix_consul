package v1

import (
	"encoding/json"
	"fmt"
	"gin_client/rpc"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/8/21 下午2:55
 * @Desc:	micro wrap 中间件 + hystrix 熔断 + redis 降级
 */

func GetUserWrapHystrixRedis(c *gin.Context) {
	fmt.Println("GetUsers - controller")
	userId := cast.ToInt64(c.Query("userId"))
	
	userList, err := rpc.GetUsersWrap(userId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	fmt.Println("userList:", userList)
	fmt.Println("userList.code:", userList.Code)
	
	data := make(map[string]interface{})
	json.Unmarshal(userList.Data, &data)
	
	fmt.Println("data:", data)
	fmt.Println("data.Username:", data["Username"])
	
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "successfully.",
		"data":    data,
	})
}

