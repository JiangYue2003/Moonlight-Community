# 单容器部署快速开始

## 🚀 一键启动

```bash
# 构建并启动（首次需要 5-10 分钟）
docker compose -f deploy/compose/docker-compose.allinone.yml up -d --build

# 查看启动日志
docker logs -f zhiguang-allinone

# 等待约 30-60 秒，所有服务启动完成
```

## ✅ 验证

```bash
# 1. 查看所有服务状态
docker exec zhiguang-allinone supervisorctl status

# 2. 测试 API
curl http://localhost:8080/api/v1/counter/knowpost/test

# 3. 查看 Etcd 服务注册
docker exec zhiguang-allinone etcdctl get --prefix ""
```

## 📝 常用命令

```bash
# 查看日志
docker logs -f zhiguang-allinone
docker exec zhiguang-allinone tail -f /var/log/supervisor/gateway.log

# 重启服务
docker exec zhiguang-allinone supervisorctl restart gateway

# 停止容器
docker compose -f deploy/compose/docker-compose.allinone.yml down

# 连接数据库
docker exec -it zhiguang-allinone mysql -uroot -pZz123456 zhiguang

# 连接 Redis
docker exec -it zhiguang-allinone redis-cli
```

## 📊 与多容器方案对比

| 特性 | 单容器方案 | 多容器方案 |
|-----|----------|----------|
| 启动速度 | ⚡ 快（一个容器） | 慢（多个容器） |
| 资源占用 | ✅ 少 | 多 |
| 管理复杂度 | ✅ 简单 | 复杂 |
| 服务隔离 | ❌ 无 | ✅ 有 |
| 水平扩展 | ❌ 不支持 | ✅ 支持 |
| 适用场景 | 开发、联调 | 测试、生产 |

## 🎯 使用建议

**推荐使用单容器方案的场景**：
- ✅ 本地开发调试
- ✅ 前后端联调
- ✅ 功能演示
- ✅ 快速验证

**不推荐使用单容器方案的场景**：
- ❌ 生产环境
- ❌ 性能测试
- ❌ 高并发场景
- ❌ 需要多实例部署

详细说明：[deploy/docker/README-allinone.md](../docker/README-allinone.md)
