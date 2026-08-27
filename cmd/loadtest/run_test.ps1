# Feed推拉结合架构 - 压力测试执行脚本

Write-Host "=== Feed推拉结合架构 - 压力测试 ===" -ForegroundColor Cyan
Write-Host ""

# 检查服务是否启动
Write-Host "检查服务状态..." -ForegroundColor Yellow
$etcdRunning = Test-Connection -ComputerName 127.0.0.1 -Port 2379 -Quiet -ErrorAction SilentlyContinue
$redisRunning = Test-Connection -ComputerName 127.0.0.1 -Port 6379 -Quiet -ErrorAction SilentlyContinue

if (-not $etcdRunning) {
    Write-Host "❌ Etcd未运行 (127.0.0.1:2379)" -ForegroundColor Red
    Write-Host "请先启动Etcd服务" -ForegroundColor Yellow
    exit 1
}

if (-not $redisRunning) {
    Write-Host "❌ Redis未运行 (127.0.0.1:6379)" -ForegroundColor Red
    Write-Host "请先启动Redis服务" -ForegroundColor Yellow
    exit 1
}

Write-Host "✓ Etcd运行中" -ForegroundColor Green
Write-Host "✓ Redis运行中" -ForegroundColor Green
Write-Host ""

# 选择测试场景
Write-Host "请选择测试场景:" -ForegroundColor Cyan
Write-Host "1. 快速验证测试 (50用户, 100帖子, 500读取, 30秒)"
Write-Host "2. 中等规模测试 (200用户, 1000帖子, 5000读取, 60秒)"
Write-Host "3. 大规模压测 (500用户, 3000帖子, 15000读取, 120秒)"
Write-Host "4. 仅读取压测 (1000次读取, 50并发)"
Write-Host "5. 自定义参数"
Write-Host ""

$choice = Read-Host "请输入选项 (1-5)"

switch ($choice) {
    "1" {
        Write-Host "`n开始快速验证测试..." -ForegroundColor Cyan
        .\feed_loadtest.exe -users 50 -bigv 5 -followings 10 -posts 100 -post-c 10 -reads 500 -read-c 20 -duration 30s
    }
    "2" {
        Write-Host "`n开始中等规模测试..." -ForegroundColor Cyan
        .\feed_loadtest.exe -users 200 -bigv 20 -followings 30 -posts 1000 -post-c 30 -reads 5000 -read-c 100 -duration 60s
    }
    "3" {
        Write-Host "`n开始大规模压测..." -ForegroundColor Cyan
        .\feed_loadtest.exe -users 500 -bigv 50 -followings 50 -posts 3000 -post-c 50 -reads 15000 -read-c 200 -duration 120s
    }
    "4" {
        Write-Host "`n开始仅读取压测..." -ForegroundColor Cyan
        .\feed_read_test.exe -users 100 -n 1000 -c 50 -d 60s
    }
    "5" {
        Write-Host "`n请输入自定义参数:" -ForegroundColor Cyan
        $users = Read-Host "用户数量 (default: 100)"
        $posts = Read-Host "发帖数量 (default: 500)"
        $reads = Read-Host "读取数量 (default: 2000)"
        $duration = Read-Host "测试时长/秒 (default: 60)"

        if ([string]::IsNullOrEmpty($users)) { $users = 100 }
        if ([string]::IsNullOrEmpty($posts)) { $posts = 500 }
        if ([string]::IsNullOrEmpty($reads)) { $reads = 2000 }
        if ([string]::IsNullOrEmpty($duration)) { $duration = 60 }

        Write-Host "`n开始自定义测试..." -ForegroundColor Cyan
        .\feed_loadtest.exe -users $users -posts $posts -reads $reads -duration "${duration}s"
    }
    default {
        Write-Host "无效选项，退出" -ForegroundColor Red
        exit 1
    }
}

Write-Host "`n测试完成！" -ForegroundColor Green
