# familyChain-contract

智能合约目录，包含FamilyChain项目的所有区块链相关代码。

## 目录结构

- `abi/` - 合约的ABI接口文件
- `solidity/` - Solidity智能合约源代码
  - `FamilyRegistry.sol` - 家庭注册管理合约
  - `RewardToken.sol` - 奖励代币合约
  - `TaskRegistry.sol` - 任务管理合约
- `deploy/` - 合约部署脚本
- `scripts/` - 合约测试和部署相关脚本
- `test/` - 合约测试用例

## 安装

```bash
npm install
```

## 编译合约

```bash
npx hardhat compile
```

## 部署合约

```bash
npx hardhat run scripts/deploy.js --network <network-name>
```

## 测试

```bash
npx hardhat test
``` 