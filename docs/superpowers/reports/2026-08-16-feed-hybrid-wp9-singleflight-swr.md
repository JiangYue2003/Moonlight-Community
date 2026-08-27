# Feed Hybrid WP9：Singleflight、SWR 与有界刷新

日期：2026-08-16

状态：完成；终审 Ready（Critical 0 / Important 0）

## 1. 结论

WP9 已把 WP7 的 Fresh PageCache 补齐为可投入正式容量测试的第一页缓存状态机：相同完整 page key 的冷请求共享一次 L2 查询和一次 Hybrid loader；Fresh 最后 20% 与合法 Stale 使用 stale-while-revalidate；后台刷新由固定 worker 和有界队列执行。刷新失败、超时或排队不会把 Feed 可用性绑定到刷新器，过期请求会接管尚未执行的 queued refresh，或在 running refresh 失败后回退原 Hybrid 冷路径。

```text
L1 Fresh
  -> 直接返回

L1 miss
  -> full-key coordinator
     -> L2 Fresh: 返回；最后 20% 尝试 enqueue refresh
     -> L2 Stale: 返回；尝试 enqueue refresh
     -> expired/miss:
          queued refresh  -> foreground takeover -> bounded cold loader
          running refresh -> wait bounded result -> failure => bounded cold loader
          no refresh      -> bounded cold loader
```

这保持了用户已接受的 Feed 取舍：普通发布允许数秒短暂不可见；删除、转私密等敏感撤销继续由 content-safety epoch fail closed，不允许 stale 跨 epoch 返回。

## 2. 实现与边界

- 缓存范围：仅 `hybrid`、第一页、`size=20`。
- full key 包含 strategy version、user ID、relation epoch、safety epoch、page 和 size。
- L2 value：固定二进制时间头 + Protobuf `FeedPage`，不使用 JSON。
- nominal TTL：L1 `800ms`、L2 `4s`；jitter `±20%` 后有效硬边界分别为 `500ms~1s`、`3s~5s`。
- Stale window：最多 `10s`；epoch 变化、SafetyPending 或超过窗口时禁止返回 stale。
- 刷新：`32` workers、`1024` queue、每次 loader 独立 `2s` Context。
- overflow：不启动额外 goroutine；Fresh/Stale 继续返回，真正 cold miss 走同步的有界 loader。
- 多实例：本阶段不引入 Redis 分布式锁；同 key 合并范围明确为单 KnowPost 进程。

## 3. 审查闭环

首次复审识别三个 Important：

1. running refresh 与 cold load 可能并发写回。修复为同 full key promise 协调，过期请求不再和已运行刷新并发写。
2. TTL 只校验 base，jitter 可突破新鲜度上限。修复为校验 jitter 后完整区间，并把配置 tag/YAML 默认统一为 `800ms/4s`。
3. 缺少刷新运行态指标与限频错误。新增 queue depth、active workers、pending keys Gauge，queue full 与 refresh load error 分操作限频。

二次复审又发现刷新器可能成为可用性单点：running refresh error 会直接传给过期请求，queued refresh 可能让前台等待超时。最终修复为：

- worker 在 `loadMu` 下把 promise 从 queued 原子切换为 started；
- 前台只接管尚未 started 的任务，worker 后续识别失效 promise 并跳过 loader/write；
- 已 started 的刷新完成后才处理结果，失败时用新的独立、有限时 Context 回退 cold loader；
- `Close` 取消 active worker、完成队列内 promise，最后等待 cold flights，无遗漏 promise 或锁环。

终审结论：Critical 0、Important 0、Ready。

## 4. 自动化验证

通过：

```text
go test ./services/knowpost/rpc/internal/cache/userfeed ./services/knowpost/rpc/internal/svc -count=1
go test ./services/knowpost/rpc/internal/cache/userfeed -count=50
go test ./services/knowpost/... ./pkg/cachex/... ./pkg/kafkax/... ./cmd/loadtest/... ./deploy/compose/... -count=1
go vet ./services/knowpost/... ./cmd/loadtest/... ./deploy/compose/...
git diff --check
```

覆盖包括：同/不同 key、waiter 独立取消、Fresh final window、Stale async、epoch 隔离、过期窗口、refresh 成功/失败/queued takeover、queue overflow、worker 上限、loader timeout、Close、TTL jitter、二进制 codec、真实 Ristretto/Redis adapter，以及刷新指标饱和与归零。

Windows 当前 `CGO_ENABLED=0`，实际执行 `go test -race` 返回 `-race requires cgo; enable cgo by setting CGO_ENABLED=1`，因此没有声称 race 门禁通过。

## 5. Docker 与一致性验收

构建显式使用 `--pull=false`，本地层命中，未发生镜像源异常。最终镜像：

```text
sha256:5a4c1e4832048c71c35fe3eb5d9e133eb7f57091174f31abcbb8d9ff4a048ef5
```

KnowPost：

- strategy=`hybrid`
- observability、relation epoch、content safety epoch、route snapshot、combined pipeline 全部启用
- page cache=`l1-l2`
- healthy、restart=0；最近日志未见 error/fatal/panic
- `feed-fanout-group`、`knowpost-feed-relation-epoch`、`knowpost-feed-content-safety-epoch` lag 均为 0

内容安全回归：

- 转私密：`95.3161ms` 内从已预热个人 Feed 消失，safety epoch `26 -> 28`。
- 删除：`93.2795ms` 内消失，safety epoch `30 -> 32`。

## 6. 运行态功能样本

热读路径：

- 报告：`results/feed-loadtest/feed-wp9-final-functional-20260816/hybrid-rpc-hot-read-c16/trial-1/report.json`
- 400/400 成功，failed=0、timeout=0、degraded=false
- QPS `10694.99`，P95 `1.3515ms`，P99 `17.6618ms`
- cold compute=1，singleflight shared=15
- 一次 Relation、Counter、Inbox+BigV Pipeline、FeedItem MGet、L2 get/set 支撑整批请求

SWR 定向路径：

- 报告：`results/feed-loadtest/feed-wp9-refresh-functional-20260816/hybrid-rpc-hot-read-c1/trial-1/report.json`
- 40/40 成功，P95 `0.804ms`
- `page_cache_refresh_enqueue=1`、`page_cache_refresh_load=1`
- 完成后 `feed_page_refresh_queue_depth=0`、`active_workers=0`、`pending_keys=0`

上述均为功能小样本，不作为最大承载能力结论。当前可信的正式容量仍是 WP6 冷路径 c128 `10661 QPS`、c64 稳定约 `9866 QPS`；WP10 将使用固定拓扑、预热 10 秒、采样 60 秒、三轮重复和多读者基数测出 PageCache 后的新稳定吞吐、拐点与资源上限。

## 7. 下一步

进入 WP10：扩展压测工具以显式区分 cold、L2 warm、L1 warm 与 expire-together，记录刷新 Gauge、PageCache source 和 pprof 路径，然后执行小/中/大读者基数的 c32/c64/c128/c256 正式矩阵。若 Docker 镜像构建或拉取异常，仍只按 registry/mirror 镜像源问题处理；必要时保留中间件容器、项目服务切换本地，并为新拓扑重建基线。
