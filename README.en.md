# familyChain

[English](README.en.md) | [中文](README.md)

## 📖 Project Introduction & Vision

familyChain is a family task management and reward system based on Ethereum blockchain, designed to cultivate children's sense of responsibility and financial literacy through blockchain technology.

### 🌟 Core Philosophy

- **Blockchain Education**: Help children naturally understand blockchain and cryptocurrency basics while completing household tasks
- **Incentive Mechanism**: Motivate children to complete tasks on time through automatically executed rewards via smart contracts
- **Family Collaboration**: Enhance interaction among family members, engaging both parents and children in setting and completing tasks
- **Financial Literacy**: Foster children's understanding and management of digital assets, preparing them for the future financial world

### 💡 Main Features v1.0

- **Task Management**: Parents can create, assign, and track household chores
- **Blockchain Rewards**: Automatic ETH rewards via smart contracts upon task completion
- **Growth Records**: Track children's task completion history and earned rewards
- **Wallet Integration**: Seamless integration with mainstream wallets like MetaMask, helping children learn to manage their digital assets

### Mind Map
![familyChain](/familyChain-pic/FamilyChain.png)

### 💡 Physical Reward Integration v2.0

- Connect digital assets with physical prizes
- Create family-exclusive "stores" with physical reward exchange mechanisms purchased by parents
- Task contracts mint tokens for children when they complete tasks, with token display pages on the children's interface
- Parents' interface adds a highly customizable physical exchange page where parents upload photos of physical items with token prices (toys, books, movie tickets, etc.)
- Children's interface adds a prize browsing page for token exchanges, with successful exchanges displayed on the child's achievement showcase page

### Mind Map
![familyChain](/familyChain-pic/familyChain_2.0.png)

### 🔮 Future Outlook

The familyChain project has broad development potential. Here are some interesting features and mechanics we've planned:

- **NFT Achievement System**: 
  - Earn unique NFT badges by completing specific types of tasks
  - Design collectible NFTs to encourage children's continued participation

- **Family DAO Governance**:
  - Create family-exclusive DAOs for children to participate in family decision-making
  - Use token voting systems to decide family activities, weekend plans, or dinner choices
  - Teach children about decentralized organizations and democratic decision processes

- **Skill Trees & Growth Paths**:
  - Create decentralized identities (DIDs) for children to record growth
  - Design skill trees in different areas (housework, learning, creativity, etc.)
  - Unlock higher-level tasks and rewards by completing specific skill paths
  - Generate visualized growth reports showing children's progress

- **Social & Competitive Elements**:
  - Allow safe comparison and competition between different families
  - Create community leaderboards showcasing children who complete the most tasks
  - Organize collaborative tasks between families to foster team spirit

- **Multi-Chain Integration & Cross-Chain Experience**:
  - Expand to multiple blockchain networks, helping children understand different blockchain characteristics
  - Provide simplified cross-chain operation experiences
  - Set up different types of tasks and rewards on different chains

These innovative features will help familyChain become not just a task management tool, but a comprehensive educational platform for children to learn about blockchain, financial knowledge, and responsibility.

## 🚀 Installation & Usage Guide

### System Requirements

- **Go** (1.19+) - Backend development
- **Node.js** (16+) - Frontend development
- **Yarn** - Package manager
- **MetaMask** - Ethereum wallet
- **Sepolia Testnet** - For development testing

### Quick Start

Use the one-click startup script to run both frontend and backend services:

```bash
# Clone repository
git clone https://github.com/yourusername/familyChain.git
cd familyChain

# Start all services
./start.sh
```

After startup, access:
- **Frontend application**: http://localhost:5173
- **Backend API**: http://localhost:8080

### Manual Installation Steps

#### 1. Backend Setup

```bash
cd familyChain-backend

# Install dependencies
go mod download

# Configure environment variables
cp .env.example .env
# Edit the .env file to set necessary configuration items

# Start backend service
go run cmd/server/main.go
```

#### 2. Frontend Setup

```bash
cd familyChain

# Install dependencies
yarn install

# Configure environment variables
cp .env.example .env
# Edit the .env file to set necessary configuration items, including WalletConnect projectId

# Start development server
yarn dev
```
#### 3. Smart Contract Setup

```bash
cd familyChain-contract

# Configure environment variables
cp .env.example .env
```

#### 4. Wallet Configuration

1. Install MetaMask browser extension
2. Create or import wallet
3. Connect to Sepolia testnet
4. Get testnet ETH (available from faucets)

### Stopping Services

```bash
./stop.sh
```

### Troubleshooting

If you encounter issues, try:

1. Check log files `logs/backend.log` and `logs/frontend.log`
2. Ensure MetaMask is connected to Sepolia testnet
3. Verify environment variable configurations are correct
4. Reinstall dependencies

```bash
# Reinstall Go dependencies
cd familyChain-backend
go mod tidy

# Reinstall frontend dependencies
cd familyChain
rm -rf node_modules
yarn install
```

## 📁 Project Structure

```
.
├── familyChain/              # Frontend project (React + TypeScript)
├── familyChain-backend/      # Backend project (Go + Gin)
├── familyChain-contract/     # Smart contract project (Solidity + Hardhat)
├── logs/                     # Running logs
├── start.sh                  # Startup script
└── stop.sh                   # Shutdown script
```

## 📄 License

MIT License 