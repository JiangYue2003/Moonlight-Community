package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"golang.org/x/crypto/bcrypt"

	"github.com/zhiguang/zhiguang-go/pkg/jwtx"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	relationpb "github.com/zhiguang/zhiguang-go/services/relation/rpc/relation"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

type setupConfig struct {
	UserConcurrency       int           `json:",default=16"`
	CursorSeedConcurrency int           `json:",default=4"`
	RelationConcurrency   int           `json:",default=32"`
	WarmupConcurrency     int           `json:",default=1"`
	StrategySeedPosts     int           `json:",default=5"`
	CounterSyncTimeout    time.Duration `json:",default=120s"`
	CounterPollInterval   time.Duration `json:",default=500ms"`
	WarmupPosts           int           `json:",default=2"`
}

type benchmarkLoadConfig struct {
	Requests         int           `json:",default=1000"`
	Duration         time.Duration `json:",default=0s"`
	RequestTimeout   time.Duration `json:",default=5s"`
	Page             int32         `json:",default=1"`
	Size             int32         `json:",default=20"`
	ConcurrencySteps []int
}

type executionConfig struct {
	Topology              string  `json:",default=compose-full"`
	LocalStateFile        string  `json:",default=.tmp/feed-local/state.json"`
	FeaturePreset         string  `json:",default=treatment"`
	CacheState            string  `json:",default=natural"`
	CacheWarmConcurrency  int     `json:",default=64"`
	ClientCPULimitPercent float64 `json:",default=90"`
	PprofDir              string
}

type benchmarkConfig struct {
	RunID             string
	ReaderCardinality string `json:",default=hot"`
	Password          string
	GatewayURL        string `json:",default=http://127.0.0.1:8080"`
	ReportDir         string `json:",default=results/feed-loadtest"`
	ManifestDir       string `json:",default=results/feed-loadtest/manifests"`
	RedisAddr         string `json:",default=127.0.0.1:6379"`

	Topology  topologyConfig
	Setup     setupConfig
	Load      benchmarkLoadConfig
	Monitor   monitorConfig
	Execution executionConfig

	KnowPostRpc zrpc.RpcClientConf
	RelationRpc zrpc.RpcClientConf
	CounterRpc  zrpc.RpcClientConf
	UserRpc     zrpc.RpcClientConf
	AuthRpc     zrpc.RpcClientConf
}

type rpcClients struct {
	knowpost knowpostpb.KnowPostClient
	relation relationpb.RelationClient
	counter  counterpb.UserCounterClient
	user     userpb.UserClient
	auth     userpb.AuthClient
}

type cliOptions struct {
	configPath        string
	phase             string
	runID             string
	reportRunID       string
	manifest          string
	entry             string
	strategy          string
	scenario          string
	concurrency       int
	concurrencies     string
	requests          int
	duration          time.Duration
	warmup            time.Duration
	trials            int
	readerCardinality string
	topology          string
	pprofDir          string
	cacheState        string
	featurePreset     string
	gatewayAuthMode   string
	jwtPrivateKey     string
	jwtPublicKey      string
	jwtIssuer         string
	confirmCleanup    string
	cleanupDryRun     bool
	checkpointPath    string
	confirmMutation   string
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "feed load test failed: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	options := parseCLI()
	var cfg benchmarkConfig
	if err := conf.Load(options.configPath, &cfg, conf.UseEnv()); err != nil {
		return fmt.Errorf("load config %s: %w", options.configPath, err)
	}
	applyConfigDefaults(&cfg)
	if options.readerCardinality != "" {
		cfg.ReaderCardinality = strings.ToLower(strings.TrimSpace(options.readerCardinality))
	}
	if err := applyReaderCardinalityPreset(cfg.ReaderCardinality, &cfg.Topology); err != nil {
		return err
	}
	if options.topology != "" {
		cfg.Execution.Topology = strings.ToLower(strings.TrimSpace(options.topology))
	}
	if options.pprofDir != "" {
		cfg.Execution.PprofDir = options.pprofDir
	}
	if options.cacheState != "" {
		cfg.Execution.CacheState = strings.ToLower(strings.TrimSpace(options.cacheState))
	}
	if options.featurePreset != "" {
		cfg.Execution.FeaturePreset = strings.ToLower(strings.TrimSpace(options.featurePreset))
	}
	if err := validateExecutionConfig(cfg.Execution); err != nil {
		return err
	}
	applyExecutionTopology(&cfg)
	if options.runID != "" {
		cfg.RunID = options.runID
	}
	if cfg.RunID == "" {
		cfg.RunID = "feed-" + time.Now().Format("20060102-150405")
	}
	if options.manifest == "" {
		options.manifest = filepath.Join(cfg.ManifestDir, cfg.RunID+".json")
	}
	if options.checkpointPath == "" {
		options.checkpointPath = defaultMutationCheckpointPath(cfg.ReportDir, cfg.RunID)
	}
	if options.concurrency > 0 && strings.TrimSpace(options.concurrencies) != "" {
		return fmt.Errorf("use either -concurrency or -concurrencies, not both")
	}
	if strings.TrimSpace(options.concurrencies) != "" {
		levels, err := parseConcurrencyList(options.concurrencies)
		if err != nil {
			return err
		}
		cfg.Load.ConcurrencySteps = levels
	} else if options.concurrency > 0 {
		cfg.Load.ConcurrencySteps = []int{options.concurrency}
	}
	if options.requests >= 0 {
		cfg.Load.Requests = options.requests
	}
	if options.duration >= 0 {
		cfg.Load.Duration = options.duration
	}

	if options.phase == "run" {
		// The go-zero duration interceptor logs full request/response payloads
		// after 500 ms. At saturation that both leaks Feed content and makes the
		// load generator part of the bottleneck, so measured runs use recorders
		// and reports as their only output channel.
		logx.Disable()
	} else {
		logx.DisableStat()
	}
	ctx := context.Background()
	if options.phase == "compare" {
		paths, err := generateComparison(cfg.ReportDir, effectiveReportRunID(cfg.RunID, options.reportRunID))
		if err != nil {
			return err
		}
		fmt.Printf("compare: json=%s csv=%s markdown=%s\n", paths.JSON, paths.CSV, paths.Markdown)
		return nil
	}
	if options.phase == "snapshot" {
		snapshot, err := captureEnvironmentSnapshot(ctx, cfg.Monitor, cfg.Execution, cfg.RunID)
		if err != nil {
			return err
		}
		path, err := writeEnvironmentSnapshot(cfg.ReportDir, snapshot)
		if err != nil {
			return err
		}
		fmt.Printf("snapshot: strategy=%s containers=%d path=%s\n", snapshot.EffectiveStrategy, len(snapshot.Containers), path)
		return nil
	}
	if options.phase == "cleanup" {
		manifest, err := loadManifest(options.manifest)
		if err != nil {
			return err
		}
		result, err := cleanupDataset(ctx, cfg, manifest, options.confirmCleanup, options.cleanupDryRun)
		if err != nil {
			return err
		}
		fmt.Printf("cleanup: run=%s dry_run=%t redis_keys=%d planned_keys=%d %s\n", result.RunID, result.DryRun, result.RedisKeys, result.RedisKeyPlan, cleanupSummary(result))
		return nil
	}
	if phaseRequiresEffectiveStrategy(options.phase) {
		runtimeState, err := readFeedRuntimeState(ctx, cfg.Execution, cfg.Monitor)
		if err != nil {
			return err
		}
		if err := validateEffectiveStrategy(options.strategy, runtimeState.Strategy); err != nil {
			return err
		}
		if err := validateRuntimeFeaturePreset(cfg.Execution.FeaturePreset, runtimeState.Features); err != nil {
			return err
		}
	}
	if options.phase == "create-mutation-checkpoint" {
		manifest, err := loadManifest(options.manifest)
		if err != nil {
			return err
		}
		checkpoint, err := createMutationCheckpointRuntime(
			ctx, cfg, manifest, options.strategy, options.confirmMutation, options.checkpointPath,
		)
		if err != nil {
			return err
		}
		fmt.Printf("create-mutation-checkpoint: run=%s id=%s path=%s posts=%d redis_keys=%d\n",
			manifest.RunID, checkpoint.ID, options.checkpointPath, len(checkpoint.MySQL.PostIDs), len(checkpoint.Redis.Keys))
		return nil
	}
	if options.phase == "verify-mutation-checkpoint" || options.phase == "restore-mutation-checkpoint" {
		manifest, err := loadManifest(options.manifest)
		if err != nil {
			return err
		}
		confirmation := manifest.RunID
		if options.phase == "restore-mutation-checkpoint" {
			confirmation = options.confirmMutation
		}
		controller, err := openMutationCheckpointController(
			ctx, cfg, manifest, options.strategy, confirmation, options.checkpointPath,
		)
		if err != nil {
			return err
		}
		defer controller.Close()
		if options.phase == "verify-mutation-checkpoint" {
			if err := controller.Verify(ctx); err != nil {
				return err
			}
			fmt.Printf("verify-mutation-checkpoint: run=%s id=%s complete=true\n", manifest.RunID, controller.checkpoint.ID)
			return nil
		}
		observation, err := controller.Restore(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("restore-mutation-checkpoint: run=%s id=%s posts=%d outbox=%d safety_epoch=%d duration=%s\n",
			manifest.RunID, controller.checkpoint.ID, observation.RemovedPosts, observation.RemovedOutbox,
			observation.SafetyEpoch, observation.Duration)
		return nil
	}

	clients := connectRPCs(cfg)
	switch options.phase {
	case "setup":
		_, err := setupDataset(ctx, cfg, clients, options.manifest, options.strategy)
		return err
	case "setup-cursor-deep":
		if cfg.ReaderCardinality != readerCardinalityCursorDeep {
			return fmt.Errorf("setup-cursor-deep requires -reader-cardinality %s", readerCardinalityCursorDeep)
		}
		cursorCfg := cfg
		cursorCfg.Setup.WarmupPosts = 0
		manifest, err := setupDataset(ctx, cursorCfg, clients, options.manifest, options.strategy)
		if err != nil {
			return err
		}
		manifest, err = seedCursorDeepDataset(ctx, cursorCfg, clients.knowpost, manifest, options.manifest)
		if err != nil {
			return err
		}
		fmt.Printf("setup-cursor-deep: manifest=%s posts=%d oracle=%d hash=%s\n", options.manifest, len(manifest.Posts), len(manifest.CursorDeep.Oracle), manifest.CursorDeep.OracleHash)
		return nil
	case "seed-cursor-deep":
		if cfg.ReaderCardinality != readerCardinalityCursorDeep {
			return fmt.Errorf("seed-cursor-deep requires -reader-cardinality %s", readerCardinalityCursorDeep)
		}
		manifest, err := loadManifest(options.manifest)
		if err != nil {
			return err
		}
		if err := validateResumableManifest(cfg, manifest, options.strategy); err != nil {
			return err
		}
		manifest, err = seedCursorDeepDataset(ctx, cfg, clients.knowpost, manifest, options.manifest)
		if err != nil {
			return err
		}
		fmt.Printf("seed-cursor-deep: manifest=%s posts=%d oracle=%d hash=%s\n", options.manifest, len(manifest.Posts), len(manifest.CursorDeep.Oracle), manifest.CursorDeep.OracleHash)
		return nil
	case "resume-setup":
		manifest, err := loadManifest(options.manifest)
		if err != nil {
			return err
		}
		if err := validateResumableManifest(cfg, manifest, options.strategy); err != nil {
			return err
		}
		manifest, err = finalizeDatasetSetup(ctx, cfg, clients.counter, clients.knowpost, manifest, options.manifest)
		if err != nil {
			return err
		}
		fmt.Printf("resume-setup: manifest=%s posts=%d\n", options.manifest, len(manifest.Posts))
		return nil
	case "smoke":
		manifest, err := loadManifest(options.manifest)
		if err != nil {
			return err
		}
		if err := validateManifestStrategy(manifest, options.strategy); err != nil {
			return err
		}
		return smokeDataset(ctx, cfg, clients, manifest, options.strategy)
	case "run":
		manifest, err := loadManifest(options.manifest)
		if err != nil {
			return err
		}
		if _, err := resolveManifestReaderCardinality(manifest, cfg.ReaderCardinality); err != nil {
			return err
		}
		if err := validateManifestSeed(manifest, cfg.Topology.Seed); err != nil {
			return err
		}
		if err := validateManifestStrategy(manifest, options.strategy); err != nil {
			return err
		}
		if scenarioRequiresCursorDeepDataset(options.scenario) {
			if err := validateCursorDeepManifest(manifest); err != nil {
				return err
			}
		}
		return runBenchmarks(ctx, cfg, clients, manifest, options)
	case "reset-feed":
		manifest, err := loadManifest(options.manifest)
		if err != nil {
			return err
		}
		deleted, err := resetFeedKeys(ctx, cfg.RedisAddr, manifest)
		if err != nil {
			return err
		}
		fmt.Printf("reset-feed: deleted=%d scoped_keys=%d run=%s\n", deleted, len(feedKeysForManifest(manifest)), manifest.RunID)
		return nil
	case "seed-feed":
		manifest, err := loadManifest(options.manifest)
		if err != nil {
			return err
		}
		if err := validateManifestStrategy(manifest, options.strategy); err != nil {
			return err
		}
		posts, err := seedFeed(ctx, clients.knowpost, manifest, options.strategy, cfg.Setup.StrategySeedPosts, cfg.Load.RequestTimeout)
		if err != nil {
			return err
		}
		fmt.Printf("seed-feed: strategy=%s posts=%d authors=%d depth=%d\n", options.strategy, len(posts), len(manifest.NormalAuthors)+len(manifest.BigVAuthors), cfg.Setup.StrategySeedPosts)
		return nil
	case "all":
		manifest, err := setupDataset(ctx, cfg, clients, options.manifest, options.strategy)
		if err != nil {
			return err
		}
		if err := smokeDataset(ctx, cfg, clients, manifest, options.strategy); err != nil {
			return err
		}
		return runBenchmarks(ctx, cfg, clients, manifest, options)
	default:
		return fmt.Errorf("invalid phase %q: expected setup, resume-setup, setup-cursor-deep, seed-cursor-deep, smoke, reset-feed, seed-feed, create-mutation-checkpoint, verify-mutation-checkpoint, restore-mutation-checkpoint, run, compare, snapshot, cleanup, or all", options.phase)
	}
}

func parseCLI() cliOptions {
	var options cliOptions
	flag.StringVar(&options.configPath, "f", "cmd/loadtest/load_test.yaml", "benchmark YAML")
	flag.StringVar(&options.phase, "phase", "run", "setup, resume-setup, setup-cursor-deep, seed-cursor-deep, smoke, reset-feed, seed-feed, create-mutation-checkpoint, verify-mutation-checkpoint, restore-mutation-checkpoint, run, compare, snapshot, cleanup, or all")
	flag.StringVar(&options.runID, "run-id", "", "stable dataset run id")
	flag.StringVar(&options.reportRunID, "report-run-id", "", "result namespace; defaults to the dataset run id")
	flag.StringVar(&options.manifest, "manifest", "", "dataset manifest path")
	flag.StringVar(&options.entry, "entry", "rpc", "rpc or gateway")
	flag.StringVar(&options.strategy, "strategy", "hybrid", "push, pull, or hybrid label")
	flag.StringVar(&options.scenario, "scenario", "distributed-read", "hot-read, distributed-read, page1/page5/page20/page50, cursor-page5/cursor-page20/cursor-page50, sequential-page-50, sequential-50, same-second-cursor, deep-page, burst-read, steady-read, publish, burst-publish, mixed-90-10, or mixed-80-20")
	flag.IntVar(&options.concurrency, "concurrency", 0, "single concurrency override")
	flag.StringVar(&options.concurrencies, "concurrencies", "", "comma-separated concurrency overrides, prepared in one process")
	flag.IntVar(&options.requests, "requests", -1, "request count override; 0 uses duration")
	flag.DurationVar(&options.duration, "duration", -1, "duration override")
	flag.DurationVar(&options.warmup, "warmup", 0, "unmeasured warmup before each trial")
	flag.IntVar(&options.trials, "trials", 1, "independent measured repetitions per concurrency")
	flag.StringVar(&options.readerCardinality, "reader-cardinality", "", "hot, distributed, high, or cursor-deep; overrides the YAML preset")
	flag.StringVar(&options.topology, "topology", "", "compose-full or compose-middleware-local-services")
	flag.StringVar(&options.pprofDir, "pprof-dir", "", "directory containing profiles captured for this report namespace")
	flag.StringVar(&options.cacheState, "cache-state", "", "natural, cold, l1-warm, l2-warm, or expire-together")
	flag.StringVar(&options.featurePreset, "feature-preset", "", "control, treatment, or cursor-ab; verified against the running KnowPost")
	flag.StringVar(&options.gatewayAuthMode, "gateway-auth-mode", "login", "Gateway token preparation: login or signed")
	flag.StringVar(&options.jwtPrivateKey, "jwt-private-key", "certs/jwt_private.pem", "private key used only by signed load-test Gateway auth")
	flag.StringVar(&options.jwtPublicKey, "jwt-public-key", "certs/jwt_public.pem", "public key paired with the load-test JWT private key")
	flag.StringVar(&options.jwtIssuer, "jwt-issuer", "zhiguang", "issuer for signed load-test Gateway auth")
	flag.StringVar(&options.confirmCleanup, "confirm-cleanup", "", "must exactly match run id before cleanup")
	flag.BoolVar(&options.cleanupDryRun, "cleanup-dry-run", false, "validate and preview cleanup without deleting MySQL or Redis data")
	flag.StringVar(&options.checkpointPath, "checkpoint", "", "mutation checkpoint path; defaults to results/feed-loadtest/checkpoints/<run-id>.json")
	flag.StringVar(&options.confirmMutation, "confirm-mutation", "", "must exactly match run id before creating/restoring mutation checkpoint state")
	flag.Parse()
	return options
}

func parseConcurrencyList(raw string) ([]int, error) {
	parts := strings.Split(raw, ",")
	levels := make([]int, 0, len(parts))
	seen := make(map[int]struct{}, len(parts))
	for _, part := range parts {
		value, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil || value <= 0 {
			return nil, fmt.Errorf("invalid concurrency %q in %q", strings.TrimSpace(part), raw)
		}
		if _, ok := seen[value]; ok {
			return nil, fmt.Errorf("duplicate concurrency %d in %q", value, raw)
		}
		seen[value] = struct{}{}
		levels = append(levels, value)
	}
	if len(levels) == 0 {
		return nil, fmt.Errorf("at least one concurrency is required")
	}
	return levels, nil
}

func loadtestJWTSigner(options cliOptions) (*jwtx.Signer, error) {
	priv, err := jwtx.LoadPrivateKey(options.jwtPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("load load-test JWT private key: %w", err)
	}
	pub, err := jwtx.LoadPublicKey(options.jwtPublicKey)
	if err != nil {
		return nil, fmt.Errorf("load load-test JWT public key: %w", err)
	}
	return jwtx.NewSigner(jwtx.Config{
		PrivateKey: priv,
		PublicKey:  pub,
		Issuer:     options.jwtIssuer,
		AccessTtl:  30 * time.Minute,
		RefreshTtl: 24 * time.Hour,
	})
}

func applyConfigDefaults(cfg *benchmarkConfig) {
	if cfg.Password == "" {
		cfg.Password = "FeedLoadTest!2026"
	}
	if cfg.Topology.Seed == 0 {
		cfg.Topology.Seed = 42
	}
	if cfg.Setup.UserConcurrency <= 0 {
		cfg.Setup.UserConcurrency = 16
	}
	if cfg.Setup.CursorSeedConcurrency <= 0 {
		cfg.Setup.CursorSeedConcurrency = 4
	}
	if cfg.Setup.RelationConcurrency <= 0 {
		cfg.Setup.RelationConcurrency = 32
	}
	if cfg.Setup.WarmupConcurrency <= 0 {
		cfg.Setup.WarmupConcurrency = 1
	}
	if cfg.Setup.StrategySeedPosts <= 0 {
		cfg.Setup.StrategySeedPosts = 5
	}
	if cfg.Setup.CounterSyncTimeout <= 0 {
		cfg.Setup.CounterSyncTimeout = 2 * time.Minute
	}
	if cfg.Setup.CounterPollInterval <= 0 {
		cfg.Setup.CounterPollInterval = 500 * time.Millisecond
	}
	if cfg.Load.RequestTimeout <= 0 {
		cfg.Load.RequestTimeout = 5 * time.Second
	}
	if cfg.Load.Page <= 0 {
		cfg.Load.Page = 1
	}
	if cfg.Load.Size <= 0 {
		cfg.Load.Size = 20
	}
	if len(cfg.Load.ConcurrencySteps) == 0 {
		cfg.Load.ConcurrencySteps = []int{1, 2, 4, 8, 16, 32, 64, 128}
	}
	if cfg.Monitor.Interval <= 0 {
		cfg.Monitor.Interval = time.Second
	}
	if cfg.Monitor.StrategyContainer == "" {
		cfg.Monitor.StrategyContainer = "zg-knowpost"
	}
	if cfg.ReaderCardinality == "" {
		cfg.ReaderCardinality = readerCardinalityHot
	}
	if cfg.Execution.Topology == "" {
		cfg.Execution.Topology = "compose-full"
	}
	if cfg.Execution.LocalStateFile == "" {
		cfg.Execution.LocalStateFile = ".tmp/feed-local/state.json"
	}
	if cfg.Execution.FeaturePreset == "" {
		cfg.Execution.FeaturePreset = "treatment"
	}
	if cfg.Execution.ClientCPULimitPercent <= 0 {
		cfg.Execution.ClientCPULimitPercent = 90
	}
	if cfg.Execution.CacheState == "" {
		cfg.Execution.CacheState = cacheStateNatural
	}
	if cfg.Execution.CacheWarmConcurrency <= 0 {
		cfg.Execution.CacheWarmConcurrency = 64
	}
}

func validateExecutionConfig(cfg executionConfig) error {
	switch cfg.Topology {
	case "compose-full", "compose-middleware-local-services":
	default:
		return fmt.Errorf("execution topology %q must be compose-full or compose-middleware-local-services", cfg.Topology)
	}
	if cfg.ClientCPULimitPercent <= 0 || cfg.ClientCPULimitPercent > 100 {
		return fmt.Errorf("client CPU limit %.2f must be within (0,100]", cfg.ClientCPULimitPercent)
	}
	if cfg.Topology == "compose-middleware-local-services" && strings.TrimSpace(cfg.LocalStateFile) == "" {
		return fmt.Errorf("local service state file is required for topology %s", cfg.Topology)
	}
	if cfg.FeaturePreset != "control" && cfg.FeaturePreset != "treatment" && cfg.FeaturePreset != "cursor-ab" {
		return fmt.Errorf("feature preset %q must be control, treatment, or cursor-ab", cfg.FeaturePreset)
	}
	if _, err := normalizeCacheState(cfg.CacheState); err != nil {
		return err
	}
	if cfg.FeaturePreset == "control" || cfg.FeaturePreset == "cursor-ab" {
		switch cfg.CacheState {
		case cacheStateNatural, cacheStateCold:
		default:
			return fmt.Errorf("%s feature preset cannot claim cache state %q while page cache is off", cfg.FeaturePreset, cfg.CacheState)
		}
	}
	if cfg.CacheWarmConcurrency <= 0 {
		return fmt.Errorf("cache warm concurrency %d must be positive", cfg.CacheWarmConcurrency)
	}
	return nil
}

func applyExecutionTopology(cfg *benchmarkConfig) {
	if cfg.Execution.Topology != "compose-middleware-local-services" {
		return
	}
	cfg.GatewayURL = "http://127.0.0.1:18080"
	cfg.Monitor.StrategyContainer = ""
	cfg.Monitor.PrometheusContainer = ""
	cfg.Monitor.PrometheusURL = "http://127.0.0.1:19104/metrics"
	cfg.Monitor.LocalStateFile = cfg.Execution.LocalStateFile
	cfg.Monitor.DockerContainers = []string{"zg-etcd", "zg-zk", "zg-kafka", "zg-canal", "zg-es"}
}

func connectRPCs(cfg benchmarkConfig) rpcClients {
	return rpcClients{
		knowpost: knowpostpb.NewKnowPostClient(zrpc.MustNewClient(cfg.KnowPostRpc).Conn()),
		relation: relationpb.NewRelationClient(zrpc.MustNewClient(cfg.RelationRpc).Conn()),
		counter:  counterpb.NewUserCounterClient(zrpc.MustNewClient(cfg.CounterRpc).Conn()),
		user:     userpb.NewUserClient(zrpc.MustNewClient(cfg.UserRpc).Conn()),
		auth:     userpb.NewAuthClient(zrpc.MustNewClient(cfg.AuthRpc).Conn()),
	}
}

func setupDataset(ctx context.Context, cfg benchmarkConfig, clients rpcClients, manifestPath, strategy string) (datasetManifest, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(cfg.Password), bcrypt.MinCost)
	if err != nil {
		return datasetManifest{}, fmt.Errorf("hash benchmark password: %w", err)
	}
	fmt.Printf("setup: creating %d users for run %s\n", cfg.Topology.RequiredUsers(), cfg.RunID)
	users, createErr := createBenchmarkUsers(ctx, clients.user, cfg.Topology.RequiredUsers(), cfg.Setup.UserConcurrency, cfg.RunID, string(hash))
	manifest := datasetManifest{
		RunID: cfg.RunID, Seed: cfg.Topology.Seed, Strategy: strategy,
		ReaderCardinality: cfg.ReaderCardinality, Users: users,
	}
	if len(users) > 0 {
		if err := saveManifest(manifestPath, manifest); err != nil {
			return datasetManifest{}, err
		}
	}
	if createErr != nil {
		return manifest, createErr
	}
	ids := make([]int64, len(users))
	for i := range users {
		ids[i] = users[i].ID
	}
	plan, err := buildTopology(ids, cfg.Topology)
	if err != nil {
		return datasetManifest{}, err
	}
	markUserRoles(users, plan)
	manifest.Users = users
	manifest.NormalAuthors = plan.NormalAuthors
	manifest.BigVAuthors = plan.BigVAuthors
	manifest.Readers = plan.Readers
	manifest.FollowerCounts = plan.FollowerCounts
	if err := saveManifest(manifestPath, manifest); err != nil {
		return datasetManifest{}, err
	}
	edges, err := applyTopology(ctx, clients.relation, plan, cfg.Setup.RelationConcurrency)
	if err != nil {
		return datasetManifest{}, err
	}
	fmt.Printf("setup: created or verified %d follow edges\n", edges)
	manifest, err = finalizeDatasetSetup(ctx, cfg, clients.counter, clients.knowpost, manifest, manifestPath)
	if err != nil {
		return manifest, err
	}
	fmt.Printf("setup: manifest=%s\n", manifestPath)
	return manifest, nil
}

func markUserRoles(users []benchmarkUser, plan *topologyPlan) {
	roles := make(map[int64]string, len(users))
	for _, id := range plan.NormalAuthors {
		roles[id] = "normal_author"
	}
	for _, id := range plan.BigVAuthors {
		roles[id] = "bigv_author"
	}
	for _, id := range plan.Readers {
		roles[id] = "reader"
	}
	for i := range users {
		if role := roles[users[i].ID]; role != "" {
			users[i].Role = role
		} else {
			users[i].Role = "follower"
		}
	}
}

func smokeDataset(ctx context.Context, cfg benchmarkConfig, clients rpcClients, manifest datasetManifest, strategy string) error {
	return runSmokeDataset(ctx, cfg, clients, manifest, strategy)
}

func feedContainsAll(page *feedPageResponse, want map[string]struct{}) bool {
	if page == nil {
		return false
	}
	remaining := make(map[string]struct{}, len(want))
	for id := range want {
		remaining[id] = struct{}{}
	}
	for _, item := range page.Items {
		delete(remaining, item.ID)
	}
	return len(remaining) == 0
}

func runBenchmarks(ctx context.Context, cfg benchmarkConfig, clients rpcClients, manifest datasetManifest, options cliOptions) error {
	readers, err := manifestReaders(manifest)
	if err != nil {
		return err
	}
	directFeedClient := newRPCFeedClient(clients.knowpost)
	var feedClient feedReader = directFeedClient
	var publisher publisherClient = clients.knowpost
	if options.entry == "gateway" {
		var signer *jwtx.Signer
		switch strings.ToLower(strings.TrimSpace(options.gatewayAuthMode)) {
		case "login":
			readers, err = loginReaders(ctx, clients.auth, readersToUsers(manifest, readers), cfg.Password, cfg.Setup.UserConcurrency)
		case "signed":
			signer, err = loadtestJWTSigner(options)
			if err == nil {
				readers, err = issueReaderTokens(ctx, signer, readersToUsers(manifest, readers), cfg.Setup.UserConcurrency)
			}
		default:
			return fmt.Errorf("gateway auth mode %q must be login or signed", options.gatewayAuthMode)
		}
		if err != nil {
			return err
		}
		feedClient = newGatewayFeedClient(cfg.GatewayURL, nil)
		if scenarioHasPublish(options.scenario) {
			authorIDs := append(append([]int64(nil), manifest.NormalAuthors...), manifest.BigVAuthors...)
			authorIdentities := make([]readerIdentity, 0, len(authorIDs))
			for _, authorID := range authorIDs {
				authorIdentities = append(authorIdentities, readerIdentity{UserID: authorID})
			}
			var authors []readerIdentity
			if signer != nil {
				authors, err = issueReaderTokens(ctx, signer, readersToUsers(manifest, authorIdentities), cfg.Setup.UserConcurrency)
			} else {
				authors, err = loginReaders(ctx, clients.auth, readersToUsers(manifest, authorIdentities), cfg.Password, cfg.Setup.UserConcurrency)
			}
			if err != nil {
				return err
			}
			tokens := make(map[int64]string, len(authors))
			for _, author := range authors {
				tokens[author.UserID] = author.AccessToken
			}
			publisher = newGatewayPublisherClient(cfg.GatewayURL, nil, tokens)
		}
	} else if options.entry != "rpc" {
		return fmt.Errorf("invalid entry %q: expected rpc or gateway", options.entry)
	}
	cacheFeedClient, err := cacheStateFeedClient(options.entry, directFeedClient)
	if err != nil {
		return err
	}
	var mutationRestorer mutationCheckpointRestorer
	if scenarioHasPublish(options.scenario) {
		controller, openErr := openMutationCheckpointController(
			ctx, cfg, manifest, options.strategy, options.confirmMutation, options.checkpointPath,
		)
		if openErr != nil {
			return fmt.Errorf("open formal mutation checkpoint: %w", openErr)
		}
		mutationRestorer = controller
		defer mutationRestorer.Close()
	}

	trials := options.trials
	if trials <= 0 {
		trials = 1
	}
	for _, concurrency := range cfg.Load.ConcurrencySteps {
		for trial := 1; trial <= trials; trial++ {
			var mutationObservation *mutationTrialObservation
			var warmup func() error
			if options.warmup > 0 {
				warmup = func() error {
					warmupCfg := cfg
					warmupCfg.Load.Requests = 0
					warmupCfg.Load.Duration = options.warmup
					warmupResult, warmupErr := runScenario(
						ctx, warmupCfg, publisher, manifest, feedClient, readers, options.scenario, concurrency, nil,
					)
					if warmupErr != nil {
						return fmt.Errorf("warmup c%d trial %d: %w", concurrency, trial, warmupErr)
					}
					if !reportIsComplete(warmupResult.Stages, nil) {
						return fmt.Errorf("warmup c%d trial %d contained failed or empty stages", concurrency, trial)
					}
					return nil
				}
			}
			if mutationRestorer != nil {
				metadata := mutationRestorer.Metadata()
				mutationObservation = &mutationTrialObservation{
					CheckpointID: metadata.ID, BaselineFingerprint: metadata.BaselineFingerprint,
					CheckpointCreatedAt: metadata.CreatedAt, ToolVersion: metadata.ToolVersion,
					SubjectBinaryIdentity: metadata.SubjectBinaryIdentity,
				}
				if _, pendingErr := writeMutationPreparationReport(
					cfg.ReportDir, manifest, options, concurrency, trial, mutationObservation, nil,
				); pendingErr != nil {
					return fmt.Errorf("write mutation preparation audit c%d trial %d: %w", concurrency, trial, pendingErr)
				}
				prepared, prepareErr := prepareMutationTrial(ctx, mutationRestorer, warmup)
				mutationObservation = &prepared
				if prepareErr != nil {
					wrapped := fmt.Errorf("prepare mutation c%d trial %d: %w", concurrency, trial, prepareErr)
					_, reportErr := writeMutationPreparationReport(
						cfg.ReportDir, manifest, options, concurrency, trial, mutationObservation, wrapped,
					)
					return errors.Join(wrapped, reportErr)
				}
			} else if warmup != nil {
				if warmupErr := warmup(); warmupErr != nil {
					return warmupErr
				}
			}
			stateReaders := cacheStateReaders(options.scenario, readers)
			stateRuntime := &liveCacheStateRuntime{
				redisAddr: cfg.RedisAddr, client: cacheFeedClient, readers: stateReaders,
				allowed:        int64Set(append(append([]int64(nil), manifest.NormalAuthors...), manifest.BigVAuthors...)),
				requestTimeout: cfg.Load.RequestTimeout, concurrency: cfg.Execution.CacheWarmConcurrency,
			}
			if err := runBenchmarkTrial(
				ctx, cfg, publisher, manifest, feedClient, readers, options, concurrency, trial, stateRuntime,
				mutationRestorer, mutationObservation,
			); err != nil {
				return err
			}
		}
	}
	return nil
}

func runBenchmarkTrial(
	ctx context.Context,
	cfg benchmarkConfig,
	publisher publisherClient,
	manifest datasetManifest,
	feedClient feedReader,
	readers []readerIdentity,
	options cliOptions,
	concurrency int,
	trial int,
	stateRuntime cacheStateRuntime,
	mutationRestorer mutationCheckpointRestorer,
	mutationObservation *mutationTrialObservation,
) error {
	failPreparation := func(operation string, err error) error {
		wrapped := fmt.Errorf("%s c%d trial %d: %w", operation, concurrency, trial, err)
		if mutationObservation == nil {
			return wrapped
		}
		_, reportErr := writeMutationPreparationReport(
			cfg.ReportDir, manifest, options, concurrency, trial, mutationObservation, wrapped,
		)
		return errors.Join(wrapped, reportErr)
	}
	authors := append(append([]int64(nil), manifest.NormalAuthors...), manifest.BigVAuthors...)
	feedMetricsPrimed, primeErr := primeFeedMetricsForScenario(
		ctx, options.scenario, feedClient, readers, int64Set(authors), cfg.Load.RequestTimeout, cfg.Load.Size,
	)
	if primeErr != nil {
		return failPreparation("prime Feed metrics", primeErr)
	}
	evidenceCtx, cancelEvidence := context.WithTimeout(context.Background(), 8*time.Second)
	environmentEvidence, environmentEvidenceErr := captureRunEnvironmentEvidence(evidenceCtx, cfg, manifest)
	cancelEvidence()
	monitor := startResourceMonitor(ctx, cfg.Monitor, cfg.RedisAddr)
	var cursorPreparationForRun *cursorPreparation
	if targetPage, ok := fixedCursorTargetPage(options.scenario); ok {
		prepared, prepareErr := prepareCursorTargets(
			ctx,
			feedClient,
			readers,
			benchmarkCursorTargetSpec(cfg, manifest, targetPage, concurrency),
		)
		if prepareErr != nil {
			monitor.Stop(context.Background())
			return failPreparation("prepare cursor targets", prepareErr)
		}
		cursorPreparationForRun = &prepared
	}
	cacheObservation, cachePrepareErr := prepareCacheState(ctx, cfg.Execution.CacheState, stateRuntime)
	if cachePrepareErr != nil {
		monitor.Stop(context.Background())
		return failPreparation("prepare cache state", cachePrepareErr)
	}
	boundaryCtx, cancelBoundary := context.WithTimeout(context.Background(), 3*time.Second)
	boundaryDuration, boundaryErr := monitor.ResetCounterBaselines(boundaryCtx)
	cancelBoundary()
	cacheObservation.BoundaryDuration = boundaryDuration
	if boundaryErr != nil {
		monitor.Stop(context.Background())
		return failPreparation("reset measurement baseline", boundaryErr)
	}
	if err := validateCacheStateMeasurementBoundary(cacheObservation, boundaryDuration); err != nil {
		monitor.Stop(context.Background())
		return failPreparation("validate cache state measurement boundary", err)
	}
	clientCPU, clientCPUStartErr := startClientCPUTimer()
	started := time.Now()
	runResult, runErr := runScenario(
		ctx, cfg, publisher, manifest, feedClient, readers, options.scenario, concurrency, cursorPreparationForRun,
	)
	stages := runResult.Stages
	if mutationObservation != nil {
		mutationObservation.PublishedPostIDs = append([]int64(nil), runResult.PublishedPostIDs...)
	}
	elapsed := time.Since(started)
	var clientCPUSample resourceSample
	var clientCPUStopErr error
	if clientCPUStartErr == nil {
		clientCPUSample, clientCPUStopErr = clientCPU.Stop()
	}
	var recovery *recoveryObservation
	if scenarioHasPublish(options.scenario) {
		recoveryStarted := time.Now()
		finalMetrics, recoveryErr := waitForMutationKafkaDrain(ctx, cfg.Monitor, 3*time.Minute)
		recovery = &recoveryObservation{
			KafkaDrain: time.Since(recoveryStarted), FinalMetrics: finalMetrics,
			Complete: recoveryErr == nil,
		}
		if recoveryErr != nil {
			recovery.Error = recoveryErr.Error()
		}
	}
	resources, missing := monitor.Stop(ctx)
	if clientCPUStartErr != nil || clientCPUStopErr != nil {
		missing = append(missing, "client_cpu")
	} else {
		resources = append(resources, clientCPUSample)
		if clientCPUSample.Values["cpu_percent_normalized"] >= cfg.Execution.ClientCPULimitPercent {
			missing = append(missing, "client_cpu_saturated")
		}
	}
	feedMetrics, feedMetricsOK := summarizeFeedMetrics(resources, successfulReadCount(stages))
	if !feedMetricsOK && !containsString(missing, "feed_metrics") {
		missing = append(missing, "feed_metrics")
	}
	if feedMetrics != nil && feedMetrics.Degraded && !containsString(missing, "feed_degraded") {
		missing = append(missing, "feed_degraded")
	}
	redisMetrics, redisMetricsOK := summarizeRedisMetrics(resources)
	if !redisMetricsOK && !containsString(missing, "redis_metrics_delta") {
		missing = append(missing, "redis_metrics_delta")
	}
	if redisMetrics != nil && redisMetrics.EvictedKeysDelta > 0 {
		missing = append(missing, "redis_evictions")
	}
	if redisMetrics != nil && redisMetrics.RejectedConnectionsDelta > 0 {
		missing = append(missing, "redis_rejections")
	}
	if redisSafetyEpochChanged(options.scenario, redisMetrics) {
		missing = append(missing, "redis_safety_epoch_changed")
	}
	if recovery != nil && !recovery.Complete {
		missing = append(missing, "kafka_recovery")
	}
	if runErr != nil {
		missing = append(missing, "run_error")
	}
	missing = append(missing, paginationMissingMetrics(stages)...)
	if environmentEvidenceErr != nil {
		missing = append(missing, "environment_evidence")
	}
	environment := map[string]string{
		"gateway_url":                   cfg.GatewayURL,
		"go_version":                    strings.TrimSpace(runtimeVersion()),
		"effective_strategy":            options.strategy,
		"strategy_source":               "unavailable",
		"dataset_run_id":                manifest.RunID,
		"reader_cardinality":            cfg.ReaderCardinality,
		"reader_count":                  fmt.Sprintf("%d", len(readers)),
		"topology":                      cfg.Execution.Topology,
		"cache_state":                   cfg.Execution.CacheState,
		"feature_preset":                cfg.Execution.FeaturePreset,
		"gateway_auth_mode":             strings.ToLower(strings.TrimSpace(options.gatewayAuthMode)),
		"cache_warm_readers":            fmt.Sprintf("%d", cacheObservation.WarmReaders),
		"cache_warm_duration":           cacheObservation.WarmDuration.String(),
		"measurement_boundary_duration": cacheObservation.BoundaryDuration.String(),
		"request_timeout":               cfg.Load.RequestTimeout.String(),
		"page":                          fmt.Sprintf("%d", cfg.Load.Page),
		"size":                          fmt.Sprintf("%d", cfg.Load.Size),
		"configured_duration":           cfg.Load.Duration.String(),
		"configured_requests":           fmt.Sprintf("%d", cfg.Load.Requests),
		"warmup":                        options.warmup.String(),
		"feed_metrics_primed":           fmt.Sprintf("%t", feedMetricsPrimed),
	}
	if host, err := os.Hostname(); err == nil {
		environment["client_host"] = host
	} else {
		missing = append(missing, "environment_evidence")
	}
	if cfg.Execution.PprofDir != "" {
		environment["pprof_dir"] = cfg.Execution.PprofDir
	}
	if mutationObservation != nil {
		environment["mutation_checkpoint_id"] = mutationObservation.CheckpointID
		environment["mutation_baseline_fingerprint"] = mutationObservation.BaselineFingerprint
		environment["mutation_checkpoint_created_at"] = mutationObservation.CheckpointCreatedAt.Format(time.RFC3339Nano)
		environment["mutation_checkpoint_tool_version"] = mutationObservation.ToolVersion
	}
	for key, value := range environmentEvidence {
		environment[key] = value
	}
	featureCtx, cancelFeatures := context.WithTimeout(context.Background(), 5*time.Second)
	runtimeState, featureErr := readFeedRuntimeState(featureCtx, cfg.Execution, cfg.Monitor)
	cancelFeatures()
	if featureErr != nil {
		if !containsString(missing, "feed_runtime_flags") {
			missing = append(missing, "feed_runtime_flags")
		}
	} else {
		featureFlags := runtimeState.Features
		logging := runtimeState.Logging
		environment["strategy_source"] = runtimeState.Source
		environment["observability_enabled"] = fmt.Sprintf("%t", featureFlags.ObservabilityEnabled)
		environment["route_snapshot_enabled"] = fmt.Sprintf("%t", featureFlags.RouteSnapshotEnabled)
		environment["relation_epoch_consumer_enabled"] = fmt.Sprintf("%t", featureFlags.RelationEpochEnabled)
		environment["safety_epoch_consumer_enabled"] = fmt.Sprintf("%t", featureFlags.SafetyEpochEnabled)
		environment["combined_pipeline_enabled"] = fmt.Sprintf("%t", featureFlags.CombinedPipelineEnabled)
		environment["cursor_pagination_enabled"] = fmt.Sprintf("%t", featureFlags.CursorPaginationEnabled)
		environment["page_cache_mode"] = featureFlags.PageCacheMode
		environment["logging_preset"] = logging.Preset
		environment["logging_level"] = logging.Level
		environment["logging_stat"] = fmt.Sprintf("%t", logging.Stat)
		environment["gateway_access_log"] = fmt.Sprintf("%t", logging.GatewayAccessLog)
		environment["rpc_stat_middleware"] = fmt.Sprintf("%t", logging.RPCStatMiddleware)
		environment["sql_statement_info"] = fmt.Sprintf("%t", logging.SQLStatementInfo)
		environment["logging_persistence"] = logging.Persistence
	}
	environment["compatibility_fingerprint"] = environmentCompatibilityFingerprint(environment)
	report := benchmarkReport{
		RunID:          effectiveReportRunID(manifest.RunID, options.reportRunID),
		Strategy:       options.strategy,
		Entry:          options.entry,
		Scenario:       fmt.Sprintf("%s-c%d", options.scenario, concurrency),
		StartedAt:      started,
		Duration:       elapsed,
		Concurrency:    concurrency,
		Trial:          trial,
		ExpectedTrials: normalizedExpectedTrials(options.trials),
		Complete:       reportIsComplete(stages, missing),
		MissingMetrics: missing,
		Environment:    environment,
		Stages:         stages,
		FeedMetrics:    feedMetrics,
		RedisMetrics:   redisMetrics,
		Resources:      resources,
		Recovery:       recovery,
		Mutation:       mutationObservation,
		Notes:          []string{"SLA values are reference lines, not pass/fail gates."},
	}
	if runErr != nil {
		report.Notes = append(report.Notes, "Scenario error: "+runErr.Error())
	}
	var paths reportPaths
	var err error
	if mutationRestorer != nil {
		restoreCtx, cancelRestore := context.WithTimeout(context.Background(), 4*time.Minute)
		paths, err = writeMutationReportWithRestore(restoreCtx, cfg.ReportDir, &report, mutationRestorer)
		cancelRestore()
	} else {
		paths, err = writeReportFiles(cfg.ReportDir, report)
	}
	if err != nil {
		return errors.Join(runErr, err)
	}
	fmt.Printf("run: strategy=%s entry=%s scenario=%s concurrency=%d trial=%d report=%s\n",
		options.strategy, options.entry, options.scenario, concurrency, trial, paths.Markdown)
	if runErr != nil {
		return runErr
	}
	return benchmarkReportCompletionError(report)
}

func benchmarkReportCompletionError(report benchmarkReport) error {
	if report.Complete {
		return nil
	}
	return fmt.Errorf("benchmark report is incomplete: missing_metrics=%v", report.MissingMetrics)
}

func paginationMissingMetrics(stages []stageResult) []string {
	missingOracle := false
	correctnessFailure := false
	for _, stage := range stages {
		observation := stage.Pagination
		if observation == nil || !strings.Contains(observation.Mode, "cursor") {
			continue
		}
		if observation.OracleHash == "" {
			missingOracle = true
		}
		if observation.ObservedHash != observation.OracleHash || observation.DuplicateItems > 0 ||
			observation.OracleMismatches > 0 || observation.CursorLoops > 0 || observation.EarlyTerminations > 0 {
			correctnessFailure = true
		}
	}
	result := make([]string, 0, 2)
	if missingOracle {
		result = append(result, "cursor_oracle")
	}
	if correctnessFailure {
		result = append(result, "cursor_correctness")
	}
	return result
}

func normalizedExpectedTrials(value int) int {
	if value <= 0 {
		return 1
	}
	return value
}

func effectiveReportRunID(datasetRunID, requested string) string {
	if value := strings.TrimSpace(requested); value != "" {
		return value
	}
	return datasetRunID
}

func successfulReadCount(stages []stageResult) int64 {
	var total int64
	for _, stage := range stages {
		if stage.Name == "read" {
			total += stage.Success
		}
	}
	return total
}

func containsString(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

type scenarioRunResult struct {
	Stages           []stageResult
	PublishedPostIDs []int64
}

func runScenario(
	ctx context.Context,
	cfg benchmarkConfig,
	publisher publisherClient,
	manifest datasetManifest,
	feedClient feedReader,
	readers []readerIdentity,
	scenario string,
	concurrency int,
	prepared *cursorPreparation,
) (scenarioRunResult, error) {
	readSpec := loadSpec{
		Requests:       cfg.Load.Requests,
		Duration:       cfg.Load.Duration,
		Concurrency:    concurrency,
		RequestTimeout: cfg.Load.RequestTimeout,
		Page:           cfg.Load.Page,
		Size:           cfg.Load.Size,
		Seed:           cfg.Topology.Seed,
	}
	authors := append(append([]int64(nil), manifest.NormalAuthors...), manifest.BigVAuthors...)
	readSpec.AllowedCreators = int64Set(authors)
	readSpec.RequireItems = true
	if targetPage, ok := offsetTargetPage(scenario); ok {
		readSpec.Page = int32(targetPage)
		result, err := runReadStage(ctx, feedClient, readers, readSpec)
		if result.Pagination == nil {
			result.Pagination = &paginationObservation{Mode: "page", TargetPage: targetPage}
		}
		return scenarioRunResult{Stages: []stageResult{result}}, err
	}
	if targetPage, ok := fixedCursorTargetPage(scenario); ok {
		cursorSpec := benchmarkCursorTargetSpec(cfg, manifest, targetPage, concurrency)
		if prepared == nil {
			cursorPreparation, err := prepareCursorTargets(ctx, feedClient, readers, cursorSpec)
			if err != nil {
				return scenarioRunResult{}, err
			}
			prepared = &cursorPreparation
		}
		result, err := runFixedCursorStage(ctx, feedClient, readers, *prepared, cursorSpec)
		return scenarioRunResult{Stages: []stageResult{result}}, err
	}
	if targetPage, ok := sequentialCursorTargetPage(scenario); ok {
		result, err := runSequentialCursorStage(
			ctx, feedClient, readers, benchmarkCursorTargetSpec(cfg, manifest, targetPage, concurrency),
		)
		return scenarioRunResult{Stages: []stageResult{result}}, err
	}
	if targetPage, ok := sequentialOffsetTargetPage(scenario); ok {
		result, err := runSequentialOffsetStage(
			ctx, feedClient, readers, benchmarkCursorTargetSpec(cfg, manifest, targetPage, concurrency),
		)
		return scenarioRunResult{Stages: []stageResult{result}}, err
	}
	switch scenario {
	case "hot-read":
		result, err := runReadStage(ctx, feedClient, readers[:1], readSpec)
		return scenarioRunResult{Stages: []stageResult{result}}, err
	case "distributed-read", "burst-read", "steady-read":
		result, err := runReadStage(ctx, feedClient, readers, readSpec)
		return scenarioRunResult{Stages: []stageResult{result}}, err
	case "deep-page":
		readSpec.Page = maxInt32(readSpec.Page, 5)
		readSpec.RequireItems = false
		result, err := runReadStage(ctx, feedClient, readers, readSpec)
		return scenarioRunResult{Stages: []stageResult{result}}, err
	case "publish", "burst-publish":
		stages, postIDs, err := runPublishStage(ctx, publisher, authors, publishSpec{
			Requests:       cfg.Load.Requests,
			Duration:       cfg.Load.Duration,
			Concurrency:    concurrency,
			RequestTimeout: cfg.Load.RequestTimeout,
			RunID:          manifest.RunID,
			Seed:           cfg.Topology.Seed,
		})
		return scenarioRunResult{Stages: stages, PublishedPostIDs: postIDs}, err
	case "mixed-90-10":
		return runMixedScenario(ctx, cfg, publisher, manifest, feedClient, readers, concurrency, 90)
	case "mixed-80-20":
		return runMixedScenario(ctx, cfg, publisher, manifest, feedClient, readers, concurrency, 80)
	default:
		return scenarioRunResult{}, fmt.Errorf("unknown scenario %q", scenario)
	}
}

func benchmarkCursorTargetSpec(
	cfg benchmarkConfig,
	manifest datasetManifest,
	targetPage int,
	concurrency int,
) cursorTargetSpec {
	authors := append(append([]int64(nil), manifest.NormalAuthors...), manifest.BigVAuthors...)
	var oracle map[int64][]string
	oracleHash := ""
	if manifest.CursorDeep != nil {
		oracleIDs := make([]string, len(manifest.CursorDeep.Oracle))
		for i, post := range manifest.CursorDeep.Oracle {
			oracleIDs[i] = strconv.FormatInt(post.PostID, 10)
		}
		oracle = make(map[int64][]string, len(manifest.Readers))
		for _, readerID := range manifest.Readers {
			oracle[readerID] = oracleIDs
		}
		oracleHash = manifest.CursorDeep.OracleHash
	}
	return cursorTargetSpec{
		TargetPage:      targetPage,
		Size:            cfg.Load.Size,
		Timeout:         cfg.Load.RequestTimeout,
		Requests:        cfg.Load.Requests,
		Duration:        cfg.Load.Duration,
		Concurrency:     concurrency,
		Seed:            cfg.Topology.Seed,
		RequireItems:    true,
		AllowedCreators: int64Set(authors),
		Oracle:          oracle,
		OracleHash:      oracleHash,
	}
}

func scenarioHasPublish(scenario string) bool {
	switch scenario {
	case "publish", "burst-publish", "mixed-90-10", "mixed-80-20":
		return true
	default:
		return false
	}
}

func primeFeedMetricsForScenario(
	ctx context.Context,
	scenario string,
	client feedReader,
	readers []readerIdentity,
	allowedCreators map[int64]struct{},
	requestTimeout time.Duration,
	size int32,
) (bool, error) {
	if scenario != "publish" && scenario != "burst-publish" {
		return false, nil
	}
	if len(readers) == 0 {
		return false, fmt.Errorf("Feed metric priming requires at least one reader")
	}
	stage, err := runReadStage(ctx, client, readers[:1], loadSpec{
		Requests: 1, Concurrency: 1, RequestTimeout: requestTimeout,
		Page: 1, Size: size, Seed: 1, RequireItems: true, AllowedCreators: allowedCreators,
	})
	if err != nil {
		return false, err
	}
	if stage.Total != 1 || stage.Success != 1 || stage.Failed != 0 {
		return false, fmt.Errorf("Feed metric priming request failed: total=%d success=%d failed=%d",
			stage.Total, stage.Success, stage.Failed)
	}
	return true, nil
}

func redisSafetyEpochChanged(scenario string, metrics *redisMetricsSummary) bool {
	return metrics != nil && !scenarioHasPublish(scenario) && metrics.SafetyEpochStart != metrics.SafetyEpochEnd
}

func runMixedScenario(ctx context.Context, cfg benchmarkConfig, publisher publisherClient, manifest datasetManifest, feedClient feedReader, readers []readerIdentity, concurrency, readPercent int) (scenarioRunResult, error) {
	authors := append(append([]int64(nil), manifest.NormalAuthors...), manifest.BigVAuthors...)
	stages, postIDs, err := runMixedStage(ctx, publisher, feedClient, authors, readers, mixedSpec{
		Requests: cfg.Load.Requests, Duration: cfg.Load.Duration, Concurrency: concurrency,
		ReadPercent: readPercent, RequestTimeout: cfg.Load.RequestTimeout,
		Page: cfg.Load.Page, Size: cfg.Load.Size, RunID: manifest.RunID, Seed: cfg.Topology.Seed,
		RequireItems: cfg.Load.Page == 1, AllowedCreators: int64Set(authors),
	})
	return scenarioRunResult{Stages: stages, PublishedPostIDs: postIDs}, err
}

func int64Set(values []int64) map[int64]struct{} {
	result := make(map[int64]struct{}, len(values))
	for _, value := range values {
		result[value] = struct{}{}
	}
	return result
}

func manifestReaders(manifest datasetManifest) ([]readerIdentity, error) {
	if len(manifest.Readers) == 0 {
		return nil, fmt.Errorf("manifest has no readers")
	}
	result := make([]readerIdentity, 0, len(manifest.Readers))
	for _, id := range manifest.Readers {
		result = append(result, readerIdentity{UserID: id})
	}
	return result, nil
}

func readersToUsers(manifest datasetManifest, readers []readerIdentity) []benchmarkUser {
	wanted := make(map[int64]struct{}, len(readers))
	for _, reader := range readers {
		wanted[reader.UserID] = struct{}{}
	}
	users := make([]benchmarkUser, 0, len(readers))
	for _, user := range manifest.Users {
		if _, ok := wanted[user.ID]; ok {
			users = append(users, user)
		}
	}
	sort.Slice(users, func(i, j int) bool { return users[i].ID < users[j].ID })
	return users
}

func runtimeVersion() string {
	return runtime.Version()
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxInt32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
