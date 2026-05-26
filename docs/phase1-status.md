# 阶段1实施快照

> 项目分析 & 设计文档：[`project-analysis.md`](./project-analysis.md)

## ✅ 已完成（端到端最小闭环：注册 → 登录 → 点赞 → 计数查询）

### 基础设施
- `go.work` + `go.mod`（module: `github.com/zhiguang/zhiguang-go`）
- `pkg/`：8 个跨服务公共库
  - `snowflakex`（与 Java 严格对齐 EPOCH，含单元测试）
  - `jwtx`（RS256，签发与验签）
  - `redisx`（go-redis 客户端封装 + Script alias）
  - `kafkax`（producer / consumer 手动提交）
  - `counterlua`（toggle / incr_field / decr_field 三份 Lua + go:embed）
  - `hotkey`（滑动窗口热点探测）
  - `errorx`（业务错误码 + BizError）
  - `responsex`（统一 HTTP 响应封装）
- `common/`：业务公共
  - `ctxdata`（userId 在 ctx / metadata 互转）
  - `middleware`（HTTP `AuthMiddleware` + `AuthRpcMiddleware`）
  - `interceptor`（gRPC userId 透传）

### 数据基础
- `db/schema.sql` 拷贝
- `db/migrations/000001~000005`：users / login_logs / know_posts / outbox / following+follower

### Proto
- `proto/auth/auth.proto`：SendCode / Register / Login / Refresh / Logout / VerifyToken / PasswordReset
- `proto/user/user.proto`：GetById / GetByIdentifier / FindByIds / Create / ExistsByIdentifier / UpdatePassword
- `proto/counter/counter.proto`：Toggle / GetCounts / IsMarked / BatchGetCounts

### 核心服务（5 进程）
| 服务 | 端口 | 角色 |
|---|---|---|
| user-rpc | 9002 | 用户 CRUD（goctl model 反向生成 + Redis 缓存） |
| auth-rpc | 9001 | 认证核心（验证码、JWT签发、刷新白名单、bcrypt） |
| auth-api | 8001 | HTTP 网关 |
| counter-rpc | 9003 | 位图 + SDS（Lua 原子操作 + Kafka 事件发布） |
| counter-api | 8002 | HTTP 网关 |
| counter-aggregator | (worker) | 消费 counter-events，每秒刷写 SDS |

### 部署 / 工具
- `deploy/compose/docker-compose.dev.yml`：mysql + redis + kafka + zk + adminer
- `scripts/gen.sh` + `gen.bat`：goctl 批量代码生成
- `scripts/migrate.sh` + `migrate.bat`：golang-migrate 包装

### 编译 & 测试
- `go build ./...`：✅ 全绿
- `go test ./pkg/...`：✅ Snowflake 单元测试通过

---

## 🚀 端到端启动流程

### 1. 启动依赖
```bash
make up
# 或: docker compose -f deploy/compose/docker-compose.dev.yml up -d
```

### 2. 数据库迁移
```bash
make migrate-up
```

### 3. 生成 RSA 密钥（auth-rpc 需要）
```bash
mkdir -p certs
openssl genpkey -algorithm RSA -pkcs8 -out certs/jwt_private.pem -pkeyopt rsa_keygen_bits:2048
openssl rsa -in certs/jwt_private.pem -pubout -out certs/jwt_public.pem
```

### 4. 启动 5 个核心服务（5 个终端）
```bash
go run ./services/user/rpc                       # :9002
go run ./services/auth/rpc                       # :9001
go run ./services/auth/api                       # :8001
go run ./services/counter/rpc                    # :9003
go run ./services/counter/api                    # :8002
go run ./services/counter/aggregator             # 后台 worker
```

### 5. 端到端 curl 验证
```bash
# 1. 发送验证码（dev 模式 code 打印在 auth-rpc 日志）
curl -XPOST http://127.0.0.1:8001/api/v1/auth/send-code \
  -H 'Content-Type: application/json' \
  -d '{"scene":"REGISTER","identifier":"13800138000"}'

# 2. 注册（用日志里的 code）
TOKEN=$(curl -s -XPOST http://127.0.0.1:8001/api/v1/auth/register \
  -H 'Content-Type: application/json' \
  -d '{"identifier":"13800138000","password":"Aa123456","code":"<日志里的码>","nickname":"u1","agreeTerms":true}' \
  | jq -r .token.accessToken)

# 3. 点赞
curl -XPOST http://127.0.0.1:8002/api/v1/action/like \
  -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"entityType":"knowpost","entityId":"123"}'

# 4. 查计数（约 1s 后聚合到位）
sleep 2
curl http://127.0.0.1:8002/api/v1/counter/knowpost/123
# 预期：{"entityType":"knowpost","entityId":"123","counts":{"like":1,"fav":0}}

# 5. Redis 旁路验证
docker exec -it zg-redis redis-cli BITCOUNT bm:like:knowpost:123:0   # = 1
docker exec -it zg-redis redis-cli STRLEN cnt:v1:knowpost:123        # = 20
```

---

## 📝 已知简化（阶段1为减少范围而留作 TODO）

| 简化项 | 说明 | 后续阶段 |
|---|---|---|
| auth-api logout | RPC LogoutReq 需要 user_id+jti，HTTP 只持有 refreshToken；阶段1暂未拆解 | 阶段2 增加 RPC `RevokeRefresh(refresh_token)` |
| counter-rpc rebuild | 阶段1未实现 SDS 缺失自动 rebuild + Redisson 锁 | 阶段2 接入 redsync + bitcount 重建 |
| counter-aggregator HA | 单实例，flushOnce 无分布式锁；多副本会重复 flush | 阶段2 加 Redis 锁选主 |
| user-rpc FindByIds | 逐 id 查（依赖 model cache）；后续可优化为 IN | 阶段3 可观测性引入后再优化 |
| auth-rpc 验证码渠道 | dev 阶段把 code 打印到日志；生产应接 SMS/邮件 | 阶段3 接入阿里云短信/SES |
| `auth-api` jwt 中间件 | 仅有占位 AccessSecret；真实校验改走 auth-rpc.VerifyToken | 已在 counter-api 用 `AuthRpcMiddleware`，后续补 auth-api 同款 |

---

## 🗂 关键文件索引

- 提示词与设计：`docs/project-analysis.md`
- 计数 Lua：`pkg/counterlua/{toggle,incr_field,decr_field}.lua`
- Snowflake：`pkg/snowflakex/snowflake.go`
- JWT：`pkg/jwtx/{rs256,claims,keys}.go`
- 错误码：`pkg/errorx/codes.go`
- HotKey：`pkg/hotkey/detector.go`
- 中间件：`common/middleware/{auth,authrpc}.go`
- gRPC 拦截器：`common/interceptor/userid.go`
- Counter 共享类型：`services/counter/shared/{schema,sds,event}/*.go`
