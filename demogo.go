package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"demoGo/internal/config"
	"demoGo/internal/handler"
	"demoGo/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/demogo-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 初始化日志配置
	logx.MustSetup(c.Log)

	// 关闭stat统计日志
	logx.DisableStat()

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
