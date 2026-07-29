package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"

	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	relationpb "github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

// LoadTestConfig 压测配置
type LoadTestConfig struct {
	// 用户配置
	UserCount     int   // 测试用户数量
	UserIDStart   int64 // 起始用户ID
	BigVUserCount int   // 大V用户数量

	// 关系配置
	FollowingsPerUser int // 每个用户关注数量
	FollowRatio       int // 关注大V的比例（0-100）

	// 发帖配置
	PostCount       int           // 总发帖数
	PostConcurrency int           // 发帖并发数
	PostInterval    time.Duration // 发帖间隔

	// 读取配置
	ReadCount       int           // 总读取数
	ReadConcurrency int           // 读取并发数
	ReadInterval    time.Duration // 读取间隔

	// 服务配置
	KnowPostRpc zrpc.RpcClientConf
	RelationRpc zrpc.RpcClientConf
	CounterRpc  zrpc.RpcClientConf
	UserRpc     zrpc.RpcClientConf
}

// LoadTestStats 压测统计
type LoadTestStats struct {
	// 发帖统计
	PostTotal     int64
	PostSuccess   int64
	PostFail      int64
	PostTotalTime int64 // 纳秒

	// 读取统计
	ReadTotal     int64
	ReadSuccess   int64
	ReadFail      int64
	ReadTotalTime int64 // 纳秒

	// 扇出统计
	PushModeCount int64
	PullModeCount int64

	mu sync.Mutex
}

type userCreator interface {
	Create(context.Context, *userpb.CreateReq, ...grpc.CallOption) (*userpb.CreateResp, error)
}

func createTestUsers(ctx context.Context, client userCreator, count int, runID string) ([]int64, error) {
	userIDs := make([]int64, 0, count)
	for i := 0; i < count; i++ {
		resp, err := client.Create(ctx, &userpb.CreateReq{
			Email:    fmt.Sprintf("loadtest+%s-%d@example.invalid", runID, i),
			Nickname: fmt.Sprintf("Load Test User %d", i),
		})
		if err != nil {
			return nil, fmt.Errorf("create load-test user %d: %w", i, err)
		}
		if resp == nil || resp.Id <= 0 {
			return nil, fmt.Errorf("create load-test user %d: invalid user id", i)
		}
		userIDs = append(userIDs, resp.Id)
	}
	return userIDs, nil
}

func newConfirmContentRequest(postID, creatorID int64) *knowpostpb.ConfirmContentReq {
	return &knowpostpb.ConfirmContentReq{
		Id:        postID,
		CreatorId: creatorID,
		ObjectKey: fmt.Sprintf("loadtest/%d/content.md", postID),
		Size:      1,
	}
}

func validateLoadTestResults(config *LoadTestConfig, stats *LoadTestStats) error {
	postTotal := atomic.LoadInt64(&stats.PostTotal)
	postSuccess := atomic.LoadInt64(&stats.PostSuccess)
	postFail := atomic.LoadInt64(&stats.PostFail)
	readTotal := atomic.LoadInt64(&stats.ReadTotal)
	readSuccess := atomic.LoadInt64(&stats.ReadSuccess)
	readFail := atomic.LoadInt64(&stats.ReadFail)

	if postTotal != int64(config.PostCount) {
		return fmt.Errorf("post requests: got %d, want %d", postTotal, config.PostCount)
	}
	if postFail > 0 {
		return fmt.Errorf("post failures: %d", postFail)
	}
	if postSuccess != int64(config.PostCount) {
		return fmt.Errorf("post successes: got %d, want %d", postSuccess, config.PostCount)
	}
	if readTotal != int64(config.ReadCount) {
		return fmt.Errorf("read requests: got %d, want %d", readTotal, config.ReadCount)
	}
	if readFail > 0 {
		return fmt.Errorf("read failures: %d", readFail)
	}
	if readSuccess != int64(config.ReadCount) {
		return fmt.Errorf("read successes: got %d, want %d", readSuccess, config.ReadCount)
	}
	return nil
}

func waitForLoadTest(done <-chan struct{}, timeout time.Duration) error {
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case <-done:
		return nil
	case <-timer.C:
		return fmt.Errorf("load test timed out after %s", timeout)
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "\n压测失败: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	var (
		configFile      = flag.String("f", "load_test.yaml", "配置文件路径")
		userCount       = flag.Int("users", 100, "测试用户数量")
		bigvCount       = flag.Int("bigv", 10, "大V用户数量")
		followings      = flag.Int("followings", 20, "每用户关注数")
		postCount       = flag.Int("posts", 1000, "总发帖数")
		postConcurrency = flag.Int("post-c", 10, "发帖并发数")
		readCount       = flag.Int("reads", 5000, "总读取数")
		readConcurrency = flag.Int("read-c", 50, "读取并发数")
		duration        = flag.Duration("duration", 60*time.Second, "测试持续时间")
		setupOnly       = flag.Bool("setup-only", false, "只执行初始化，不压测")
		skipSetup       = flag.Bool("skip-setup", false, "跳过初始化，直接压测")
	)
	flag.Parse()

	logx.DisableStat()

	// 加载配置
	var config LoadTestConfig
	if err := conf.Load(*configFile, &config, conf.UseEnv()); err != nil {
		// 使用命令行参数作为默认配置
		config = LoadTestConfig{
			UserCount:         *userCount,
			UserIDStart:       100000,
			BigVUserCount:     *bigvCount,
			FollowingsPerUser: *followings,
			FollowRatio:       30, // 30%关注大V
			PostCount:         *postCount,
			PostConcurrency:   *postConcurrency,
			PostInterval:      10 * time.Millisecond,
			ReadCount:         *readCount,
			ReadConcurrency:   *readConcurrency,
			ReadInterval:      0, // 移除限制，全速测试
			KnowPostRpc: zrpc.RpcClientConf{
				Etcd: discov.EtcdConf{
					Hosts: []string{"127.0.0.1:2379"},
					Key:   "knowpost.rpc",
				},
				NonBlock: true,
				Timeout:  5000,
			},
			RelationRpc: zrpc.RpcClientConf{
				Etcd: discov.EtcdConf{
					Hosts: []string{"127.0.0.1:2379"},
					Key:   "relation.rpc",
				},
				NonBlock: true,
				Timeout:  5000,
			},
			CounterRpc: zrpc.RpcClientConf{
				Etcd: discov.EtcdConf{
					Hosts: []string{"127.0.0.1:2379"},
					Key:   "counter.rpc",
				},
				NonBlock: true,
				Timeout:  5000,
			},
			UserRpc: zrpc.RpcClientConf{
				Etcd: discov.EtcdConf{
					Hosts: []string{"127.0.0.1:2379"},
					Key:   "user.rpc",
				},
				NonBlock: true,
				Timeout:  5000,
			},
		}
	}

	fmt.Println("=== Feed推拉结合架构 - 端到端压力测试 ===")
	fmt.Printf("测试配置:\n")
	fmt.Printf("  用户数: %d (其中大V: %d)\n", config.UserCount, config.BigVUserCount)
	fmt.Printf("  每用户关注: %d\n", config.FollowingsPerUser)
	fmt.Printf("  发帖: %d条 (并发: %d)\n", config.PostCount, config.PostConcurrency)
	fmt.Printf("  读取: %d次 (并发: %d)\n", config.ReadCount, config.ReadConcurrency)
	fmt.Printf("  测试时长: %v\n", *duration)
	fmt.Println()

	// 连接RPC服务
	fmt.Println("连接RPC服务...")
	knowpostClient := knowpostpb.NewKnowPostClient(zrpc.MustNewClient(config.KnowPostRpc).Conn())
	relationClient := relationpb.NewRelationClient(zrpc.MustNewClient(config.RelationRpc).Conn())
	counterClient := counterpb.NewUserCounterClient(zrpc.MustNewClient(config.CounterRpc).Conn())
	userClient := userpb.NewUserClient(zrpc.MustNewClient(config.UserRpc).Conn())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	userIDs := make([]int64, config.UserCount)
	for i := range userIDs {
		userIDs[i] = config.UserIDStart + int64(i)
	}

	// 初始化测试数据
	if !*skipSetup {
		fmt.Println("初始化测试数据...")
		fmt.Println("  创建测试用户...")
		var err error
		userIDs, err = createTestUsers(
			ctx,
			userClient,
			config.UserCount,
			strconv.FormatInt(time.Now().UnixNano(), 10),
		)
		if err != nil {
			return fmt.Errorf("initialize users: %w", err)
		}
		fmt.Printf("  ✓ 测试用户创建完成: %d\n", len(userIDs))

		if err := setupTestData(ctx, &config, userIDs, relationClient, counterClient); err != nil {
			return fmt.Errorf("initialize relations: %w", err)
		}
		fmt.Println("✓ 初始化完成")
		fmt.Println()

		if *setupOnly {
			fmt.Println("仅初始化模式，测试结束")
			return nil
		}
	}

	// 统计对象
	stats := &LoadTestStats{}

	// 启动统计goroutine
	stopStats := make(chan struct{})
	go printStats(stats, stopStats)

	fmt.Println("开始压力测试...")
	startTime := time.Now()

	// 启动发帖压测
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		runPostLoadTest(ctx, &config, userIDs, knowpostClient, counterClient, stats)
	}()

	// 启动读取压测
	wg.Add(1)
	go func() {
		defer wg.Done()
		runReadLoadTest(ctx, &config, userIDs, knowpostClient, relationClient, counterClient, stats)
	}()

	// 等待测试完成或超时
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	completionErr := waitForLoadTest(done, *duration)
	if completionErr == nil {
		fmt.Println("\n测试任务完成")
	} else {
		cancel()
		fmt.Println("\n测试时间到达")
	}

	elapsed := time.Since(startTime)
	close(stopStats)
	time.Sleep(100 * time.Millisecond) // 等待统计输出

	// 打印最终报告
	printFinalReport(stats, elapsed)
	if completionErr != nil {
		return completionErr
	}
	return validateLoadTestResults(&config, stats)
}

// setupTestData 初始化测试数据
func setupTestData(ctx context.Context, config *LoadTestConfig, userIDs []int64,
	relationClient relationpb.RelationClient, counterClient counterpb.UserCounterClient) error {

	fmt.Printf("  创建关注关系...\n")

	// 为每个用户创建关注关系
	var successCount, errorCount int64
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 20) // 并发限制

	for _, userID := range userIDs {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(userID int64) {
			defer wg.Done()
			defer func() { <-semaphore }()

			// 随机选择关注对象
			followings := make([]int64, 0, config.FollowingsPerUser)
			for j := 0; j < config.FollowingsPerUser; j++ {
				var targetID int64

				// 按比例关注大V或普通用户
				if rand.Intn(100) < config.FollowRatio && config.BigVUserCount > 0 {
					// 关注大V
					bigvIdx := rand.Intn(config.BigVUserCount)
					targetID = userIDs[bigvIdx]
				} else {
					// 关注普通用户
					targetID = userIDs[rand.Intn(len(userIDs))]
				}

				if targetID == userID {
					continue // 不关注自己
				}

				// 去重
				exists := false
				for _, fid := range followings {
					if fid == targetID {
						exists = true
						break
					}
				}
				if !exists {
					followings = append(followings, targetID)
				}
			}

			// 批量创建关注
			for _, targetID := range followings {
				_, err := relationClient.Follow(ctx, &relationpb.FollowReq{
					FromUserId: userID,
					ToUserId:   targetID,
				})
				if err != nil {
					atomic.AddInt64(&errorCount, 1)
				} else {
					atomic.AddInt64(&successCount, 1)
				}
			}
		}(userID)
	}

	wg.Wait()
	fmt.Printf("  ✓ 关注关系创建完成: 成功=%d, 失败=%d\n", successCount, errorCount)
	if errorCount > 0 {
		return fmt.Errorf("create followings: %d requests failed", errorCount)
	}

	// 等待Counter同步（通过Kafka消费）
	fmt.Printf("  等待Counter同步...\n")
	time.Sleep(2 * time.Second)
	fmt.Printf("  ✓ Counter同步完成\n")

	return nil
}

// runPostLoadTest 发帖压测
func runPostLoadTest(ctx context.Context, config *LoadTestConfig, userIDs []int64,
	knowpostClient knowpostpb.KnowPostClient, counterClient counterpb.UserCounterClient,
	stats *LoadTestStats) {

	semaphore := make(chan struct{}, config.PostConcurrency)

	for i := 0; i < config.PostCount; i++ {
		semaphore <- struct{}{}

		go func(idx int) {
			defer func() { <-semaphore }()

			atomic.AddInt64(&stats.PostTotal, 1)

			// 随机选择一个用户发帖
			userID := userIDs[rand.Intn(len(userIDs))]

			start := time.Now()

			// 1. 创建草稿
			draftResp, err := knowpostClient.CreateDraft(ctx, &knowpostpb.CreateDraftReq{
				CreatorId: userID,
			})
			if err != nil {
				recordPostFailure(stats, "CreateDraft", err)
				return
			}

			// 转换ID为int64
			postID, err := strconv.ParseInt(draftResp.Id, 10, 64)
			if err != nil {
				recordPostFailure(stats, "ParseDraftID", err)
				return
			}

			// 2. 更新元数据
			_, err = knowpostClient.PatchMetadata(ctx, &knowpostpb.PatchMetadataReq{
				Id:        postID,
				CreatorId: userID,
				Title:     fmt.Sprintf("Load Test Post #%d", idx),
				TitleSet:  true,
			})
			if err != nil {
				recordPostFailure(stats, "PatchMetadata", err)
				return
			}

			// 3. 确认内容
			_, err = knowpostClient.ConfirmContent(ctx, newConfirmContentRequest(postID, userID))
			if err != nil {
				recordPostFailure(stats, "ConfirmContent", err)
				return
			}

			// 4. 发布
			_, err = knowpostClient.Publish(ctx, &knowpostpb.PublishReq{
				Id:        postID,
				CreatorId: userID,
			})
			if err != nil {
				recordPostFailure(stats, "Publish", err)
				return
			}

			elapsed := time.Since(start)
			atomic.AddInt64(&stats.PostTotalTime, int64(elapsed))
			atomic.AddInt64(&stats.PostSuccess, 1)

			// 判断推拉模式（通过获取粉丝数）
			snapshot, err := counterClient.GetUserSnapshot(ctx, &counterpb.GetUserSnapshotReq{
				UserId: userID,
			})
			if err == nil && snapshot.Snapshot != nil {
				if snapshot.Snapshot.Followers > 1000 {
					atomic.AddInt64(&stats.PullModeCount, 1)
				} else {
					atomic.AddInt64(&stats.PushModeCount, 1)
				}
			}

		}(i)

		if config.PostInterval > 0 {
			time.Sleep(config.PostInterval)
		}
	}

	// 等待所有发帖完成
	for i := 0; i < config.PostConcurrency; i++ {
		semaphore <- struct{}{}
	}
}

func recordPostFailure(stats *LoadTestStats, stage string, err error) {
	failures := atomic.AddInt64(&stats.PostFail, 1)
	if failures <= 5 {
		fmt.Printf("\n[ERROR] %s failed: %v\n", stage, err)
	}
}

// runReadLoadTest 读取Feed压测
func runReadLoadTest(ctx context.Context, config *LoadTestConfig, userIDs []int64,
	knowpostClient knowpostpb.KnowPostClient, relationClient relationpb.RelationClient,
	counterClient counterpb.UserCounterClient, stats *LoadTestStats) {

	// 等待一些帖子发布
	time.Sleep(2 * time.Second)

	semaphore := make(chan struct{}, config.ReadConcurrency)

	for i := 0; i < config.ReadCount; i++ {
		semaphore <- struct{}{}

		go func() {
			defer func() { <-semaphore }()

			atomic.AddInt64(&stats.ReadTotal, 1)

			// 随机选择一个用户读取Feed
			userID := userIDs[rand.Intn(len(userIDs))]

			start := time.Now()

			// 读取用户个性化Feed（推拉混合）
			_, err := knowpostClient.GetUserFeed(ctx, &knowpostpb.GetUserFeedReq{
				UserId: userID,
				Page:   1,
				Size:   20,
			})

			elapsed := time.Since(start)
			atomic.AddInt64(&stats.ReadTotalTime, int64(elapsed))

			if err != nil {
				atomic.AddInt64(&stats.ReadFail, 1)
				// 只打印前几个错误
				if atomic.LoadInt64(&stats.ReadFail) <= 5 {
					fmt.Printf("\n[ERROR] GetUserFeed failed: %v\n", err)
				}
			} else {
				atomic.AddInt64(&stats.ReadSuccess, 1)
			}
		}()

		if config.ReadInterval > 0 {
			time.Sleep(config.ReadInterval)
		}
	}

	// 等待所有读取完成
	for i := 0; i < config.ReadConcurrency; i++ {
		semaphore <- struct{}{}
	}
}

// printStats 定期打印统计
func printStats(stats *LoadTestStats, stop chan struct{}) {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			postTotal := atomic.LoadInt64(&stats.PostTotal)
			postSuccess := atomic.LoadInt64(&stats.PostSuccess)
			postFail := atomic.LoadInt64(&stats.PostFail)

			readTotal := atomic.LoadInt64(&stats.ReadTotal)
			readSuccess := atomic.LoadInt64(&stats.ReadSuccess)
			readFail := atomic.LoadInt64(&stats.ReadFail)

			pushMode := atomic.LoadInt64(&stats.PushModeCount)
			pullMode := atomic.LoadInt64(&stats.PullModeCount)

			fmt.Printf("\r[进度] 发帖: %d/%d/%d (总/成功/失败) | 读取: %d/%d/%d | 推拉: %d/%d     ",
				postTotal, postSuccess, postFail, readTotal, readSuccess, readFail, pushMode, pullMode)

		case <-stop:
			return
		}
	}
}

// printFinalReport 打印最终报告
func printFinalReport(stats *LoadTestStats, elapsed time.Duration) {
	fmt.Print("\n\n")
	fmt.Println("=== 压测报告 ===")
	fmt.Println()

	// 发帖统计
	postTotal := atomic.LoadInt64(&stats.PostTotal)
	postSuccess := atomic.LoadInt64(&stats.PostSuccess)
	postFail := atomic.LoadInt64(&stats.PostFail)
	postTotalTime := atomic.LoadInt64(&stats.PostTotalTime)

	fmt.Println("【发帖性能】")
	fmt.Printf("  总请求数: %d\n", postTotal)
	fmt.Printf("  成功: %d (%.1f%%)\n", postSuccess, float64(postSuccess)/float64(postTotal)*100)
	fmt.Printf("  失败: %d (%.1f%%)\n", postFail, float64(postFail)/float64(postTotal)*100)

	if postSuccess > 0 {
		avgPostTime := time.Duration(postTotalTime / postSuccess)
		fmt.Printf("  平均延迟: %v\n", avgPostTime)
		fmt.Printf("  QPS: %.1f\n", float64(postSuccess)/elapsed.Seconds())
	}

	pushMode := atomic.LoadInt64(&stats.PushModeCount)
	pullMode := atomic.LoadInt64(&stats.PullModeCount)
	fmt.Printf("  推模式: %d (%.1f%%)\n", pushMode, float64(pushMode)/float64(postSuccess)*100)
	fmt.Printf("  拉模式: %d (%.1f%%)\n", pullMode, float64(pullMode)/float64(postSuccess)*100)
	fmt.Println()

	// 读取统计
	readTotal := atomic.LoadInt64(&stats.ReadTotal)
	readSuccess := atomic.LoadInt64(&stats.ReadSuccess)
	readFail := atomic.LoadInt64(&stats.ReadFail)
	readTotalTime := atomic.LoadInt64(&stats.ReadTotalTime)

	fmt.Println("【读取性能】")
	fmt.Printf("  总请求数: %d\n", readTotal)
	fmt.Printf("  成功: %d (%.1f%%)\n", readSuccess, float64(readSuccess)/float64(readTotal)*100)
	fmt.Printf("  失败: %d (%.1f%%)\n", readFail, float64(readFail)/float64(readTotal)*100)

	if readSuccess > 0 {
		avgReadTime := time.Duration(readTotalTime / readSuccess)
		fmt.Printf("  平均延迟: %v\n", avgReadTime)
		fmt.Printf("  QPS: %.1f\n", float64(readSuccess)/elapsed.Seconds())
	}
	fmt.Println()

	// 总体统计
	fmt.Println("【总体统计】")
	fmt.Printf("  测试时长: %v\n", elapsed)
	fmt.Printf("  总请求数: %d\n", postTotal+readTotal)
	fmt.Printf("  总成功数: %d\n", postSuccess+readSuccess)
	fmt.Printf("  总QPS: %.1f\n", float64(postSuccess+readSuccess)/elapsed.Seconds())
	fmt.Println()
}
