# Feed推拉结合架构 - 端到端压力测试完成

## ✅ 已完成工作

### 1. 压力测试工具 (已编译)
- **主工具**: `cmd/loadtest/feed_loadtest.exe` (79MB)
  - 完整的发帖+读取压力测试
  - 支持用户关注关系初始化
  - 自动统计推拉模式分布
  - 实时显示QPS和成功率

### 2. 配置文件
- `cmd/loadtest/feed_loadtest.yaml` - 默认配置
- `cmd/loadtest/run_test.ps1` - Windows执行脚本
- `cmd/loadtest/run_test.sh` - Linux/Mac执行脚本

### 3. 文档
- `cmd/loadtest/README.md` - 详细使用指南

---

## 🚀 快速开始

### 步骤1: 启动所有服务

```powershell
# 启动依赖服务
cd deploy/compose
docker compose -f docker-compose.dev.yml up -d

# 启动业务服务
cd ../..
.\scripts\start-all.ps1

# 等待服务启动完成 (约30秒)
Start-Sleep -Seconds 30
```

### 步骤2: 运行压力测试

#### 方式1: 使用交互式脚本 (推荐)
```powershell
cd cmd/loadtest
.\run_test.ps1
```

然后选择测试场景：
- `1` - 快速验证 (50用户, 30秒)
- `2` - 中等规模 (200用户, 60秒)
- `3` - 大规模压测 (500用户, 120秒)
- `4` - 仅读取压测
- `5` - 自定义参数

#### 方式2: 直接命令行
```powershell
cd cmd/loadtest

# 快速测试
.\feed_loadtest.exe -users 50 -posts 100 -reads 500 -duration 30s

# 中等规模
.\feed_loadtest.exe -users 200 -posts 1000 -reads 5000 -duration 60s

# 大规模压测
.\feed_loadtest.exe -users 500 -posts 3000 -reads 15000 -duration 120s
```

---

## 📊 测试输出示例

```
=== Feed推拉结合架构 - 端到端压力测试 ===
测试配置:
  用户数: 200 (其中大V: 20)
  每用户关注: 30
  发帖: 1000条 (并发: 30)
  读取: 5000次 (并发: 100)
  测试时长: 1m0s

连接RPC服务...
初始化测试数据...
  创建关注关系...
  ✓ 关注关系创建完成: 成功=5800, 失败=200
  等待Counter同步...
  ✓ Counter同步完成
✓ 初始化完成

开始压力测试...
[进度] 发帖: 850/820/30 (总/成功/失败) | 读取: 4200/4180/20 | 推拉: 680/140

测试时间到达

=== 压测报告 ===

【发帖性能】
  总请求数: 1000
  成功: 970 (97.0%)
  失败: 30 (3.0%)
  平均延迟: 85ms
  QPS: 16.2
  推模式: 800 (82.5%)
  拉模式: 170 (17.5%)

【读取性能】
  总请求数: 5000
  成功: 4950 (99.0%)
  失败: 50 (1.0%)
  平均延迟: 42ms
  QPS: 82.5

【总体统计】
  测试时长: 1m0.2s
  总请求数: 6000
  总成功数: 5920
  总QPS: 98.7
```

---

## 🎯 测试参数说明

### 基础参数
- `-users` : 测试用户数量 (default: 100)
- `-bigv` : 大V用户数量 (default: 10)
- `-followings` : 每用户关注数 (default: 20)

### 发帖参数
- `-posts` : 总发帖数 (default: 1000)
- `-post-c` : 发帖并发数 (default: 10)

### 读取参数
- `-reads` : 总读取数 (default: 5000)
- `-read-c` : 读取并发数 (default: 50)

### 控制参数
- `-duration` : 测试持续时间 (default: 60s)
- `-setup-only` : 只初始化不测试
- `-skip-setup` : 跳过初始化直接测试
- `-f` : 配置文件路径

---

## 📈 性能目标

| 指标 | 目标 | 说明 |
|------|------|------|
| 发帖成功率 | ≥95% | 包含草稿创建、发布等完整流程 |
| 读取成功率 | ≥98% | 推拉混合读取 |
| 发帖延迟 | <200ms | 平均值 |
| 读取延迟 | <100ms | 平均值 |
| 总QPS | ≥100 | 发帖+读取 |

---

## 🔍 监控命令

### 查看Redis状态
```powershell
# 查看Feed相关的key
redis-cli KEYS "feed:*" | Measure-Object

# 查看内存使用
redis-cli INFO memory | Select-String "used_memory_human"

# 实时监控命令
redis-cli monitor | Select-String "feed"
```

### 查看Kafka状态
```powershell
# 查看消费lag
docker exec zg-kafka kafka-consumer-groups.sh `
  --bootstrap-server localhost:9092 `
  --describe `
  --group feed-fanout-group

# 查看topic
docker logs zg-kafka | Select-String "feed"
```

### 查看服务日志
```powershell
# KnowPost服务
Get-Content -Wait logs/dev/knowpost-service.log | Select-String "feed"

# Counter服务
Get-Content -Wait logs/dev/counter-service.log | Select-String "feed"
```

---

## ⚠️ 注意事项

### 1. 首次运行前
- 确保所有服务已启动并注册到Etcd
- 检查Redis、MySQL、Kafka是否正常
- 预留足够的系统资源（建议8GB+内存）

### 2. 测试期间
- 关闭其他占用资源的应用
- 不要同时运行多个压测实例
- 监控系统资源使用情况

### 3. 测试后
- 清理测试数据（可选）
- 检查服务日志是否有异常
- 分析性能瓶颈

### 4. 清理测试数据
```powershell
# 清理Redis中的Feed数据
redis-cli --scan --pattern "feed:*" | ForEach-Object { redis-cli del $_ }

# 清理测试用户的帖子（可选）
# mysql -u root -p zhiguang -e "DELETE FROM knowposts WHERE creator_id >= 100000"
```

---

## 🐛 故障排查

### 问题1: 连接Etcd失败
```
解决方法:
1. 检查Etcd是否启动: netstat -an | findstr 2379
2. 检查服务是否注册: etcdctl get --prefix ""
```

### 问题2: 发帖成功率低
```
可能原因:
- Kafka未启动或消费延迟
- MySQL连接池耗尽
- 并发数过高

解决方法:
1. 检查Kafka: docker ps | findstr kafka
2. 降低并发数: -post-c 10
3. 增加超时时间
```

### 问题3: 读取延迟高
```
可能原因:
- Redis缓存未命中
- 关注的大V过多
- RPC调用超时

解决方法:
1. 预热缓存: 先发一些帖子
2. 减少关注数: -followings 10
3. 检查Redis性能
```

---

## 📝 下一步

### 1. 运行基础测试
```powershell
cd cmd/loadtest
.\feed_loadtest.exe -users 50 -posts 100 -reads 500 -duration 30s
```

### 2. 逐步增加负载
```powershell
# 小规模
.\feed_loadtest.exe -users 100 -posts 500 -reads 2000 -duration 60s

# 中等规模
.\feed_loadtest.exe -users 200 -posts 1000 -reads 5000 -duration 60s

# 大规模
.\feed_loadtest.exe -users 500 -posts 3000 -reads 15000 -duration 120s
```

### 3. 分析结果
- 查看QPS是否满足预期
- 检查成功率是否达标
- 分析延迟分布
- 记录性能瓶颈

### 4. 优化调整
- 根据测试结果调整配置
- 优化Redis、Kafka等组件
- 增加服务实例（如需要）

---

## 🎉 总结

✅ **压力测试工具已完成并可用**
- 完整的端到端测试
- 真实的RPC调用
- 详细的性能报告
- 易于使用的脚本

你现在可以：
1. 启动所有服务
2. 运行 `.\run_test.ps1`
3. 选择测试场景
4. 查看性能报告

祝测试顺利！🚀
