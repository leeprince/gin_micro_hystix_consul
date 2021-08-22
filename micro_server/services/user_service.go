package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/leeprince/protobuf/grpc/gin_micro"
	"github.com/spf13/cast"
	"micro_server/dao"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/8/13 下午11:41
 * @Desc:
 */

type UserService struct {
	
}

func (s *UserService) GetUsers(ctx context.Context, req *gin_micro.GetUsersReq, rsp *gin_micro.GetUsersRsp) error  {
	fmt.Println(">>>>>>>=====>>>>>>>>>>:8081-----------start")
	
	// 开始延迟, 并在rpc客户端测试熔断
	fmt.Println(">>>sleep=========测试熔断----start")
	time.Sleep(time.Second * 3)
	fmt.Println(">>>sleep=========测试熔断----end")
	
	userId := req.UserId
	
	userDao := dao.NewUserDao()
	user, err := userDao.GetUserByUserId(userId)
	if err != nil {
		return err
	}
	jsonBytes,_ := json.Marshal(user)
	
	// nTime := time.Now().Format("2006-01-02 15:04:05") // Y-m-d H:i:s 的固定格式"2006-01-02 15:04:05"
	nTime := time.Now().UnixNano()
	rsp.Code = 0
	rsp.Message = "micro successful."+cast.ToString(nTime)
	rsp.Data = jsonBytes
	
	fmt.Println("GetUsers.rsp:", rsp)
	fmt.Println(">>>>>>>=====>>>>>>>>>>:8081-----------end")
	return nil
}


 