# Feed Hybrid WP8：Content Safety Epoch 写路径检查点

日期：2026-08-15

状态：完成；二次复审 Ready（Critical 0 / Important 0）

## 1. 结论

WP8 已把删除、转私密、置顶和元数据变更接入可持久补偿的 content-safety epoch，并封闭了“旧 FeedItem 重建新页面”的一致性缺口。Feed 仍接受发布后数秒内短暂不可见，但敏感撤销采用 fail-closed：无法确认 safety 状态时不读取整页缓存，也不读取个人 FeedItem cache，直接走数据库冷路径。

```text
敏感事务 commit
  -> 既有缓存失效
  -> 同步 BumpSafety（低延迟）
  -> canal-outbox 独立 consumer 再次 BumpSafety（持久补偿）

GetUserFeed with PageCache
  -> flush/check pending + read relation/safety epoch
  -> PageCache / safety-versioned FeedItem cache
  -> return 前重新校验 pending 和 epochs
       changed/error -> discard -> DB-backed cold path
```

正常敏感 mutation 会因同步与异步链路产生两次单调 bump。这会增加一次全局缓存冷却，但不影响正确性；当前以低频敏感写的额外冷流量换取低延迟和跨实例、重启后的可恢复性。

## 2. 实现边界

敏感写路径：

- `Delete`、`UpdateVisibility`、`UpdateTop`、`PatchMetadata`：数据库提交后同步 bump。
- `Publish`：不 bump，新内容仍遵循 Fresh TTL 的短暂不可见取舍。
- 事务未提交：不 bump。
- 已提交但既有模型缓存失效失败：仍 bump，随后返回原缓存错误。
- bump 失败：已提交业务结果不被改写；设置进程内 pending，后续读取尝试 flush。

持久补偿：

- 独立 consumer group `knowpost-feed-content-safety-epoch` 消费 `canal-outbox`。
- KnowPost Updated/Deleted 执行 bump；Published 明确跳过。
- Redis 错误返回 handler，由共享 consumer 对同一消息重试，成功前不拉取并提交后续 offset。
- PageCache 非 `off` 时强制要求 `Feed.Epoch.SafetyConsumerEnabled=true`。

读路径防线：

- 个人 FeedItem key 为 `feed:item:personal:s{safetyEpoch}:{postID}`。
- 旧无版本 key、旧 safety 版本 key，以及 pending 期间写入的当前版本 key均不可跨安全边界复用。
- pending flush 失败、epoch 读取失败、PageCache 不准入或返回前 epoch 变化时，PageCache 与个人 FeedItem cache 同时 bypass。
- L1/L2 页面仍携带完整 relation/safety epoch；任何不匹配页面均丢弃。

## 3. 复审闭环

首次审查识别出两个问题：

1. Critical：只失效页面 key 不足以阻止失败的旧 `feed:item` 删除重新水合敏感帖子。修复为 safety-versioned 个人 FeedItem key，并在 safety 不可信时完全绕过该缓存。
2. Important：只依赖进程内 pending，实例崩溃或其他实例无法获得可靠补偿。修复为事务 outbox 的独立 Kafka consumer，并让 PageCache 配置强制依赖该 consumer。

新增测试覆盖旧 unversioned key、旧 epoch key、pending 当前 epoch key、commit 后缓存删除失败、同步 bump 失败与合并补偿。二次复审结论为 Ready，Critical 0、Important 0。

非阻塞后续项：校验 outbox row type 与 payload event type 一致；增加“同步 bump 故障后仅由 Kafka 最终补偿”的故障注入门禁；监控双 bump 引起的全局冷流量。

## 4. 自动化验证

已通过：

```text
go test ./services/knowpost/... ./pkg/cachex/... ./pkg/kafkax/... ./cmd/loadtest/... ./deploy/compose/... -count=1
go vet ./services/knowpost/... ./cmd/loadtest/... ./deploy/compose/...
go test -tags=integration ./cmd/loadtest -run '^$' -count=1
go test -tags=integration ./cmd/loadtest -run TestContentSafetyMutationsInvalidateWarmPersonalFeedPage -count=1 -v
git diff --check
```

真实开发栈集成结果：

- 转私密后 `95.4161ms` 内从已预热的个人 Feed 第一页消失，safety epoch `10 -> 12`。
- 恢复并再次预热后，删除在 `92.5074ms` 内生效，safety epoch `14 -> 16`。
- 测试显式等待 `before+2`，因此不仅证明同步 bump，也证明独立 outbox consumer 完成补偿。

Windows 当前 `CGO_ENABLED=0` 且无 gcc，不能构建 Go race runtime，因此没有声称 `go test -race` 通过。

## 5. Docker 与运行态验证

构建使用已有本地层并成功完成，没有发生镜像下载问题。运行镜像：

```text
sha256:f604129f626e4259ac1bc76b39aac71a1e842ad82f6c725624608d005a457cca
```

KnowPost 运行态：

- strategy=`hybrid`
- observability、relation epoch、route snapshot、combined pipeline、content safety epoch 均启用
- page cache=`l1-l2`
- container healthy，restart=0；最近日志未见 error/fatal/panic
- Kafka group `knowpost-feed-content-safety-epoch` lag=0

小样本路径验证位于 `results/feed-loadtest/feed-wp8-safety-validation-20260815`：40/40 成功、失败 0、P95 `9.5568ms`，环境快照明确记录 `safety_epoch_consumer_enabled=true` 与 `page_cache_mode=l1-l2`。该样本只证明新 key 的冷填充、开关和监控路径，不作为 WP8 的 QPS 或容量结论。

本轮构建未出现镜像源异常。后续若 Docker 拉取或构建缓慢，仍按 registry/mirror 镜像源问题处理，不能归因于网络；必要时保留中间件容器、项目服务切换到宿主机，并重新建立混合拓扑基线。

## 6. 下一步

进入 WP9：在不放宽上述 safety 门禁的前提下实现完整 page key 的 Singleflight、TTL jitter、SWR 与有界刷新。WP9 必须证明同 key 合并、不同 key 隔离、waiter 独立取消、epoch/pending 禁止 stale、队列满不增生 goroutine，并在正式性能结论前继续保留 WP6 冷路径基线。
