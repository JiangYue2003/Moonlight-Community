# Docker 环境快速启动脚本 (Windows PowerShell)

Write-Host "=========================================" -ForegroundColor Cyan
Write-Host "  知光 Docker 环境启动" -ForegroundColor Cyan
Write-Host "=========================================" -ForegroundColor Cyan
Write-Host ""

# 检查 Docker 是否运行
try {
    docker info | Out-Null
    Write-Host "✅ Docker 运行正常" -ForegroundColor Green
} catch {
    Write-Host "❌ Docker 未运行，请先启动 Docker Desktop" -ForegroundColor Red
    exit 1
}

Write-Host ""

# 选择部署模式
Write-Host "请选择部署模式："
Write-Host "1) 测试模式（基础设施 + 核心服务）"
Write-Host "2) 完整模式（所有服务）"
$mode = Read-Host "请输入选择 [1/2]"

if ($mode -eq "2") {
    $ComposeFile = "deploy/compose/docker-compose.full.yml"
    Write-Host ""
    Write-Host "📦 使用完整部署模式" -ForegroundColor Yellow
} else {
    $ComposeFile = "deploy/compose/docker-compose.test.yml"
    Write-Host ""
    Write-Host "📦 使用测试部署模式" -ForegroundColor Yellow
}

Write-Host ""
Write-Host "🔨 开始构建镜像..." -ForegroundColor Yellow
docker compose -f $ComposeFile build

if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ 构建失败" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "🚀 启动服务..." -ForegroundColor Yellow
docker compose -f $ComposeFile up -d

if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ 启动失败" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "⏳ 等待服务就绪（30秒）..." -ForegroundColor Yellow
Start-Sleep -Seconds 30

Write-Host ""
Write-Host "📊 检查服务状态..." -ForegroundColor Yellow
docker compose -f $ComposeFile ps

Write-Host ""
Write-Host "✅ 部署完成！" -ForegroundColor Green
Write-Host ""
Write-Host "📖 常用命令：" -ForegroundColor Cyan
Write-Host "  查看日志: docker compose -f $ComposeFile logs -f gateway"
Write-Host "  查看状态: docker compose -f $ComposeFile ps"
Write-Host "  停止服务: docker compose -f $ComposeFile down"
Write-Host ""
Write-Host "🔗 服务地址：" -ForegroundColor Cyan
Write-Host "  Gateway: http://localhost:8080"
Write-Host "  MySQL:   localhost:3306 (root/Zz123456)"
Write-Host "  Redis:   localhost:6379"
Write-Host "  Etcd:    localhost:2379"
Write-Host ""
Write-Host "🧪 测试服务发现：" -ForegroundColor Cyan
Write-Host '  curl http://localhost:8080/api/v1/counter/knowpost/test'
Write-Host ""
