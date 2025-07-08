import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAccount, useSwitchChain } from 'wagmi';
import { Upload, X, Calendar, PlusCircle, Award, Users } from 'lucide-react';
import Button from '../components/common/Button';
import Card, { CardBody, CardFooter } from '../components/common/Card';
import { useTask } from '../context/TaskContext';
import { useFamily } from '../context/FamilyContext';
import { ethers } from 'ethers';
import { getTaskContract, createTask as createTaskOnChain, TaskContractABI } from '../contracts/TaskContract';

// Get contract address from environment variables
const TASK_CONTRACT_ADDRESS = import.meta.env.VITE_TASK_CONTRACT_ADDRESS || '0x11dB634CFD2f58967e472a179ebDbaF8AB067144'; // Replace placeholder with actual fallback address

const CreateTask = () => {
  const navigate = useNavigate();
  const { address, chainId } = useAccount();
  const { addTask } = useTask();
  const { getAllChildren, selectedChild } = useFamily();
  const [imagePreview, setImagePreview] = useState<string | null>(null);
  const { switchChain } = useSwitchChain();
  const [localIsCreating, setLocalIsCreating] = useState(false);
  const [currentProvider, setCurrentProvider] = useState<any>(null);
  const [contractTaskId, setContractTaskId] = useState<number | null>(null);
  
  // 获取当前钱包地址
  useEffect(() => {
    const checkWallet = async () => {
      if ((window as any).ethereum) {
        try {
          const provider = new ethers.BrowserProvider((window as any).ethereum);
          const signer = await provider.getSigner();
          const signerAddress = await signer.getAddress();
          setCurrentProvider({
            providerAddress: signerAddress,
            wagmiAddress: address
          });
        } catch (err) {
          console.error("获取钱包地址失败:", err);
        }
      }
    };
    
    checkWallet();
  }, [address]);
  
  const children = getAllChildren();
  
  const [formData, setFormData] = useState({
    title: '',
    description: '',
    deadline: '',
    difficulty: 'easy',
    reward: '0.001',
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
      // 检查文件类型是否为图片
      if (!file.type.startsWith('image/')) {
        alert('请上传有效的图片文件');
        return;
      }

      console.log('上传图片类型:', file.type);
      console.log('上传图片大小:', file.size / 1024, 'KB');

      // 如果图片大于1MB，进行压缩
      if (file.size > 1024 * 1024) {
        console.log('图片过大，开始压缩...');
        const reader = new FileReader();
        reader.onloadend = () => {
          const img = new Image();
          img.onload = () => {
            console.log('原始图片尺寸:', img.width, 'x', img.height);
            const canvas = document.createElement('canvas');
            // 保持原始宽高比，但最大限制为800px
            const MAX_WIDTH = 800;
            const MAX_HEIGHT = 800;
            let width = img.width;
            let height = img.height;
            
            if (width > height) {
              if (width > MAX_WIDTH) {
                height *= MAX_WIDTH / width;
                width = MAX_WIDTH;
              }
            } else {
              if (height > MAX_HEIGHT) {
                width *= MAX_HEIGHT / height;
                height = MAX_HEIGHT;
              }
            }
            
            console.log('压缩后图片尺寸:', width, 'x', height);
            canvas.width = width;
            canvas.height = height;
            const ctx = canvas.getContext('2d');
            ctx?.drawImage(img, 0, 0, width, height);
            
            // 转换为Blob
            canvas.toBlob((blob) => {
              if (blob) {
                // 将Blob转换为data URL
                const reader = new FileReader();
                reader.onloadend = () => {
                  const compressedDataURL = reader.result as string;
                  console.log('压缩后图片长度:', compressedDataURL.length);
                  console.log('压缩后图片类型:', compressedDataURL.substring(0, 30) + '...');
                  setImagePreview(compressedDataURL);
                  console.log('图片已压缩，大小约为:', Math.round(blob.size / 1024), 'KB');
                };
                reader.readAsDataURL(blob);
              }
            }, file.type, 0.7); // 0.7的质量比较好地平衡了大小和质量
          };
          img.src = reader.result as string;
        };
        reader.readAsDataURL(file);
      } else {
        // 小图片直接使用
        const reader = new FileReader();
        reader.onloadend = () => {
          if (reader.result && typeof reader.result === 'string') {
            const dataURL = reader.result;
            console.log('小图片长度:', dataURL.length);
            console.log('小图片类型:', dataURL.substring(0, 30) + '...');
            setImagePreview(dataURL);
            console.log('使用原始图片，大小约为:', Math.round(file.size / 1024), 'KB');
          }
        };
        reader.readAsDataURL(file);
      }
    }
  };
  
  const removeImage = () => {
    setImagePreview(null);
  };
  
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    // 验证必填字段
    if (!formData.title || !formData.description || !formData.reward || !formData.deadline || !formData.completionCriteria) {
      alert('请填写所有必填字段');
      return;
    }

    // 设置创建中状态为true
    setLocalIsCreating(true);

    try {
      // 确保钱包连接
      if (!address) {
        alert('请先连接您的钱包');
        setLocalIsCreating(false);
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
                console.log('网络可能已存在或其他错误:', addError);
              }
            }

            // 然后尝试切换到 Sepolia 网络
            try {
              await ethereum.request({
                method: 'wallet_switchEthereumChain',
                params: [{ chainId: '0xaa36a7' }] // 11155111 in hex
              });
            } catch (switchError: any) {
              console.error('切换网络失败:', switchError);
              throw switchError;
            }
          }

          // 使用 wagmi 的 switchChain 作为备选方案
          if (switchChain) {
            try {
              await switchChain({ chainId: 11155111 });
            } catch (wagmiError) {
              console.log('Wagmi切换失败，但手动切换可能已成功:', wagmiError);
            }
          }

          // 重要：轮询以确保chainId在切换尝试后已更新
          let attempts = 0;
          const maxAttempts = 30; // 尝试长达15秒（30 * 500毫秒）
          const delayMs = 500;

          while (attempts < maxAttempts) {
            if (!(window as any).ethereum) {
              alert('网络切换尝试后未找到以太坊提供者。');
              setLocalIsCreating(false);
              return;
            }
            const tempProvider = new ethers.BrowserProvider((window as any).ethereum);
            const networkFromTempProvider = await tempProvider.getNetwork();
            if (networkFromTempProvider.chainId === 11155111n) {
              console.log('钱包在切换尝试后成功检测到Sepolia。');
              break; // 退出轮询循环
            }
            await new Promise(resolve => setTimeout(resolve, delayMs));
            attempts++;
          }

          // 如果轮询后，网络仍然不是Sepolia
          if ((await new ethers.BrowserProvider((window as any).ethereum).getNetwork()).chainId !== 11155111n) {
            alert('无法自动切换到Sepolia网络。请在您的钱包中手动切换并重试。');
            setLocalIsCreating(false);
            return;
          }
          console.log('成功切换到Sepolia或已经在Sepolia上。');

        } catch (switchError) {
          console.error('切换网络失败:', switchError);
          alert('切换到Sepolia网络失败。请在您的钱包中手动切换。');
          setLocalIsCreating(false);
          return;
        }
      }

      // 检查以太坊提供者
      if (!(window as any).ethereum) {
        alert('未找到以太坊提供者');
        setLocalIsCreating(false);
        return;
      }      

      // 检查是否有MetaMask
      if (!((window as any).ethereum.isMetaMask)) {
        alert('请使用MetaMask钱包进行交易。如果已安装MetaMask，请确保它是您的默认钱包。');
        setLocalIsCreating(false);
        return;
      }      

      try {
        // 1. 直接与智能合约交互创建任务
        console.log('开始与合约交互，创建任务...');
        const provider = new ethers.BrowserProvider((window as any).ethereum);
        const signer = await provider.getSigner();
        
        // 创建合约实例
        const taskContract = new ethers.Contract(
          TASK_CONTRACT_ADDRESS,
          TaskContractABI,
          signer
        );
        
        // 将奖励金额转换为wei
        const rewardWei = ethers.parseEther(formData.reward);
        
        // 调用合约创建任务
        console.log('调用合约创建任务:', {
          title: formData.title,
          description: formData.description,
          reward: rewardWei.toString()
        });
        
        // 首先检查账户余额
        const balance = await provider.getBalance(address);
        console.log('账户余额:', ethers.formatEther(balance), 'ETH');
        console.log('需要的金额:', ethers.formatEther(rewardWei), 'ETH');
        
        if (balance < rewardWei) {
          throw new Error(`账户余额不足。当前余额: ${ethers.formatEther(balance)} ETH，需要: ${ethers.formatEther(rewardWei)} ETH`);
        }
        
        // 调用合约方法，并发送ETH
        const tx = await taskContract.createTask(
          formData.title,
          formData.description,
          rewardWei,
          {value: rewardWei}
        );
        
        console.log('交易已发送:', tx.hash);
        
        // 等待交易被确认
        const receipt = await tx.wait();
        console.log('交易已确认:', receipt);
        
        // 解析事件日志获取任务ID
        const taskCreatedEvent = receipt.logs
          .map((log: any) => {
            try {
              return taskContract.interface.parseLog({
                topics: log.topics as string[],
                data: log.data
              });
            } catch (e) {
              return null;
            }
          })
          .find((event: any) => event && event.name === 'TaskCreated');
        
        let contractTaskId = null;
        if (taskCreatedEvent) {
          contractTaskId = taskCreatedEvent.args.taskId.toString();
          console.log('从区块链获取到任务ID:', contractTaskId);
          setContractTaskId(Number(contractTaskId));
        }
        
        // 2. 然后调用后端API，让后端存储任务信息并关联区块链任务ID
        const result = await addTask({
          title: formData.title,
          description: formData.description,
          reward: formData.reward,
          deadline: formData.deadline,
          difficulty: formData.difficulty as 'easy' | 'medium' | 'hard',
          completionCriteria: formData.completionCriteria,
          createdBy: address,
          assignedChildId: selectedChild?.id || undefined,
          assignedTo: selectedChild?.walletAddress || undefined,
          contractTaskId: contractTaskId, // 将区块链任务ID传递给后端
          imageUrl: imagePreview // 添加图片URL
        });

        // 3. 如果选中了子账户，分配任务给子账户
        if (selectedChild && selectedChild.walletAddress && contractTaskId) {
          try {
            // 调用合约的assignTask方法
            const assignTx = await taskContract.assignTask(
              contractTaskId,
              selectedChild.walletAddress
            );
            
            console.log('分配任务交易已发送:', assignTx.hash);
            await assignTx.wait();
            console.log('成功分配任务给子账户:', selectedChild.walletAddress);
          } catch (assignError) {
            console.error('分配任务给子账户失败:', assignError);
            // 不阻止流程继续，因为后端会处理任务分配的数据库记录
          }
        }

        alert('任务创建成功！');
        navigate('/parent');
      } catch (contractError: any) {
        console.error('合约交互失败:', contractError);
        
        // 解析具体的错误类型
        let errorMessage = '创建任务失败: ';
        
        if (contractError.code === 'INSUFFICIENT_FUNDS') {
          errorMessage += '账户余额不足，请确保有足够的 ETH 支付交易费用和奖励金额';
        } else if (contractError.code === 'UNPREDICTABLE_GAS_LIMIT') {
          errorMessage += '无法估算 gas 费用，请检查合约地址和网络连接';
        } else if (contractError.code === 'NETWORK_ERROR') {
          errorMessage += '网络连接错误，请检查网络连接并重试';
        } else if (contractError.message?.includes('missing revert data')) {
          errorMessage += '交易被拒绝，可能是合约地址错误或网络问题。请检查合约地址: ' + TASK_CONTRACT_ADDRESS;
        } else if (contractError.message?.includes('execution reverted')) {
          errorMessage += '合约执行失败，可能是合约逻辑错误。请检查合约代码。';
        } else if (contractError.message?.includes('insufficient funds')) {
          errorMessage += '账户余额不足，请确保有足够的 ETH 支付交易费用和奖励金额';
        } else {
          errorMessage += contractError.message || '未知的合约错误';
        }
        
        alert(errorMessage);
        throw contractError;
      }
    } catch (error) {
      console.error('创建任务时出错:', error);
      const errorMessage = error instanceof Error ? error.message : '创建任务失败。请重试。';
      alert(errorMessage);
    } finally {
      setLocalIsCreating(false);
    }
  };

  return (
    <div className="max-w-2xl mx-auto">
      <div className="mb-6">
        <h1 className="text-3xl font-bold text-gray-900 mb-2">Create New Task</h1>
        <p className="text-gray-600">Define a task for your child to complete and earn rewards</p>
        
        {/* 钱包地址调试信息 */}
        <div className="mt-4 p-3 bg-blue-50 border border-blue-200 rounded-md">
          <h3 className="text-sm font-medium text-blue-800">钱包连接信息</h3>
          <p className="text-xs mt-1">Wagmi 钱包地址: {address || '未连接'}</p>
          <p className="text-xs mt-1">合约地址: {TASK_CONTRACT_ADDRESS}</p>
          <p className="text-xs mt-1">当前网络 Chain ID: {chainId}</p>
          {currentProvider && (
            <>
              <p className="text-xs mt-1">Provider 钱包地址: {currentProvider.providerAddress}</p>
              {currentProvider.providerAddress !== address && (
                <div className="mt-2 p-2 bg-red-50 border border-red-300 rounded-md">
                  <p className="text-xs text-red-500 font-bold">
                    警告：检测到钱包地址不匹配！
                  </p>
                  <p className="text-xs mt-1 text-red-500">
                    请在MetaMask扩展中切换到地址为 {address} 的账户，然后刷新页面。
                  </p>
                  <button 
                    className="mt-2 px-3 py-1 text-xs bg-blue-500 text-white rounded-md hover:bg-blue-600"
                    onClick={async () => {
                      try {
                        // 请求用户切换账户
                        await (window as any).ethereum.request({
                          method: 'wallet_requestPermissions',
                          params: [{ eth_accounts: {} }]
                        });
                        // 刷新页面
                        window.location.reload();
                      } catch (err) {
                        console.error("请求切换账户失败:", err);
                        alert("请手动在MetaMask中切换到正确的账户");
                      }
                    }}
                  >
                    切换钱包账户
                  </button>
                </div>
              )}
            </>
          )}
        </div>
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
                      className="w-full object-contain rounded-md"
                      style={{ maxHeight: "300px", width: "100%" }}
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
              disabled={localIsCreating}
            >
              Cancel
            </Button>
            <Button 
              type="submit" 
              leftIcon={<PlusCircle className="h-5 w-5" />}
              disabled={localIsCreating}
            >
              {localIsCreating ? "Creating..." : "Create Task"}
            </Button>
          </CardFooter>
        </form>
      </Card>
    </div>
  );
};

export default CreateTask;