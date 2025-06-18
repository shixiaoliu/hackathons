// This is a simplified version of what would be a Solidity contract
// In a real application, this would be a proper Solidity file deployed to the blockchain

import { ethers } from 'ethers';

// Example ABI for a task contract
export const TaskContractABI = [
  // Task creation
  "function createTask(string title, string description, uint256 reward) payable returns (uint256)",
  
  // Task management
  "function acceptTask(uint256 taskId)",
  "function submitTaskCompletion(uint256 taskId, string proofDescription, string[] proofImages)",
  "function approveTask(uint256 taskId, string feedback)",
  "function rejectTask(uint256 taskId, string feedback)",
  
  // View functions
  "function getTask(uint256 taskId) view returns (tuple(uint256 id, string title, string description, uint256 deadline, uint8 difficulty, uint256 reward, uint8 status, address createdBy, address assignedTo, string completionCriteria, uint256 createdAt, uint256 updatedAt))",
  "function getTasksByCreator(address creator) view returns (uint256[])",
  "function getTasksByAssignee(address assignee) view returns (uint256[])",
  
  // Events
  "event TaskCreated(uint256 indexed taskId, address indexed creator, uint256 reward)",
  "event TaskAccepted(uint256 indexed taskId, address indexed assignee)",
  "event TaskCompleted(uint256 indexed taskId, address indexed assignee)",
  "event TaskApproved(uint256 indexed taskId, address indexed approver, address indexed assignee, uint256 reward)",
  "event TaskRejected(uint256 indexed taskId, address indexed approver, string feedback)"
];

// Example code to interact with the contract
export const getTaskContract = async (provider: ethers.BrowserProvider, contractAddress: string) => {
  const signer = await provider.getSigner();
  const network = await provider.getNetwork();
  
  console.log('Current network chainId:', network.chainId);
  console.log('Full network object:', network);

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
              rpcUrls: ['https://sepolia.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161'],
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
        
        // 重新获取网络信息
        const newProvider = new ethers.BrowserProvider(ethereum);
        const newSigner = await newProvider.getSigner();
        const newNetwork = await newProvider.getNetwork();
        
        if (newNetwork.chainId === 11155111n) {
          console.log('Successfully switched to Sepolia network');
          return new ethers.Contract(contractAddress, TaskContractABI, newSigner);
        }
      }
    } catch (switchError) {
      console.error('Failed to switch network automatically:', switchError);
    }
    
    throw new Error('Please switch to Sepolia network in your wallet');
  }
  
  return new ethers.Contract(contractAddress, TaskContractABI, signer);
};

export const createTask = async (
  contractPromise: Promise<ethers.Contract>,
  title: string,
  description: string,
  rewardAmount: string
) => {
  const contract = await contractPromise;
  const reward = ethers.parseEther(rewardAmount);
  const tx = await contract.createTask(
    title,
    description,
    reward,
    { value: reward.toString() }
  );
  return tx.wait();
};

export const acceptTask = async (contract: ethers.Contract, taskId: number) => {
  const tx = await contract.acceptTask(taskId);
  return tx.wait();
};

export const submitTaskCompletion = async (
  contract: ethers.Contract,
  taskId: number,
  description: string,
  imageUrls: string[]
) => {
  const tx = await contract.submitTaskCompletion(taskId, description, imageUrls);
  return tx.wait();
};

export const approveTask = async (
  contract: ethers.Contract,
  taskId: number,
  reward: string
) => {
  const value = ethers.parseEther(reward);
  const tx = await contract.approveTask(
    taskId,
    { value: value.toString() }
  );
  return tx.wait();
};

export const rejectTask = async (
  contract: ethers.Contract,
  taskId: number,
  feedback: string
) => {
  const tx = await contract.rejectTask(taskId, feedback);
  return tx.wait();
};