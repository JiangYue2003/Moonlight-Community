package svc

import (
	"context"
	"fmt"
	"strings"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/queue"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhiguang/zhiguang-go/pkg/cachex"
	"github.com/zhiguang/zhiguang-go/pkg/hotkey"
	"github.com/zhiguang/zhiguang-go/pkg/snowflakex"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache/detail"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache/mine"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache/userfeed"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/config"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/feed"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/feedepoch"
	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	model "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
	relationpb "github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
	relationmodel "github.com/zhiguang/zhiguang-go/services/relation/shared/model"
)

// feedRouteCacheTTL 限制作者跨越大V阈值后的路由最终一致性窗口。
// 在引入粉丝数变更事件驱动的精确失效前，不应放大该窗口。
const (
	feedRouteCacheTTL            = 5 * time.Second
	maxCombinedPipelineBatchSize = 128
)

type ServiceContext struct {
	Config config.Config

	Db             sqlx.SqlConn
	KnowPostsModel model.KnowPostsModel
	FeedPostLoader model.FeedPostLoader
	OutboxModel    relationmodel.OutboxModel

	Redis goredis.UniversalClient

	UserCounterRpc counterpb.UserCounterClient
	CounterRpc     counterpb.CounterClient
	RelationRpc    relationpb.RelationClient

	DetailCache   cachex.Cache[*pb.KnowPostDetail]
	FeedMineCache cachex.Cache[*pb.FeedPage]

	// FeedItem 与 FeedPublic ids 路径较特殊，logic 内直接操作 L1/L2。
	L1FeedPublic *cachex.L1
	L1FeedItem   *cachex.L1
	L1FeedEpoch  *cachex.L1
	L1FeedRoute  *cachex.L1
	L1FeedPage   *cachex.L1
	L2           *cachex.L2

	HotDetail     *hotkey.Detector
	HotFeedPublic *hotkey.Detector
	HotFeedItem   *hotkey.Detector
	HotFeedMine   *hotkey.Detector

	Snowflake *snowflakex.Generator

	// Feed推拉架构
	FeedWriter    *feed.FeedWriter
	FeedReader    *feed.FeedReader
	FeedObserver  feed.FeedObserver
	FeedEpochs    feedepoch.Store
	FeedPageCache userfeed.Cache

	feedPusher       *kq.Pusher
	feedFanoutWorker queue.MessageQueue
}

func NewServiceContext(c config.Config) *ServiceContext {
	applyFeedEpochDefaults(&c.Feed.Epoch)
	if err := applyFeedRouteSnapshotDefaults(&c.Feed.RouteSnapshot); err != nil {
		panic(err)
	}
	if err := applyFeedCombinedPipelineDefaults(&c.Feed.CombinedPipeline); err != nil {
		panic(err)
	}
	if err := applyFeedPageCacheDefaults(&c.Feed.PageCache); err != nil {
		panic(err)
	}
	if err := validateFeedSafetyConfiguration(c.Feed); err != nil {
		panic(err)
	}
	feedStrategy, err := resolveFeedStrategy(c.Feed)
	if err != nil {
		panic(err)
	}
	authorTierMode, err := resolveAuthorTierMode(c.Feed)
	if err != nil {
		panic(err)
	}

	conn := sqlx.NewMysql(c.Mysql.DataSource)
	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{
		Addrs:    []string{c.Redis.Host},
		Password: c.Redis.Pass,
	})
	feedObserver := feed.NewMetricsObserver(c.Feed.Observability.Enabled)

	mb := int64(1) << 20
	l1Detail := mustL1(c.L1.DetailNumCounters, c.L1.DetailMaxCostMB*mb)
	l1FeedPublic := mustL1(c.L1.FeedPublicNumCounters, c.L1.FeedPublicMaxCostMB*mb)
	l1FeedItem := mustL1(c.L1.FeedItemNumCounters, c.L1.FeedItemMaxCostMB*mb)
	l1FeedMine := mustL1(c.L1.FeedMineNumCounters, c.L1.FeedMineMaxCostMB*mb)
	l1FeedEpoch := mustL1(c.L1.FeedEpochNumCounters, c.L1.FeedEpochMaxCostMB*mb)
	l1FeedRoute := mustL1(c.L1.FeedRouteNumCounters, c.L1.FeedRouteMaxCostMB*mb)
	var l1FeedPage *cachex.L1
	if userfeed.Mode(c.Feed.PageCache.Mode) == userfeed.ModeL1L2 {
		l1FeedPage = mustL1(c.L1.FeedPageNumCounters, c.L1.FeedPageMaxCostMB*mb)
	}
	l2 := cachex.NewL2(rdb)
	feedPageCache, err := userfeed.New(userfeed.Config{
		Mode:           userfeed.Mode(c.Feed.PageCache.Mode),
		KeyPrefix:      c.Feed.PageCache.KeyPrefix,
		L1FreshTTL:     c.Feed.PageCache.L1FreshTTL,
		L2FreshTTL:     c.Feed.PageCache.L2FreshTTL,
		StaleTTL:       c.Feed.PageCache.StaleTTL,
		JitterPercent:  c.Feed.PageCache.JitterPercent,
		RefreshWorkers: c.Feed.PageCache.RefreshWorkers,
		RefreshQueue:   c.Feed.PageCache.RefreshQueue,
		LoaderTimeout:  c.Feed.PageCache.LoaderTimeout,
		Observe: func(operation userfeed.Operation, err error) {
			dependency, feedOperation := pageCacheDependencyOperation(operation)
			feed.RecordDependency(feedObserver, dependency, feedOperation, feed.OutcomeFromError(err))
		},
		ObserveRefreshState: func(state userfeed.RefreshState) {
			feed.RecordPageRefreshState(feedObserver, state.QueueDepth, state.ActiveWorkers, state.PendingKeys)
		},
		ReportError: func(operation userfeed.Operation, err error) {
			logx.Errorf("Feed PageCache %s failure: %v", operation, err)
		},
	}, l1FeedPage, l2)
	if err != nil {
		panic(err)
	}
	feedEpochs, err := feedepoch.NewStore(rdb, l1FeedEpoch, feedepoch.Config{
		KeyPrefix:     c.Feed.Epoch.KeyPrefix,
		RelationL1TTL: c.Feed.Epoch.RelationL1TTL,
		SafetyL1TTL:   c.Feed.Epoch.SafetyL1TTL,
	})
	if err != nil {
		panic(err)
	}

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

	userCounterClient := counterpb.NewUserCounterClient(zrpc.MustNewClient(c.UserCounterRpc).Conn())
	counterClient := counterpb.NewCounterClient(zrpc.MustNewClient(c.CounterRpc).Conn())
	relationClient := relationpb.NewRelationClient(zrpc.MustNewClient(c.RelationRpc).Conn())
	redisAdapter := feed.NewRedisAdapter(rdb)
	relationAdapter := feed.NewRelationClientAdapter(relationClient)
	feedCounterAdapter := feed.NewCounterClientAdapter(userCounterClient)
	authorTierResolver := feed.NewAuthorTierResolver(
		model.NewFeedAuthorDeliveryTierModel(conn),
		feedCounterAdapter,
		relationmodel.NewActiveFollowerCountModel(conn),
		feedObserver,
	)
	feedPusher := kq.NewPusher(
		c.Kafka.Brokers,
		feed.FEED_FANOUT_TOPIC,
		kq.WithAllowAutoTopicCreation(),
	)
	feedFanoutWorker := feed.NewFeedFanoutWorker(
		c.Kafka.Brokers,
		redisAdapter,
		relationAdapter,
		logx.WithContext(context.Background()),
	)
	knowPostsModel := model.NewKnowPostsModel(conn, c.CacheRedis)
	return &ServiceContext{
		Config:         c,
		Db:             conn,
		KnowPostsModel: knowPostsModel,
		FeedPostLoader: knowPostsModel,
		OutboxModel:    relationmodel.NewOutboxModel(conn, c.CacheRedis),
		Redis:          rdb,
		UserCounterRpc: userCounterClient,
		CounterRpc:     counterClient,
		RelationRpc:    relationClient,

		DetailCache:   detail.New(l1Detail, l2, hotDetail),
		FeedMineCache: mine.New(l1FeedMine, l2, hotFeedMine),
		L1FeedPublic:  l1FeedPublic,
		L1FeedItem:    l1FeedItem,
		L1FeedEpoch:   l1FeedEpoch,
		L1FeedRoute:   l1FeedRoute,
		L1FeedPage:    l1FeedPage,
		L2:            l2,

		HotDetail:     hotDetail,
		HotFeedPublic: hotFeedPublic,
		HotFeedItem:   hotFeedItem,
		HotFeedMine:   hotFeedMine,

		Snowflake: snowflakex.MustNew(c.Snowflake.DatacenterId, c.Snowflake.WorkerId),

		FeedWriter: feed.NewFeedWriterWithOptions(
			redisAdapter,
			feed.NewKafkaProducerAdapter(feedPusher),
			logx.WithContext(context.Background()),
			feed.FeedWriterOptions{
				Strategy:     feedStrategy,
				TierMode:     authorTierMode,
				TierResolver: authorTierResolver,
			},
		),

		// 初始化 FeedReader
		FeedReader: feed.NewFeedReaderWithOptions(
			redisAdapter,
			relationAdapter,
			feedCounterAdapter,
			logx.WithContext(context.Background()),
			feed.FeedReaderOptions{
				Strategy:                  feedStrategy,
				TierMode:                  authorTierMode,
				TierResolver:              authorTierResolver,
				RouteCache:                l1FeedRoute,
				RouteCacheTTL:             c.Feed.RouteSnapshot.TTL,
				RouteSnapshotEnabled:      c.Feed.RouteSnapshot.Enabled,
				RelationEpochs:            feedEpochs,
				CombinedPipelineEnabled:   c.Feed.CombinedPipeline.Enabled,
				CombinedPipelineBatchSize: c.Feed.CombinedPipeline.BatchSize,
				Observer:                  feedObserver,
			},
		),
		FeedObserver:  feedObserver,
		FeedEpochs:    feedEpochs,
		FeedPageCache: feedPageCache,

		feedPusher:       feedPusher,
		feedFanoutWorker: feedFanoutWorker,
	}
}

func applyFeedEpochDefaults(c *config.FeedEpochConf) {
	if c.KeyPrefix == "" {
		c.KeyPrefix = "feed"
	}
	if c.RelationL1TTL <= 0 {
		c.RelationL1TTL = time.Second
	}
	if c.SafetyL1TTL <= 0 {
		c.SafetyL1TTL = time.Second
	}
}

func applyFeedRouteSnapshotDefaults(c *config.FeedRouteSnapshotConf) error {
	if c.TTL <= 0 {
		c.TTL = feedRouteCacheTTL
	}
	if c.TTL > feedRouteCacheTTL {
		return fmt.Errorf("feed RouteSnapshot TTL %s must not exceed %s", c.TTL, feedRouteCacheTTL)
	}
	return nil
}

func applyFeedCombinedPipelineDefaults(c *config.FeedCombinedPipelineConf) error {
	if c.BatchSize == 0 {
		c.BatchSize = maxCombinedPipelineBatchSize
	}
	if c.BatchSize < 2 || c.BatchSize > maxCombinedPipelineBatchSize {
		return fmt.Errorf("feed CombinedPipeline BatchSize %d must be between 2 and %d", c.BatchSize, maxCombinedPipelineBatchSize)
	}
	return nil
}

func applyFeedPageCacheDefaults(c *config.FeedPageCacheConf) error {
	c.Mode = strings.ToLower(strings.TrimSpace(c.Mode))
	if c.Mode == "" {
		c.Mode = string(userfeed.ModeOff)
	}
	switch userfeed.Mode(c.Mode) {
	case userfeed.ModeOff, userfeed.ModeL2, userfeed.ModeL1L2:
	default:
		return fmt.Errorf("feed PageCache Mode %q must be off, l2, or l1-l2", c.Mode)
	}
	if c.KeyPrefix == "" {
		c.KeyPrefix = "feed"
	}
	if c.Page == 0 {
		c.Page = 1
	}
	if c.Size == 0 {
		c.Size = 20
	}
	if c.L1FreshTTL <= 0 {
		c.L1FreshTTL = 800 * time.Millisecond
	}
	if c.L2FreshTTL <= 0 {
		c.L2FreshTTL = 4 * time.Second
	}
	if c.StaleTTL == 0 {
		c.StaleTTL = 10 * time.Second
	}
	if c.JitterPercent == 0 {
		c.JitterPercent = 20
	}
	if c.RefreshWorkers == 0 {
		c.RefreshWorkers = 32
	}
	if c.RefreshQueue == 0 {
		c.RefreshQueue = 1024
	}
	if c.LoaderTimeout == 0 {
		c.LoaderTimeout = 2 * time.Second
	}
	if c.Page != 1 || c.Size != 20 {
		return fmt.Errorf("feed PageCache only supports page=1,size=20, got page=%d,size=%d", c.Page, c.Size)
	}
	if c.StaleTTL <= 0 || c.StaleTTL > 10*time.Second {
		return fmt.Errorf("feed PageCache StaleTTL %s must be between 0 and 10s", c.StaleTTL)
	}
	if c.JitterPercent < 0 || c.JitterPercent > 20 {
		return fmt.Errorf("feed PageCache JitterPercent %d must be between 0 and 20", c.JitterPercent)
	}
	if err := validateJitteredFreshTTL("L1", c.L1FreshTTL, c.JitterPercent, 500*time.Millisecond, time.Second); err != nil {
		return err
	}
	if err := validateJitteredFreshTTL("L2", c.L2FreshTTL, c.JitterPercent, 3*time.Second, 5*time.Second); err != nil {
		return err
	}
	if c.RefreshWorkers <= 0 {
		return fmt.Errorf("feed PageCache RefreshWorkers %d must be positive", c.RefreshWorkers)
	}
	if c.RefreshQueue <= 0 {
		return fmt.Errorf("feed PageCache RefreshQueue %d must be positive", c.RefreshQueue)
	}
	if c.LoaderTimeout <= 0 {
		return fmt.Errorf("feed PageCache LoaderTimeout %s must be positive", c.LoaderTimeout)
	}
	return nil
}

func validateJitteredFreshTTL(
	name string,
	base time.Duration,
	jitterPercent int,
	minimum time.Duration,
	maximum time.Duration,
) error {
	lower := base * time.Duration(100-jitterPercent) / 100
	upper := base * time.Duration(100+jitterPercent) / 100
	if lower < minimum || upper > maximum {
		return fmt.Errorf(
			"feed PageCache %s effective Fresh TTL [%s,%s] must stay within [%s,%s]",
			name, lower, upper, minimum, maximum,
		)
	}
	return nil
}

func validateFeedSafetyConfiguration(c config.FeedConf) error {
	mode := userfeed.Mode(strings.ToLower(strings.TrimSpace(c.PageCache.Mode)))
	if mode != userfeed.ModeOff && !c.Epoch.SafetyConsumerEnabled {
		return fmt.Errorf("feed PageCache Mode %q requires Feed.Epoch.SafetyConsumerEnabled", c.PageCache.Mode)
	}
	return nil
}

func resolveFeedStrategy(c config.FeedConf) (feed.Strategy, error) {
	return feed.ParseStrategy(c.Strategy)
}

func resolveAuthorTierMode(c config.FeedConf) (feed.AuthorTierMode, error) {
	return feed.ParseAuthorTierMode(c.AuthorTier.Mode)
}

func pageCacheDependencyOperation(operation userfeed.Operation) (feed.FeedDependency, feed.FeedOperation) {
	switch operation {
	case userfeed.OperationL2Get:
		return feed.DependencyRedis, feed.OperationPageCacheL2Get
	case userfeed.OperationL2Decode:
		return feed.DependencyCache, feed.OperationPageCacheL2Decode
	case userfeed.OperationL2Encode:
		return feed.DependencyCache, feed.OperationPageCacheL2Encode
	case userfeed.OperationL2Set:
		return feed.DependencyRedis, feed.OperationPageCacheL2Set
	case userfeed.OperationSingleflightShared:
		return feed.DependencyCache, feed.OperationPageCacheSFShared
	case userfeed.OperationRefreshEnqueue:
		return feed.DependencyCache, feed.OperationPageCacheRefreshEnqueue
	case userfeed.OperationRefreshLoad:
		return feed.DependencyCache, feed.OperationPageCacheRefreshLoad
	default:
		return feed.DependencyUnknown, feed.OperationUnknown
	}
}

func (s *ServiceContext) StartFeedFanoutWorker() {
	go s.feedFanoutWorker.Start()
	logx.Info("feed fanout worker started")
}

func (s *ServiceContext) Close() {
	s.feedFanoutWorker.Stop()
	if err := s.feedPusher.Close(); err != nil {
		logx.Errorf("close feed kafka producer: %v", err)
	}
	if s.FeedPageCache != nil {
		if err := s.FeedPageCache.Close(); err != nil {
			logx.Errorf("close feed page cache: %v", err)
		}
	}
}

func mustL1(numCounters, maxCost int64) *cachex.L1 {
	l1, err := cachex.NewL1(cachex.L1Config{NumCounters: numCounters, MaxCost: maxCost})
	if err != nil {
		panic(err)
	}
	return l1
}
