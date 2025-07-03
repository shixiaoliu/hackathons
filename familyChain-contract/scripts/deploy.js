const { ethers } = require("hardhat");
const fs = require("fs");
const path = require("path");

async function main() {
    const [deployer] = await ethers.getSigners();
    console.log("Deploying contracts with the account:", deployer.address);
    
    // 存储所有已部署合约地址
    const deployedAddresses = {};
    
    // 1. 部署 TaskRegistry
    console.log("\n--- Deploying TaskRegistry ---");
    const TaskRegistry = await ethers.getContractFactory("TaskRegistry");
    const taskRegistry = await TaskRegistry.deploy();
    await taskRegistry.deployed();
    console.log("TaskRegistry deployed to:", taskRegistry.address);
    deployedAddresses.TaskRegistry = taskRegistry.address;
    
    // // 2. 部署 RewardToken
    // console.log("\n--- Deploying RewardToken ---");
    // const RewardToken = await ethers.getContractFactory("RewardToken");
    // const rewardToken = await RewardToken.deploy("RewardToken", "RWT");
    // await rewardToken.deployed();
    // console.log("RewardToken deployed to:", rewardToken.address);
    // deployedAddresses.RewardToken = rewardToken.address;
    
    // // 3. 部署 RewardRegistry (使用已部署的 RewardToken 地址)
    // console.log("\n--- Deploying RewardRegistry ---");
    // const RewardRegistry = await ethers.getContractFactory("RewardRegistry");
    // const rewardRegistry = await RewardRegistry.deploy(rewardToken.address);
    // await rewardRegistry.deployed();
    // console.log("RewardRegistry deployed to:", rewardRegistry.address);
    // deployedAddresses.RewardRegistry = rewardRegistry.address;
    
    // // 4. 设置 RewardRegistry 为 RewardToken 的授权铸造者
    // console.log("\n--- Setting up contract permissions ---");
    // console.log("Setting RewardRegistry as authorized minter for RewardToken...");
    // const tx = await rewardToken.addMinter(rewardRegistry.address);
    // await tx.wait();
    // console.log(`RewardRegistry (${rewardRegistry.address}) is now an authorized minter for RewardToken`);
    
    // // 保存所有合约地址到 JSON 文件
    // const deployedAddressesPath = path.join(__dirname, "..", "deployed_addresses.json");
    // fs.writeFileSync(
    //     deployedAddressesPath,
    //     JSON.stringify(deployedAddresses, null, 2)
    // );
    // console.log("\n--- All contract addresses saved to deployed_addresses.json ---");
    // console.log(deployedAddresses);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });