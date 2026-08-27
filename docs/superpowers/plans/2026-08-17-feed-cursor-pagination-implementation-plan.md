# Feed Cursor/seek 深分页实施计划

日期：2026-08-17

依据：`docs/superpowers/specs/2026-08-15-feed-hybrid-read-performance-evolution-design.md` 第 19 节

状态：C0–C5 已完成。专用深数据、真实中间件/本地服务冒烟、RPC/Gateway 正式 A/B、连续滚动严格配对、Runtime/GC 遥测和最终报告均已通过；建议保持开关可回滚并先灰度启用

## 1. 范围与执行规则

本计划只实现个性化 `GetUserFeed` 的 Cursor/seek 分页、专用深分页数据集和页码/Cursor A/B。以下内容不在范围内：Feed Head 物化、任意深页 PageCache、写入/混合 checkpoint/restore、公共 Feed 和 Relation 分页改造。

执行规则：

- 当前分支为 `codex/feed-wp11`，不在 `master/main` 上直接开发。
- 工作区已有大量未提交修改；开始每个工作包前保存相关文件的 `git status` 和 diff，不覆盖、不回滚既有改动。
- 所有生产行为先写最小失败测试并确认失败原因，再写最小实现；生成代码除外，但生成后的契约必须先有失败测试。
- 新能力默认关闭；关闭后旧 page API 的请求、响应和性能路径保持不变。
- Cursor 精确读取任一必需 Redis 命令失败时整页失败，不返回可能导致永久漏帖的部分页。
- C1–C3 正确性回归未全绿前不启动服务或长压；C4 工具验证未通过不创建正式报告。
- Docker 镜像拉取/构建异常按镜像源问题处理；优先复用既有中间件容器并在宿主机运行项目服务，不反复重建依赖镜像。
- 本计划不自动提交脏文件；只有用户明确要求时才按可审计范围提交。

## 2. 成功标准

功能成功：

- Proto/Gateway 向后兼容，旧客户端可忽略 `nextCursor`。
- Cursor v1 非法输入 fail closed；开关关闭时不生成 Cursor、带 Cursor 请求明确失败。
- 静态 1,500 候选 oracle 连续翻页无重复、无遗漏，排序与 `(sortTime DESC, postID DESC)` 完全一致。
- 同秒 tie group、插入、删除、转私密、取关、回填和依赖失败测试通过。

性能成功：

- page1 同口径 QPS/P95 回退不超过 5%。
- Cursor page50 的 Redis 返回成员或详情装载 ID 相对 page50 至少降低 70%。
- Cursor page50 P95 不高于 Cursor page5 的 1.5 倍。
- page20/page50 P95 改善至少 30%，或 sequential-50 的 CPU/request、返回成员或详情装载量降低至少 50%，且总耗时/QPS 不回退。

## 3. 工作包总览

| 工作包 | 目标 | 主要门禁 |
|---|---|---|
| C0 | 锁定基线、接口和 oracle | 现有 Feed/Proto/Gateway 测试可复现 |
| C1 | Cursor codec 与 Proto 契约 | codec 边界、生成代码和兼容测试通过 |
| C2 | 精确 Redis seek 与 FeedReader | 同秒、跨来源、失败原子性和 oracle 通过 |
| C3 | Logic、配置、Gateway 接线 | 可见性回填、开关、PageCache兼容通过 |
| C4 | loadtest Cursor/深数据能力 | 固定深度与连续翻页报告可校验 |
| C5 | 专用数据集和正式 A/B | 报告完整、三轮可比、结论满足决策规则 |

## 4. C0：基线与契约锁定

### 4.1 只读快照

记录：

```powershell
git status --short
git diff -- proto/knowpost/knowpost.proto services/knowpost/rpc/knowpost services/knowpost/rpc/internal/feed services/knowpost/rpc/internal/logic/knowpost/getuserfeedlogic.go services/gateway/internal/handler/router.go cmd/loadtest
git rev-parse --abbrev-ref HEAD
git rev-parse HEAD
go version
goctl --version
protoc --version
```

### 4.2 相关基线

先运行：

```powershell
go test ./services/knowpost/rpc/internal/feed/... -count=1
go test ./services/knowpost/rpc/internal/logic/knowpost -count=1
go test ./services/knowpost/rpc/knowpost ./services/knowpost/rpc/knowpostclient ./services/knowpost/rpc/client/knowpost -count=1
go test ./services/gateway/internal/handler -count=1
go test ./cmd/loadtest -count=1
```

若存在与 Cursor 无关的失败，保存原始输出并先判断是否属于当前脏工作区；不能用修改 Cursor 代码掩盖。

### 4.3 C0 产物

- 本计划和已批准 Spec。
- 明确 Proto 只追加字段 4/5，不改 RPC 方法名。
- 明确 Cursor 规范全序与 seek 不等式。
- 记录改造前 page1 与深页微基准；正式数据不足时只作为代码基线，不声称容量结论。

## 5. C1：Cursor codec 与 Proto 契约

### 5.1 Cursor codec RED

新增：

- `services/knowpost/rpc/internal/feedcursor/cursor_test.go`

测试期望 API：

```go
type Cursor struct {
    SortTime int64
    PostID   int64
}

func Encode(Cursor) (string, error)
func Decode(string) (Cursor, error)
```

覆盖：

- v1 round-trip 和重复编码结果完全一致。
- 最大 256 字节。
- 非规范 Base64URL、截断 JSON、尾随值、未知字段、未知版本。
- `SortTime<=0`、`PostID<=0`。
- 解码后重新编码必须等于原字符串，拒绝同义但非规范的输入。

运行并确认因实现缺失而失败：

```powershell
go test ./services/knowpost/rpc/internal/feedcursor -run TestCursor -count=1
```

### 5.2 Cursor codec GREEN

新增：

- `services/knowpost/rpc/internal/feedcursor/cursor.go`

实现边界：

- 使用标准库 `encoding/base64` 的 Raw URL encoding 和紧凑 JSON 字段 `v/t/p`。
- Decoder 拒绝未知字段和尾随数据。
- 不引入第三方库、HMAC、加密或配置依赖。

验证：

```powershell
go test ./services/knowpost/rpc/internal/feedcursor -count=10
go vet ./services/knowpost/rpc/internal/feedcursor
```

### 5.3 Proto RED

新增：

- `services/knowpost/rpc/knowpost/cursor_contract_test.go`

测试构造 `GetUserFeedReq.Cursor` 和 `FeedPage.NextCursor`，执行 protobuf marshal/unmarshal 并断言字段保留。先运行并确认因字段不存在而编译失败。

### 5.4 Proto GREEN 与可重复生成

修改：

- `proto/knowpost/knowpost.proto`
- `services/knowpost/rpc/knowpost/knowpost.pb.go`（由 protoc 生成）
- `services/knowpost/rpc/knowpost/knowpost_grpc.pb.go`（只允许生成器机械更新）
- `scripts/gen.sh`
- `scripts/gen.bat`

先把生成结果输出到 `.tmp/cursor-proto-gen` 并检查路径和 diff，再覆盖正式生成文件。推荐生成命令：

```powershell
protoc --go_out=services/knowpost/rpc --go-grpc_out=services/knowpost/rpc proto/knowpost/knowpost.proto
```

若当前 `go_package` 导致路径不匹配，停止并修正生成参数，不移动生成文件冒充正确生成。`rpc/knowpostclient` 与 `rpc/client/knowpost` 通过类型别名自动获得新增字段，禁止手改生成客户端。

验证：

```powershell
go test ./services/knowpost/rpc/knowpost ./services/knowpost/rpc/knowpostclient ./services/knowpost/rpc/client/knowpost -count=1
go vet ./services/knowpost/rpc/knowpost ./services/knowpost/rpc/knowpostclient ./services/knowpost/rpc/client/knowpost
```

## 6. C2：精确 Redis seek 与 FeedReader

### 6.1 Adapter RED

修改测试：

- `services/knowpost/rpc/internal/feed/redis_adapter_test.go`

新增测试用请求/结果结构，覆盖：

- 一个 Pipeline 内的 `score == cursorTime` 完整 tie range。
- `score < cursorTime` 的 exclusive older range 和 count 限制。
- 多 key 结果顺序与请求下标一致。
- 单个命令错误保留在对应结果中。
- Context 取消。

先运行定向测试，确认因接口/实现缺失而失败。

### 6.2 Adapter GREEN

修改：

- `services/knowpost/rpc/internal/feed/redis_adapter.go`
- `services/knowpost/rpc/internal/feed/writer.go`（只增加独立请求/结果类型时修改，不扩展 Writer `RedisClient`）

新增窄接口 `redisCursorBatchReader`，生产 Adapter 用 go-redis Pipeline 实现 score range 批量读取。每个结果必须带返回成员和独立错误；调用方负责整页错误策略。

验证：

```powershell
go test ./services/knowpost/rpc/internal/feed -run 'TestRedisAdapter.*Cursor|TestRedisAdapter.*Score' -count=20
```

### 6.3 FeedReader RED

修改：

- `services/knowpost/rpc/internal/feed/reader_test.go`
- 必要时新增 `services/knowpost/rpc/internal/feed/cursor_reader_test.go`

期望 API：

```go
func (s *FeedReadSnapshot) GetFeedPosts(ctx context.Context, page, size int) ([]Post, bool, error)
func (s *FeedReadSnapshot) GetFeedAfter(ctx context.Context, after feedcursor.Cursor, limit int) ([]Post, bool, error)
```

逐个 RED/GREEN：

1. `GetFeedPosts` 与旧 `GetFeed` ID/hasMore 完全等价。
2. 单 Inbox 不同时间的 seek。
3. 相同时间按数值 post ID seek；特意使用位数不同的 ID 证明不能依赖字符串顺序。
4. `score<T` 限量读取在更早同秒组中间截断时，第二轮 Pipeline 补齐完整边界组。
5. Inbox + 5 BigV 全局归并、跨来源去重和 Top-N。
6. tie group 100/1000 上界无遗漏。
7. 任一必需 range 失败则整页失败，不返回部分 Post。
8. Adapter 不支持精确 Cursor 接口时返回确定错误，旧 page 仍成功。
9. Context 取消后停止后续批次。
10. `Prepare` 在回填式重复读取中只调用一次 Relation/Counter。

### 6.4 FeedReader GREEN

修改：

- `services/knowpost/rpc/internal/feed/reader.go`

实现：

- 旧 `GetFeed` 委托 `GetFeedPosts` 再映射 ID，避免两套页码逻辑。
- Cursor 第一轮为每个来源生成 cursor-tie/older 两条范围请求；older 满额时第二轮按来源补齐 boundary-tie，两个 Pipeline 都按现有 batch 上限分批。
- 应用层按数值 seek 不等式过滤，归并后维护 `limit+1` Top-N，返回 limit 条与 raw hasMore。
- Cursor 路径严格错误；旧 page 的既有 Inbox/BigV 部分降级保持不变。
- 不改 FeedWriter score 和容量常量。

验证：

```powershell
go test ./services/knowpost/rpc/internal/feed -run 'Cursor|GetFeedPosts|Pagination|CombinedPipeline' -count=50
go test ./services/knowpost/rpc/internal/feed -count=1
go vet ./services/knowpost/rpc/internal/feed
```

## 7. C3：Logic、配置和 Gateway

### 7.1 配置 RED/GREEN

修改测试：

- `services/knowpost/rpc/internal/config/config_test.go`
- `services/knowpost/rpc/internal/svc/feed_strategy_test.go` 或新增 Cursor 接线测试

修改实现：

- `services/knowpost/rpc/internal/config/config.go`
- `services/knowpost/rpc/internal/svc/servicecontext.go`
- `services/knowpost/rpc/etc/knowpost.yaml`
- `services/knowpost/cmd/knowpost/etc/knowpost.yaml`
- `services/knowpost/cmd/knowpost/etc/knowpost-docker.yaml`
- `deploy/compose/docker-compose.dev.yml`

增加 `Feed.CursorPagination.Enabled`，默认 false；Compose 只透传明确环境变量，默认仍 false。配置测试先失败再实现。

### 7.2 Logic RED

修改：

- `services/knowpost/rpc/internal/logic/knowpost/getuserfeedlogic_test.go`
- `services/knowpost/rpc/internal/logic/knowpost/getuserfeed_observer_test.go`

逐个覆盖：

1. 开关关闭：page 成功且 `NextCursor` 为空；Cursor 请求返回 InvalidArgument/明确禁用错误。
2. 开关开启：首页 page 结果带精确 `NextCursor`；旧 items/page/size/hasMore 不变。
3. Cursor 请求忽略 page，响应 page=0。
4. 非法 Cursor 返回 InvalidArgument，不调用 Relation、Redis 或 DB。
5. 删除/缺失/转私密后有限回填；Cursor 来自最后一个真实返回的可见 Post。
6. 返回 size+1 个可见候选时 `hasMore=true`；末页 `hasMore=false,nextCursor=""`。
7. PageCache 首页缓存包含稳定 `NextCursor`；非空 Cursor 永远 bypass PageCache。
8. safety/relation 当前态仍在返回前生效，不因旧 Cursor 越权。
9. 指标区分 page/cursor，非法游标不记录高基数标签。

### 7.3 Logic GREEN

修改：

- `services/knowpost/rpc/internal/logic/knowpost/getuserfeedlogic.go`
- `services/knowpost/rpc/internal/feed/observer.go`
- `services/knowpost/rpc/internal/feed/observer_test.go`

将候选处理内部统一为 `[]feed.Post`，保留排序键直到可见性过滤结束。Cursor 候选不足时按 21、42、84……扩大，但复用同一 `FeedReadSnapshot`，最大不超过 5,000。非法 Cursor 在 Prepare 前失败。

### 7.4 Gateway RED/GREEN

修改：

- `services/gateway/internal/handler/router_test.go`
- `services/gateway/internal/handler/router.go`

测试先覆盖：

- `cursor` 原样透传，用户 ID 仍来自认证 Context。
- Cursor 模式允许省略 page，size 校验不变。
- JSON 返回 `nextCursor`。
- 超长 Cursor 在 RPC 前返回 400；结构合法性由 RPC 统一验证。
- 无 Cursor 的旧 URL 和响应字段继续可用。

验证：

```powershell
go test ./services/knowpost/rpc/internal/config ./services/knowpost/rpc/internal/svc -count=1
go test ./services/knowpost/rpc/internal/logic/knowpost -run 'GetUserFeed.*Cursor|GetUserFeed.*PageCache|GetUserFeed.*Backfill' -count=30
go test ./services/gateway/internal/handler -run 'FollowingFeed|Cursor' -count=30
go test ./services/knowpost/... ./services/gateway/... -count=1
go vet ./services/knowpost/... ./services/gateway/...
```

## 8. C4：loadtest Cursor 与报告

### 8.1 客户端契约 RED/GREEN

修改：

- `cmd/loadtest/runner.go`
- `cmd/loadtest/rpc_client.go`
- `cmd/loadtest/gateway_client.go`
- 对应 `_test.go`

将读取参数收敛为带 `Page/Size/Cursor` 的请求结构，响应增加 `NextCursor`。先用测试证明 RPC 字段、Gateway query 和 JSON 字段正确，再改实现；现有 page 场景调用保持兼容 helper，避免一次重写全部 runner。

### 8.2 深分页场景 RED/GREEN

新增或修改：

- `cmd/loadtest/cursor_runner.go`
- `cmd/loadtest/cursor_runner_test.go`
- `cmd/loadtest/feed_loadtest.go`
- `cmd/loadtest/report.go`
- `cmd/loadtest/report_test.go`
- `cmd/loadtest/monitor.go`

新增场景：

- `cursor-page5`、`cursor-page20`、`cursor-page50`：warmup 生成并校验每个 reader 的目标前序 Cursor，只测目标页。
- `sequential-page-50`：每个虚拟用户按页码从第 1 页连续读取到第 50 页，作为 Offset 深分页 Control，全部请求计入结果。
- `sequential-50`：每个虚拟用户维护自己的 Cursor 链连续读取 50 页，作为 Cursor Treatment，全部请求计入结果。
- `same-second-cursor`：正确性/最坏 tie 读取 probe，不混入普通容量曲线。

报告增加分页模式、目标深度、前序 Cursor 准备耗时（单列、不计目标页）、重复、遗漏、候选/成员/详情装载计数和 oracle hash。任一 Cursor 缺失、循环、提前结束、重复、越权作者或 oracle 不匹配都令报告 incomplete。

验证：

```powershell
go test ./cmd/loadtest -run 'Cursor|Sequential|Report' -count=20
go test ./cmd/loadtest -count=1
go vet ./cmd/loadtest
```

## 9. C5：深数据集和正式 A/B

### 9.1 数据集工具 RED/GREEN

修改或新增：

- `cmd/loadtest/topology.go`
- `cmd/loadtest/topology_test.go`
- `cmd/loadtest/setup.go`
- `cmd/loadtest/setup_test.go`
- `cmd/loadtest/dataset.go`
- `cmd/loadtest/dataset_test.go`
- 必要时新增 `cmd/loadtest/cursor_dataset.go` 与测试

增加独立 `cursor-deep` cardinality/template：20 readers、20 normal authors×50 posts、5 BigV×100 posts，每读者 1,000 Inbox + 500 BigV 候选。固定 seed，构造至少 100 条同秒 tie，保存 MySQL/ZSet/oracle 指纹。

setup 必须在进入测试前验证：

- 1,500 DB 基线帖子及发布/可见状态。
- 每个 Inbox 1,000、每个 BigV Outbox 100。
- page1/page5/page20/page50 非空。
- 全序 oracle hash、首尾 Cursor 和重复 ID 预期。

### 9.2 功能与集成验证

优先使用既有 Redis/MySQL/Kafka/etcd 中间件；KnowPost/Gateway/Relation/Counter 等项目服务本地运行。先跑：

```powershell
go test ./services/knowpost/rpc/internal/feed -tags=integration -run Cursor -count=1
go test ./cmd/loadtest -tags=integration -run Cursor -count=1
```

再执行单用户连续翻页 smoke，确认 oracle、删除/转私密/插入语义。失败时不开始压力测试。

### 9.3 Scout 与正式点

固定 PageCache=off、RouteSnapshot=on、CombinedPipeline=on、CursorPagination=on、Hybrid、benchmark-error-only。RPC/Gateway 分开执行：

1. page/cursor 的 page1、page5、page20、page50，c16/c32/c64，每点 10 秒 scout。
2. 每个入口选择稳定最优档及一个相邻档。
3. 正式点预热 10 秒、采样 60 秒、三轮。
4. `sequential-page-50` Control 与 `sequential-50` Cursor Treatment 在相同选定档、同一 patch 下单列三轮。
5. same-second 场景以正确性和最坏成员数为主，不混入普通容量比较。

每轮持续轮询进度和完整性；长命令产生 cell ID 时按不超过 60 秒间隔调用 wait，不让测试在无状态反馈下运行。

### 9.4 C5 报告

新增：

- `docs/superpowers/reports/2026-08-17-feed-cursor-pagination-ab.md`

报告必须：

- 先声明旧 WP11 deep-page 数据太浅、不能作为有效深页基线。
- 展示同口径 page/Cursor 三轮中位数/min/max，而不是挑最好一轮。
- 分开单目标页成本和 sequential-50 累计成本。
- 同时展示 QPS、P95/P99、CPU/RSS/GC、Redis ops/returned members、候选/详情装载和正确性。
- 按 Spec 19.11 明确回答“默认启用、保持 opt-in、还是未证明必要”，不因已经实现而预设结论。

## 10. 阶段检查点

每个检查点完成后先暂停长压并报告：

- Checkpoint A（C1）：Codec/Proto 兼容完成。
- Checkpoint B（C2）：Reader 精确 seek 与 oracle 完成。
- Checkpoint C（C3）：RPC/Gateway 功能路径完成，全量相关回归通过。
- Checkpoint D（C4/C5 smoke）：深数据与压测工具可用，提供短测结果和预计正式耗时。
- Checkpoint E（C5）：正式 A/B 报告完成，再决定是否默认开启。

## 11. 最终回归

在 C3 与 C4 完成后运行：

```powershell
go test ./services/knowpost/... ./services/gateway/... ./cmd/loadtest/... -count=1
go vet ./services/knowpost/... ./services/gateway/... ./cmd/loadtest/...
git diff --check
```

在 C5 前后分别记录 Git patch hash、被测二进制 SHA、运行参数、数据 manifest 指纹和 Redis identity。任何正式报告完整性门禁失败都必须排除，不能通过人工修改 `complete=true` 接纳。

## 12. 完成记录

2026-08-17 完成 Checkpoint E：

- 深数据 manifest Ready，Redis/MySQL/oracle 指纹和 Kafka lag 门禁通过。
- c16/c32 scout 选出 c16 稳定档；c32 作为吞吐不涨、尾延迟上升的相邻边界。
- RPC/Gateway 固定 page20/page50、同秒 tie、连续 50 页严格配对均完成；正式有效轮次错误、超时、重复和 oracle mismatch 为 0。
- 发现并补齐 `sequential-page-50` Control，以及 Go Runtime GC/分配采集；一次将 heap Gauge 误作 Counter 的无效报告被 fail-closed 排除，修复后在新命名空间完整重跑。
- Spec 19.11 的深页和连续滚动性能门槛通过；因缺少改造前同 manifest 的 page1 二进制基线，决定保持功能开关可回滚并先灰度，不直接全量默认启用。
- 最终证据与原始命名空间见 `docs/superpowers/reports/2026-08-17-feed-cursor-pagination-ab.md`。
