package bootstrap

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhiguang/zhiguang-go/services/counter/internal/application"
)

// Config counter-rpc 配置；Redis 用于位图与 SDS，Kafka 用于事件发布。
type Config struct {
	zrpc.RpcServerConf
	Kafka   application.KafkaConfig
	Rebuild application.RebuildConfig `json:",optional"`
}
