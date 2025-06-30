const { ethers } = require("hardhat");
const fs = require("fs");
const path = require("path");

async function main() {
  console.log("Deploying RewardRegistry contract...");

  // 获取 RewardToken 合约地址
  let rewardTokenAddress;
  try {
    const deployedAddressesPath = path.join(__dirname, "..", "deployed_addresses.json");
    const deployedAddresses = JSON.parse(fs.readFileSync(deployedAddressesPath, "utf8"));
    rewardTokenAddress = deployedAddresses.RewardToken;
    
    if (!rewardTokenAddress) {
      throw new Error("RewardToken address not found in deployed_addresses.json");
    }
    
    console.log(`Using existing RewardToken at address: ${rewardTokenAddress}`);
  } catch (error) {
    console.error("Error reading RewardToken address:", error.message);
    console.log("Please deploy RewardToken first using deploy.js script");
    process.exit(1);
  }

  // 获取部署者账户
  const [deployer] = await ethers.getSigners();
  console.log(`Deploying with account: ${deployer.address}`);

  // 部署 RewardRegistry 合约
  const RewardRegistry = await ethers.getContractFactory("RewardRegistry");
  const rewardRegistry = await RewardRegistry.deploy(rewardTokenAddress);
  await rewardRegistry.deployed();

  console.log(`RewardRegistry deployed to: ${rewardRegistry.address}`);

  // 更新部署地址文件
  const deployedAddressesPath = path.join(__dirname, "..", "deployed_addresses.json");
  let deployedAddresses = {};
  
  if (fs.existsSync(deployedAddressesPath)) {
    deployedAddresses = JSON.parse(fs.readFileSync(deployedAddressesPath, "utf8"));
  }
  
  deployedAddresses.RewardRegistry = rewardRegistry.address;
  
  fs.writeFileSync(
    deployedAddressesPath,
    JSON.stringify(deployedAddresses, null, 2)
  );
  
  console.log("Deployment address saved to deployed_addresses.json");

  // 将 RewardRegistry 设置为 RewardToken 的授权铸造者
  console.log("Setting RewardRegistry as authorized minter for RewardToken...");
  const RewardToken = await ethers.getContractFactory("RewardToken");
  const rewardToken = await RewardToken.attach(rewardTokenAddress);
  
  const tx = await rewardToken.addMinter(rewardRegistry.address);
  await tx.wait();
  
  console.log(`RewardRegistry (${rewardRegistry.address}) is now an authorized minter for RewardToken`);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  }); 