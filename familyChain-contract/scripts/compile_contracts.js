const fs = require('fs');
const path = require('path');
const solc = require('solc');

function ensureDirectoryExistence(dirPath) {
  if (!fs.existsSync(dirPath)) {
    fs.mkdirSync(dirPath, { recursive: true });
  }
}

// 读取合约文件
function readContract(contractPath) {
  return fs.readFileSync(contractPath, 'utf8');
}

// 编译合约
function compileContract(contractName, contractContent, dependencies = {}) {
  const input = {
    language: 'Solidity',
    sources: {
      [contractName]: {
        content: contractContent,
      },
    },
    settings: {
      outputSelection: {
        '*': {
          '*': ['*'],
        },
      },
    },
  };

  // 添加依赖
  for (const [name, content] of Object.entries(dependencies)) {
    input.sources[name] = { content };
  }

  const output = JSON.parse(solc.compile(JSON.stringify(input)));

  if (output.errors) {
    console.error('Compilation errors:');
    output.errors.forEach((error) => {
      console.error(error.formattedMessage);
    });
    if (output.errors.some(error => error.severity === 'error')) {
      throw new Error('Compilation failed');
    }
  }

  return output.contracts[contractName][path.basename(contractName, '.sol')];
}

// 保存ABI到文件
function saveAbi(contractName, abi) {
  const abiDir = path.join(__dirname, '..', 'abi');
  ensureDirectoryExistence(abiDir);
  
  const abiPath = path.join(abiDir, `${contractName}.json`);
  fs.writeFileSync(abiPath, JSON.stringify(abi, null, 2));
  console.log(`ABI saved to ${abiPath}`);
}

// 主函数
async function main() {
  try {
    const contractsDir = path.join(__dirname, '..', 'solidity');
    
    // 编译RewardToken
    console.log('Compiling RewardToken...');
    const rewardTokenPath = path.join(contractsDir, 'RewardToken.sol');
    const rewardTokenContent = readContract(rewardTokenPath);
    const rewardTokenOutput = compileContract('RewardToken.sol', rewardTokenContent);
    saveAbi('RewardToken', rewardTokenOutput.abi);
    
    // 编译TaskRegistry
    console.log('Compiling TaskRegistry...');
    const taskRegistryPath = path.join(contractsDir, 'TaskRegistry.sol');
    const taskRegistryContent = readContract(taskRegistryPath);
    const taskRegistryOutput = compileContract('TaskRegistry.sol', taskRegistryContent);
    saveAbi('TaskRegistry', taskRegistryOutput.abi);
    
    // 编译FamilyRegistry
    console.log('Compiling FamilyRegistry...');
    const familyRegistryPath = path.join(contractsDir, 'FamilyRegistry.sol');
    const familyRegistryContent = readContract(familyRegistryPath);
    const familyRegistryOutput = compileContract('FamilyRegistry.sol', familyRegistryContent);
    saveAbi('FamilyRegistry', familyRegistryOutput.abi);
    
    // 编译RewardRegistry（新添加的合约）
    console.log('Compiling RewardRegistry...');
    const rewardRegistryPath = path.join(contractsDir, 'RewardRegistry.sol');
    const rewardRegistryContent = readContract(rewardRegistryPath);
    const dependencies = {
      'RewardToken.sol': rewardTokenContent
    };
    const rewardRegistryOutput = compileContract('RewardRegistry.sol', rewardRegistryContent, dependencies);
    saveAbi('RewardRegistry', rewardRegistryOutput.abi);
    
    console.log('All contracts compiled successfully!');
  } catch (error) {
    console.error('Error compiling contracts:', error);
    process.exit(1);
  }
}

main(); 