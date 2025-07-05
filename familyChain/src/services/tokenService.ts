import { ethers } from 'ethers';
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
    if (!TOKEN_CONTRACT_ADDRESS || TOKEN_CONTRACT_ADDRESS === '0x0000000000000000000000000000000000000000') {
      console.error('FCT contract address not configured');
      return null;
    }

    // 检查是否有MetaMask
    if (!(window as any).ethereum) {
      console.error('MetaMask not found');
      return null;
    }

    // 创建provider和signer
    const provider = new ethers.BrowserProvider((window as any).ethereum);
    const signer = await provider.getSigner();
    
    // 创建合约实例
    const contract = new ethers.Contract(
      TOKEN_CONTRACT_ADDRESS,
      RewardTokenABI,
      signer
    ) as unknown as RewardTokenContract;
    
    return contract;
  } catch (error) {
    console.error('Error getting FCT contract:', error);
    return null;
  }
};

/**
 * 获取钱包FCT余额
 * @param walletAddress 钱包地址
 * @returns 格式化的余额字符串
 */
export const getTokenBalance = async (walletAddress: string): Promise<string> => {
  try {
    const contract = await getRewardTokenContract();
    if (!contract) {
      return '0';
    }

    // 获取余额和小数位
    const balance = await contract.balanceOf(walletAddress);
    const decimals = await contract.decimals();
    
    // 格式化余额
    const formattedBalance = ethers.formatUnits(balance, decimals);
    console.log(`[Blockchain] Raw balance: ${balance}, Decimals: ${decimals}, Formatted: ${formattedBalance}`);
    
    return formattedBalance;
  } catch (error) {
    console.error('Error fetching FCT balance:', error);
    return '0';
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
    
    // 确保调用者是合约授权的minter或owner
    const tokenContract = new ethers.Contract(TOKEN_CONTRACT_ADDRESS, RewardTokenABI, signer);
    
    // 将数量转换为wei (考虑18位小数)
    const amountInWei = ethers.parseUnits(amount.toString(), 18);
    console.log('销毁数量(wei):', amountInWei.toString());
    
    // 调用合约的burn方法
    const tx = await tokenContract.burn(address, amountInWei);
    console.log('销毁交易已提交:', tx.hash);
    
    // 等待交易确认
    const receipt = await tx.wait();
    console.log('销毁交易已确认:', receipt);
    
    return true;
  } catch (error) {
    console.error('销毁代币失败:', error);
    return false;
  }
}; 