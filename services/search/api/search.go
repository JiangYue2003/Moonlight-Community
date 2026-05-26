// search-api：知识帖搜索 HTTP 入口（function_score + multi_match + Suggest）。
package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	"github.com/zhiguang/zhiguang-go/services/search/api/internal/config"
	"github.com/zhiguang/zhiguang-go/services/search/api/internal/handler"
	"github.com/zhiguang/zhiguang-go/services/search/api/internal/svc"
)

var configFile = flag.String("f", "etc/search-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("search-api listening at %s:%d\n", c.Host, c.Port)
	server.Start()
}
