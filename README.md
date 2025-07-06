# familyChain

## 📖 项目介绍与愿景

familyChain 是一个基于以太坊区块链的家庭任务管理和奖励系统，旨在通过区块链技术培养孩子的责任感和金融素养。

### 🌟 核心理念

- **区块链教育**: 让孩子在完成家务任务的同时，自然地了解区块链和加密货币的基础知识
- **激励机制**: 通过智能合约自动执行的奖励系统，激励孩子按时完成任务
- **家庭协作**: 增强家庭成员之间的互动，让家长和孩子共同参与到任务的设定和完成过程中
- **金融素养**: 培养孩子对数字资产的理解和管理能力，为未来的金融世界做准备

### 💡 主要功能 v1.0

- **任务管理**: 家长可以创建、分配和跟踪家务任务
- **区块链奖励**: 任务完成后通过智能合约自动发放ETH奖励
- **成长记录**: 记录孩子的任务完成历史和获得的奖励
- **钱包集成**: 与MetaMask等主流钱包无缝集成，让孩子学习管理自己的数字资产

### 思维导图
![familyChain](/familyChain-pic/FamilyChain.png)

### 🔮 未来展望

familyChain 项目有着广阔的发展空间，以下是我们规划的一些有趣功能和玩法：

- **实物奖励集成**:v2.0
  - 连接数字资产与实物奖品
  - 创建家庭专属"商店"，设置实物奖励兑换机制，家长购买给孩子
  - 在孩子完成任务的时候，由任务合约来铸造代币，发放给孩子，孩子端要求有代币展示页面
  - 家长端新增 创建实物兑换页面，自由度极高，由家长将实物照片，标明代币价格后，上传到系统中，比如玩具、书籍、电影票等等
  - 孩子端新增奖品浏览页面，允许用代币兑换，兑换成功后，需要存放到孩子的兑换成果展示页面

- **NFT成就系统**: 
  - 完成特定类型任务可获得独特的NFT徽章
  - 设计收藏NFT，鼓励孩子持续参与

- **家庭DAO治理**:
  - 创建家庭专属DAO，让孩子参与家庭决策投票
  - 使用代币投票系统决定家庭活动、周末计划或晚餐选择
  - 教导孩子理解去中心化组织和民主决策过程

- **技能树与成长路径**:
  - 为孩子创建去中心化身份(DID)，记录成长
  - 设计不同领域的技能树（家务、学习、创意等）
  - 完成特定技能路径解锁更高级任务和奖励
  - 生成可视化成长报告，展示孩子的进步历程

- **社交与竞争元素**:
  - 允许不同家庭之间安全地比较和竞争
  - 创建社区排行榜，展示完成任务最多的孩子
  - 组织家庭间协作任务，培养团队合作精神

- **多链集成与跨链体验**:
  - 扩展到多个区块链网络，让孩子了解不同区块链特性
  - 提供简化的跨链操作体验
  - 在不同链上设置不同类型的任务和奖励

这些创新功能将帮助familyChain不仅成为一个任务管理工具，更成为孩子了解区块链、金融知识和责任感的综合教育平台。

## 🚀 安装与使用指南

### 系统要求

- **Go** (1.19+) - 后端开发
- **Node.js** (16+) - 前端开发
- **Yarn** - 包管理器
- **MetaMask** - 以太坊钱包
- **Sepolia测试网** - 用于开发测试

### 快速启动

使用一键启动脚本同时运行前端和后端服务：

```bash
# 克隆仓库
git clone https://github.com/yourusername/familyChain.git
cd familyChain

# 启动所有服务
./start.sh
```

启动后，访问：
- **前端应用**: http://localhost:5173
- **后端 API**: http://localhost:8080

### 手动安装步骤

#### 1. 后端设置

```bash
cd familyChain-backend

# 安装依赖
go mod download

# 配置环境变量
cp .env.example .env
# 编辑.env文件，设置必要的配置项

# 启动后端服务
go run cmd/server/main.go
```

#### 2. 前端设置

```bash
cd familyChain

# 安装依赖
yarn install

# 配置环境变量
cp .env.example .env
# 编辑.env文件，设置必要的配置项，包括WalletConnect projectId

# 启动开发服务器
yarn dev
```
#### 3. 智能合约设置

```bash
cd familyChain-contract

# 配置环境变量
cp .env.example .env

```

#### 4. 钱包配置

1. 安装MetaMask浏览器扩展
2. 创建或导入钱包
3. 连接到Sepolia测试网
4. 获取测试网ETH (可从水龙头获取)

### 停止服务

```bash
./stop.sh
```

### 故障排除

如遇到问题，请尝试：

1. 检查日志文件 `logs/backend.log` 和 `logs/frontend.log`
2. 确保MetaMask已连接到Sepolia测试网
3. 验证环境变量配置是否正确
4. 重新安装依赖

```bash
# 重新安装Go依赖
cd familyChain-backend
go mod tidy

# 重新安装前端依赖
cd familyChain
rm -rf node_modules
yarn install
```

## 📁 项目结构

```
.
├── familyChain/              # 前端项目 (React + TypeScript)
├── familyChain-backend/      # 后端项目 (Go + Gin)
├── familyChain-contract/     # 智能合约项目 (Solidity + Hardhat)
├── logs/                     # 运行日志
├── start.sh                  # 启动脚本
└── stop.sh                   # 停止脚本
```

## �� 许可证

MIT License
