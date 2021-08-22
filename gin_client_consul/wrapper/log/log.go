package log

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3/client"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/8/21 下午1:21
 * @Desc:	micro 中间件
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
	fmt.Println("====>>> wrapper log ===>> start")
	// 具体的业务代码
	err := c.Client.Call(ctx, req, rsp, opts...)

	fmt.Println("====>>> wrapper log ===>> end")
	return err
}
