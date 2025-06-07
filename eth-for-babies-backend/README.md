# Family Task Chain Backend

一个基于区块链的家庭任务管理系统后端API，使用Go语言和Gin框架开发。

## 功能特性

- 🔐 基于以太坊钱包的身份认证
- 👨‍👩‍👧‍👦 家庭管理系统
- 📝 任务创建和分配
- 🎯 任务完成和奖励机制
- 💰 区块链代币奖励集成
- 📊 进度统计和报告
- 🔒 JWT身份验证
- 🗄️ SQLite数据库存储

## 技术栈

- **语言**: Go 1.21+
- **框架**: Gin Web Framework
- **数据库**: SQLite (GORM ORM)
- **区块链**: Ethereum (go-ethereum)
- **认证**: JWT + 以太坊签名验证
- **配置**: 环境变量

## 项目结构

```
eth-for-babies-backend/
├── cmd/
│   └── server/
│       └── main.go              # 应用程序入口
├── internal/
│   ├── api/
│   │   ├── handlers/            # HTTP处理器
│   │   ├── middleware/          # 中间件
│   │   └── routes/              # 路由定义
│   ├── config/                  # 配置管理
│   ├── models/                  # 数据模型
│   ├── repository/              # 数据访问层
│   ├── services/                # 业务逻辑层
│   └── utils/                   # 工具函数
├── .env.example                 # 环境变量示例
├── go.mod                       # Go模块定义
└── README.md                    # 项目文档
```

## 快速开始

### 1. 环境要求

- Go 1.21 或更高版本
- Git

### 2. 克隆项目

```bash
git clone <repository-url>
cd eth-for-babies-backend
```

### 3. 安装依赖

```bash
go mod download
```

### 4. 配置环境变量

```bash
cp .env.example .env
```

编辑 `.env` 文件，配置必要的环境变量：

```env
# 服务器配置
PORT=8080
GIN_MODE=debug
JWT_SECRET=your-super-secret-jwt-key

# 数据库配置
DB_DRIVER=sqlite
DB_DSN=./data/family_task_chain.db

# 区块链配置
BLOCKCHAIN_RPC_URL=https://sepolia.infura.io/v3/your-infura-project-id
BLOCKCHAIN_PRIVATE_KEY=your-private-key
BLOCKCHAIN_CONTRACT_ADDRESS=0x...
BLOCKCHAIN_CHAIN_ID=11155111
```

### 5. 运行应用

```bash
go run cmd/server/main.go
```

服务器将在 `http://localhost:8080` 启动。

### 6. 健康检查

```bash
curl http://localhost:8080/health
```

## API 文档

### 认证相关

#### 获取登录随机数
```http
GET /api/v1/auth/nonce/:wallet_address
```

#### 用户登录
```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "wallet_address": "0x...",
  "signature": "0x...",
  "nonce": "random-nonce"
}
```

#### 用户注册
```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "wallet_address": "0x...",
  "role": "parent",
  "signature": "0x...",
  "nonce": "random-nonce"
}
```

### 家庭管理

#### 创建家庭
```http
POST /api/v1/families
Authorization: Bearer <jwt-token>
Content-Type: application/json

{
  "name": "我的家庭"
}
```

#### 获取家庭列表
```http
GET /api/v1/families
Authorization: Bearer <jwt-token>
```

### 孩子管理

#### 添加孩子
```http
POST /api/v1/children
Authorization: Bearer <jwt-token>
Content-Type: application/json

{
  "name": "小明",
  "age": 8,
  "wallet_address": "0x...",
  "family_id": 1
}
```

#### 获取孩子进度
```http
GET /api/v1/children/:id/progress
Authorization: Bearer <jwt-token>
```

### 任务管理

#### 创建任务
```http
POST /api/v1/tasks
Authorization: Bearer <jwt-token>
Content-Type: application/json

{
  "title": "整理房间",
  "description": "把房间收拾干净",
  "reward_amount": 10.0,
  "difficulty": "easy",
  "assigned_child_id": 1,
  "due_date": "2024-01-15T18:00:00Z"
}
```

#### 完成任务
```http
POST /api/v1/tasks/:id/complete
Authorization: Bearer <jwt-token>
Content-Type: application/json

{
  "proof": "任务完成证明"
}
```

#### 批准任务
```http
POST /api/v1/tasks/:id/approve
Authorization: Bearer <jwt-token>
```

### 智能合约交互

#### 获取余额
```http
GET /api/v1/contracts/balance/:address
Authorization: Bearer <jwt-token>
```

#### 转账
```http
POST /api/v1/contracts/transfer
Authorization: Bearer <jwt-token>
Content-Type: application/json

{
  "to": "0x...",
  "amount": "10.0"
}
```

## 数据模型

### 用户 (User)
- ID
- 钱包地址
- 角色 (parent/child)
- 登录随机数
- 创建时间

### 家庭 (Family)
- ID
- 家庭名称
- 家长地址
- 创建时间

### 孩子 (Child)
- ID
- 姓名
- 年龄
- 钱包地址
- 头像
- 家长地址
- 家庭ID
- 完成任务数
- 总奖励

### 任务 (Task)
- ID
- 标题
- 描述
- 奖励金额
- 难度等级
- 状态
- 分配的孩子ID
- 创建者地址
- 截止日期
- 完成证明

## 开发指南

### 添加新的API端点

1. 在 `internal/models/` 中定义数据模型
2. 在 `internal/repository/` 中实现数据访问层
3. 在 `internal/services/` 中实现业务逻辑
4. 在 `internal/api/handlers/` 中实现HTTP处理器
5. 在 `internal/api/routes/` 中注册路由

### 数据库迁移

应用启动时会自动执行数据库迁移。如需手动迁移：

```go
db.AutoMigrate(&models.User{}, &models.Family{}, &models.Child{}, &models.Task{})
```

### 测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./internal/services/

# 运行测试并显示覆盖率
go test -cover ./...
```

## 部署

### Docker 部署

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
```

### 生产环境配置

1. 设置 `GIN_MODE=release`
2. 使用强密码作为 `JWT_SECRET`
3. 配置适当的数据库连接
4. 设置正确的区块链网络配置
5. 配置HTTPS和反向代理

## 安全注意事项

- 🔐 私钥和JWT密钥必须安全存储
- 🛡️ 生产环境中禁用调试模式
- 🔒 使用HTTPS传输敏感数据
- ✅ 验证所有用户输入
- 🚫 不要在日志中记录敏感信息

## 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 联系方式

如有问题或建议，请创建 Issue 或联系项目维护者。