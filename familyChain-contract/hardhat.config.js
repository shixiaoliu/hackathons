require("@nomicfoundation/hardhat-toolbox");
require("@nomiclabs/hardhat-ethers");

// 加载环境变量
require("dotenv").config();

// 检查是否存在必要的环境变量
const SEPOLIA_RPC_URL = process.env.BLOCKCHAIN_RPC_URL;
const SEPOLIA_PRIVATE_KEY_PARENT = process.env.BLOCKCHAIN_PRIVATE_KEY_PARENT;
const SEPOLIA_PRIVATE_KEY_CHILD = process.env.BLOCKCHAIN_PRIVATE_KEY_CHILD;

// 检查环境变量是否存在
if (!SEPOLIA_RPC_URL || !SEPOLIA_PRIVATE_KEY_PARENT) {
  console.warn(
    "\n⚠️ 警告: 缺少 Sepolia 网络配置的环境变量! ⚠️\n" +
    "请创建 .env 文件并添加以下内容:\n" +
    "BLOCKCHAIN_RPC_URL=https://sepolia.infura.io/v3/YOUR_INFURA_KEY\n" +
    "BLOCKCHAIN_PRIVATE_KEY_PARENT=your_private_key\n" +
    "BLOCKCHAIN_PRIVATE_KEY_CHILD=child_private_key (可选)\n"
  );
}

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  networks: {
    hardhat: {
      chainId: 1337
    },
    // 仅在环境变量存在的情况下配置 Sepolia
    ...(SEPOLIA_RPC_URL && SEPOLIA_PRIVATE_KEY_PARENT ? {
      sepolia: {
        url: SEPOLIA_RPC_URL,
        accounts: [
          SEPOLIA_PRIVATE_KEY_PARENT,
          // 如果有子账户私钥，则添加，否则使用父账户私钥
          SEPOLIA_PRIVATE_KEY_CHILD || SEPOLIA_PRIVATE_KEY_PARENT
        ],
        chainId: 11155111
      }
    } : {})
  },
  paths: {
    sources: "./solidity",
    artifacts: "./artifacts",
    cache: "./cache"
  },
  mocha: {
    timeout: 40000
  }
};
