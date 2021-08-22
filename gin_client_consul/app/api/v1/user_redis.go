package v1

import (
	"encoding/json"
	"fmt"
	"gin_client/consts"
	"gin_client/rpc"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/8/19 下午10:37
 * @Desc:	redis缓存用户数据
 */

func GetUserRedis(c *gin.Context) {
	fmt.Println("GetUserRedis - controller")
	userId := cast.ToInt64(c.Query("userId"))
	data := make(map[string]interface{})
	
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	userKey := fmt.Sprintf(consts.USER_KEY, userId)
	val, err := rdb.Get(c, userKey).Bytes()
	fmt.Println("redis 返回：", string(val), err)
	if err != nil && err != redis.Nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -3,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	if string(val) != "" {
		json.Unmarshal(val, &data)
		
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "redis successfully.",
			"data":    data,
		})
		return
	}
	
	userList, err := rpc.GetUsers(userId)
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
	
	json.Unmarshal(userList.Data, &data)
	
	expireTime := consts.USER_EXPIE
	err = rdb.Set(c, userKey, userList.Data, expireTime).Err()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -2,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	
	fmt.Println("data:", data)
	fmt.Println("data.Username:", data["Username"])
	
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "successfully.",
		"data":    data,
	})
}

func demotion(c *gin.Context, userId int64) (map[string]interface{}, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	userKey := fmt.Sprintf(consts.USER_KEY, userId)
	val, err := rdb.Get(c, userKey).Bytes()
	fmt.Println("redis 返回：", string(val), err)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	
	data := make(map[string]interface{})
	if string(val) == "" {
		return nil, nil
	}
	json.Unmarshal(val, &data)
	return data, nil
}
