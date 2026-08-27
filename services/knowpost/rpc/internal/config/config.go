package config

import (
	"time"

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
	Feed  FeedConf

	UserCounterRpc zrpc.RpcClientConf
	CounterRpc     zrpc.RpcClientConf
	RelationRpc    zrpc.RpcClientConf // 关系服务

	L1 L1Conf
	// HotKeyDetector 与 Java cache.hotkey.* 参数对齐
	HotKey HotKeyConf `json:",optional"`

	Snowflake SnowflakeConf
}

type MysqlConf struct {
	DataSource string
}

type KafkaConf struct {
	Brokers              []string
	CounterEventsTopic   string `json:",default=counter-events"`
	GroupId              string `json:",default=knowpost-cache-invalidation"`
	CanalOutboxTopic     string `json:",default=canal-outbox"`
	RelationEpochGroupId string `json:",default=knowpost-feed-relation-epoch"`
	SafetyEpochGroupId   string `json:",default=knowpost-feed-content-safety-epoch"`
}

type FeedConf struct {
	Strategy         string `json:",default=hybrid"`
	AuthorTier       FeedAuthorTierConf
	Observability    FeedObservabilityConf
	Epoch            FeedEpochConf
	RouteSnapshot    FeedRouteSnapshotConf
	CombinedPipeline FeedCombinedPipelineConf
	CursorPagination FeedCursorPaginationConf
	PageCache        FeedPageCacheConf
}

type FeedAuthorTierConf struct {
	Mode string `json:",default=off"`
}

type FeedObservabilityConf struct {
	Enabled bool `json:",default=false"`
}

type FeedEpochConf struct {
	KeyPrefix               string        `json:",default=feed"`
	RelationL1TTL           time.Duration `json:",default=1s"`
	SafetyL1TTL             time.Duration `json:",default=1s"`
	RelationConsumerEnabled bool          `json:",default=false"`
	SafetyConsumerEnabled   bool          `json:",default=false"`
}

type FeedRouteSnapshotConf struct {
	Enabled bool          `json:",default=false"`
	TTL     time.Duration `json:",default=5s"`
}

type FeedCombinedPipelineConf struct {
	Enabled   bool `json:",default=false"`
	BatchSize int  `json:",default=128"`
}

type FeedCursorPaginationConf struct {
	Enabled bool `json:",default=false"`
}

type FeedPageCacheConf struct {
	Mode           string        `json:",default=off"`
	KeyPrefix      string        `json:",default=feed"`
	Page           int32         `json:",default=1"`
	Size           int32         `json:",default=20"`
	L1FreshTTL     time.Duration `json:",default=800ms"`
	L2FreshTTL     time.Duration `json:",default=4s"`
	StaleTTL       time.Duration `json:",default=10s"`
	JitterPercent  int           `json:",default=20"`
	RefreshWorkers int           `json:",default=32"`
	RefreshQueue   int           `json:",default=1024"`
	LoaderTimeout  time.Duration `json:",default=2s"`
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
	FeedEpochNumCounters  int64 `json:",default=200000"`
	FeedEpochMaxCostMB    int64 `json:",default=4"`
	FeedRouteNumCounters  int64 `json:",default=100000"`
	FeedRouteMaxCostMB    int64 `json:",default=32"`
	FeedPageNumCounters   int64 `json:",default=100000"`
	FeedPageMaxCostMB     int64 `json:",default=128"`
}

type HotKeyConf struct {
	WindowSeconds       int   `json:",default=60"`
	SegmentSeconds      int   `json:",default=10"`
	LevelLow            int64 `json:",default=50"`
	LevelMedium         int64 `json:",default=200"`
	LevelHigh           int64 `json:",default=500"`
	ExtendLowSeconds    int   `json:",default=20"`
	ExtendMediumSeconds int   `json:",default=60"`
	ExtendHighSeconds   int   `json:",default=120"`
}

type SnowflakeConf struct {
	WorkerId     int64 `json:",default=1"`
	DatacenterId int64 `json:",default=2"`
}
