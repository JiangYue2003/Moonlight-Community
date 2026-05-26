// llm-api：DeepSeek 描述 + 阿里通义嵌入 + ES knn 流式问答（SSE）。
package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	"github.com/zhiguang/zhiguang-go/services/llm/api/internal/config"
	"github.com/zhiguang/zhiguang-go/services/llm/api/internal/handler"
	"github.com/zhiguang/zhiguang-go/services/llm/api/internal/svc"
)

var configFile = flag.String("f", "etc/llm-api.yaml", "config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("llm-api listening at %s:%d\n", c.Host, c.Port)
	server.Start()
}
