# ETH for Babies 项目

一个基于以太坊的家庭任务管理和奖励系统，帮助家长通过区块链技术激励孩子完成任务。

## 🚀 快速启动

### 前提条件

确保您的系统已安装以下软件：

- **Go** (1.19+) - 后端开发
- **Node.js** (16+) - 前端开发
- **Yarn** - 包管理器
- **Git** - 版本控制

### 一键启动

使用提供的启动脚本可以同时启动前端和后端服务：

```bash
# 启动所有服务
./start.sh
```

启动后，您可以访问：

- **前端应用**: http://localhost:5173
- **后端 API**: http://localhost:8080
- **健康检查**: http://localhost:8080/health

### 停止服务

```bash
# 停止所有服务
./stop.sh
```

## 📁 项目结构

```
.
├── eth-for-babies/              # 前端项目 (React + TypeScript)
│   ├── src/
│   │   ├── components/          # React 组件
│   │   ├── pages/              # 页面组件
│   │   ├── hooks/              # 自定义 Hooks
│   │   ├── services/           # API 服务
│   │   ├── types/              # TypeScript 类型定义
│   │   └── utils/              # 工具函数
│   ├── package.json
│   └── vite.config.ts
│
├── eth-for-babies-backend/      # 后端项目 (Go + Gin)
│   ├── cmd/server/             # 应用入口
│   ├── internal/
│   │   ├── api/                # API 路由和处理器
│   │   ├── config/             # 配置管理
│   │   ├── models/             # 数据模型
│   │   ├── repository/         # 数据访问层
│   │   ├── services/           # 业务逻辑层
│   │   └── utils/              # 工具函数
│   ├── go.mod
│   └── README.md
│
├── logs/                        # 运行日志 (自动创建)
│   ├── backend.log             # 后端日志
│   ├── frontend.log            # 前端日志
│   ├── backend.pid             # 后端进程 ID
│   └── frontend.pid            # 前端进程 ID
│
├── start.sh                     # 启动脚本
├── stop.sh                      # 停止脚本
└── README.md                    # 项目说明
```

## 🛠️ 手动启动（开发模式）

如果您需要单独启动某个服务或进行开发调试：

### 启动后端

```bash
cd eth-for-babies-backend

# 安装依赖
go mod download

# 启动开发服务器
go run cmd/server/main.go

# 或者构建后运行
go build -o main cmd/server/main.go
./main
```

### 启动前端

```bash
cd eth-for-babies

# 安装依赖
yarn install

# 启动开发服务器
yarn dev

# 构建生产版本
yarn build
```

## 🔧 配置

### 后端配置

复制环境配置文件并根据需要修改：

```bash
cd eth-for-babies-backend
cp .env.example .env
```

主要配置项：

```env
# 服务器配置
PORT=8080
ENVIRONMENT=development

# 数据库配置
DB_DRIVER=sqlite
DB_DSN=./data/app.db

# JWT 密钥
JWT_SECRET=your-secret-key-change-in-production

# 区块链配置
BLOCKCHAIN_RPC_URL=http://localhost:8545
BLOCKCHAIN_PRIVATE_KEY=your-private-key
BLOCKCHAIN_CONTRACT_ADDRESS=your-contract-address
BLOCKCHAIN_CHAIN_ID=1337
```

### 前端配置

```bash
cd eth-for-babies
cp .env.example .env
```

## 📊 日志管理

启动脚本会自动创建日志文件：

```bash
# 实时查看后端日志
tail -f logs/backend.log

# 实时查看前端日志
tail -f logs/frontend.log

# 清理所有日志
rm -rf logs/
```

## 🐛 故障排除

### 端口冲突

如果遇到端口被占用的错误：

```bash
# 查看端口占用情况
lsof -i :8080  # 后端端口
lsof -i :5173  # 前端端口

# 杀死占用端口的进程
kill -9 <PID>

# 或使用停止脚本清理
./stop.sh
```

### 依赖问题

```bash
# 重新安装 Go 依赖
cd eth-for-babies-backend
go mod tidy
go mod download

# 重新安装前端依赖
cd eth-for-babies
rm -rf node_modules yarn.lock
yarn install
```

### 权限问题

```bash
# 确保脚本有执行权限
chmod +x start.sh stop.sh
```

## 🚀 部署

### Docker 部署

后端项目包含 Docker 配置：

```bash
cd eth-for-babies-backend

# 构建镜像
docker build -t eth-for-babies-backend .

# 运行容器
docker run -p 8080:8080 eth-for-babies-backend
```

### 生产环境

1. 设置环境变量为生产模式
2. 配置反向代理（如 Nginx）
3. 设置 HTTPS 证书
4. 配置数据库（PostgreSQL 推荐）
5. 设置监控和日志收集

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

MIT License

---

**快速命令参考：**

```bash
./start.sh          # 启动所有服务
./stop.sh           # 停止所有服务
tail -f logs/*.log  # 查看日志
```