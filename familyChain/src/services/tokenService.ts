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
    console.log(`已清除地址 ${walletAddress} 的余额缓存`);
  } else {
    cachedBalances = {};
    console.log('已清除所有余额缓存');
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
    
    // 尝试不同的销毁方法
    let tx;
    try {
      // 首先尝试burn方法
      console.log('尝试burn方法...');
      tx = await tokenContract.burn(address, amountInWei);
    } catch (burnError) {
      console.error('burn方法失败:', burnError);
      
      try {
        // 尝试transfer方法，将代币转移到0地址
        console.log('尝试transfer方法，转移到0地址...');
        const zeroAddress = '0x0000000000000000000000000000000000000000';
        tx = await tokenContract.transfer(zeroAddress, amountInWei);
      } catch (transferError) {
        console.error('transfer方法失败:', transferError);
        
        // 尝试通过API销毁代币
        console.log('尝试通过API销毁代币...');
        try {
          const response = await contractApi.transfer('0x0000000000000000000000000000000000000000', amount.toString());
          if (response.success) {
            console.log('通过API销毁代币成功');
            return true;
          } else {
            console.error('通过API销毁代币失败:', response.error);
            return false;
          }
        } catch (apiError) {
          console.error('API销毁代币失败:', apiError);
          return false;
        }
      }
    }
    
    console.log('销毁交易已提交:', tx.hash);
    
    // 等待交易确认
    const receipt = await tx.wait();
    console.log('销毁交易已确认:', receipt);
    
    // 更新本地缓存的余额
    const newBalance = (parseFloat(await getTokenBalance(address)) - amount).toString();
    updateLocalBalance(address, newBalance);
    
    return true;
  } catch (error) {
    console.error('销毁代币失败:', error);
    return false;
  }
}; 