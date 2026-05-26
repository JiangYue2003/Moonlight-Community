# 阶段2 实施快照（内容主线）

> 项目分析 & 设计：[`project-analysis.md`](./project-analysis.md)
> 阶段1 状态：[`phase1-status.md`](./phase1-status.md)

## ✅ 已完成（端到端：草稿 → OSS 直传 → 确认 → 元数据 → 发布 → Feed/详情命中三级缓存 → 点赞触发反向索引失效 → 个人主页计数 +1）

### 新增 pkg 库（5 个）
- `pkg/cachex` 三级缓存抽象（L1 ristretto + L2 Redis + null sentinel + 单飞 + jitter + 热点延长）
- `pkg/ossx` 阿里云 OSS 封装（Presign / PutObject / ext/objectKey 推导）
- `pkg/lockx` redsync 分布式锁（TryAcquire 立即返回，不重试）
- `pkg/sfx` SingleFlight 泛型包装
- `pkg/hotkey` 补全 `TTLForPublic / TTLForMine / Extension`

### shared 扩展
- `services/counter/shared/sds/sds.go` 新增 `DecodeN(raw, n)` 变长解码（user counter 复用）
- `services/counter/shared/schema/user_keys.go` 用户维度 SDS 字段索引（与 Java 严格对齐）

### 新增 proto / api
- `proto/usercounter/usercounter.proto` 新
- `proto/knowpost/knowpost.proto` 新
- `proto/auth/auth.proto` 增 `RevokeRefresh`
- `proto/user/user.proto` 增 `UpdateProfile` / `ExistsByZgIdExceptId`
- `services/storage/api/storage.api`、`profile/api/profile.api`、`knowpost/api/knowpost.api` 全新

### 新增 / 补丁服务（8 项）
| 服务 | 端口 | 角色 |
|---|---|---|
| usercounter-rpc | 9005 | 用户维度 SDS（incr/get/batchGet）|
| knowpost-rpc | 9004 | 帖子核心 + 三级缓存 + 进程内 Kafka invalidation listener |
| knowpost-api | 8005 | HTTP 网关（13 端点）|
| storage-api | 8003 | OSS 预签名 + 归属校验 |
| profile-api | 8004 | 资料 PATCH + 头像中转上传 |
| auth-rpc 补丁 | 9001 | 增 `RevokeRefresh` |
| user-rpc 补丁 | 9002 | 增 `UpdateProfile` + `ExistsByZgIdExceptId` |
| counter-aggregator 补丁 | (worker) | flushOnce 加 redsync 选主锁 |

### 三级缓存实现（与 Java 严格对齐）
- L1：4 个 ristretto 实例（detail / feed_public / feed_item / feed_mine）
- L2：Redis String/List/Set 片段（feed:public:ids:* + feed:item:* + feed:public:index:*）
- L3：DB 回源（go-zero sqlx + 自定义 ListPublicFeed/ListMyFeed）
- 防护：SingleFlight + null sentinel + TTL 抖动 + 热点 TTL 延长 + 双删

### 测试覆盖
新增 8 个测试文件：
- `pkg/cachex/cachex_test.go`（10 case，含三级缓存全路径）
- `pkg/ossx/ossx_test.go`（10 case，ext / objectKey / BuildContentUrl）
- `pkg/lockx/redsync_test.go`（3 case，互斥 / 释放 / 并发）
- `pkg/sfx/group_test.go`（4 case，去重 / cancel / 错误传播）
- `pkg/hotkey/ttl_test.go`（3 case，4 档延长）
- `services/counter/shared/schema/user_keys_test.go`
- `services/counter/shared/sds/sds_decoden_test.go`
- `services/usercounter/rpc/internal/logic/usercounter/logic_test.go`（8 case，增减 / saturate / 解码）
- `services/storage/api/internal/logic/presignlogic_test.go`（6 case，归属 / 校验 / 边界）
- `services/profile/api/internal/logic/validate_test.go`（10 case，PatchProfileReq 全字段校验）
- `services/auth/rpc/internal/logic/auth/revokerefreshlogic_test.go`（4 case，撤销 / 幂等失败 / 跨 token）
- `services/counter/aggregator/internal/flusher/lock_test.go`（1 集成 case，互斥临界区）
- `services/knowpost/rpc/internal/cache/keys_test.go`（3 case，key 模板 + hour slot）
- `services/knowpost/rpc/internal/logic/knowpost/helper_test.go`（7 case，row→pb 映射 + JSON 编解码）
- `services/user/rpc/internal/logic/user/updateprofilelogic_test.go`（3 case）

**累计阶段1+2 测试：80+ 个用例，覆盖 15 个测试包，全部通过。**

### 编译/校验
- `go fmt ./...` ✅ 全部归一
- `go vet ./...` ✅ 无告警
- `go build ./...` ✅ 全绿
- `go test ./...` ✅ 0 FAIL，20 个测试包全绿

---

## ⚠️ 阶段2 显式不做（移交阶段3）
- 搜索（ES / IK 分词）
- RAG 真实实现（向量库 + SSE 流式）
- relation 完整链路（关注/取关 / Outbox / Canal / Kafka 同步 worker）
- counter-rpc.GetCounts 的 SDS 缺失重建（待 lockx 落地后引入）
- knowpost AI 摘要持久化与流式输出
- favs_received 字段（usercounter 仅 likes_received）

## 🚀 端到端启动

```bash
make up                     # docker compose mysql/redis/kafka/zk
make migrate-up             # 跑迁移（含 know_posts 表）

# 启动 9 个进程：
go run ./services/user/rpc                       # :9002
go run ./services/auth/rpc                       # :9001
go run ./services/auth/api                       # :8001
go run ./services/counter/rpc                    # :9003
go run ./services/counter/api                    # :8002
go run ./services/counter/aggregator             # worker（已加锁）
go run ./services/usercounter/rpc                # :9005
go run ./services/knowpost/rpc                   # :9004（含 invalidation listener）
go run ./services/knowpost/api                   # :8005
go run ./services/storage/api                    # :8003
go run ./services/profile/api                    # :8004
```

13 条端到端 curl 用例见 `plans/refactored-dazzling-sunset.md` §8（实跑校验为下一步交付项）。

---

## 📊 阶段对比

| 维度 | 阶段1 | 阶段2 |
|---|---|---|
| 服务数 | 5 + 1 worker | 8 + 1 worker（+3 新 + 3 补丁） |
| pkg 数 | 8 | 13（+5 新） |
| proto 数 | 3 | 5（+2 新 + 2 补丁） |
| 测试包 | 6 | 20（+14） |
| 测试用例 | 41 | 80+ |
| 模块覆盖 | auth / user / counter | + knowpost / storage / profile / usercounter |
