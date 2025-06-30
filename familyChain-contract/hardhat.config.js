require("@nomicfoundation/hardhat-toolbox");
require("@nomiclabs/hardhat-ethers");
require("dotenv").config();

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.19",
  networks: {
    hardhat: {
      chainId: 1337
    },
    localhost: {
      url: "http://127.0.0.1:8545",
      accounts: [
        "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" // 默认测试账户私钥
      ]
    },
    sepolia: {
      url: process.env.BLOCKCHAIN_RPC_URL,
      accounts: [process.env.BLOCKCHAIN_PRIVATE_KEY_PARENT, process.env.BLOCKCHAIN_PRIVATE_KEY_CHILD], // 推荐用 .env 管理私钥
      chainId: 11155111
    },
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
