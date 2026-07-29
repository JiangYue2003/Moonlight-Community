#!/bin/bash

# Docker 环境快速启动脚本

set -e

echo "========================================="
echo "  知光 Docker 环境启动"
echo "========================================="
echo ""

# 检查 Docker 是否运行
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker 未运行，请先启动 Docker"
    exit 1
fi

echo "✅ Docker 运行正常"
echo ""

# 选择部署模式
echo "请选择部署模式："
echo "1) 测试模式（基础设施 + 核心服务）"
echo "2) 完整模式（所有服务）"
read -p "请输入选择 [1/2]: " mode

if [ "$mode" = "2" ]; then
    COMPOSE_FILE="deploy/compose/docker-compose.full.yml"
    echo ""
    echo "📦 使用完整部署模式"
else
    COMPOSE_FILE="deploy/compose/docker-compose.test.yml"
    echo ""
    echo "📦 使用测试部署模式"
fi

echo ""
echo "🔨 开始构建镜像..."
docker compose -f $COMPOSE_FILE build

echo ""
echo "🚀 启动服务..."
docker compose -f $COMPOSE_FILE up -d

echo ""
echo "⏳ 等待服务就绪（30秒）..."
sleep 30

echo ""
echo "📊 检查服务状态..."
docker compose -f $COMPOSE_FILE ps

echo ""
echo "✅ 部署完成！"
echo ""
echo "📖 常用命令："
echo "  查看日志: docker compose -f $COMPOSE_FILE logs -f gateway"
echo "  查看状态: docker compose -f $COMPOSE_FILE ps"
echo "  停止服务: docker compose -f $COMPOSE_FILE down"
echo ""
echo "🔗 服务地址："
echo "  Gateway: http://localhost:8080"
echo "  MySQL:   localhost:3306 (root/Zz123456)"
echo "  Redis:   localhost:6379"
echo "  Etcd:    localhost:2379"
echo ""
echo "🧪 测试服务发现："
echo "  curl http://localhost:8080/api/v1/counter/knowpost/test"
echo ""
