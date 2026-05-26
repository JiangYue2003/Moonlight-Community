package svc

import (
	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhiguang/zhiguang-go/pkg/cachex"
	"github.com/zhiguang/zhiguang-go/pkg/hotkey"
	"github.com/zhiguang/zhiguang-go/pkg/snowflakex"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache/detail"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache/mine"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/config"
	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	model "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
	outboxmodel "github.com/zhiguang/zhiguang-go/services/relation/shared/model"
)

type ServiceContext struct {
	Config config.Config

	Db             sqlx.SqlConn
	KnowPostsModel model.KnowPostsModel
	OutboxModel    outboxmodel.OutboxModel

	Redis goredis.UniversalClient

	UserCounterRpc counterpb.UserCounterClient
	CounterRpc     counterpb.CounterClient

	DetailCache   cachex.Cache[*pb.KnowPostDetail]
	FeedMineCache cachex.Cache[*pb.FeedPage]

	// FeedItem 与 FeedPublic ids 路径较特殊，logic 内直接操作 L1/L2。
	L1FeedPublic *cachex.L1
	L1FeedItem   *cachex.L1
	L2           *cachex.L2

	HotDetail     *hotkey.Detector
	HotFeedPublic *hotkey.Detector
	HotFeedItem   *hotkey.Detector
	HotFeedMine   *hotkey.Detector

	Snowflake *snowflakex.Generator
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{
		Addrs:    []string{c.Redis.Host},
		Password: c.Redis.Pass,
	})

	mb := int64(1) << 20
	l1Detail := mustL1(c.L1.DetailNumCounters, c.L1.DetailMaxCostMB*mb)
	l1FeedPublic := mustL1(c.L1.FeedPublicNumCounters, c.L1.FeedPublicMaxCostMB*mb)
	l1FeedItem := mustL1(c.L1.FeedItemNumCounters, c.L1.FeedItemMaxCostMB*mb)
	l1FeedMine := mustL1(c.L1.FeedMineNumCounters, c.L1.FeedMineMaxCostMB*mb)
	l2 := cachex.NewL2(rdb)

	hotCfg := hotkey.Config{
		WindowSeconds:  c.HotKey.WindowSeconds,
		SegmentSeconds: c.HotKey.SegmentSeconds,
		LevelLow:       c.HotKey.LevelLow,
		LevelMedium:    c.HotKey.LevelMedium,
		LevelHigh:      c.HotKey.LevelHigh,
	}
	hotkey.ConfigureExtensions(c.HotKey.ExtendLowSeconds, c.HotKey.ExtendMediumSeconds, c.HotKey.ExtendHighSeconds)
	hotDetail := hotkey.New(hotCfg)
	hotFeedPublic := hotkey.New(hotCfg)
	hotFeedItem := hotkey.New(hotCfg)
	hotFeedMine := hotkey.New(hotCfg)

	return &ServiceContext{
		Config:         c,
		Db:             conn,
		KnowPostsModel: model.NewKnowPostsModel(conn, c.CacheRedis),
		OutboxModel:    outboxmodel.NewOutboxModel(conn, c.CacheRedis),
		Redis:          rdb,
		UserCounterRpc: counterpb.NewUserCounterClient(zrpc.MustNewClient(c.UserCounterRpc).Conn()),
		CounterRpc:     counterpb.NewCounterClient(zrpc.MustNewClient(c.CounterRpc).Conn()),

		DetailCache:   detail.New(l1Detail, l2, hotDetail),
		FeedMineCache: mine.New(l1FeedMine, l2, hotFeedMine),
		L1FeedPublic:  l1FeedPublic,
		L1FeedItem:    l1FeedItem,
		L2:            l2,

		HotDetail:     hotDetail,
		HotFeedPublic: hotFeedPublic,
		HotFeedItem:   hotFeedItem,
		HotFeedMine:   hotFeedMine,

		Snowflake: snowflakex.MustNew(c.Snowflake.DatacenterId, c.Snowflake.WorkerId),
	}
}

func mustL1(numCounters, maxCost int64) *cachex.L1 {
	l1, err := cachex.NewL1(cachex.L1Config{NumCounters: numCounters, MaxCost: maxCost})
	if err != nil {
		panic(err)
	}
	return l1
}
