package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

// Config knowpost-rpc 配置：MySQL（写帖子表）+ Redis（缓存）+ Kafka（监听 counter-events）
// + 下游 usercounter-rpc / counter-rpc + 三个 ristretto L1 缓存参数 + 雪花 ID。
type Config struct {
	zrpc.RpcServerConf

	Mysql      MysqlConf
	CacheRedis cache.CacheConf

	Kafka KafkaConf

	UserCounterRpc zrpc.RpcClientConf
	CounterRpc     zrpc.RpcClientConf

	L1 L1Conf
	// HotKeyDetector 与 Java cache.hotkey.* 参数对齐
	HotKey HotKeyConf `json:",optional"`

	Snowflake SnowflakeConf
}

type MysqlConf struct {
	DataSource string
}

type KafkaConf struct {
	Brokers            []string
	CounterEventsTopic string `json:",default=counter-events"`
	GroupId            string `json:",default=knowpost-cache-invalidation"`
}

type L1Conf struct {
	DetailNumCounters     int64 `json:",default=50000"`
	DetailMaxCostMB       int64 `json:",default=100"`
	FeedPublicNumCounters int64 `json:",default=10000"`
	FeedPublicMaxCostMB   int64 `json:",default=50"`
	FeedItemNumCounters   int64 `json:",default=50000"`
	FeedItemMaxCostMB     int64 `json:",default=50"`
	FeedMineNumCounters   int64 `json:",default=10000"`
	FeedMineMaxCostMB     int64 `json:",default=50"`
}

type HotKeyConf struct {
	WindowSeconds      int   `json:",default=60"`
	SegmentSeconds     int   `json:",default=10"`
	LevelLow           int64 `json:",default=50"`
	LevelMedium        int64 `json:",default=200"`
	LevelHigh          int64 `json:",default=500"`
	ExtendLowSeconds   int   `json:",default=20"`
	ExtendMediumSeconds int  `json:",default=60"`
	ExtendHighSeconds  int   `json:",default=120"`
}

type SnowflakeConf struct {
	WorkerId     int64 `json:",default=1"`
	DatacenterId int64 `json:",default=2"`
}
