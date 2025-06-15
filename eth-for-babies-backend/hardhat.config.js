require("@nomicfoundation/hardhat-toolbox");
require("@nomiclabs/hardhat-ethers");
require("dotenv").config();

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  networks: {
    sepolia: {
      url: process.env.BLOCKCHAIN_RPC_URL || "https://sepolia.infura.io/v3/YOUR-PROJECT-ID",
      accounts: [process.env.BLOCKCHAIN_PRIVATE_KEY], // 推荐用 .env 管理私钥
      chainId: 11155111
    },
  }
};
