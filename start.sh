#!/bin/bash

# familyChain 项目启动脚本
# 同时启动前端和后端服务

echo "🚀 启动 familyChain 项目..."
echo "================================"

# 检查是否安装了必要的依赖
echo "📋 检查依赖..."

# 检查 Go
if ! command -v go &> /dev/null; then
    echo "❌ Go 未安装，请先安装 Go"
    exit 1
fi

# 检查 Node.js
if ! command -v node &> /dev/null; then
    echo "❌ Node.js 未安装，请先安装 Node.js"
    exit 1
fi

# 检查 yarn
if ! command -v yarn &> /dev/null; then
    echo "❌ Yarn 未安装，请先安装 Yarn"
    exit 1
fi

echo "✅ 依赖检查完成"
echo ""

# 获取脚本所在目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BACKEND_DIR="$SCRIPT_DIR/familyChain-backend"
FRONTEND_DIR="$SCRIPT_DIR/familyChain"

# 检查项目目录是否存在
if [ ! -d "$BACKEND_DIR" ]; then
    echo "❌ 后端项目目录不存在: $BACKEND_DIR"
    exit 1
fi

if [ ! -d "$FRONTEND_DIR" ]; then
    echo "❌ 前端项目目录不存在: $FRONTEND_DIR"
    exit 1
fi

# 创建日志目录
LOG_DIR="$SCRIPT_DIR/logs"
mkdir -p "$LOG_DIR"

# 启动后端服务
echo "🔧 启动后端服务 (端口: 8080)..."
cd "$BACKEND_DIR"

# 检查并安装 Go 依赖
if [ ! -f "go.sum" ] || [ "go.mod" -nt "go.sum" ]; then
    echo "📦 安装 Go 依赖..."
    go mod download
fi

# 启动后端服务（后台运行）
nohup go run cmd/server/main.go > "$LOG_DIR/backend.log" 2>&1 &
BACKEND_PID=$!
echo "✅ 后端服务已启动 (PID: $BACKEND_PID)"

# 等待后端服务启动
echo "⏳ 等待后端服务启动..."
sleep 3

# 检查后端服务是否正常运行
if ! kill -0 $BACKEND_PID 2>/dev/null; then
    echo "❌ 后端服务启动失败，请检查日志: $LOG_DIR/backend.log"
    exit 1
fi

# 启动前端服务
echo "🎨 启动前端服务 (端口: 5173)..."
cd "$FRONTEND_DIR"

# 检查并安装前端依赖
if [ ! -d "node_modules" ] || [ "package.json" -nt "node_modules" ]; then
    echo "📦 安装前端依赖..."
    yarn install
fi

# 启动前端服务（后台运行）
nohup yarn dev > "$LOG_DIR/frontend.log" 2>&1 &
FRONTEND_PID=$!
echo "✅ 前端服务已启动 (PID: $FRONTEND_PID)"

# 等待前端服务启动
echo "⏳ 等待前端服务启动..."
sleep 5

# 检查前端服务是否正常运行
if ! kill -0 $FRONTEND_PID 2>/dev/null; then
    echo "❌ 前端服务启动失败，请检查日志: $LOG_DIR/frontend.log"
    # 清理后端进程
    kill $BACKEND_PID 2>/dev/null
    exit 1
fi

echo ""
echo "🎉 所有服务启动成功！"
echo "================================"
echo "📱 前端地址: http://localhost:5173"
echo "🔧 后端地址: http://localhost:8080"
echo "📋 健康检查: http://localhost:8080/health"
echo "📝 日志目录: $LOG_DIR"
echo ""
echo "进程信息:"
echo "  后端 PID: $BACKEND_PID"
echo "  前端 PID: $FRONTEND_PID"
echo ""
echo "💡 使用以下命令停止服务:"
echo "  kill $BACKEND_PID $FRONTEND_PID"
echo ""
echo "📊 实时查看日志:"
echo "  后端: tail -f $LOG_DIR/backend.log"
echo "  前端: tail -f $LOG_DIR/frontend.log"
echo ""
echo "按 Ctrl+C 停止所有服务..."

# 创建 PID 文件用于后续管理
echo $BACKEND_PID > "$LOG_DIR/backend.pid"
echo $FRONTEND_PID > "$LOG_DIR/frontend.pid"

# 等待用户中断
trap 'echo "\n🛑 正在停止服务..."; kill $BACKEND_PID $FRONTEND_PID 2>/dev/null; rm -f "$LOG_DIR/backend.pid" "$LOG_DIR/frontend.pid"; echo "✅ 所有服务已停止"; exit 0' INT

# 监控进程状态
while true; do
    if ! kill -0 $BACKEND_PID 2>/dev/null; then
        echo "❌ 后端服务异常退出"
        kill $FRONTEND_PID 2>/dev/null
        rm -f "$LOG_DIR/backend.pid" "$LOG_DIR/frontend.pid"
        exit 1
    fi
    
    if ! kill -0 $FRONTEND_PID 2>/dev/null; then
        echo "❌ 前端服务异常退出"
        kill $BACKEND_PID 2>/dev/null
        rm -f "$LOG_DIR/backend.pid" "$LOG_DIR/frontend.pid"
        exit 1
    fi
    
    sleep 5
done