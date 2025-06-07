#!/bin/bash

# ETH for Babies 项目停止脚本
# 停止前端和后端服务

echo "🛑 停止 ETH for Babies 项目..."
echo "================================"

# 获取脚本所在目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
LOG_DIR="$SCRIPT_DIR/logs"

# 检查 PID 文件是否存在
BACKEND_PID_FILE="$LOG_DIR/backend.pid"
FRONTEND_PID_FILE="$LOG_DIR/frontend.pid"

STOPPED_COUNT=0

# 停止后端服务
if [ -f "$BACKEND_PID_FILE" ]; then
    BACKEND_PID=$(cat "$BACKEND_PID_FILE")
    if kill -0 $BACKEND_PID 2>/dev/null; then
        echo "🔧 停止后端服务 (PID: $BACKEND_PID)..."
        kill $BACKEND_PID
        sleep 2
        
        # 如果进程仍在运行，强制杀死
        if kill -0 $BACKEND_PID 2>/dev/null; then
            echo "⚠️  强制停止后端服务..."
            kill -9 $BACKEND_PID 2>/dev/null
        fi
        
        echo "✅ 后端服务已停止"
        STOPPED_COUNT=$((STOPPED_COUNT + 1))
    else
        echo "ℹ️  后端服务未运行"
    fi
    rm -f "$BACKEND_PID_FILE"
else
    echo "ℹ️  未找到后端服务 PID 文件"
fi

# 停止前端服务
if [ -f "$FRONTEND_PID_FILE" ]; then
    FRONTEND_PID=$(cat "$FRONTEND_PID_FILE")
    if kill -0 $FRONTEND_PID 2>/dev/null; then
        echo "🎨 停止前端服务 (PID: $FRONTEND_PID)..."
        kill $FRONTEND_PID
        sleep 2
        
        # 如果进程仍在运行，强制杀死
        if kill -0 $FRONTEND_PID 2>/dev/null; then
            echo "⚠️  强制停止前端服务..."
            kill -9 $FRONTEND_PID 2>/dev/null
        fi
        
        echo "✅ 前端服务已停止"
        STOPPED_COUNT=$((STOPPED_COUNT + 1))
    else
        echo "ℹ️  前端服务未运行"
    fi
    rm -f "$FRONTEND_PID_FILE"
else
    echo "ℹ️  未找到前端服务 PID 文件"
fi

# 额外检查：通过端口杀死可能遗留的进程
echo "🔍 检查端口占用情况..."

# 检查 8080 端口（后端）
BACKEND_PORT_PID=$(lsof -ti:8080 2>/dev/null)
if [ ! -z "$BACKEND_PORT_PID" ]; then
    echo "⚠️  发现 8080 端口被占用 (PID: $BACKEND_PORT_PID)，正在清理..."
    kill -9 $BACKEND_PORT_PID 2>/dev/null
    echo "✅ 8080 端口已清理"
    STOPPED_COUNT=$((STOPPED_COUNT + 1))
fi

# 检查 5173 端口（前端）
FRONTEND_PORT_PID=$(lsof -ti:5173 2>/dev/null)
if [ ! -z "$FRONTEND_PORT_PID" ]; then
    echo "⚠️  发现 5173 端口被占用 (PID: $FRONTEND_PORT_PID)，正在清理..."
    kill -9 $FRONTEND_PORT_PID 2>/dev/null
    echo "✅ 5173 端口已清理"
    STOPPED_COUNT=$((STOPPED_COUNT + 1))
fi

echo ""
if [ $STOPPED_COUNT -gt 0 ]; then
    echo "🎉 成功停止 $STOPPED_COUNT 个服务"
else
    echo "ℹ️  没有发现正在运行的服务"
fi

echo "================================"
echo "✅ 停止操作完成"

# 显示日志文件位置（如果存在）
if [ -d "$LOG_DIR" ]; then
    echo ""
    echo "📝 日志文件保留在: $LOG_DIR"
    echo "   - backend.log: 后端日志"
    echo "   - frontend.log: 前端日志"
    echo ""
    echo "💡 清理日志文件: rm -rf $LOG_DIR"
fi