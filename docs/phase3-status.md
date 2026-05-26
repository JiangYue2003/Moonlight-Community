# 阶段3 实施快照（社交闭环 + 事件驱动）

> 项目分析：[`project-analysis.md`](./project-analysis.md)
> 阶段1 状态：[`phase1-status.md`](./phase1-status.md)
> 阶段2 状态：[`phase2-status.md`](./phase2-status.md)

## ✅ 已完成（端到端：follow → outbox → Canal binlog → Kafka → relation-syncer → ZSet/usercounter）

### 新增 pkg 库（3 个）
- `pkg/canalx` Canal flatMessage 解析（替代 Java OutboxMessageUtil）
- `pkg/ratelimit` Lua 令牌桶（与原 Java RelationServiceImpl 内联脚本等价）
- `pkg/txx` go-zero sqlx 事务封装

### shared 扩展
- `services/relation/shared/event/event.go` RelationEvent 跨进程结构
- `services/relation/shared/zset/zset.go` ZSet 列表读取（FollowingKey / FollowerKey / PageByOffset / PageByCursor）
- `services/relation/shared/model/` 三个 model（following / follower / outbox）+ 业务方法（UpsertActive / MarkInactive / PageActive / InsertInTx）

### 新增 proto / api
- `proto/relation/relation.proto` 5 方法
- `services/relation/api/relation.api` 5 端点

### 新增 / 补丁服务（5 项）
| 服务 | 端口 | 角色 |
|---|---|---|
| relation-rpc | 9006 | follow/unfollow（Outbox 事务）+ status + 列表（ZSet/DB） |
| relation-api | 8006 | HTTP 网关，JWT 强制鉴权 |
| relation-syncer | (worker) | 消费 canal-outbox → follower 反查表 + ZSet + usercounter |
| counter-rpc 补丁 | 9003 | GetCounts SDS miss 走 lockx + BITCOUNT 重建（清阶段1 技术债） |
| usercounter-rpc | 9005 | 阶段2 已就绪，本期被 syncer 真实调用 |

### 部署
- `deploy/compose/docker-compose.dev.yml`：新增 canal-server 容器；MySQL 开 binlog ROW；Kafka 双 listener（host 9092 / containers 29092）
- `deploy/compose/canal-conf/canal.properties` Canal Server 主配置（serverMode=kafka, flatMessage=true, partitionHash=zhiguang.outbox:id）
- `deploy/compose/canal-conf/instance.properties` instance 配置（filter.regex=zhiguang\.outbox, slaveId=1234）
- `db/migrations/000006_canal_user.up.sql`：dev 环境的 canal 复制账号

### 测试覆盖（新增）
- `pkg/canalx/flatmessage_test.go`（8 case，覆盖 INSERT/UPDATE/DELETE/DDL 路径 + 坏 JSON 防御 + ts/es/pkNames 字段保留）
- `pkg/ratelimit/tokenbucket_test.go`（5 case，覆盖容量耗尽 / refill / 多 key 隔离 / 并发 / 时钟回拨）
- `pkg/txx/tx_test.go`（3 case，COMMIT / ROLLBACK / 空 fn 仍 BEGIN+COMMIT）
- `services/counter/rpc/internal/logic/counter/getcountslogic_test.go`（5 case，覆盖 SDS hit / miss 重建 / hit 不重建 / 并发一致性 / 入参校验）
- `services/relation/rpc/internal/logic/relation/followlogic_test.go`（4 case，自关注拒绝 / 零 id 拒绝 / 限流命中 / 并发限流）
- `services/relation/syncer/internal/processor/dedup_test.go`（3 case，首次成功 / 多 key 隔离 / retry loop 单次执行）

**累计阶段1+2+3 测试：100+ 用例 / 25 个测试包，全部通过。**

### 编译/校验
- `go fmt ./...` ✅ 全部归一（最后一次格式化 9 文件）
- `go vet ./...` ✅ 无告警
- `go build ./...` ✅ 全绿
- `go test ./...` ✅ **0 FAIL**

---

## ⚠️ 阶段3 显式不做（移交阶段4/5）
- 搜索（ES IK 分词、`completion suggester`）
- RAG 真实实现
- 大V Top-N Caffeine 热缓存（V > 50 万时启用，阶段3 简化为"全部走 ZSet + DB"）
- Outbox GC worker
- 计数对账 worker（定时 BITCOUNT 校准 SDS）
- Canal Server HA 与 DLQ
- counter-rpc.GetCounts 重建扫全部 chunk（当前限定前 64 chunk，覆盖 userId<2M）

---

## 🚀 端到端启动

```bash
make up                            # 起 mysql/redis/kafka/zk + canal-server
sleep 30                           # canal-server 首启较慢
docker compose -f deploy/compose/docker-compose.dev.yml logs canal-server | tail
make migrate-up                    # 跑迁移（含 000006_canal_user）

# 13 个进程
go run ./services/user/rpc                        # :9002
go run ./services/auth/rpc                        # :9001
go run ./services/auth/api                        # :8001
go run ./services/counter/rpc                     # :9003 (含 SDS 重建)
go run ./services/counter/api                     # :8002
go run ./services/counter/aggregator              # worker
go run ./services/usercounter/rpc                 # :9005
go run ./services/knowpost/rpc                    # :9004
go run ./services/knowpost/api                    # :8005
go run ./services/storage/api                     # :8003
go run ./services/profile/api                     # :8004
go run ./services/relation/rpc                    # :9006   ← 新
go run ./services/relation/api                    # :8006   ← 新
go run ./services/relation/syncer                 # worker  ← 新
```

详见 `plans/refactored-dazzling-sunset.md` §11 验证脚本。

---

## 📊 阶段对比

| 维度 | 阶段1 | 阶段2 | 阶段3 |
|---|---|---|---|
| 服务 | 5 + 1 worker | 8 + 1 worker | 11 + 2 worker（+3 新 + 1 补丁）|
| pkg | 8 | 13 | 16（+3 新）|
| proto | 3 | 5 | 6（+1 新）|
| 测试包 | 6 | 20 | 25（+5）|
| 测试用例 | 41 | 80+ | 100+ |
| 模块覆盖 | auth / user / counter | + knowpost / storage / profile / usercounter | + relation / canal-bridge（删，由 Canal Server 替代）|
| 原项目 18 服务进度 | 6/18 | 9/18 | **12/18** |
