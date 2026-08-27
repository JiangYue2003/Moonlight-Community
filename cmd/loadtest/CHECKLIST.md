# Feed推拉结合架构 - 压力测试执行清单

## ✅ 工具完成确认

- [x] 压力测试工具已编译：`feed_loadtest.exe` (79MB)
- [x] 配置文件已创建：`feed_loadtest.yaml`
- [x] Windows脚本已创建：`run_test.ps1`
- [x] Linux/Mac脚本已创建：`run_test.sh`
- [x] 快速指南已创建：`QUICKSTART.md`
- [x] 详细文档已创建：`README.md`
- [x] 完成报告已创建：`docs/feed-push-pull-loadtest-completion.md`

---

## 🚀 立即执行（3步）

### 1. 启动所有服务
```powershell
# 当前目录: F:\zhiguang_be\zhiguang-go

# 启动依赖服务（如果Docker未启动）
cd deploy\compose
docker compose -f docker-compose.dev.yml up -d

# 启动业务服务
cd ..\..
.\scripts\start-all.ps1

# 等待30秒让服务完全启动
Start-Sleep -Seconds 30
```

### 2. 运行压力测试
```powershell
# 进入测试目录
cd cmd\loadtest

# 运行交互式测试脚本
.\run_test.ps1

# 选择场景（推荐首次使用场景1）
# 1 - 快速验证 (50用户, 30秒)
```

### 3. 查看结果
测试完成后会自动显示性能报告。

---

## 📋 执行前检查清单

### 依赖服务状态
```powershell
# 检查Redis
redis-cli ping
# 应返回: PONG

# 检查Etcd
netstat -an | findstr 2379
# 应看到: 0.0.0.0:2379

# 检查Kafka
docker ps | findstr kafka
# 应看到: zg-kafka

# 检查MySQL
netstat -an | findstr 3306
# 应看到: 0.0.0.0:3306
```

### 业务服务状态
```powershell
# 检查服务日志
Get-Content logs\dev\gateway.log -Tail 10
Get-Content logs\dev\knowpost-service.log -Tail 10
Get-Content logs\dev\relation-service.log -Tail 10
Get-Content logs\dev\counter-service.log -Tail 10

# 应该看到服务已启动的日志
```

---

## 🎯 推荐的测试顺序

### 第1次：快速验证（5分钟）
```powershell
cd cmd\loadtest
.\feed_loadtest.exe -users 50 -posts 100 -reads 500 -duration 30s
```
**目的**: 验证工具可以正常运行，所有服务正常工作

**预期结果**:
- 成功率 ≥ 95%
- 无报错
- 能看到完整的测试报告

---

### 第2次：中等负载（10分钟）
```powershell
.\feed_loadtest.exe -users 200 -posts 1000 -reads 5000 -duration 60s
```
**目的**: 测试正常业务负载下的性能

**预期结果**:
- 发帖QPS: 15-30
- 读取QPS: 80-150
- 总QPS: 100-180

---

### 第3次：大规模压测（20分钟）
```powershell
.\feed_loadtest.exe -users 500 -posts 3000 -reads 15000 -duration 120s
```
**目的**: 找到系统瓶颈和极限

**预期结果**:
- 发现性能瓶颈
- 识别需要优化的地方
- 为容量规划提供数据

---

## 📊 性能数据记录模板

### 测试记录表

| 运行时间 | 用户数 | 发帖数 | 读取数 | 发帖QPS | 读取QPS | 发帖成功率 | 读取成功率 | 平均延迟 | 备注 |
|---------|-------|-------|-------|---------|---------|-----------|-----------|---------|------|
| 2026-07-29 17:30 | 50 | 100 | 500 | - | - | - | - | - | 首次测试 |
| | 200 | 1000 | 5000 | - | - | - | - | - | 中等负载 |
| | 500 | 3000 | 15000 | - | - | - | - | - | 大规模 |

---

## 🐛 常见问题快速处理

### 问题1: "连接Etcd失败"
```powershell
# 检查Etcd
netstat -an | findstr 2379

# 如果没有，启动Etcd
# （你的Etcd应该已经在运行）
```

### 问题2: "初始化失败"
```powershell
# 检查Relation和Counter服务
Get-Content logs\dev\relation-service.log -Tail 20
Get-Content logs\dev\counter-service.log -Tail 20
```

### 问题3: "发帖失败率高"
```powershell
# 检查KnowPost服务
Get-Content logs\dev\knowpost-service.log -Tail 50

# 检查Kafka
docker logs zg-kafka --tail 50
```

### 问题4: "读取延迟高"
```powershell
# 检查Redis
redis-cli INFO stats | Select-String "instantaneous_ops"

# 检查Redis内存
redis-cli INFO memory | Select-String "used_memory"
```

---

## 💾 测试数据管理

### 测试前
```powershell
# 清理旧的测试数据（可选）
redis-cli --scan --pattern "feed:*" | ForEach-Object { redis-cli del $_ }
```

### 测试后
```powershell
# 导出Redis数据（可选，用于分析）
redis-cli --scan --pattern "feed:*" > feed_keys.txt

# 查看数据量
redis-cli KEYS "feed:*" | Measure-Object

# 清理测试数据（可选）
redis-cli --scan --pattern "feed:*" | ForEach-Object { redis-cli del $_ }
```

---

## 📈 性能优化建议

### 如果QPS低于预期
1. 检查CPU使用率（任务管理器）
2. 检查Redis延迟（redis-cli --latency）
3. 检查MySQL连接数
4. 考虑增加并发数

### 如果成功率低于预期
1. 检查服务日志
2. 降低并发数
3. 增加超时时间
4. 检查网络连接

### 如果延迟高于预期
1. 检查Redis缓存命中率
2. 检查数据库慢查询
3. 检查RPC调用链路
4. 优化数据结构

---

## 🎉 成功标准

测试被认为成功，当：
- [x] 工具正常运行，无崩溃
- [x] 发帖成功率 ≥ 95%
- [x] 读取成功率 ≥ 98%
- [x] 平均延迟符合预期
- [x] 能生成完整的测试报告
- [x] 推拉模式分类正确

---

## 📞 需要帮助？

如果遇到问题：
1. 查看 `cmd\loadtest\QUICKSTART.md` - 快速开始指南
2. 查看 `cmd\loadtest\README.md` - 详细文档
3. 查看服务日志查找错误信息
4. 联系开发团队

---

## 🎯 现在开始！

```powershell
# 1. 确认当前目录
pwd
# 应该是: F:\zhiguang_be\zhiguang-go

# 2. 进入测试目录
cd cmd\loadtest

# 3. 运行测试
.\run_test.ps1

# 4. 选择场景1（快速验证）

# 5. 等待测试完成

# 6. 查看报告
```

**预计用时**: 3-5分钟（快速验证）

祝测试顺利！🚀
