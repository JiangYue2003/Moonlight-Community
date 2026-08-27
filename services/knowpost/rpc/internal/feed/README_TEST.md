# Feed推拉结合架构 - 测试说明

## 快速开始

### 1. 环境准备
```bash
# 确保服务运行
redis-cli ping                    # Redis
docker ps | grep kafka            # Kafka

# 如果Kafka未运行
cd deploy/compose
docker compose -f docker-compose.dev.yml up -d kafka
```

### 2. 运行测试
```bash
cd services/knowpost/rpc/internal/feed

# 集成测试
go test -tags=integration -v

# 性能测试
go test -tags=integration -bench=. -benchtime=2s -run=^$
```

## 测试结果

### ✅ 集成测试: 7/7 通过
- 普通用户发帖（推模式）
- 大V发帖（拉模式）
- 阈值边界测试
- 推拉混合读取
- 容量限制测试
- 扇出Worker测试

### ✅ 性能测试: 全部达标
- 推模式写入: <50ms ✅
- 拉模式写入: <10ms ✅
- 混合读取: <100ms ✅
- 扇出1000粉丝: ~35ms ✅

## 文档

- **测试报告**: [feed-push-pull-test-report.md](./feed-push-pull-test-report.md)
- **测试指南**: [feed-push-pull-testing-guide.md](./feed-push-pull-testing-guide.md)
- **实施总结**: [feed-push-pull-test-summary.md](./feed-push-pull-test-summary.md)
- **架构设计**: [feed-push-pull-implementation-report.md](./feed-push-pull-implementation-report.md)

## 常见问题

**Q: Redis连接失败？**  
A: `redis-cli ping` 检查Redis是否运行

**Q: Kafka超时？**  
A: `docker ps | grep kafka` 检查Kafka容器

**Q: 测试失败？**  
A: 查看 [测试指南](./feed-push-pull-testing-guide.md) 的故障排查章节

## 生产部署

推拉结合Feed架构已通过完整测试，**可以部署到生产环境**。

部署步骤请参考：[测试总结 - 生产部署建议](./feed-push-pull-test-summary.md#生产部署建议)

---

**测试时间**: 2026-07-29  
**测试状态**: ✅ 全部通过  
**生产就绪**: ✅ 是
