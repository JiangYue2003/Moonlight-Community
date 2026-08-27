#!/bin/bash
# Feed推拉结合架构 - 压力测试执行脚本 (Linux/Mac)

echo "=== Feed推拉结合架构 - 压力测试 ==="
echo ""

# 检查服务是否启动
echo "检查服务状态..."
if ! nc -z 127.0.0.1 2379 2>/dev/null; then
    echo "❌ Etcd未运行 (127.0.0.1:2379)"
    echo "请先启动Etcd服务"
    exit 1
fi

if ! nc -z 127.0.0.1 6379 2>/dev/null; then
    echo "❌ Redis未运行 (127.0.0.1:6379)"
    echo "请先启动Redis服务"
    exit 1
fi

echo "✓ Etcd运行中"
echo "✓ Redis运行中"
echo ""

# 选择测试场景
echo "请选择测试场景:"
echo "1. 快速验证测试 (50用户, 100帖子, 500读取, 30秒)"
echo "2. 中等规模测试 (200用户, 1000帖子, 5000读取, 60秒)"
echo "3. 大规模压测 (500用户, 3000帖子, 15000读取, 120秒)"
echo "4. 仅读取压测 (1000次读取, 50并发)"
echo "5. 自定义参数"
echo ""

read -p "请输入选项 (1-5): " choice

case $choice in
    1)
        echo -e "\n开始快速验证测试..."
        ./feed_loadtest -users 50 -bigv 5 -followings 10 -posts 100 -post-c 10 -reads 500 -read-c 20 -duration 30s
        ;;
    2)
        echo -e "\n开始中等规模测试..."
        ./feed_loadtest -users 200 -bigv 20 -followings 30 -posts 1000 -post-c 30 -reads 5000 -read-c 100 -duration 60s
        ;;
    3)
        echo -e "\n开始大规模压测..."
        ./feed_loadtest -users 500 -bigv 50 -followings 50 -posts 3000 -post-c 50 -reads 15000 -read-c 200 -duration 120s
        ;;
    4)
        echo -e "\n开始仅读取压测..."
        ./feed_read_test -users 100 -n 1000 -c 50 -d 60s
        ;;
    5)
        echo -e "\n请输入自定义参数:"
        read -p "用户数量 (default: 100): " users
        read -p "发帖数量 (default: 500): " posts
        read -p "读取数量 (default: 2000): " reads
        read -p "测试时长/秒 (default: 60): " duration

        users=${users:-100}
        posts=${posts:-500}
        reads=${reads:-2000}
        duration=${duration:-60}

        echo -e "\n开始自定义测试..."
        ./feed_loadtest -users $users -posts $posts -reads $reads -duration "${duration}s"
        ;;
    *)
        echo "无效选项，退出"
        exit 1
        ;;
esac

echo -e "\n测试完成！"
