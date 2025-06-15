const { ethers } = require("hardhat");

async function main() {
    const [deployer] = await ethers.getSigners();
    console.log("Deploying contracts with the account:", deployer.address);
  
    // 部署 TaskRegistry
    const TaskRegistry = await ethers.getContractFactory("TaskRegistry");
    const taskRegistry = await TaskRegistry.deploy();
  
    await taskRegistry.deployed();
  
    console.log("TaskRegistry deployed to:", taskRegistry.address);
  
    // // 如需部署 FamilyRegistry，取消注释
    // const FamilyRegistry = await ethers.getContractFactory("FamilyRegistry");
    // const familyRegistry = await FamilyRegistry.deploy();
    // await familyRegistry.deployed();
    // console.log("FamilyRegistry deployed to:", familyRegistry.address);
  
    // // 如需部署 RewardToken，需传入 name 和 symbol
    // const RewardToken = await ethers.getContractFactory("RewardToken");
    // const rewardToken = await RewardToken.deploy("RewardToken", "RWT");
    // await rewardToken.deployed();
    // console.log("RewardToken deployed to:", rewardToken.address);
  }
  
  main().then(() => process.exit(0)).catch((error) => {
    console.error(error);
    process.exit(1);
  });