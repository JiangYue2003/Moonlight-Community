# Feed Hybrid WP1 pprof 检查点

日期：2026-08-15

对应计划：`docs/superpowers/plans/2026-08-15-feed-hybrid-read-performance-evolution-plan.md` WP1

## 结论

WP1 门禁通过：KnowPost 与 Relation 均能在开发配置下通过独立私有 HTTP mux 提供 pprof，Docker 仅向宿主机回环地址发布端口；上下文取消、端口冲突、禁用模式和配置契约均有自动化测试。Hybrid Feed 功能烟测无回归。

## 运行拓扑

- 拓扑：`compose-full`。
- 应用镜像：使用本地基础镜像和 BuildKit 缓存构建 `zhiguang-go-core:dev`，构建命令显式使用 `--pull=false`。
- 容器内监听：KnowPost `0.0.0.0:6064`，Relation `0.0.0.0:6066`。
- 宿主机回环映射：KnowPost `127.0.0.1:16064`，Relation `127.0.0.1:16066`。
- 端口偏移原因：Windows/Hyper-V 当前排除 TCP `6034-6133`；`netsh interface ipv4 show excludedportrange protocol=tcp` 可复验。6064/6066 写入 Docker HostConfig 后不会形成实际宿主机发布，因此只调整宿主机端口，不改变容器内端口。

## 自动化验证

通过：

```text
go test ./pkg/debughttp/...
go test ./services/knowpost/...
go test ./services/relation/...
go test ./deploy/compose/...
go vet ./pkg/debughttp/... ./services/knowpost/... ./services/relation/...
git diff --check
```

覆盖项：

- 私有 mux 提供 pprof 首页与 goroutine profile。
- debug HTTP 实际 Handler 不是 `http.DefaultServeMux`，仓库不存在使用默认 mux/`Handler=nil` 的 listener。
- `Enabled=false` 不监听。
- 监听冲突返回错误。
- Context 取消后服务退出并释放端口。
- 30 秒 CPU profile 进行中取消 Context 时，profile 请求与组件均在 1 秒内正常结束。
- 本地与 Docker 配置地址正确。
- Compose 只向 `127.0.0.1` 发布 pprof。

`go test -race` 未执行：当前 Windows Go 环境 `CGO_ENABLED=0` 且没有可用 C 编译器。普通测试、vet 和真实容器运行均通过；后续可在 Linux CI 或具备 CGO 工具链的环境补跑 race，不修改全局开发工具链。

## 运行态验证

- 两个容器均为 healthy。
- 容器内 pprof 首页均可访问。
- 宿主机 `16064/16066` pprof 首页均返回 HTTP 200。
- Hybrid 烟测结果：`reader=3099 normal_fanout=100 bigv_fanout=0 kafka_events=1`。

## c128 Profile 检查点

场景：`hybrid / rpc / distributed-read / c128 / 30s`。

```text
QPS       2636.79
success   79185
failed    0
timeout   0
P50       47.900 ms
P95       62.435 ms
P99       69.700 ms
```

这是带 CPU profile 采样的单次检查点，不替代 WP2 的三次 60 秒正式基线。

Profile：

- `results/feed-loadtest/feedbench-20260814c/profiles/knowpost-cpu-c128-25s.pb.gz`
- `results/feed-loadtest/feedbench-20260814c/profiles/relation-cpu-c128-25s.pb.gz`

两份 CPU profile 分别在两次同拓扑、同数据集、同 c128 参数的 30 秒负载中采集，避免本机工具调用被串行调度后产生空闲样本；不能把两份 profile 的 CPU 秒数相加解释为单次运行资源消耗。

主要累计热点：

- KnowPost：`GetUserFeed` 49.32%，`loadFeedItems` 24.22%，Hybrid 冷读 15.70%，Relation 拉取 7.87%，Inbox/BigV Redis 读取占据下一层主要路径。
- Relation：`ListFollowing` 52.10%，`PageActive` 44.91%，底层主要为 MySQL/sqlx 查询、行反射解码和网络写入。

热点证据支持既定顺序：先建立 WP2 阶段指标，再通过 RouteSnapshot 消除每请求 Relation 回源，随后验证 Redis 合并 Pipeline 与第一页缓存；不先做切片池或归并微优化。

## 工作区说明

仓库在 WP1 开始前已有大量未提交和未跟踪修改。本轮没有清理、覆盖或整文件提交既有改动；实施文件保持在工作区，待后续按所有权边界拆分提交。
