import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { useAccount } from 'wagmi';
import { useAuthContext } from './AuthContext';
import { taskApi, childApi, apiClient, ApiResponse } from '../services/api';
import type { Task as ApiTask, Child as ApiChild } from '../services/api';
import { mockTasks } from '../data/mockTasks';
import { ethers } from 'ethers';
import { TaskContractABI } from '../contracts/TaskContract';

// Get contract address from environment variables - fallback to deployed address
const TASK_CONTRACT_ADDRESS = import.meta.env.VITE_TASK_CONTRACT_ADDRESS || '0x11dB634CFD2f58967e472a179ebDbaF8AB067144'; // Deployed TaskRegistry address

export interface Task {
  id: string;
  title: string;
  description: string;
  reward: string;
  deadline: string;
  difficulty: 'easy' | 'medium' | 'hard';
  status: 'open' | 'in-progress' | 'completed' | 'approved' | 'rejected';
  assignedTo?: string; // child wallet address
  assignedChildId?: string; // child id
  createdBy: string; // parent wallet address
  createdAt: string;
  updatedAt: string;
  completionCriteria: string;
  imageUrl?: string;
  contractTaskId?: string; // Add blockchain contract task ID
  submissionProof?: {
    description: string;
    imageUrl?: string;
    submittedAt: string;
  };
}

interface TaskContextType {
  tasks: Task[];
  addTask: (task: Omit<Task, 'id' | 'createdAt' | 'updatedAt' | 'status'>) => Promise<Task>;
  updateTask: (taskId: string, updates: Partial<Task>) => Promise<void>;
  assignTask: (taskId: string, childWalletAddress: string, childId: string) => Promise<void>;
  submitTask: (taskId: string, proof: Task['submissionProof']) => Promise<void>;
  approveTask: (taskId: string) => Promise<void>;
  rejectTask: (taskId: string) => Promise<void>;
  getTasksForChild: (childWalletAddress: string) => Task[];
  getTasksForParent: (parentWalletAddress: string) => Task[];
  getAvailableTasks: () => Task[];
  refreshTasks: () => Promise<void>;
  getAllTasks: () => Promise<{ id: number; title: string; description: string; }[]>;
}

const TaskContext = createContext<TaskContextType | undefined>(undefined);

export const TaskProvider = ({ children }: { children: ReactNode }) => {
  const { address } = useAccount();
  const { user, isAuthenticated } = useAuthContext();
  const [tasks, setTasks] = useState<Task[]>(mockTasks); // 初始化为mock数据

  // 当用户登录时，从API获取任务数据
  const fetchTasks = async () => {
    if (isAuthenticated && user && address) {
      console.log('[TaskContext] 从API获取任务数据，用户地址:', address);
      try {
        const response = await taskApi.getAll();
        if (response.success && response.data) {
          // 获取所有children数据以便映射assigned_child_id到wallet地址
          let childrenMap: { [key: string]: string } = {};
          try {
            const childrenResponse = await childApi.getAll();
            if (childrenResponse.success && childrenResponse.data) {
              childrenMap = childrenResponse.data.reduce((map: { [key: string]: string }, child: ApiChild) => {
                map[child.id.toString()] = child.wallet_address;
                return map;
              }, {});
              console.log('[TaskContext] 获取到children映射:', childrenMap);
            }
          } catch (error) {
            console.warn('[TaskContext] 获取children数据失败:', error);
          }

          // 转换API任务数据格式到本地Task格式
          const apiTasks = response.data.map((apiTask: ApiTask) => {
            const assignedChildId = apiTask.assigned_child_id ? apiTask.assigned_child_id.toString() : undefined;
            const assignedTo = assignedChildId ? childrenMap[assignedChildId] : undefined;
            
            return {
              id: apiTask.id.toString(),
              title: apiTask.title,
              description: apiTask.description,
              reward: apiTask.reward_amount,
              deadline: apiTask.due_date || new Date(Date.now() + 7 * 24 * 60 * 60 * 1000).toISOString(),
              difficulty: apiTask.difficulty || 'medium',
              status: apiTask.status === 'pending' ? 'open' as const :
                     apiTask.status === 'in_progress' ? 'in-progress' as const :
                     apiTask.status === 'completed' ? 'completed' as const :
                     apiTask.status === 'approved' ? 'approved' as const :
                     apiTask.status === 'rejected' ? 'rejected' as const : 'open' as const,
              assignedTo: assignedTo, // 根据assigned_child_id查找child的wallet地址
              assignedChildId: assignedChildId,
              createdBy: apiTask.created_by,
              createdAt: apiTask.created_at,
              updatedAt: apiTask.updated_at,
              completionCriteria: apiTask.description, // 使用description作为完成标准
              contractTaskId: apiTask.contract_task_id ? apiTask.contract_task_id.toString() : undefined // 添加合约任务ID
            };
          });
          console.log('[TaskContext] 成功获取任务:', apiTasks.length, '个任务');
          console.log('[TaskContext] 任务详情:', apiTasks.map(t => ({ id: t.id, title: t.title, assignedTo: t.assignedTo, assignedChildId: t.assignedChildId })));
          
          // 合并本地存储的任务状态
          const tasksWithLocalStatus = apiTasks.map(task => {
            const localStatus = localStorage.getItem(`task_${task.id}_status`);
            if (localStatus === 'completed' && task.status !== 'completed' && task.status !== 'approved' && task.status !== 'rejected') {
              console.log(`[TaskContext] 任务 ${task.id} 在本地存储中标记为已完成`);
              return {
                ...task,
                status: 'completed' as const
              };
            }
            return task;
          });
          
          setTasks(tasksWithLocalStatus);
        } else {
          console.log('[TaskContext] 获取任务失败，使用mock数据:', response.error);
          setTasks(mockTasks);
        }
      } catch (error) {
        console.error('[TaskContext] 获取任务时发生错误，使用mock数据:', error);
        setTasks(mockTasks);
      }
    } else {
      // 用户未登录或child用户，使用mock数据
      console.log('[TaskContext] 用户未认证，使用mock数据');
      setTasks(mockTasks);
    }
  };

  // 添加定时刷新数据的功能
  useEffect(() => {
    // 首次加载数据
    fetchTasks();

    // 设置定时器，每分钟自动刷新一次数据
    const intervalId = setInterval(() => {
      console.log('[TaskContext] 定时刷新任务数据');
      fetchTasks();
    }, 60000); // 60秒刷新一次

    // 清理函数
    return () => {
      clearInterval(intervalId);
    };
  }, [isAuthenticated, user, address]);
  
  // 在任务状态变化时刷新数据
  const refreshTasks = async () => {
    if (isAuthenticated && user) {
      console.log('[TaskContext] 手动刷新任务数据');
      await fetchTasks();
    }
  };

  const addTask = async (taskData: Omit<Task, 'id' | 'createdAt' | 'updatedAt' | 'status'>) => {
    if (!address) {
      throw new Error('无法创建任务：钱包地址不可用。');
    }
    const creatorAddress = address as string; // 明确断言 address 为 string

    try {
      // 转换本地Task格式到API格式
      // 将datetime-local格式转换为RFC3339格式
      let formattedDueDate: string | undefined = undefined;
      if (taskData.deadline) {
        // datetime-local格式: "2025-06-07T16:13"
        // RFC3339格式: "2025-06-07T16:13:00Z"
        formattedDueDate = taskData.deadline + ':00Z';
      }
      
      const apiTaskData = {
        title: taskData.title,
        description: taskData.description,
        reward_amount: taskData.reward,
        difficulty: taskData.difficulty,
        assigned_child_id: taskData.assignedChildId ? parseInt(taskData.assignedChildId) : undefined,
        due_date: formattedDueDate,
        created_by: creatorAddress,
        status: 'pending' as const,
        contract_task_id: taskData.contractTaskId ? Number(taskData.contractTaskId) : undefined // Add contract task ID
      };
      
      console.log('Sending API data:', apiTaskData);
      
      const response = await taskApi.create(apiTaskData);
      if (response.success && response.data) {
        // 转换API响应到本地格式并添加到状态
        const newTask: Task = {
          id: response.data.id.toString(),
          title: response.data.title,
          description: response.data.description,
          reward: response.data.reward_amount,
          deadline: response.data.due_date || taskData.deadline,
          difficulty: taskData.difficulty,
          status: 'open',
          assignedTo: taskData.assignedTo,
          assignedChildId: response.data.assigned_child_id?.toString(),
          createdBy: response.data.created_by,
          createdAt: response.data.created_at,
          updatedAt: response.data.updated_at,
          completionCriteria: taskData.completionCriteria,
          contractTaskId: response.data.contract_task_id?.toString() // Add contract task ID to response
        };
        setTasks(prev => [...prev, newTask]);
        console.log('Task added via API:', newTask);
        return newTask;
      } else {
        const errorMessage = response.error || '创建任务失败';
        console.error('Failed to create task:', errorMessage);
        throw new Error(errorMessage);
      }
    } catch (error) {
      console.error('Error creating task:', error);
      if (error instanceof Error) {
        throw error;
      }
      throw new Error('创建任务时发生未知错误');
    }
  };

  const updateTask = async (taskId: string, updates: Partial<Task>) => {
    console.log('updateTask', taskId, updates);
    // 立即更新本地状态
    setTasks(prev => {
      const newTasks = prev.map(task => 
        task.id === taskId 
          ? { ...task, ...updates, updatedAt: new Date().toISOString() }
          : task
      );
      
      // 找到更新后的任务进行日志输出
      const updatedTask = newTasks.find(t => t.id === taskId);
      console.log('任务状态已更新为:', updatedTask?.status);
      
      return newTasks;
    });
    
    // 同步到后端
    try {
      // 后端API期望的状态格式
      let apiStatus: 'pending' | 'in_progress' | 'completed' | 'approved' | 'rejected' | undefined;
      
      switch(updates.status) {
        case 'in-progress':
          apiStatus = 'in_progress';
          break;
        case 'completed':
          apiStatus = 'completed';
          break;
        case 'approved':
          apiStatus = 'approved';
          break;
        case 'rejected':
          apiStatus = 'rejected';
          break;
        case 'open':
          apiStatus = 'pending';
          break;
        default:
          apiStatus = undefined;
      }
      
      // 确保ID是数字类型
      const taskIdNumber = parseInt(taskId);
      
      // 创建API请求数据对象
      const apiData: any = {};
      
      // 只添加需要更新的字段
      if (apiStatus) {
        apiData.status = apiStatus;
      }
      
      // 处理任务提交相关信息
      if (updates.submissionProof) {
        apiData.submission_note = updates.submissionProof.description;
      }
      
      console.log('正在更新任务，ID:', taskIdNumber, '数据:', apiData);
      
      // 使用完成任务的专用API端点，而不是普通更新
      if (apiStatus === 'completed') {
        const response = await taskApi.complete(taskIdNumber);
        if (response.success) {
          console.log('任务已成功标记为完成');
        } else {
          console.error('标记任务完成失败:', response.error);
          // 失败时重新从后端获取数据
          await fetchTasks();
        }
      } else if (apiStatus === 'approved') {
        const response = await taskApi.approve(taskIdNumber);
        if (response.success) {
          console.log('任务已成功批准');
          alert('任务已成功批准');
        } else {
          console.error('批准任务失败:', response.error);
          // 失败时重新从后端获取数据
          await fetchTasks();
        }
      } else if (apiStatus === 'rejected') {
        const response = await taskApi.reject(taskIdNumber);
        if (response.success) {
          console.log('任务已成功拒绝');
          alert('任务已成功拒绝');
        } else {
          console.error('拒绝任务失败:', response.error);
          // 失败时重新从后端获取数据
          await fetchTasks();
        }
      } else if (Object.keys(apiData).length > 0) {
        // 其他普通更新
        const response = await taskApi.update(taskIdNumber, apiData);
        if (response.success) {
          console.log('任务已成功更新');
        } else {
          console.error('更新任务失败:', response.error);
          // 失败时重新从后端获取数据
          await fetchTasks();
        }
      }
    } catch (error) {
      console.error('更新任务状态失败:', error);
      // 出错时重新从后端获取数据
      await fetchTasks();
    }
  };

  const assignTask = async (taskId: string, childWalletAddress: string, childId: string) => {
    console.log('assignTask', taskId, childWalletAddress, childId);
    
    // 获取任务信息
    const task = tasks.find(t => t.id === taskId);
    if (!task) {
      console.error('任务不存在:', taskId);
      alert('任务不存在');
      return;
    }
    
    // 检查是否有合约任务ID
    if (task.contractTaskId) {
      // 检查以太坊提供者
      if (!(window as any).ethereum) {
        alert('未找到以太坊提供者，请安装MetaMask');
        return;
      }
      
      try {
        console.log('开始与合约交互，分配任务...');
        const provider = new ethers.BrowserProvider((window as any).ethereum);
        
        // 检查网络是否为Sepolia
        const network = await provider.getNetwork();
        console.log('当前网络:', network.name, '链ID:', network.chainId.toString());
        
        if (network.chainId !== 11155111n) {
          console.log('当前不在Sepolia网络，尝试切换...');
          try {
            await (window as any).ethereum.request({
              method: 'wallet_switchEthereumChain',
              params: [{ chainId: '0xaa36a7' }] // Sepolia chainId
            });
            console.log('已切换到Sepolia网络');
          } catch (switchError) {
            console.error('切换网络失败:', switchError);
            alert('请在钱包中手动切换到Sepolia测试网络');
            return;
          }
        }
        
        const signer = await provider.getSigner();
        const signerAddress = await signer.getAddress();
        console.log('签名者地址:', signerAddress);
        
        // 创建合约实例
        const taskContract = new ethers.Contract(
          TASK_CONTRACT_ADDRESS,
          TaskContractABI,
          signer
        );
        
        // 检查合约方法
        console.log('合约方法:', Object.keys(taskContract.interface));
        
        console.log('调用合约分配任务:', {
          taskId: Number(task.contractTaskId),
          childAddress: childWalletAddress
        });
        
        // 调用合约方法分配任务
        const tx = await taskContract.assignTask(
          Number(task.contractTaskId),
          childWalletAddress
        );
        
        console.log('交易已发送:', tx.hash);
        
        // 等待交易被确认
        const receipt = await tx.wait();
        console.log('交易已确认:', receipt);
        
        console.log('成功在区块链上分配任务');
      } catch (contractError) {
        console.error('与智能合约交互时出错:', contractError);
        alert(`与智能合约交互失败: ${contractError instanceof Error ? contractError.message : '未知错误'}`);
        return;
      }
    } else {
      console.log('任务没有区块链ID，跳过合约交互');
    }
    
    // 更新后端状态
    await updateTask(taskId, {
      assignedTo: childWalletAddress,
      assignedChildId: childId,
      status: 'in-progress'
    });
    
    // 确保数据与后端同步
    await refreshTasks();
  };

  const submitTask = async (taskId: string, proof: Task['submissionProof']) => {
    console.log('提交任务前状态:', tasks.find(t => t.id === taskId)?.status);
    console.log('使用的合约地址:', TASK_CONTRACT_ADDRESS);
    
    if (!proof) {
      console.error('提交证明不能为空');
      alert('提交证明不能为空');
      return;
    }
    
    // 确保子女用户才能完成任务
    if (!isAuthenticated || !user || user.role !== 'child') {
      console.error('只有子女才能完成任务');
      alert('权限错误：只有子女可以完成任务');
      return;
    }
    
    try {
      const taskIdNumber = parseInt(taskId);
      const task = tasks.find(t => t.id === taskId);
      
      if (!task) {
        console.error('任务不存在:', taskId);
        return;
      }
      
      console.log('==== 任务信息 ====');
      console.log('任务ID:', taskIdNumber);
      console.log('任务标题:', task.title);
      console.log('当前状态:', task.status);
      console.log('合约任务ID:', task.contractTaskId);
      console.log('合约任务ID类型:', typeof task.contractTaskId);
      console.log('合约任务ID类型:', typeof task.contractTaskId);

      // 检查以太坊提供者
      if (!(window as any).ethereum) {
        alert('未找到以太坊提供者，请安装MetaMask');
        return;
      }

      // 检查是否有合约任务ID
      if (!task.contractTaskId) {
        console.warn('任务没有关联的区块链合约ID，只更新数据库');
      } else {
        // 1. 直接与智能合约交互，提交任务完成证明
        try {
          console.log('开始与合约交互，提交任务完成证明...');
          const provider = new ethers.BrowserProvider((window as any).ethereum);

          // 检查网络是否为Sepolia
          const network = await provider.getNetwork();
          console.log('当前网络:', network.name, '链ID:', network.chainId.toString());
          
          if (network.chainId !== 11155111n) {
            console.log('当前不在Sepolia网络，尝试切换...');
            try {
              await (window as any).ethereum.request({
                method: 'wallet_switchEthereumChain',
                params: [{ chainId: '0xaa36a7' }] // Sepolia chainId
              });
              console.log('已切换到Sepolia网络');
            } catch (switchError) {
              console.error('切换网络失败:', switchError);
              alert('请在钱包中手动切换到Sepolia测试网络');
              return;
            }
          }

          const signer = await provider.getSigner();
          const signerAddress = await signer.getAddress();
          console.log('签名者地址:', signerAddress);
          
          // 创建合约实例
          const taskContract = new ethers.Contract(
            TASK_CONTRACT_ADDRESS,
            TaskContractABI,
            signer
          );
          
          // 检查合约方法
          console.log('合约方法:', Object.keys(taskContract.interface));
          
          console.log('调用合约提交任务完成:', {
            taskId: Number(task.contractTaskId)
          });
          
          // 调用合约方法提交任务
          const tx = await taskContract.completeTask(
            Number(task.contractTaskId)
          );
          
          console.log('交易已发送:', tx.hash);
          
          // 等待交易被确认
          const receipt = await tx.wait();
          console.log('交易已确认:', receipt);
          
          console.log('成功在区块链上提交任务完成证明');
        } catch (contractError) {
          console.error('与智能合约交互时出错:', contractError);
          alert(`与智能合约交互失败: ${contractError instanceof Error ? contractError.message : '未知错误'}`);
          return;
        }
      }

      // 从localStorage获取认证令牌
      const token = localStorage.getItem('auth_token');
      if (!token) {
        console.error('无法完成任务: 未找到认证令牌');
        alert('未找到认证令牌，请重新登录');
        return;
      }
      
      // 获取API基础URL
      const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1';
      
      // 检查用户角色
      const userRole = localStorage.getItem('user_role') || '';
      console.log('当前用户角色:', userRole);
      
      if (userRole !== 'child') {
        console.warn('警告: 非子女角色尝试完成任务');
      }
      
      // 构建请求数据
      const requestData = {
        completion_proof: proof.description || '任务已完成'
      };
      
      console.log('==== 请求详情 ====');
      console.log('请求URL:', `${API_BASE_URL}/tasks/${taskIdNumber}/complete`);
      console.log('请求方法:', 'POST');
      console.log('请求头:', {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token.substring(0, 15)}...`
      });
      console.log('请求体:', requestData);
      
      // 暂时更新本地状态，提供即时反馈，但依赖后端确认
      setTasks(prev => {
        const newTasks = prev.map(task => 
          task.id === taskId 
            ? { 
                ...task, 
                status: 'completed' as const, 
                submissionProof: proof,
                updatedAt: new Date().toISOString() 
              }
            : task
        );
        return newTasks;
      });
      
      // 调用完成任务API
      const response = await fetch(`${API_BASE_URL}/tasks/${taskIdNumber}/complete`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify(requestData)
      });
      
      if (response.ok) {
        const result = await response.json();
        console.log('任务完成成功:', result);
        alert('任务已成功提交并更新到数据库！');
        
        // 成功后重新获取任务列表以确保数据同步
        await fetchTasks();
      } else {
        const errorData = await response.text();
        let parsedError;
        try {
          parsedError = JSON.parse(errorData);
        } catch (e) {
          parsedError = { error: errorData || '未知错误' };
        }
        
        console.error('任务完成失败:', response.status);
        console.error('错误详情:', errorData);
        
        // 根据错误状态码提供更具体的错误信息
        let errorMessage = '任务提交失败';
        if (response.status === 401) {
          errorMessage = '认证失败，请重新登录';
        } else if (response.status === 403) {
          errorMessage = '权限不足，只有被分配的孩子可以完成任务';
        } else if (response.status === 404) {
          errorMessage = '任务不存在';
        } else if (response.status === 400) {
          errorMessage = parsedError.error || '任务状态不正确，只能完成进行中的任务';
        }
        
        alert(`任务提交失败: ${errorMessage}`);
        
        // API调用失败时，回滚前端状态
        await fetchTasks();
      }
    } catch (error) {
      console.error('网络错误:', error);
      alert('网络连接失败，请检查网络连接后重试');
      
      // 网络错误时，回滚前端状态
      await fetchTasks();
    }
  };

  const approveTask = async (taskId: string) => {
    try {
      // 获取任务信息
      const task = tasks.find(t => t.id === taskId);
      if (!task) {
        console.error('找不到任务:', taskId);
        alert('找不到任务');
        return;
      }
      
      const taskIdNumber = parseInt(taskId);
      if (isNaN(taskIdNumber)) {
        console.error('无效的任务ID:', taskId);
        alert('无效的任务ID');
        return;
      }
      
      console.log('==== 批准任务 ====');
      console.log('任务ID:', taskIdNumber);
      console.log('任务标题:', task.title);
      console.log('当前状态:', task.status);
      console.log('合约任务ID:', task.contractTaskId);
      console.log('合约任务ID类型:', typeof task.contractTaskId);

      // 检查以太坊提供者
      if (!(window as any).ethereum) {
        alert('未找到以太坊提供者，请安装MetaMask');
        return;
      }

      // 检查是否有合约任务ID
      if (!task.contractTaskId) {
        console.warn('任务没有关联的区块链合约ID，只更新数据库');
      } else {
        // 与智能合约交互，批准任务
        try {
          console.log('开始与合约交互，批准任务...');
          const provider = new ethers.BrowserProvider((window as any).ethereum);

          // 检查网络是否为Sepolia
          const network = await provider.getNetwork();
          console.log('当前网络:', network.name, '链ID:', network.chainId.toString());
          
          if (network.chainId !== 11155111n) {
            console.log('当前不在Sepolia网络，尝试切换...');
            try {
              await (window as any).ethereum.request({
                method: 'wallet_switchEthereumChain',
                params: [{ chainId: '0xaa36a7' }] // Sepolia chainId
              });
              console.log('已切换到Sepolia网络');
            } catch (switchError) {
              console.error('切换网络失败:', switchError);
              alert('请在钱包中手动切换到Sepolia测试网络');
              return;
            }
          }

          const signer = await provider.getSigner();
          const signerAddress = await signer.getAddress();
          console.log('签名者地址:', signerAddress);
          
          // 创建合约实例
          const taskContract = new ethers.Contract(
            TASK_CONTRACT_ADDRESS,
            TaskContractABI,
            signer
          );
          
          // 检查合约方法
          console.log('合约方法:', Object.keys(taskContract.interface));
          
          console.log('调用合约批准任务:', {
            taskId: Number(task.contractTaskId)
          });
          
          // 调用合约方法批准任务
          const tx = await taskContract.approveTask(
            Number(task.contractTaskId)
          );
          
          console.log('交易已发送:', tx.hash);
          
          // 等待交易被确认
          const receipt = await tx.wait();
          console.log('交易已确认:', receipt);
          
          console.log('成功在区块链上批准任务');
        } catch (contractError) {
          console.error('与智能合约交互时出错:', contractError);
          alert(`与智能合约交互失败: ${contractError instanceof Error ? contractError.message : '未知错误'}`);
          return;
        }
      }

      // 调用后端API更新状态
      await updateTask(taskId, { status: 'approved' });
      
      // 刷新数据确保同步
      await refreshTasks();
    } catch (error) {
      console.error('批准任务时发生错误:', error);
      alert('批准任务失败，请重试');
    }
  };

  const rejectTask = async (taskId: string) => {
    try {
      // 获取任务信息
      const task = tasks.find(t => t.id === taskId);
      if (!task) {
        console.error('找不到任务:', taskId);
        alert('找不到任务');
        return;
      }
      
      const taskIdNumber = parseInt(taskId);
      if (isNaN(taskIdNumber)) {
        console.error('无效的任务ID:', taskId);
        alert('无效的任务ID');
        return;
      }
      
      console.log('==== 拒绝任务 ====');
      console.log('任务ID:', taskIdNumber);
      console.log('任务标题:', task.title);
      console.log('当前状态:', task.status);
      console.log('合约任务ID:', task.contractTaskId);
      console.log('合约任务ID类型:', typeof task.contractTaskId);

      // 检查以太坊提供者
      if (!(window as any).ethereum) {
        alert('未找到以太坊提供者，请安装MetaMask');
        return;
      }

      // 检查是否有合约任务ID
      if (!task.contractTaskId) {
        console.warn('任务没有关联的区块链合约ID，只更新数据库');
      } else {
        // 与智能合约交互，拒绝任务
        try {
          console.log('开始与合约交互，拒绝任务...');
          const provider = new ethers.BrowserProvider((window as any).ethereum);

          // 检查网络是否为Sepolia
          const network = await provider.getNetwork();
          console.log('当前网络:', network.name, '链ID:', network.chainId.toString());
          
          if (network.chainId !== 11155111n) {
            console.log('当前不在Sepolia网络，尝试切换...');
            try {
              await (window as any).ethereum.request({
                method: 'wallet_switchEthereumChain',
                params: [{ chainId: '0xaa36a7' }] // Sepolia chainId
              });
              console.log('已切换到Sepolia网络');
            } catch (switchError) {
              console.error('切换网络失败:', switchError);
              alert('请在钱包中手动切换到Sepolia测试网络');
              return;
            }
          }

          const signer = await provider.getSigner();
          const signerAddress = await signer.getAddress();
          console.log('签名者地址:', signerAddress);
          
          // 创建合约实例
          const taskContract = new ethers.Contract(
            TASK_CONTRACT_ADDRESS,
            TaskContractABI,
            signer
          );
          
          // 检查合约方法
          console.log('合约方法:', Object.keys(taskContract.interface));
          
          console.log('调用合约拒绝任务:', {
            taskId: Number(task.contractTaskId)
          });
          
          // 调用合约方法拒绝任务
          const tx = await taskContract.rejectTask(
            Number(task.contractTaskId)
          );
          
          console.log('交易已发送:', tx.hash);
          
          // 等待交易被确认
          const receipt = await tx.wait();
          console.log('交易已确认:', receipt);
          
          console.log('成功在区块链上拒绝任务');
        } catch (contractError) {
          console.error('与智能合约交互时出错:', contractError);
          alert(`与智能合约交互失败: ${contractError instanceof Error ? contractError.message : '未知错误'}`);
          return;
        }
      }

      // 调用后端API更新状态
      await updateTask(taskId, { status: 'rejected' });
      
      // 刷新数据确保同步
      await refreshTasks();
    } catch (error) {
      console.error('拒绝任务时发生错误:', error);
      alert('拒绝任务失败，请重试');
    }
  };

  const getTasksForChild = (childWalletAddress: string): Task[] => {
    console.log('[getTasksForChild] 查找child任务，地址:', childWalletAddress);
    console.log('[getTasksForChild] 所有任务:', tasks.map(t => ({ id: t.id, title: t.title, assignedTo: t.assignedTo, status: t.status })));
    const filteredTasks = tasks.filter(task => task.assignedTo?.toLowerCase() === childWalletAddress.toLowerCase());
    console.log('[getTasksForChild] 过滤后的任务:', filteredTasks.map(t => ({ id: t.id, title: t.title, assignedTo: t.assignedTo, status: t.status })));
    return filteredTasks;
  };

  const getTasksForParent = (parentWalletAddress: string): Task[] => {
    console.log('[getTasksForParent] 查找父母地址:', parentWalletAddress);
    console.log('[getTasksForParent] 所有任务:', tasks.map(t => ({ id: t.id, title: t.title, createdBy: t.createdBy })));
    
    // 获取父母创建的任务
    const filteredTasks = tasks.filter(task => task.createdBy.toLowerCase() === parentWalletAddress.toLowerCase());
    
    console.log('[getTasksForParent] 过滤后的任务:', filteredTasks.map(t => ({ id: t.id, title: t.title, createdBy: t.createdBy, status: t.status })));
    return filteredTasks;
  };

  const getAvailableTasks = (): Task[] => {
    console.log('[getAvailableTasks] 查找可用任务');
    const availableTasks = tasks.filter(task => task.status === 'open' || task.status === 'in-progress');
    console.log('[getAvailableTasks] 可用任务:', availableTasks.map(t => ({ id: t.id, title: t.title, status: t.status })));
    return availableTasks;
  };

  const getAllTasks = async (): Promise<Task[]> => {
    try {
      const response = await taskApi.getAll();
      console.log('API 响应:', response); // 添加调试信息
      if (!response.success) {
        throw new Error(response.error || '获取任务失败');
      }
      const data = response.data || [];
      console.log('获取到的任务数据:', data); // 添加调试信息
      
      // 获取所有children数据以便映射assigned_child_id到wallet地址
      let childrenMap: { [key: string]: string } = {};
      try {
        const childrenResponse = await childApi.getAll();
        if (childrenResponse.success && childrenResponse.data) {
          childrenMap = childrenResponse.data.reduce((map: { [key: string]: string }, child: ApiChild) => {
            map[child.id.toString()] = child.wallet_address;
            return map;
          }, {});
        }
      } catch (error) {
        console.warn('[TaskContext] 获取children数据失败:', error);
      }
      
      return data.map((apiTask: ApiTask) => {
        const assignedChildId = apiTask.assigned_child_id ? apiTask.assigned_child_id.toString() : undefined;
        const assignedTo = assignedChildId ? childrenMap[assignedChildId] : undefined;
        
        return {
          id: apiTask.id.toString(),
          title: apiTask.title,
          description: apiTask.description,
          reward: apiTask.reward_amount,
          deadline: apiTask.due_date || new Date(Date.now() + 7 * 24 * 60 * 60 * 1000).toISOString(),
          difficulty: apiTask.difficulty || 'medium',
          status: apiTask.status === 'pending' ? 'open' as const :
                 apiTask.status === 'in_progress' ? 'in-progress' as const :
                 apiTask.status === 'completed' ? 'completed' as const :
                 apiTask.status === 'approved' ? 'approved' as const :
                 apiTask.status === 'rejected' ? 'rejected' as const : 'open' as const,
          assignedTo: assignedTo,
          assignedChildId: assignedChildId,
          createdBy: apiTask.created_by,
          createdAt: apiTask.created_at,
          updatedAt: apiTask.updated_at,
          completionCriteria: apiTask.description,
          contractTaskId: apiTask.contract_task_id ? apiTask.contract_task_id.toString() : undefined
        };
      });
    } catch (error) {
      console.error('获取任务失败:', error);
      return []; // 返回空数组以防止崩溃
    }
  };

  return (
    <TaskContext.Provider value={{
      tasks,
      addTask,
      updateTask,
      assignTask,
      submitTask,
      approveTask,
      rejectTask,
      getTasksForChild,
      getTasksForParent,
      getAvailableTasks,
      refreshTasks,
      getAllTasks
    }}>
      {children}
    </TaskContext.Provider>
  );
};

export const useTask = (): TaskContextType => {
  const context = useContext(TaskContext);
  if (context === undefined) {
    throw new Error('useTask must be used within a TaskProvider');
  }
  return context;
};