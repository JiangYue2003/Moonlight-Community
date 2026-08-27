# Feed推拉结合架构 - 压力测试指南

## 快速开始

### 1. 启动所有服务

```bash
# 启动依赖服务（MySQL、Redis、Etcd、Kafka）
cd deploy/compose
docker compose -f docker-compose.dev.yml up -d

# 启动业务服务
cd ../..
.\scripts\start-all.ps1
```

### 2. 运行压力测试

```bash
cd cmd/loadtest

# 快速测试（小规模）
go run feed_loadtest.go -users 50 -posts 200 -reads 1000 -duration 30s

# 中等规模测试
go run feed_loadtest.go -users 200 -posts 2000 -reads 10000 -duration 60s

# 大规模测试（使用配置文件）
go run feed_loadtest.go -f feed_loadtest.yaml -duration 120s
```

---

## 测试场景

### 场景1：小规模验证测试
```bash
go run feed_loadtest.go \
  -users 50 \
  -bigv 5 \
  -followings 10 \
  -posts 100 \
  -post-c 10 \
  -reads 500 \
  -read-c 20 \
  -duration 30s
```

**预期**:
- 发帖成功率: >95%
- 读取成功率: >95%
- 发帖平均延迟: <100ms
- 读取平均延迟: <50ms

---

### 场景2：中等规模压力测试
```bash
go run feed_loadtest.go \
  -users 200 \
  -bigv 20 \
  -followings 30 \
  -posts 2000 \
  -post-c 30 \
  -reads 10000 \
  -read-c 100 \
  -duration 60s
```

**预期**:
- 发帖成功率: >90%
- 读取成功率: >95%
- 发帖平均延迟: <200ms
- 读取平均延迟: <100ms
- 总QPS: >200

---

### 场景3：大规模饱和测试
```bash
go run feed_loadtest.go \
  -users 500 \
  -bigv 50 \
  -followings 50 \
  -posts 5000 \
  -post-c 50 \
  -reads 20000 \
  -read-c 200 \
  -duration 120s
```

**预期**:
- 发帖成功率: >85%
- 读取成功率: >90%
- 发帖平均延迟: <500ms
- 读取平均延迟: <200ms
- 总QPS: >500

---

## 参数说明

### 用户相关
- `-users`: 测试用户总数（默认100）
- `-bigv`: 大V用户数量（默认10）
- `-followings`: 每用户关注数（默认20）

### 发帖相关
- `-posts`: 总发帖数（默认1000）
- `-post-c`: 发帖并发数（默认10）

### 读取相关
- `-reads`: 总读取数（默认5000）
- `-read-c`: 读取并发数（默认50）

### 其他
- `-duration`: 测试持续时间（默认60s）
- `-f`: 配置文件路径
- `-setup-only`: 只执行初始化，不压测
- `-skip-setup`: 跳过初始化，直接压测

---

## 测试流程

### 阶段1：初始化（setup）
1. 创建测试用户的关注关系
2. 等待Counter同步（通过Kafka）
3. 验证初始数据

**时间**: 约10-30秒（取决于用户数）

### 阶段2：发帖压测
1. 并发创建草稿
2. 更新元数据
3. 确认内容
4. 发布帖子
5. 统计推拉模式分布

**指标**: QPS、延迟、成功率

### 阶段3：读取压测
1. 并发读取公共Feed
2. 触发推拉混合逻辑
3. 测试收件箱+大V发件箱读取

**指标**: QPS、延迟、成功率

---

## 测试报告示例

```
=== 压测报告 ===

【发帖性能】
  总请求数: 2000
  成功: 1950 (97.5%)
  失败: 50 (2.5%)
  平均延迟: 87ms
  QPS: 32.5
  推模式: 1560 (80.0%)
  拉模式: 390 (20.0%)

【读取性能】
  总请求数: 10000
  成功: 9850 (98.5%)
  失败: 150 (1.5%)
  平均延迟: 45ms
  QPS: 164.2

【总体统计】
  测试时长: 1m0s
  总请求数: 12000
  总成功数: 11800
  总QPS: 196.7
```

---

## 监控指标

### Redis监控
```bash
# 监控Redis命令
redis-cli monitor | grep feed

# 查看内存使用
redis-cli INFO memory | grep used_memory_human

# 查看key数量
redis-cli DBSIZE
```

### Kafka监控
```bash
# 查看消费lag
docker exec zg-kafka kafka-consumer-groups.sh \
  --bootstrap-server localhost:9092 \
  --describe \
  --group feed-fanout-group

# 查看topic
docker exec zg-kafka kafka-topics.sh \
  --list \
  --bootstrap-server localhost:9092
```

### 服务日志
```bash
# Gateway日志
Get-Content -Wait logs/dev/gateway.log | Select-String "feed"

# KnowPost日志
Get-Content -Wait logs/dev/knowpost-service.log | Select-String "feed"

# Counter日志
Get-Content -Wait logs/dev/counter-service.log
```

---

## 性能调优建议

### 1. Redis优化
```yaml
# redis.conf
maxmemory 2gb
maxmemory-policy allkeys-lru
```

### 2. Kafka优化
```yaml
# server.properties
num.network.threads=8
num.io.threads=8
socket.send.buffer.bytes=102400
socket.receive.buffer.bytes=102400
```

### 3. Go服务优化
```go
// 增加GOMAXPROCS
runtime.GOMAXPROCS(runtime.NumCPU())

// 增加连接池
redis.PoolSize = 100
```

---

## 故障排查

### 问题1: 发帖失败率高
**可能原因**:
- Kafka未启动或消费延迟
- MySQL连接池耗尽
- Redis内存不足

**排查步骤**:
```bash
# 检查Kafka
docker logs zg-kafka | tail -50

# 检查MySQL连接数
mysql -u root -p -e "SHOW PROCESSLIST;"

# 检查Redis内存
redis-cli INFO memory
```

---

### 问题2: 读取延迟高
**可能原因**:
- 关注的大V过多，并行拉取慢
- Redis缓存未命中
- RPC调用超时

**排查步骤**:
```bash
# 查看Redis命中率
redis-cli INFO stats | grep keyspace

# 检查RPC超时
grep "timeout" logs/dev/knowpost-service.log
```

---

### 问题3: 服务OOM
**可能原因**:
- 并发过高，goroutine泄漏
- 缓存未清理
- 连接未关闭

**排查步骤**:
```bash
# 查看goroutine数量
curl http://localhost:9104/debug/pprof/goroutine?debug=1

# 查看内存profile
go tool pprof http://localhost:9104/debug/pprof/heap
```

---

## 性能基准

### 硬件配置: AMD Ryzen 7 9700X (8核), 32GB RAM

| 场景 | 用户数 | 并发数 | 发帖QPS | 读取QPS | 总QPS |
|------|-------|--------|---------|---------|-------|
| 小规模 | 50 | 30 | 10-20 | 50-100 | 60-120 |
| 中等规模 | 200 | 130 | 30-50 | 150-250 | 180-300 |
| 大规模 | 500 | 250 | 40-70 | 300-500 | 340-570 |

---

## 持续压测

### 使用while循环持续压测
```bash
# 持续压测30分钟
$endTime = (Get-Date).AddMinutes(30)
while ((Get-Date) -lt $endTime) {
    go run feed_loadtest.go -users 200 -posts 500 -reads 2000 -duration 60s
    Start-Sleep -Seconds 10
}
```

### 使用不同负载模式
```bash
# 写多读少
go run feed_loadtest.go -posts 5000 -post-c 50 -reads 2000 -read-c 20

# 读多写少
go run feed_loadtest.go -posts 500 -post-c 10 -reads 20000 -read-c 200

# 均衡负载
go run feed_loadtest.go -posts 2000 -post-c 30 -reads 10000 -read-c 100
```

---

## 清理测试数据

```bash
# 清理Redis中的Feed数据
redis-cli --scan --pattern "feed:*" | xargs redis-cli del

# 清理数据库中的测试帖子（可选）
# mysql -u root -p zhiguang -e "DELETE FROM knowposts WHERE user_id >= 100000"
```

---

## 下一步

1. **分析瓶颈**: 使用pprof分析CPU/内存热点
2. **优化配置**: 根据压测结果调整参数
3. **扩容方案**: 设计多实例部署方案
4. **监控告警**: 配置Prometheus + Grafana

---

**文档版本**: 1.0  
**最后更新**: 2026-07-29
