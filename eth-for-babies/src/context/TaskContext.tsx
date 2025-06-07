import React, { createContext, useContext, useState, ReactNode, useEffect } from 'react';
import { useAccount } from 'wagmi';
import { useAuthContext } from './AuthContext';
import { taskApi } from '../services/api';
import type { Task as ApiTask } from '../services/api';

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
  submissionProof?: {
    description: string;
    imageUrl?: string;
    submittedAt: string;
  };
}

interface TaskContextType {
  tasks: Task[];
  addTask: (task: Omit<Task, 'id' | 'createdAt' | 'updatedAt' | 'status'>) => Promise<Task>;
  updateTask: (taskId: string, updates: Partial<Task>) => void;
  assignTask: (taskId: string, childWalletAddress: string, childId: string) => void;
  submitTask: (taskId: string, proof: Task['submissionProof']) => void;
  approveTask: (taskId: string) => void;
  rejectTask: (taskId: string) => void;
  getTasksForChild: (childWalletAddress: string) => Task[];
  getTasksForParent: (parentWalletAddress: string) => Task[];
  getAvailableTasks: () => Task[];
}

const TaskContext = createContext<TaskContextType | undefined>(undefined);

export const TaskProvider = ({ children }: { children: ReactNode }) => {
  const { address } = useAccount();
  const { user, isAuthenticated } = useAuthContext();
  const [tasks, setTasks] = useState<Task[]>([]);

  // 当用户登录时，从API获取任务数据
  useEffect(() => {
    const fetchTasks = async () => {
      if (isAuthenticated && user && address) {
        console.log('[TaskContext] 从API获取任务数据，用户地址:', address);
        try {
          const response = await taskApi.getAll();
          if (response.success && response.data) {
            // 转换API任务数据格式到本地Task格式
            const apiTasks = response.data.map((apiTask: ApiTask) => ({
              id: apiTask.id.toString(),
              title: apiTask.title,
              description: apiTask.description,
              reward: apiTask.reward_amount,
              deadline: apiTask.due_date || new Date(Date.now() + 7 * 24 * 60 * 60 * 1000).toISOString(),
              difficulty: 'medium' as const, // API中没有difficulty字段，设置默认值
              status: apiTask.status === 'pending' ? 'open' as const :
                     apiTask.status === 'in_progress' ? 'in-progress' as const :
                     apiTask.status === 'completed' ? 'completed' as const :
                     apiTask.status === 'approved' ? 'approved' as const :
                     apiTask.status === 'rejected' ? 'rejected' as const : 'open' as const,
              assignedTo: undefined, // 需要根据assigned_child_id查找child的wallet地址
              assignedChildId: apiTask.assigned_child_id ? apiTask.assigned_child_id.toString() : undefined,
              createdBy: apiTask.created_by,
              createdAt: apiTask.created_at,
              updatedAt: apiTask.updated_at,
              completionCriteria: apiTask.description // 使用description作为完成标准
            }));
            console.log('[TaskContext] 成功获取任务:', apiTasks.length, '个任务');
            setTasks(apiTasks);
          } else {
            console.log('[TaskContext] 获取任务失败:', response.error);
            setTasks([]);
          }
        } catch (error) {
          console.error('[TaskContext] 获取任务时发生错误:', error);
          setTasks([]);
        }
      } else if (!isAuthenticated) {
         // 用户未登录时清空任务
         setTasks([]);
       }
     };

     fetchTasks();
   }, [isAuthenticated, user, address]);

  const addTask = async (taskData: Omit<Task, 'id' | 'createdAt' | 'updatedAt' | 'status'>) => {
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
        due_date: formattedDueDate
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
          completionCriteria: taskData.completionCriteria
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

  const updateTask = (taskId: string, updates: Partial<Task>) => {
    setTasks(prev => prev.map(task => 
      task.id === taskId 
        ? { ...task, ...updates, updatedAt: new Date().toISOString() }
        : task
    ));
  };

  const assignTask = (taskId: string, childWalletAddress: string, childId: string) => {
    updateTask(taskId, {
      assignedTo: childWalletAddress,
      assignedChildId: childId,
      status: 'in-progress'
    });
  };

  const submitTask = (taskId: string, proof: Task['submissionProof']) => {
    updateTask(taskId, {
      status: 'completed',
      submissionProof: proof
    });
  };

  const approveTask = (taskId: string) => {
    updateTask(taskId, { status: 'approved' });
  };

  const rejectTask = (taskId: string) => {
    updateTask(taskId, { status: 'rejected' });
  };

  const getTasksForChild = (childWalletAddress: string): Task[] => {
    return tasks.filter(task => task.assignedTo?.toLowerCase() === childWalletAddress.toLowerCase());
  };

  const getTasksForParent = (parentWalletAddress: string): Task[] => {
    console.log('[getTasksForParent] 查找父母地址:', parentWalletAddress);
    console.log('[getTasksForParent] 所有任务:', tasks.map(t => ({ id: t.id, title: t.title, createdBy: t.createdBy })));
    const filteredTasks = tasks.filter(task => task.createdBy.toLowerCase() === parentWalletAddress.toLowerCase());
    console.log('[getTasksForParent] 过滤后的任务:', filteredTasks.map(t => ({ id: t.id, title: t.title, createdBy: t.createdBy })));
    return filteredTasks;
  };

  const getAvailableTasks = (): Task[] => {
    return tasks.filter(task => task.status === 'open');
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
      getAvailableTasks
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