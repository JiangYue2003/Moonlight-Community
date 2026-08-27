# 推拉结合Feed架构 - 测试快速指南

## 一、环境准备

### 1.1 启动必需服务

```bash
# 1. 确保MySQL已启动
# 本地MySQL: 127.0.0.1:3306

# 2. 确保Redis已启动
redis-cli ping  # 应返回 PONG

# 3. 确保Etcd已启动
# 本地Etcd: 127.0.0.1:2379

# 4. 启动Kafka（Docker）
cd deploy/compose
docker compose -f docker-compose.dev.yml up -d kafka

# 验证Kafka
docker ps | grep zg-kafka  # 应看到容器运行
```

---

## 二、运行测试

### 2.1 集成测试

#### 运行全部集成测试
```bash
cd services/knowpost/rpc/internal/feed
go test -tags=integration -v -timeout 3m
```

**预期输出**:
```
=== RUN   TestPushPullFeedIntegration
=== RUN   TestPushPullFeedIntegration/NormalUser_PushMode
    ✓ Normal user post: postID=2001, creatorID=1001...
=== RUN   TestPushPullFeedIntegration/BigV_PullMode
    ✓ Big V post: postID=2002, creatorID=1002...
...
--- PASS: TestPushPullFeedIntegration (0.XX s)
--- PASS: TestFanoutWorkerIntegration (0.XX s)
PASS
```

#### 运行特定测试场景
```bash
# 只测试推拉混合读取
go test -tags=integration -v -run=TestPushPullFeedIntegration/PushPull_MixedRead

# 只测试扇出Worker
go test -tags=integration -v -run=TestFanoutWorkerIntegration
```

---

### 2.2 性能测试

#### 快速性能测试（2秒基准）
```bash
cd services/knowpost/rpc/internal/feed
go test -tags=integration -bench=. -benchmem -benchtime=2s -run=^$ -timeout 5m
```

**预期输出**:
```
BenchmarkFeedWriter_PushMode-16         XXXX    XXXX ns/op    XXX B/op    XX allocs/op
BenchmarkFeedWriter_PullMode-16         XXXX    XXXX ns/op    XXX B/op    XX allocs/op
BenchmarkFeedReader_SmallFollowings-16  XXXX    XXXX ns/op    XXX B/op    XX allocs/op
...
PASS
```

#### 完整性能测试（10秒基准）
```bash
go test -tags=integration -bench=. -benchmem -benchtime=10s -run=^$ -timeout 10m
```

#### 测试特定组件
```bash
# 只测试写入性能
go test -tags=integration -bench=BenchmarkFeedWriter -benchmem -benchtime=5s -run=^$

# 只测试读取性能
go test -tags=integration -bench=BenchmarkFeedReader -benchmem -benchtime=5s -run=^$

# 只测试Redis性能
go test -tags=integration -bench=BenchmarkRedis -benchmem -benchtime=5s -run=^$
```

---

## 三、测试数据清理

### 3.1 自动清理

测试会自动清理数据，无需手动操作。每个测试在结束时会清理：
- `feed:inbox:*`
- `feed:bigv:*`
- `feed:fanout:processing:*`

### 3.2 手动清理（如需要）

```bash
# 清理所有Feed相关数据
redis-cli --scan --pattern "feed:*" | xargs redis-cli del

# 清理特定模式
redis-cli del $(redis-cli keys "feed:inbox:*")
redis-cli del $(redis-cli keys "feed:bigv:*")
redis-cli del $(redis-cli keys "feed:fanout:processing:*")
```

---

## 四、故障排查

### 4.1 Redis连接失败

**错误**: `dial tcp 127.0.0.1:6379: connect: connection refused`

**解决方法**:
```bash
# 检查Redis是否运行
redis-cli ping

# 如果未运行，启动Redis
# Windows: 运行redis-server.exe
# Linux: sudo systemctl start redis
```

---

### 4.2 Kafka未启动

**错误**: Kafka相关测试超时或失败

**解决方法**:
```bash
# 启动Kafka
cd deploy/compose
docker compose -f docker-compose.dev.yml up -d kafka

# 检查状态
docker ps | grep kafka
docker logs zg-kafka  # 查看日志
```

---

### 4.3 测试超时

**错误**: `panic: test timed out after Xm`

**解决方法**:
```bash
# 增加超时时间
go test -tags=integration -v -timeout 10m

# 或者只运行部分测试
go test -tags=integration -v -run=TestPushPullFeedIntegration/NormalUser_PushMode
```

---

### 4.4 端口占用

**错误**: `bind: address already in use`

**解决方法**:
```bash
# Windows查看端口占用
netstat -ano | findstr :6379
netstat -ano | findstr :9092

# 结束占用进程（根据PID）
taskkill /PID <pid> /F
```

---

## 五、性能基准参考

### 5.1 写入性能目标

| 操作 | 目标延迟 | 说明 |
|------|---------|------|
| 推模式写入 | <50ms | 写收件箱+发Kafka |
| 拉模式写入 | <10ms | 只写大V发件箱 |
| 扇出1000粉丝 | <100ms | 批量写收件箱 |

### 5.2 读取性能目标

| 场景 | 目标延迟 | 说明 |
|------|---------|------|
| 10个关注 | <100ms | 推拉混合 |
| 50个关注 | <200ms | 推拉混合 |
| 100个关注 | <500ms | 推拉混合 |

### 5.3 如何解读性能测试结果

```
BenchmarkFeedWriter_PushMode-16    50000    35000 ns/op    500 B/op    10 allocs/op
```

- `50000`: 执行次数
- `35000 ns/op`: 单次操作耗时35微秒（0.035毫秒）
- `500 B/op`: 单次操作分配500字节内存
- `10 allocs/op`: 单次操作分配10次内存
- `-16`: 使用16个并发goroutine

**性能评级**:
- ✅ 优秀: <10ms
- ✅ 良好: 10-50ms
- ⚠️  可接受: 50-100ms
- ❌ 需优化: >100ms

---

## 六、持续集成（CI）集成

### 6.1 GitHub Actions示例

```yaml
name: Feed Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    
    services:
      redis:
        image: redis:7
        ports:
          - 6379:6379
      
      kafka:
        image: bitnami/kafka:3.7
        ports:
          - 9092:9092
        env:
          KAFKA_CFG_ZOOKEEPER_CONNECT: zookeeper:2181
    
    steps:
      - uses: actions/checkout@v3
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Run integration tests
        run: |
          cd services/knowpost/rpc/internal/feed
          go test -tags=integration -v -timeout 5m
      
      - name: Run benchmarks
        run: |
          cd services/knowpost/rpc/internal/feed
          go test -tags=integration -bench=. -benchtime=2s -run=^$
```

---

## 七、开发调试

### 7.1 启用详细日志

测试使用go-zero的logx，日志级别为info。如需更多调试信息，可以修改：

```go
// 在测试文件中临时添加
logx.SetLevel(logx.DebugLevel)
```

### 7.2 使用Redis Monitor

```bash
# 实时监控Redis命令
redis-cli monitor

# 过滤Feed相关命令
redis-cli monitor | grep feed
```

### 7.3 检查Kafka消息

```bash
# 查看Kafka topics
docker exec zg-kafka kafka-topics.sh --list --bootstrap-server localhost:9092

# 消费feed-fanout topic
docker exec zg-kafka kafka-console-consumer.sh \
  --bootstrap-server localhost:9092 \
  --topic feed-fanout \
  --from-beginning
```

---

## 八、常用命令速查

```bash
# === 运行测试 ===
# 集成测试
go test -tags=integration -v

# 性能测试
go test -tags=integration -bench=. -benchmem

# 特定测试
go test -tags=integration -v -run=TestPushPullFeedIntegration

# === Redis操作 ===
# 查看所有Feed keys
redis-cli keys "feed:*"

# 查看ZSet内容
redis-cli ZRANGE "feed:inbox:3001" 0 -1 WITHSCORES

# 清理测试数据
redis-cli --scan --pattern "feed:*" | xargs redis-cli del

# === Kafka操作 ===
# 查看容器状态
docker ps | grep kafka

# 查看日志
docker logs zg-kafka

# 重启Kafka
docker restart zg-kafka

# === 性能分析 ===
# CPU profile
go test -tags=integration -bench=. -cpuprofile=cpu.prof

# 内存 profile
go test -tags=integration -bench=. -memprofile=mem.prof

# 查看profile
go tool pprof cpu.prof
```

---

## 九、测试最佳实践

### 9.1 测试前检查清单

- [ ] Redis已启动并可连接
- [ ] Kafka已启动（如需测试Worker）
- [ ] 端口无占用
- [ ] 有足够的磁盘空间（日志）

### 9.2 性能测试建议

1. **多次运行取平均值**
   ```bash
   for i in {1..3}; do
     go test -tags=integration -bench=BenchmarkFeedWriter_PushMode -benchtime=5s
   done
   ```

2. **关闭其他应用**
   - 关闭浏览器、IDE等占用资源的应用
   - 确保系统负载较低

3. **使用一致的环境**
   - 相同的硬件配置
   - 相同的Go版本
   - 相同的Redis/Kafka版本

---

## 十、下一步

### 10.1 集成到项目

1. **启动Worker**
   ```go
   // 在knowpost服务启动时添加
   go feed.StartFeedFanoutWorker(
       kafkaBrokers,
       redisClient,
       relationClient,
       logger,
   )
   ```

2. **集成到API**
   ```go
   // 发帖时调用
   feedWriter.OnPostPublished(ctx, postID, creatorID, followerCount)
   
   // 获取Feed时调用
   posts, hasMore, err := feedReader.GetFeed(ctx, userID, page, size)
   ```

### 10.2 监控部署

1. 添加Prometheus metrics
2. 配置Grafana dashboard
3. 设置告警规则

### 10.3 压力测试

使用工具如Apache Bench、wrk或自定义脚本进行压力测试：

```bash
# 使用wrk压测
wrk -t12 -c400 -d30s --latency http://localhost:8080/api/v1/feed
```

---

**文档版本**: 1.0  
**最后更新**: 2026-07-29  
**维护者**: 项目团队
