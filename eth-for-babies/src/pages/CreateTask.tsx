import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAccount, useSwitchChain } from 'wagmi';
import { Upload, X, Calendar, PlusCircle, Award, Users } from 'lucide-react';
import Button from '../components/common/Button';
import Card, { CardBody, CardFooter } from '../components/common/Card';
import { useTask } from '../context/TaskContext';
import { useFamily } from '../context/FamilyContext';
import { ethers } from 'ethers';
import { getTaskContract, createTask as createTaskOnChain } from '../contracts/TaskContract';

const TASK_CONTRACT_ADDRESS = import.meta.env.VITE_TASK_CONTRACT_ADDRESS || '0xYourContractAddressHere';

const CreateTask = () => {
  const navigate = useNavigate();
  const { address, chainId } = useAccount();
  const { addTask } = useTask();
  const { getAllChildren, selectedChild } = useFamily();
  const [imagePreview, setImagePreview] = useState<string | null>(null);
  const { switchChain } = useSwitchChain();
  
  const children = getAllChildren();
  
  const [formData, setFormData] = useState({
    title: '',
    description: '',
    deadline: '',
    difficulty: 'easy',
    reward: '0.01',
    completionCriteria: '',
  });
  
  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value,
    }));
  };
  
  const handleImageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        setImagePreview(reader.result as string);
      };
      reader.readAsDataURL(file);
    }
  };
  
  const removeImage = () => {
    setImagePreview(null);
  };
  
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    // 验证必填字段
    if (!formData.title || !formData.description || !formData.reward || !formData.deadline || !formData.completionCriteria) {
      alert('Please fill in all required fields');
      return;
    }

    try {
      // 确保钱包连接
      if (!address) {
        alert('Please connect your wallet first');
        return;
      }

      // 检查当前网络是否为 Sepolia，如果不是则尝试切换
      if (chainId !== 11155111) { // Sepolia 的 chainId 为 11155111
        try {
          // 首先尝试添加 Sepolia 网络到钱包（如果还没有的话）
          const ethereum = (window as any).ethereum;
          if (ethereum) {
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
              // 如果网络已经存在，这个错误是正常的
              if (addError.code !== 4902) {
                console.log('Network may already exist or other error:', addError);
              }
            }

            // 然后尝试切换到 Sepolia 网络
            try {
              await ethereum.request({
                method: 'wallet_switchEthereumChain',
                params: [{ chainId: '0xaa36a7' }] // 11155111 in hex
              });
            } catch (switchError: any) {
              console.error('Failed to switch network:', switchError);
              throw switchError;
            }
          }

          // 使用 wagmi 的 switchChain 作为备选方案
          if (switchChain) {
            try {
              await switchChain({ chainId: 11155111 });
            } catch (wagmiError) {
              console.log('Wagmi switch failed, but manual switch may have worked:', wagmiError);
            }
          }

          // IMPORTANT: Poll to ensure chainId has updated after switch attempt
          let attempts = 0;
          const maxAttempts = 30; // Try for up to 15 seconds (30 * 500ms)
          const delayMs = 500;

          while (attempts < maxAttempts) {
            if (!(window as any).ethereum) {
              alert('No Ethereum provider found after network switch attempt.');
              return;
            }
            const tempProvider = new ethers.BrowserProvider((window as any).ethereum);
            const networkFromTempProvider = await tempProvider.getNetwork();
            if (networkFromTempProvider.chainId === 11155111n) {
              console.log('Wallet successfully detected Sepolia after switch attempt.');
              break; // Exit polling loop
            }
            await new Promise(resolve => setTimeout(resolve, delayMs));
            attempts++;
          }

          // If after polling, network is still not Sepolia
          if ((await new ethers.BrowserProvider((window as any).ethereum).getNetwork()).chainId !== 11155111n) {
            alert('Failed to switch to Sepolia network automatically. Please switch manually in your wallet and try again.');
            return;
          }
          console.log('Successfully switched to Sepolia or already on it.');

        } catch (switchError) {
          console.error('Failed to switch network:', switchError);
          alert('Failed to switch to Sepolia network. Please switch manually in your wallet.');
          return;
        }
      }

      // 1. 调用合约，发送ETH锁定奖励
      if (!(window as any).ethereum) {
        alert('No Ethereum provider found');
        return;
      }
      // ethers v6: 使用 BrowserProvider
      const provider = new ethers.BrowserProvider((window as any).ethereum);
      const contractPromise = getTaskContract(provider, TASK_CONTRACT_ADDRESS);
      // 转换时间为秒
      const deadlineTimestamp = Math.floor(new Date(formData.deadline).getTime() / 1000);
      // 难度映射
      const difficultyMap = { easy: 0, medium: 1, hard: 2 };
      const difficultyNum = difficultyMap[formData.difficulty as keyof typeof difficultyMap] ?? 0;
      // 调用合约
      const txReceipt = await createTaskOnChain(
        contractPromise,
        formData.title,
        formData.description,
        formData.reward
      );
      // 2. 合约成功后再调用后端API
      await addTask({
        title: formData.title,
        description: formData.description,
        reward: formData.reward,
        deadline: formData.deadline,
        difficulty: formData.difficulty as 'easy' | 'medium' | 'hard',
        completionCriteria: formData.completionCriteria,
        createdBy: address,
        assignedChildId: selectedChild?.id || undefined,
        assignedTo: selectedChild?.walletAddress || undefined
      });

      alert('Task created successfully!');
      navigate('/parent');
    } catch (error) {
      console.error('Error creating task:', error);
      const errorMessage = error instanceof Error ? error.message : 'Failed to create task. Please try again.';
      alert(errorMessage);
    }
  };

  return (
    <div className="max-w-2xl mx-auto">
      <div className="mb-6">
        <h1 className="text-3xl font-bold text-gray-900 mb-2">Create New Task</h1>
        <p className="text-gray-600">Define a task for your child to complete and earn rewards</p>
      </div>
      
      <Card>
        <form onSubmit={handleSubmit}>
          <CardBody>
            <div className="space-y-6">
              <div>
                <label htmlFor="title" className="block text-sm font-medium text-gray-700 mb-1">
                  Task Title*
                </label>
                <input
                  type="text"
                  id="title"
                  name="title"
                  value={formData.title}
                  onChange={handleChange}
                  className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                  placeholder="E.g., Clean the kitchen"
                  required
                />
              </div>
              
              <div>
                <label htmlFor="description" className="block text-sm font-medium text-gray-700 mb-1">
                  Description*
                </label>
                <textarea
                  id="description"
                  name="description"
                  value={formData.description}
                  onChange={handleChange}
                  rows={4}
                  className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                  placeholder="Provide details about what needs to be done"
                  required
                ></textarea>
              </div>
              
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label htmlFor="deadline" className="block text-sm font-medium text-gray-700 mb-1">
                    Deadline*
                  </label>
                  <div className="relative">
                    <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                      <Calendar className="h-5 w-5 text-gray-400" />
                    </div>
                    <input
                      type="datetime-local"
                      id="deadline"
                      name="deadline"
                      value={formData.deadline}
                      onChange={handleChange}
                      className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                      required
                    />
                  </div>
                </div>
                
                <div>
                  <label htmlFor="difficulty" className="block text-sm font-medium text-gray-700 mb-1">
                    Difficulty Level*
                  </label>
                  <select
                    id="difficulty"
                    name="difficulty"
                    value={formData.difficulty}
                    onChange={handleChange}
                    className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                    required
                  >
                    <option value="easy">Easy</option>
                    <option value="medium">Medium</option>
                    <option value="hard">Hard</option>
                  </select>
                </div>
              </div>
              
              <div>
                <label htmlFor="reward" className="block text-sm font-medium text-gray-700 mb-1">
                  Reward (ETH)*
                </label>
                <div className="relative">
                  <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <Award className="h-5 w-5 text-gray-400" />
                  </div>
                  <input
                    type="number"
                    id="reward"
                    name="reward"
                    value={formData.reward}
                    onChange={handleChange}
                    step="0.001"
                    min="0.001"
                    className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                    required
                  />
                </div>
                <p className="mt-1 text-xs text-gray-500">
                  This amount will be locked in the smart contract until the task is approved.
                </p>
              </div>
              
              <div>
                <label htmlFor="image" className="block text-sm font-medium text-gray-700 mb-1">
                  Task Image (Optional)
                </label>
                {!imagePreview ? (
                  <div className="mt-1 flex justify-center px-6 py-10 border-2 border-gray-300 border-dashed rounded-md">
                    <div className="text-center">
                      <Upload className="mx-auto h-12 w-12 text-gray-400" />
                      <div className="mt-2 flex text-sm text-gray-600">
                        <label
                          htmlFor="image"
                          className="relative cursor-pointer bg-white rounded-md font-medium text-primary-600 hover:text-primary-500"
                        >
                          <span>Upload a file</span>
                          <input 
                            id="image" 
                            name="image" 
                            type="file" 
                            className="sr-only" 
                            accept="image/*"
                            onChange={handleImageChange}
                          />
                        </label>
                        <p className="pl-1">or drag and drop</p>
                      </div>
                      <p className="text-xs text-gray-500">
                        PNG, JPG, GIF up to 5MB
                      </p>
                    </div>
                  </div>
                ) : (
                  <div className="relative mt-1">
                    <img
                      src={imagePreview}
                      alt="Task preview"
                      className="w-full h-48 object-cover rounded-md"
                    />
                    <button
                      type="button"
                      onClick={removeImage}
                      className="absolute top-2 right-2 p-1 bg-white rounded-full shadow-md"
                    >
                      <X className="h-5 w-5 text-gray-600" />
                    </button>
                  </div>
                )}
              </div>
              
              <div>
                <label htmlFor="completionCriteria" className="block text-sm font-medium text-gray-700 mb-1">
                  Completion Criteria*
                </label>
                <textarea
                  id="completionCriteria"
                  name="completionCriteria"
                  value={formData.completionCriteria}
                  onChange={handleChange}
                  rows={3}
                  className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                  placeholder="Describe what the completed task should look like"
                  required
                ></textarea>
              </div>
            </div>
          </CardBody>
          
          <CardFooter className="flex justify-end space-x-4">
            <Button 
              type="button" 
              variant="outline" 
              onClick={() => navigate('/parent')}
            >
              Cancel
            </Button>
            <Button 
              type="submit" 
              leftIcon={<PlusCircle className="h-5 w-5" />}
            >
              Create Task
            </Button>
          </CardFooter>
        </form>
      </Card>
    </div>
  );
};

export default CreateTask;