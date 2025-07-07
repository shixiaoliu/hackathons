import { ethers } from 'ethers';
import { contractApi } from './api';
import RewardTokenABI from '../../../familyChain-contract/abi/RewardToken.json';

// FCT代币合约地址
const TOKEN_CONTRACT_ADDRESS = import.meta.env.VITE_TOKEN_CONTRACT_ADDRESS || '0x0000000000000000000000000000000000000000';

// 定义RewardToken合约接口
interface RewardTokenContract extends ethers.BaseContract {
  balanceOf(account: string): Promise<bigint>;
  decimals(): Promise<number>;
  symbol(): Promise<string>;
  name(): Promise<string>;
}

/**
 * 获取FCT代币合约实例
 */
export const getRewardTokenContract = async (): Promise<RewardTokenContract | null> => {
  try {
    console.log('开始获取FCT代币合约...');
    
    if (!TOKEN_CONTRACT_ADDRESS || TOKEN_CONTRACT_ADDRESS === '0x0000000000000000000000000000000000000000') {
      console.error('FCT合约地址未配置，尝试使用硬编码地址');
      // 使用硬编码的合约地址作为备选
      const hardcodedAddress = '0x5FbDB2315678afecb367f032d93F642f64180aa3'; // 本地开发环境的默认合约地址
      console.log('使用硬编码合约地址:', hardcodedAddress);
      
      // 创建provider和signer
      const provider = new ethers.JsonRpcProvider('http://localhost:8545'); // 本地开发环境
      
      // 创建合约实例
      const contract = new ethers.Contract(
        hardcodedAddress,
        RewardTokenABI,
        provider
      ) as unknown as RewardTokenContract;
      
      // 测试合约连接
      try {
        const symbol = await contract.symbol();
        console.log('成功连接到FCT合约，代币符号:', symbol);
        return contract;
      } catch (testError) {
        console.error('测试合约连接失败:', testError);
      }
    }

    // 检查是否有MetaMask
    if (!(window as any).ethereum) {
      console.error('未找到MetaMask，尝试使用HTTP提供者');
      // 尝试使用HTTP提供者
      const provider = new ethers.JsonRpcProvider('http://localhost:8545');
      
      const contract = new ethers.Contract(
        TOKEN_CONTRACT_ADDRESS,
        RewardTokenABI,
        provider
      ) as unknown as RewardTokenContract;
      
      return contract;
    }

    // 创建provider和signer
    const provider = new ethers.BrowserProvider((window as any).ethereum);
    const signer = await provider.getSigner();
    
    console.log('已获取以太坊提供者和签名者');
    
    // 创建合约实例
    const contract = new ethers.Contract(
      TOKEN_CONTRACT_ADDRESS,
      RewardTokenABI,
      signer
    ) as unknown as RewardTokenContract;
    
    // 测试合约连接
    try {
      const symbol = await contract.symbol();
      console.log('成功连接到FCT合约，代币符号:', symbol);
    } catch (testError) {
      console.warn('合约方法调用测试失败:', testError);
      // 继续使用合约实例，即使测试失败
    }
    
    return contract;
  } catch (error) {
    console.error('获取FCT合约失败:', error);
    return null;
  }
};

// 添加本地缓存，提高余额查询的可靠性
let cachedBalances: {[address: string]: {balance: string, timestamp: number}} = {};

/**
 * 获取钱包FCT余额
 * @param walletAddress 钱包地址
 * @returns 格式化的余额字符串
 */
export const getTokenBalance = async (walletAddress: string): Promise<string> => {
  try {
    // 尝试从区块链获取余额
    console.log('从区块链获取代币余额...');
    
    // 检查缓存是否有最近的余额数据
    const cachedData = cachedBalances[walletAddress];
    const now = Date.now();
    const cacheValidTime = 10000; // 10秒内的缓存有效
    
    if (cachedData && (now - cachedData.timestamp < cacheValidTime)) {
      console.log('使用缓存的余额数据:', cachedData.balance);
      return cachedData.balance;
    }
    
    // 首先尝试直接从区块链获取余额
    try {
      const contract = await getRewardTokenContract();
      if (contract) {
        // 获取余额和小数位
        const balance = await contract.balanceOf(walletAddress);
        const decimals = await contract.decimals();
        
        // 格式化余额
        const formattedBalance = ethers.formatUnits(balance, decimals);
        console.log(`[Blockchain] 直接获取余额: ${balance}, 小数位: ${decimals}, 格式化后: ${formattedBalance}`);
        
        // 确保余额不为零或无效值
        if (formattedBalance && parseFloat(formattedBalance) > 0) {
          // 更新缓存
          cachedBalances[walletAddress] = {
            balance: formattedBalance,
            timestamp: now
          };
          
          return formattedBalance;
        } else {
          console.log('区块链返回的余额为零或无效，尝试其他方式获取');
        }
      }
    } catch (blockchainError) {
      console.error('直接从区块链获取余额失败:', blockchainError);
      // 区块链获取失败，继续尝试API
    }
    
    // 如果区块链获取失败，尝试通过API获取
    console.log('尝试通过API获取余额...');
    const response = await contractApi.getBalance(walletAddress);
    
    if (response.success && response.data) {
      console.log('API获取余额成功:', response.data.balance);
      
      // 确保API返回的余额不为零或无效值
      if (response.data.balance && parseFloat(response.data.balance) > 0) {
        // 更新缓存
        cachedBalances[walletAddress] = {
          balance: response.data.balance,
          timestamp: now
        };
        return response.data.balance;
      } else {
        console.log('API返回的余额为零或无效');
      }
    }
    
    // 如果API调用失败但有缓存，返回缓存的余额
    if (cachedData) {
      console.log('API调用失败，使用缓存的余额数据:', cachedData.balance);
      return cachedData.balance;
    }
    
    // 移除测试余额，始终返回真实数据
    console.log('无法获取余额，返回0');
    return '0';
  } catch (error) {
    console.error('获取代币余额失败:', error);
    
    // 如果有缓存，在出错时返回缓存的余额
    if (cachedBalances[walletAddress]) {
      return cachedBalances[walletAddress].balance;
    }
    
    // 移除测试余额，始终返回真实数据
    return '0';
  }
};

// 更新本地缓存的余额（用于兑换后立即更新）
export const updateLocalBalance = (walletAddress: string, newBalance: string): void => {
  cachedBalances[walletAddress] = {
    balance: newBalance,
    timestamp: Date.now()
  };
  console.log(`已更新本地缓存的余额: ${walletAddress} -> ${newBalance}`);
};

// 清除余额缓存
export const clearBalanceCache = (walletAddress?: string): void => {
  if (walletAddress) {
    delete cachedBalances[walletAddress];
    console.log(`已清除钱包地址 ${walletAddress} 的余额缓存`);
  } else {
    cachedBalances = {};
    console.log('已清除所有余额缓存');
  }
};

/**
 * 通过合约兑换奖品
 * @param rewardId 奖品ID
 * @returns 交易结果
 */
export const exchangeRewardWithContract = async (rewardId: number): Promise<boolean> => {
  try {
    console.log(`通过合约兑换奖品: rewardId=${rewardId}`);
    
    // 检查是否有window.ethereum
    if (!window.ethereum) {
      console.error('未找到MetaMask或其他以太坊提供者');
      return false;
    }
    
    // 创建provider和signer
    const provider = new ethers.BrowserProvider(window.ethereum);
    const signer = await provider.getSigner();
    const signerAddress = await signer.getAddress();
    
    console.log('当前签名者地址:', signerAddress);
    
    // 获取RewardRegistry合约地址
    const rewardRegistryAddress = import.meta.env.VITE_REWARD_CONTRACT_ADDRESS || '0x0000000000000000000000000000000000000000';
    if (!rewardRegistryAddress || rewardRegistryAddress === '0x0000000000000000000000000000000000000000') {
      console.error('RewardRegistry合约地址未配置，请检查环境变量VITE_REWARD_CONTRACT_ADDRESS');
      return false;
    }
    
    console.log('使用RewardRegistry合约地址:', rewardRegistryAddress);
    
    try {
      // 使用硬编码的ABI，包含exchangeReward方法
      const rewardRegistryABI = [
        {
          "inputs": [
            {
              "internalType": "uint256",
              "name": "_rewardId",
              "type": "uint256"
            }
          ],
          "name": "exchangeReward",
          "outputs": [
            {
              "internalType": "uint256",
              "name": "",
              "type": "uint256"
            }
          ],
          "stateMutability": "nonpayable",
          "type": "function"
        }
      ];
      
      // 创建合约实例
      const rewardRegistryContract = new ethers.Contract(rewardRegistryAddress, rewardRegistryABI, signer);
      
      // 在调用前查询用户余额
      let beforeBalance = "0";
      let decimals = 18;
      try {
        const tokenContract = await getRewardTokenContract();
        if (tokenContract) {
          const balance = await tokenContract.balanceOf(signerAddress);
          decimals = await tokenContract.decimals();
          beforeBalance = ethers.formatUnits(balance, decimals);
          console.log(`交易前余额: ${beforeBalance} FCT`);
        }
      } catch (balanceError) {
        console.error('查询交易前余额失败:', balanceError);
      }
      
      // 调用exchangeReward方法
      console.log('调用exchangeReward方法, rewardId:', rewardId);
      const tx = await rewardRegistryContract.exchangeReward(
        rewardId,
        {
          gasLimit: 3000000 // 设置较大的gas限制
        }
      );
      console.log('exchangeReward方法调用成功，交易哈希:', tx.hash);
      
      // 等待交易确认
      const receipt = await tx.wait();
      console.log('兑换交易已确认:', receipt);
      
      // 在交易后查询用户余额
      try {
        const tokenContract = await getRewardTokenContract();
        if (tokenContract) {
          const balance = await tokenContract.balanceOf(signerAddress);
          const afterBalance = ethers.formatUnits(balance, decimals);
          console.log(`交易后余额: ${afterBalance} FCT`);
          
          // 计算消耗的代币数量
          const consumedTokens = parseFloat(beforeBalance) - parseFloat(afterBalance);
          console.log(`消耗的代币数量: ${consumedTokens.toFixed(18)} FCT`);
        }
      } catch (balanceError) {
        console.error('查询交易后余额失败:', balanceError);
      }
      
      // 清除余额缓存，确保下次获取最新余额
      clearBalanceCache();
      
      return true;
    } catch (contractError) {
      console.error('调用合约方法失败:', contractError);
      return false;
    }
  } catch (error) {
    console.error('通过合约兑换奖品失败:', error);
    return false;
  }
};

/**
 * 获取FCT代币信息
 */
export const getTokenInfo = async (): Promise<{ symbol: string; name: string; decimals: number } | null> => {
  try {
    const contract = await getRewardTokenContract();
    if (!contract) {
      return null;
    }

    const [symbol, name, decimals] = await Promise.all([
      contract.symbol(),
      contract.name(),
      contract.decimals()
    ]);

    return { symbol, name, decimals };
  } catch (error) {
    console.error('Error fetching FCT info:', error);
    return null;
  }
}; 

/**
 * 销毁代币 - 用于兑换奖品时
 * @param address 用户钱包地址
 * @param amount 要销毁的代币数量
 * @returns 交易结果
 */
export const burnTokens = async (address: string, amount: number): Promise<boolean> => {
  try {
    console.log(`正在销毁代币: 地址=${address}, 数量=${amount}`);
    
    // 检查是否有window.ethereum
    if (!window.ethereum) {
      console.error('未找到MetaMask或其他以太坊提供者');
      return false;
    }
    
    // 创建provider和signer
    const provider = new ethers.BrowserProvider(window.ethereum);
    const signer = await provider.getSigner();
    const signerAddress = await signer.getAddress();
    
    console.log('当前签名者地址:', signerAddress);
    
    // 获取合约地址
    let contractAddress = TOKEN_CONTRACT_ADDRESS;
    if (!contractAddress || contractAddress === '0x0000000000000000000000000000000000000000') {
      console.warn('合约地址未配置，尝试使用硬编码地址');
      contractAddress = '0x5FbDB2315678afecb367f032d93F642f64180aa3'; // 本地开发环境的默认合约地址
    }
    
    console.log('使用合约地址:', contractAddress);
    
    // 创建合约实例
    const tokenContract = new ethers.Contract(contractAddress, RewardTokenABI, signer);
    
    // 检查用户余额
    try {
      const balance = await tokenContract.balanceOf(address);
      const decimals = await tokenContract.decimals();
      const formattedBalance = ethers.formatUnits(balance, decimals);
      
      console.log(`当前余额: ${formattedBalance}`);
      
      if (parseFloat(formattedBalance) < amount) {
        console.error(`余额不足: ${formattedBalance} < ${amount}`);
        return false;
      }
    } catch (balanceError) {
      console.error('检查余额失败:', balanceError);
    }
    
    // 将数量转换为wei (考虑18位小数)
    const amountInWei = ethers.parseUnits(amount.toString(), 18);
    console.log('销毁数量(wei):', amountInWei.toString());
    
    // 检查当前用户是否有权限销毁代币
    let isMinter = false;
    try {
      if (typeof tokenContract.authorizedMinters === 'function') {
        isMinter = await tokenContract.authorizedMinters(signerAddress);
        console.log('当前用户是否为授权铸币者:', isMinter);
      }
    } catch (error) {
      console.log('检查铸币权限失败:', error);
    }
    
    // 尝试不同的销毁方法
    let tx;
    try {
      // 检查合约方法
      console.log('检查合约方法...');
      try {
        // 使用正确的方式获取合约函数
        const functionFragments = tokenContract.interface.fragments.filter(f => f.type === 'function');
        const functionNames = functionFragments.map(f => {
          // 使用类型断言或检查属性存在性
          const fragment = f as { name?: string };
          return fragment.name || 'unknown';
        });
        console.log('可用合约方法:', functionNames);
      } catch (error) {
        console.error('获取合约方法失败:', error);
      }
      
      // 如果当前用户是铸币者，尝试直接burn
      if (isMinter) {
        console.log('尝试burn方法(作为铸币者)...');
        // 根据RewardToken.sol合约，burn方法需要两个参数: from和amount
        tx = await tokenContract.burn(address, amountInWei);
        console.log('burn方法调用成功，交易哈希:', tx.hash);
      } else {
        // 如果当前用户不是铸币者，尝试使用标准ERC20的transfer方法
        throw new Error('当前用户不是授权铸币者，尝试其他方法');
      }
    } catch (burnError) {
      console.error('burn方法失败:', burnError);
      
      // 检查当前签名者是否是代币持有者
      const isTokenHolder = address.toLowerCase() === signerAddress.toLowerCase();
      console.log('当前签名者是否是代币持有者:', isTokenHolder);
      
      if (isTokenHolder) {
        // 如果当前签名者就是代币持有者，尝试使用标准ERC20的transfer方法
        try {
          console.log('尝试使用标准ERC20 transfer方法，转移到0地址...');
          const zeroAddress = '0x0000000000000000000000000000000000000000';
          tx = await tokenContract.transfer(zeroAddress, amountInWei);
          console.log('transfer方法调用成功，交易哈希:', tx.hash);
        } catch (transferError) {
          console.error('transfer方法失败:', transferError);
          return false;
        }
      } else {
        console.error('当前签名者不是代币持有者，无法销毁代币');
        return false;
      }
    }
    
    // 等待交易确认
    try {
      const receipt = await tx.wait();
      console.log('销毁代币交易已确认:', receipt);
      
      // 清除余额缓存，确保下次获取最新余额
      clearBalanceCache(address);
      
      return true;
    } catch (waitError) {
      console.error('等待交易确认失败:', waitError);
      return false;
    }
  } catch (error) {
    console.error('销毁代币失败:', error);
    return false;
  }
}; 

/**
 * 通过合约创建奖品
 * @param familyId 家庭ID
 * @param name 奖品名称
 * @param description 奖品描述
 * @param imageURI 奖品图片URI
 * @param tokenPrice 代币价格
 * @param stock 库存数量
 * @returns 创建的奖品ID，如果失败则返回0
 */
export const createRewardWithContract = async (
  familyId: number,
  name: string,
  description: string,
  imageURI: string,
  tokenPrice: number,
  stock: number
): Promise<number> => {
  try {
    console.log(`通过合约创建奖品: 家庭ID=${familyId}, 名称=${name}, 价格=${tokenPrice}, 库存=${stock}`);
    
    // 检查是否有window.ethereum
    if (!(window as any).ethereum) {
      console.error('未找到MetaMask或其他以太坊提供者');
      return 0;
    }
    
    // 创建provider和signer
    const provider = new ethers.BrowserProvider((window as any).ethereum);
    const signer = await provider.getSigner();
    const signerAddress = await signer.getAddress();
    
    console.log('当前签名者地址:', signerAddress);
    
    // 获取RewardRegistry合约地址
    const rewardRegistryAddress = import.meta.env.VITE_REWARD_CONTRACT_ADDRESS || '0x0000000000000000000000000000000000000000';
    if (!rewardRegistryAddress || rewardRegistryAddress === '0x0000000000000000000000000000000000000000') {
      console.error('RewardRegistry合约地址未配置，请检查环境变量VITE_REWARD_CONTRACT_ADDRESS');
      return 0;
    }
    
    console.log('使用RewardRegistry合约地址:', rewardRegistryAddress);
    
    try {
      // 定义扩展的ABI，包含createReward方法和rewardCount方法
      const rewardRegistryABI = [
        {
          "inputs": [
            {
              "internalType": "uint256",
              "name": "_familyId",
              "type": "uint256"
            },
            {
              "internalType": "string",
              "name": "_name",
              "type": "string"
            },
            {
              "internalType": "string",
              "name": "_description",
              "type": "string"
            },
            {
              "internalType": "string",
              "name": "_imageURI",
              "type": "string"
            },
            {
              "internalType": "uint256",
              "name": "_tokenPrice",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "_stock",
              "type": "uint256"
            }
          ],
          "name": "createReward",
          "outputs": [
            {
              "internalType": "uint256",
              "name": "",
              "type": "uint256"
            }
          ],
          "stateMutability": "nonpayable",
          "type": "function"
        },
        {
          "inputs": [],
          "name": "rewardCount",
          "outputs": [
            {
              "internalType": "uint256",
              "name": "",
              "type": "uint256"
            }
          ],
          "stateMutability": "view",
          "type": "function"
        }
      ];
      
      // 创建合约实例
      const rewardRegistryContract = new ethers.Contract(rewardRegistryAddress, rewardRegistryABI, signer);
      
      // 先获取当前的rewardCount，用于后续比较
      const beforeCount = await rewardRegistryContract.rewardCount();
      console.log('创建奖品前的rewardCount:', Number(beforeCount));
      
      // 将tokenPrice转换为合适的单位，确保正确地乘以10^18
      const tokenPriceValue = typeof tokenPrice === 'string' ? parseFloat(tokenPrice) : tokenPrice;
      
      // 将代币价格乘以10^18，以符合ERC20代币的小数位数
      const tokenPriceForContract = ethers.parseUnits(tokenPriceValue.toString(), 18);
      
      console.log('处理后的代币价格:', tokenPriceForContract.toString());
      
      // 处理图片URI，避免数据过大
      let processedImageURI = imageURI;
      
      // 检查是否为base64编码的图片数据
      if (imageURI && imageURI.startsWith('data:image')) {
        console.log('检测到base64编码的图片，使用外部URL替代');
        // 使用占位图片URL替代base64数据
        processedImageURI = 'https://via.placeholder.com/400x300?text=Reward';
      }
      
      // 限制description长度
      const shortDescription = description.length > 100 ? description.substring(0, 100) + '...' : description;
      
      // 调用createReward方法，增加gas限制
      console.log('调用createReward方法，参数:', {
        familyId,
        name,
        description: shortDescription,
        imageURI: '图片URL已处理', // 不打印完整URL
        tokenPrice: tokenPriceForContract.toString(),
        stock
      });
      
      const tx = await rewardRegistryContract.createReward(
        familyId,
        name,
        shortDescription,
        processedImageURI,
        tokenPriceForContract,
        stock,
        { 
          gasLimit: 3000000 // 设置较大的gas限制
        }
      );
      
      console.log('createReward方法调用成功，交易哈希:', tx.hash);
      
      // 等待交易确认
      const receipt = await tx.wait();
      console.log('创建奖品交易已确认:', receipt);
      
      // 直接调用rewardCount方法获取当前计数器值
      // 这个值应该就是我们刚刚创建的奖品ID
      const afterCount = await rewardRegistryContract.rewardCount();
      const rewardId = Number(afterCount);
      console.log('创建奖品后的rewardCount:', rewardId);
      
      // 验证rewardCount是否增加了1
      if (Number(afterCount) === Number(beforeCount) + 1) {
        console.log('rewardCount成功增加，新奖品ID应为:', rewardId);
      } else {
        console.warn(`rewardCount变化异常: ${Number(beforeCount)} -> ${rewardId}`);
      }
      
      return rewardId;
    } catch (contractError) {
      console.error('调用合约方法失败:', contractError);
      return 0;
    }
  } catch (error) {
    console.error('通过合约创建奖品失败:', error);
    return 0;
  }
}; 