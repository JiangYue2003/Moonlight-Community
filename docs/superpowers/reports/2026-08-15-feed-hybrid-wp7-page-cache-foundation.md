# Feed Hybrid WP7：第一页 L1/L2 PageCache 基础路径检查点

日期：2026-08-15

状态：完成；复审 Ready（Critical 0 / Important 0）；默认保持关闭，等待 WP8 Content Safety Epoch 写路径

## 1. 结论

WP7 已实现 Hybrid `page=1,size=20` 的 Fresh 两级整页缓存基础路径，同时保留 WP6 作为无条件可回退的冷路径：

```text
GetUserFeed
  -> PageCache mode / strategy / page / size gate
  -> relation epoch + safety epoch
  -> L1 Fresh
       miss -> Redis L2 Protobuf
                  miss -> WP6 cold compute -> L2 -> L1
  -> before return: pending + relation epoch + safety epoch revalidation
       changed/error -> discard cached result -> WP6 cold compute
```

功能支持三种显式模式：`off`、`l2`、`l1-l2`，默认 `off`。只有 `l1-l2` 才创建独立 Page L1；`off` 和 `l2` 不承担 Ristretto 元数据或后台资源。

WP7 不是新的正式容量结论。WP6 已在无整页缓存的冷路径上达到 `10661.31 QPS`；本工作包只建立缓存正确性、故障回退和运行态证据。Singleflight、SWR、有界刷新属于 WP9，敏感内容变更后的 safety bump 属于 WP8，因此当前容器已恢复 `FEED_PAGE_CACHE_MODE=off`。

## 2. 实现边界

新增：

- `services/knowpost/rpc/internal/cache/userfeed/key.go`
- `services/knowpost/rpc/internal/cache/userfeed/codec.go`
- `services/knowpost/rpc/internal/cache/userfeed/cache.go`
- `services/knowpost/rpc/internal/cache/userfeed/cache_test.go`

主要修改：

- `services/knowpost/rpc/internal/logic/knowpost/getuserfeedlogic.go`
- `services/knowpost/rpc/internal/svc/servicecontext.go`
- Feed observer、Config、两份 KnowPost YAML、Compose 环境变量。
- `cmd/loadtest` 运行态 flag、环境快照和报告环境字段。

核心语义：

- 完整 key：`feed:page:v1:{strategy}:{user}:r{relation}:s{safety}:p{page}:n{size}`。
- 仅 `hybrid-v1`、page=1、size=20 可缓存；其他请求记录 bypass 并走 WP6。
- L2 使用 Protobuf bytes，不使用 JSON。
- L1 cost 初值为 `2*proto.Size(page) + len(key) + 256`。
- loader 错误或 nil page 不写缓存；L2 get/decode/encode/set 错误不影响本次正确业务结果。
- L2 写成功后再写 L1；L1/L2 失败均不会成为 Feed 可用性的单点。
- loader 使用 Cache 提供的 Context，GetUserFeed 冷计算不再闭包绑定原始请求 Context。
- 共享 `FeedPage` 只读返回；真实 L1 与 64 路并发 marshal 测试通过。

## 3. 一致性与复审闭环

首次复审发现三个 Important，均已修复：

1. 只在 lookup 前检查 SafetyPending 存在 TOCTOU 窗口。现在缓存返回后再次检查 pending，并重新读取 relation/safety epoch；状态变化、读取错误或版本不一致时丢弃缓存结果并走 WP6。测试在 cache lookup 中途分别触发 pending、relation bump、safety bump，旧 sentinel 均未返回。
2. L2 Redis/codec 故障曾被误记为普通 miss。现在 `l2_get/l2_decode/l2_encode/l2_set` 分别以固定低基数 operation 观测；Redis round trip 与本地 codec 使用不同 dependency，错误回退不被隐藏。
3. epoch 与 L2 故障日志可能形成请求级风暴。现在 epoch 使用进程级 atomic 10 秒限频；L2 错误按 operation 独立限频，日志回调不会接收 key、payload 或 Feed 内容。

最终复审：Ready，Critical 0 / Important 0。

该二次校验关闭了“lookup 进行中发生本实例 epoch/pending 变化”的可验证窗口；跨实例关系事件仍遵循既有约 1 秒 epoch L1 与 5 秒 TTL 的最终一致性设计，不宣称全局线性一致。

## 4. 自动化验证

覆盖：

- key 的 user、strategy、relation epoch、safety epoch、page、size 隔离。
- 默认 off 与合法/非法 mode、TTL 上界、固定 page/size 配置。
- L1 Fresh 不访问 L2/loader；L2 Fresh 解码并回填 L1。
- miss 写 L2 和 L1；L2 read/write 与 corrupt payload 正确回退。
- L2 get/decode/set 错误的独立观测和按 operation 限频。
- EpochStore error、SafetyPending、lookup 中途 epoch/pending 变化全部 bypass。
- Hybrid 以外、其他页码和尺寸不读取 epoch、不访问 PageCache。
- cache 提供的 loader Context 贯穿 FeedReader、FeedItem MGET、DB loader 和回写。
- 真实 Ristretto + miniredis 填充后，关闭 Redis 仍从 L1 命中。
- 64 路并发只读 Protobuf marshal。

已通过：

```text
go test -count=1 ./services/knowpost/rpc/... ./services/knowpost/cmd/knowpost/... ./cmd/loadtest ./deploy/compose
go test -count=1 -tags=integration ./services/knowpost/rpc/internal/feed/...
go test -count=50 ./services/knowpost/rpc/internal/cache/userfeed
go test -count=50 GetUserFeed PageCache 定向用例
go vet 相关 KnowPost、loadtest 与 compose 包
git diff --check
```

Windows 当前 `CGO_ENABLED=0` 且没有 gcc，仍不能构建 Go race runtime，因此没有声称 `go test -race` 通过。共享对象风险由只读约束、真实 L1、并发 marshal 重复测试与后续高并发压测继续验证。

## 5. Docker 与运行态验证

构建继续使用已有本地层：

```text
docker compose -f deploy/compose/docker-compose.dev.yml build --pull=false knowpost
```

构建成功，没有镜像下载问题。最新运行镜像：

```text
sha256:203349ab6e77343a5eb328149da87ac5e664075be9631556813696d84d3f6fb6
```

验证顺序：

1. `PageCache=off`：容器 healthy，Hybrid 路由冒烟通过。
2. `PageCache=l2`：容器 healthy，Hybrid 路由冒烟通过；具体 L2 get/decode/fill 由真实 RedisAdapter 自动化测试覆盖。
3. `PageCache=l1-l2`：容器 healthy，Hybrid 冒烟通过；随后执行 40 个 RPC page1/size20 运行态验证。
4. 最终恢复 `PageCache=off`；容器 running=true、healthy、restart=0，最近日志未发现 error/fatal/panic。

L1+L2 小样本报告：

- `results/feed-loadtest/feed-wp7-l1l2-validation-20260815`
- 40/40 成功，失败 0，timeout 0，报告 complete=true。
- 报告环境明确记录 `page_cache_mode=l1-l2` 及 WP3-WP6 开关。
- 容器累计 PageCache 指标：`l1_fresh=1729`、`miss=38`。
- L2 get/set 与 Protobuf encode 均有独立成功指标。

40 请求仅用于证明真实路径和指标，不作为 QPS/P95 容量结论；样本时长只有约 133ms，且分散读者访问分布不固定。

本轮构建没有镜像源异常。后续若拉取或构建缓慢，仍必须按 Docker registry/mirror 问题处理，不能归因于网络；必要时保留中间件容器、项目服务切换到宿主机，并为混合拓扑重建基线。

## 6. 下一步

进入 WP8：把 Delete、UpdateVisibility、UpdateTop、PatchMetadata 等敏感写路径接入 content-safety epoch。写成功后 bump；bump 失败时 MarkSafetyPending，页面缓存保持 bypass，直到补偿成功。

在 WP8 完成删除、转私密、下架和置顶变化的回归/集成测试之前：

- PageCache 必须保持默认 `off`。
- 不执行 PageCache 正式性能门禁。
- 不把 WP7 小样本命中数字写成容量提升。

WP8 通过后再按 `l2 -> l1-l2` 重新启用，并进入 WP9 Singleflight/SWR/有界刷新。
