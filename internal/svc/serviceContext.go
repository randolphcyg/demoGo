package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"demoGo/internal/config"
	"demoGo/javastub/democlient"
)

type ServiceContext struct {
	Config config.Config
	// rpc客户端
	DemoRpcClient democlient.DemoClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//	rpc客户端
		DemoRpcClient: democlient.NewDemoClient(zrpc.MustNewClient(c.DemoRpcConf)),
	}
}
