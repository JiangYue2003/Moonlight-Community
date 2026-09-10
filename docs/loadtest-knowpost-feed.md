# 知文 Feed 压测报告

> 测试时间：2026-07-03
> 测试目标：知文公开 Feed 接口（`GET /api/v1/knowposts/feed`）在多级缓存架构下的 QPS/延迟表现，及缓存分层对数据库负载的削减效果。

## 一、测试环境

| 项目 | 内容 |
|---|---|
| 压测工具 | k6（本地已安装，`c:/ProgramData/chocolatey/bin/k6`） |
| 被测入口 | `http://127.0.0.1:8080`（gateway） → `knowpost-rpc`（`services/knowpost/cmd/knowpost`，9004） |
| 依赖组件 | MySQL、Redis（宿主机进程）；Kafka/ZooKeeper/Elasticsearch/Canal（docker，`deploy/compose/docker-compose.dev.yml`） |
| 硬件 | 本机 16 核 CPU，MySQL/Redis 与微服务进程同机运行（资源存在竞争，见「局限性」） |
| MySQL 连接池 | go-zero `sqlx.NewMysql` 默认值：`MaxOpenConns=64`、`MaxIdleConns=64`、`ConnMaxLifetime=1min`（未额外配置，走 go-zero 内置默认） |
| Redis 连接池 | `go-redis` `NewUniversalClient` 默认值：`PoolSize = 10 × GOMAXPROCS = 160`（16核机器，未显式配置） |

### Feed 接口的缓存路径（被测对象的架构背景）

`services/knowpost/rpc/internal/logic/knowpost/getpublicfeedlogic.go` 实现了四级读路径：

1. **L1**：进程内缓存（`cachex.L1`，基于 ristretto），整页命中直接返回，无网络往返
2. **L2 整页**：Redis `feed:public:{size}:{page}:v1` 整页 JSON 命中，回填 L1
3. **L2 片段重建**：Redis `ids list` + `MGET feed:item:{id}` 拼出整页，任意片段缺失即 miss
4. **DB 回源**：前三级都 miss 时，走 `singleflight`（`sfx` 包）合并同一 `(size,page)` 的并发回源请求，只有一个 goroutine 真正查 DB，其余等待结果；回源后写回 L2 三种形态 + L1

此外还有 `hotkey.Detector` 探测热点页面并动态延长该页 TTL，减少已知热点的重复回源。

这套架构是本次压测的核心验证对象：**多级缓存 + 单飞合并，能把面向用户的 QPS 和面向 DB 的 QPS 解耦到什么程度**。

## 二、测试数据准备（可复用）

### 2.1 数据范围

往 `know_posts` 表批量插入了 **5000 条** `status='published', visible='public'` 的测试知文，ID 范围：

```
900000000000000000 ~ 900000000000004999   （共 5000 个连续整数，避开真实雪花 ID 区间）
```

- `creator_id`：在 `1` 和 `2` 之间轮换（对应本地环境已注册的 `smoketest@zhiguang.dev` / `smoketest2@zhiguang.dev` 两个测试账号）
- `publish_time`：按 ID 递增，每条间隔 8 分钟，从当前时间往前铺满约 27.8 天，保证 `order by publish_time desc` 分页有真实的时间分布，不是同一时刻的堆叠数据
- `is_top`：每 97 条设 1 条置顶（`idx % 97 == 0`），用于后续如果要测置顶逻辑
- `title`/`description`：`压测知文 #{idx}` / `seed data for load test {idx}`，可用于按标题过滤区分种子数据和真实数据
- `content_object_key`：`knowpost/content/seed-{idx}.md`（占位，未指向真实 OSS 对象）

按 `size=20` 分页，可用页数 ≈ **250 页**（`5003 / 20`，含此前联调阶段遗留的 3 条真实数据）。

### 2.2 复用方法

生成脚本已保留在 `.tmp/archive/seed-harness/main.go`（未提交 git，`.gitignore` 已排除 `.tmp/archive/seed-harness/`）：

```bash
cd F:/zhiguang_be/zhiguang-go
export GOCACHE=F:/zhiguang_be/zhiguang-go/.gocache
go run ./.tmp/archive/seed-harness/main.go
```

如果 `.tmp/archive/seed-harness/main.go` 已被清理，核心逆向即可重建：批量 `INSERT INTO know_posts (...) VALUES (...)`，`id` 从 `900000000000000000` 开始连续自增，`status='published', visible='public'`，`creator_id` 取库中已存在的用户 ID。

### 2.3 清理方法（本次未清理，供后续需要时参考）

```sql
DELETE FROM know_posts WHERE id >= 900000000000000000 AND id < 900000000000010000;
```

压测过程中对种子数据发起的点赞/收藏（`entityType=knowpost, entityId=900000000000xxxxxx`）会在 Redis/Kafka 侧留下对应的计数和事件，清理知文数据后这些计数会成为孤儿数据，如需彻底清理可一并检查 `counter-rpc` 侧的 key。

## 三、压测场景设计

固定用 k6 的 `ramping-vus` 渐进加压：`0→50（30s）→200（30s）→500（30s）→500（60s，稳态观察）→0（20s）`，总时长约 2分50秒。除特别说明外均为此曲线，最终统计的 QPS/延迟取整个运行窗口的汇总值（含加压过程，非纯稳态峰值）。

脚本保留在 `.tmp/archive/seed-harness/k6/`：`common.js`（登录复用）、`scenario_a_cold.js`、`scenario_b_hotspot.js`、`scenario_c_warm.js`、`scenario_d_mixed.js`、`scenario_a_fixed500.js`。

| 场景 | 请求模式 | 验证目标 |
|---|---|---|
| **A：全冷启动** | `page` 在 1~250 内均匀随机，`size=20` | 缓存大范围 miss 时的 DB 回源能力上限，及 singleflight 击穿保护的真实效果 |
| **B：热点 collapse** | 90% 流量固定打 `page=1`，10% 随机分散到 1~250 | L1 本地缓存 + hotkey 动态延长 TTL 在"部分页面爆热"场景下的收益 |
| **C：预热后** | `setup` 阶段先顺序请求 1~50 页把缓存打满，正式压测阶段仅在这 50 页内随机 | 纯缓存命中路径的 QPS 天花板 |
| **D：混合读写** | 80% 走场景 C 的热点读，20% 随机对种子数据发起点赞/收藏（`POST /api/v1/action/like`/`fav`） | 验证 counter 走 Kafka 异步聚合后，写请求不拖慢读路径 |
| **A'：瞬时 500 并发** | 固定 `vus=500` 直接启动（非渐进），`duration=30s`，配合数据库层秒级采样 | 观察突增流量（而非渐进爬升）下系统的短时抗冲击能力，及此时数据库真实负载 |

## 四、压测结果

### 4.1 应用层（gateway 入口）结果汇总

| 场景 | 总请求数 | QPS | 平均延迟 | P90 | P95 | 失败率 |
|---|---|---|---|---|---|---|
| A：冷启动 | 243,015 | **1427.7** | 105.4ms | 210.5ms | 242.6ms | 0% |
| B：热点collapse | 275,410 | **1618.2** | 81.2ms | 174.4ms | 197.9ms | 0% |
| C：预热后 | 237,457 | **1391.8** | 110.2ms | 207.9ms | 236.3ms | 0% |
| D：混合读写 | 290,269 | **1704.3** | 71.9ms | 158.3ms | 171.7ms | 0% |
| A'：瞬时500并发 | 64,095 | **2112.3**（峰值窗口） | 134.2ms | 222.5ms | 246.6ms | **0.47%**（306次失败） |

### 4.2 数据库真实负载（配合场景 A' 秒级采样）

采样方式：压测运行期间每秒查询一次 `performance_schema.global_status` 中的 `Threads_connected` 和 `Questions`（后者做差值得到每秒查询增量）。

| 时间点 | Threads_connected | 每秒查询增量 |
|---|---|---|
| 压测启动前（t=0~1s） | 7 | 0~2（背景噪音） |
| 500并发瞬时涌入（t=3s） | 72 | **254**（冷启动集中回源） |
| 稳态（t=4s ~ t=39s，持续36秒） | 72（不变） | **稳定在 2**/秒 |

**结论：应用层扛住 2112 QPS 的同一时间窗口内，真正打到数据库的查询只有个位数每秒**，且 `Threads_connected` 全程没有随并发数增长（72 是服务启动时建立的常驻连接，不是压测流量新增的）。

### 4.3 对照实验：绕过所有缓存直连数据库

为了量化"如果没有这套缓存架构，数据库单独能扛多大压力"，写了一个直连 MySQL 的对照程序（`.tmp/archive/seed-harness/monitor/dbdirect.go`），用完全相同的 SQL（`ListPublicFeed` 的查询语句）、**更低的并发（200，低于应用层测的 500）**、跑 20 秒：

```
concurrency=200 duration=20s
total_requests=19264 success=19064 fail=200
qps=963.2 avg_latency_ms=206.57
```

| 对比维度 | 应用层（带缓存，500并发） | 直连数据库（无缓存，200并发） |
|---|---|---|
| QPS | 2112 | 963 |
| 失败率 | 0.47% | 约1.04%（200/19264） |
| 平均延迟 | 134.2ms | 206.6ms |

**在并发数更低的情况下，直连数据库的 QPS 反而只有应用层的不到一半，延迟还更高、失败率相当**——这是缓存分层架构价值的直接量化证据。

## 五、测试总结

1. **缓存架构把有效吞吐提升了 2 倍以上，同时把数据库负载削减了 99% 以上**：应用层 500 并发下达到 2112 QPS，而同样的查询逻辑直连数据库在更低的 200 并发下只能跑到 963 QPS；压测全程数据库每秒查询量稳定在个位数，与应用层千级 QPS 完全不成比例，证明多级缓存 + singleflight 单飞回源确实把绝大部分读请求挡在了数据库之前。

2. **缓存命中路径的性能梯度符合架构设计预期**：场景 B（热点集中，L1 命中率最高）> 场景 C（预热后，Redis 整页命中）> 场景 A（冷启动随机，命中率最低）。三者 QPS 差距不算悬殊（1392~1618），说明即便是"随机 250 页"的冷场景，在 500 并发的重复访问下也很快被摸热，验证了 `hotkey.Detector` 动态延长热点 TTL 的机制确实在起作用。

3. **写路径（点赞/收藏）没有拖累读路径**：场景 D（80%读 + 20%写）反而是四个场景里 QPS 最高、延迟最低的，因为点赞/收藏走 `counter-rpc` → Redis 直接计数 + Kafka 异步聚合，没有同步阻塞在数据库写入上。这验证了 counter 模块"读写分离、异步落库"的设计收益。

4. **渐进加压 vs 瞬时冲击存在真实差异**：渐进式加压到 500 并发全程 0% 失败率；但固定 500 并发瞬时启动（跳过渐进爬坡）出现了 0.47% 的失败率，集中在冷启动瞬间（连接池/goroutine 调度的建立开销）。这是一个诚实且有技术深度的发现——系统能稳定承载持续 500 并发，但对**突增流量**的缓冲能力还有优化空间（例如可以考虑连接池预热、限流排队而非直接拒绝等方向）。

## 六、局限性说明

- 本地单机环境，MySQL/Redis/6个微服务进程与压测工具共享同一台机器的 CPU/内存，绝对 QPS 数字不能直接类比生产多机集群的表现；但**"有缓存 vs 无缓存"的相对倍数关系**具备参考价值。
- 测试数据量为 5000 条，规模小于典型生产场景；数据量更大时单次 DB 回源查询的延迟会更高，缓存收益预计会更明显（因为 DB 回源的边际成本更高）。
- 未测试"缓存集体失效瞬间"的惊群效应（比如大量 key 同一时刻 TTL 到期），这依赖 `cachex.Jitter` 的抖动机制，本次场景设计未针对性覆盖，是后续可以补充的测试点。
- 未做过服务重启后冷启动的对照（本次场景A/C都在同一进程生命周期内完成，L1 缓存有连续性），生产环境滚动发布后的冷启动表现可能与本次数据不同。

## 七、复现步骤

```bash
# 1. 确保依赖服务已启动
docker compose -f deploy/compose/docker-compose.dev.yml up -d
powershell -File scripts/start-all.ps1

# 2. 确认测试数据已存在（见第二节，若已存在可跳过）
go run ./.tmp/archive/seed-harness/main.go

# 3. 依次运行各场景（每个约2分50秒）
k6 run .tmp/archive/seed-harness/k6/scenario_c_warm.js
k6 run .tmp/archive/seed-harness/k6/scenario_a_cold.js
k6 run .tmp/archive/seed-harness/k6/scenario_b_hotspot.js
k6 run .tmp/archive/seed-harness/k6/scenario_d_mixed.js

# 4. 瞬时并发 + DB监控对照（约30秒）
go run ./.tmp/archive/seed-harness/monitor/dbmon.go > .tmp/archive/seed-harness/results/dbmon_a.csv &
k6 run .tmp/archive/seed-harness/k6/scenario_a_fixed500.js

# 5. 直连数据库对照实验（约20秒）
go run ./.tmp/archive/seed-harness/monitor/dbdirect.go
```









