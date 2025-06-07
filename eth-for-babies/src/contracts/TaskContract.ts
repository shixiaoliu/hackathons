// This is a simplified version of what would be a Solidity contract
// In a real application, this would be a proper Solidity file deployed to the blockchain

import { ethers } from 'ethers';

// Example ABI for a task contract
export const TaskContractABI = [
  // Task creation
  "function createTask(string title, string description, uint256 deadline, uint8 difficulty, string completionCriteria) payable returns (uint256)",
  
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
export const getTaskContract = (provider: ethers.providers.Web3Provider, contractAddress: string) => {
  const signer = provider.getSigner();
  return new ethers.Contract(contractAddress, TaskContractABI, signer);
};

export const createTask = async (
  contract: ethers.Contract,
  title: string,
  description: string,
  deadline: number,
  difficulty: number,
  completionCriteria: string,
  rewardAmount: string
) => {
  const tx = await contract.createTask(
    title,
    description,
    deadline,
    difficulty,
    completionCriteria,
    { value: ethers.utils.parseEther(rewardAmount) }
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
  feedback: string
) => {
  const tx = await contract.approveTask(taskId, feedback);
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