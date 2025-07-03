// This is a simplified version of what would be a Solidity contract
// In a real application, this would be a proper Solidity file deployed to the blockchain

import { ethers } from 'ethers';

// Define the type for our Task Contract based on TaskRegistry ABI
interface TaskContract extends ethers.BaseContract {
  createTask(title: string, description: string, reward: ethers.BigNumberish, overrides?: {value: ethers.BigNumberish}): Promise<ethers.ContractTransactionResponse>;
  assignTask(taskId: number, childAddress: string): Promise<ethers.ContractTransactionResponse>;
  completeTask(taskId: number): Promise<ethers.ContractTransactionResponse>;
  approveTask(taskId: number): Promise<ethers.ContractTransactionResponse>;
  rejectTask(taskId: number): Promise<ethers.ContractTransactionResponse>;
  getTask(taskId: number): Promise<[bigint, string, string, string, string, bigint, boolean, boolean, boolean]>;
  owner(): Promise<string>;
  taskCount(): Promise<bigint>;
  tasks(taskId: number): Promise<[bigint, string, string, string, string, bigint, boolean, boolean, boolean]>;
  withdraw(): Promise<ethers.ContractTransactionResponse>;
}

// Import the actual TaskRegistry ABI
import TaskRegistryABI from '../../../familyChain-contract/abi/TaskRegistry.json';

// Use the real TaskRegistry ABI
export const TaskContractABI = TaskRegistryABI;

// Example code to interact with the contract
export const getTaskContract = async (provider: ethers.BrowserProvider, contractAddress: string): Promise<TaskContract> => {
  try {
    // 优先使用 MetaMask 提供的 ethereum 对象
    if ((window as any).ethereum && (window as any).ethereum.isMetaMask) {
      console.log('Using MetaMask provider');
      const metaMaskProvider = new ethers.BrowserProvider((window as any).ethereum);
      const signer = await metaMaskProvider.getSigner();
      const signerAddress = await signer.getAddress();
      const network = await metaMaskProvider.getNetwork();
      
      console.log('Current network chainId:', network.chainId);
      console.log('Using signer address:', signerAddress);

      // 检查是否在 Sepolia 网络上
      if (network.chainId !== 11155111n) { // Sepolia 的 chainId
        // 尝试自动切换到 Sepolia 网络
        try {
          const ethereum = (window as any).ethereum;
          if (ethereum) {
            // 首先尝试添加 Sepolia 网络
            try {
              await ethereum.request({
                method: 'wallet_addEthereumChain',
                params: [{
                  chainId: '0xaa36a7', // 11155111 in hex
                  chainName: 'Sepolia Test Network',
                  nativeCurrency: {
                    name: 'Sepolia ETH',
                    symbol: 'SEP',
                    decimals: 18
                  },
                  rpcUrls: ['https://eth-sepolia.g.alchemy.com/v2/ISsWLMLFTjBF1rFC4G9R3'],
                  blockExplorerUrls: ['https://sepolia.etherscan.io/']
                }]
              });
            } catch (addError: any) {
              // 网络可能已经存在
              if (addError.code !== 4902) {
                console.log('Network may already exist:', addError);
              }
            }

            // 然后切换到 Sepolia 网络
            await ethereum.request({
              method: 'wallet_switchEthereumChain',
              params: [{ chainId: '0xaa36a7' }]
            });

            // 等待网络切换完成
            await new Promise(resolve => setTimeout(resolve, 1000));
            
            // 重新获取网络信息和签名者
            const newProvider = new ethers.BrowserProvider(ethereum);
            const newSigner = await newProvider.getSigner();
            const newSignerAddress = await newSigner.getAddress();
            const newNetwork = await newProvider.getNetwork();
            
            console.log('After network switch, using signer address:', newSignerAddress);
            
            if (newNetwork.chainId === 11155111n) {
              console.log('Successfully switched to Sepolia network');
              return new ethers.Contract(contractAddress, TaskContractABI, newSigner) as unknown as TaskContract;
            }
          }
        } catch (switchError) {
          console.error('Failed to switch network automatically:', switchError);
        }
        
        throw new Error('Please switch to Sepolia network in your wallet');
      }
      
      return new ethers.Contract(contractAddress, TaskContractABI, signer) as unknown as TaskContract;
    } else {
      // 如果没有 MetaMask，则使用提供的 provider
      console.log('Using provided provider (non-MetaMask)');
      const signer = await provider.getSigner();
      return new ethers.Contract(contractAddress, TaskContractABI, signer) as unknown as TaskContract;
    }
  } catch (error) {
    console.error('Error getting task contract:', error);
    throw error;
  }
};

export const createTask = async (
  contractPromise: Promise<TaskContract>,
  title: string,
  description: string,
  rewardAmount: string
) => {
  try {
    const contract = await contractPromise;
    
    // 优先使用 MetaMask 提供的 ethereum 对象
    if ((window as any).ethereum && (window as any).ethereum.isMetaMask) {
      // 获取当前连接的钱包地址
      const provider = new ethers.BrowserProvider((window as any).ethereum);
      const accounts = await provider.send("eth_requestAccounts", []);
      const currentAddress = accounts[0];
      
      console.log("Transaction will be sent from address:", currentAddress);
      
      // 确保合约使用正确的钱包地址
      const signer = await provider.getSigner(currentAddress);
      const contractWithSigner = contract.connect(signer) as TaskContract;
      
      const reward = ethers.parseEther(rewardAmount);
      const tx = await contractWithSigner.createTask(
        title,
        description,
        reward,
        { value: reward }
      );
      return tx.wait();
    } else {
      // 如果没有 MetaMask，使用合约的默认签名者
      console.log("Using contract's default signer");
      const reward = ethers.parseEther(rewardAmount);
      const tx = await contract.createTask(
        title,
        description,
        reward,
        { value: reward }
      );
      return tx.wait();
    }
  } catch (error) {
    console.error('Error creating task:', error);
    throw error;
  }
};

// acceptTask function removed - not available in TaskRegistry contract

export const assignTask = async (contract: TaskContract, taskId: number, childAddress: string) => {
  console.log(`准备分配任务 ${taskId} 给子账户 ${childAddress}`);
  const tx = await contract.assignTask(taskId, childAddress);
  console.log('任务分配交易已发送:', tx.hash);
  return tx.wait();
};

export const completeTask = async (
  contract: TaskContract,
  taskId: number
) => {
  const tx = await contract.completeTask(taskId);
  return tx.wait();
};

export const approveTask = async (
  contract: TaskContract,
  taskId: number
) => {
  const tx = await contract.approveTask(taskId);
  return tx.wait();
};

export const rejectTask = async (
  contract: TaskContract,
  taskId: number
) => {
  const tx = await contract.rejectTask(taskId);
  return tx.wait();
};