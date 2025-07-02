const fs = require('fs');
const path = require('path');

// 定义路径
const artifactsDir = path.join(__dirname, '../artifacts/solidity');
const abiDir = path.join(__dirname, '../abi');

// 确保 abi 目录存在
if (!fs.existsSync(abiDir)) {
  fs.mkdirSync(abiDir, { recursive: true });
}

// 合约名称列表
const contracts = [
  'TaskRegistry',
  'FamilyRegistry',
  'RewardToken',
  'RewardRegistry'
];

console.log('开始提取 ABI 文件...');

// 遍历处理每个合约
contracts.forEach(contractName => {
  const contractDir = path.join(artifactsDir, `${contractName}.sol`);
  const contractJsonPath = path.join(contractDir, `${contractName}.json`);
  
  try {
    // 检查合约文件是否存在
    if (fs.existsSync(contractJsonPath)) {
      console.log(`处理合约: ${contractName}`);
      
      // 读取合约 JSON 文件
      const contractJson = require(contractJsonPath);
      
      // 提取 ABI
      const abi = contractJson.abi;
      
      // 保存 ABI 到单独文件
      const abiPath = path.join(abiDir, `${contractName}.json`);
      fs.writeFileSync(abiPath, JSON.stringify(abi, null, 2));
      
      console.log(`✅ 成功保存 ${contractName} 的 ABI 到: ${abiPath}`);
    } else {
      console.log(`⚠️ 找不到合约文件: ${contractJsonPath}`);
    }
  } catch (error) {
    console.error(`❌ 处理 ${contractName} 时出错:`, error);
  }
});

console.log('ABI 提取完成! 文件保存在:', abiDir); 