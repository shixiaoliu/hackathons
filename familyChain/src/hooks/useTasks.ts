import { useState, useEffect, useCallback } from 'react';
import { taskApi } from '../services/api';
import type { Task } from '../services/api';
import { useAuth } from './useAuth';

interface TasksState {
  tasks: Task[];
  loading: boolean;
  error: string | null;
}

export const useTasks = (childId?: number) => {
  const { user, isAuthenticated } = useAuth();
  const [tasksState, setTasksState] = useState<TasksState>({
    tasks: [],
    loading: false,
    error: null,
  });

  // 加载任务列表
  const loadTasks = useCallback(async (params?: { child_id?: number; status?: string }) => {
    if (!isAuthenticated) return;

    setTasksState(prev => ({ ...prev, loading: true, error: null }));

    try {
      const response = await taskApi.getAll(params);
      if (response.success && response.data) {
        setTasksState({
          tasks: response.data,
          loading: false,
          error: null,
        });
      } else {
        setTasksState(prev => ({
          ...prev,
          loading: false,
          error: response.error || '加载任务失败',
        }));
      }
    } catch (error) {
      setTasksState(prev => ({
        ...prev,
        loading: false,
        error: error instanceof Error ? error.message : '加载任务失败',
      }));
    }
  }, [isAuthenticated]);

  // 创建任务
  const createTask = useCallback(async (taskData: {
    title: string;
    description: string;
    reward_amount: string;
    difficulty: 'easy' | 'medium' | 'hard';
    assigned_child_id: number;
    due_date?: string;
  }) => {
    if (!user || user.role !== 'parent') {
      throw new Error('只有家长可以创建任务');
    }

    try {
      const response = await taskApi.create({
        ...taskData,
        status: 'pending',
        created_by: user.wallet_address,
      });

      if (response.success && response.data) {
        // 重新加载任务列表
        await loadTasks(childId ? { child_id: childId } : undefined);
        return response.data;
      } else {
        throw new Error(response.error || '创建任务失败');
      }
    } catch (error) {
      throw error;
    }
  }, [user, loadTasks, childId]);

  // 更新任务
  const updateTask = useCallback(async (taskId: number, updates: Partial<Task>) => {
    try {
      const response = await taskApi.update(taskId, updates);
      if (response.success) {
        // 更新本地状态
        setTasksState(prev => ({
          ...prev,
          tasks: prev.tasks.map(task =>
            task.id === taskId ? { ...task, ...updates } : task
          ),
        }));
        return response.data;
      } else {
        throw new Error(response.error || '更新任务失败');
      }
    } catch (error) {
      throw error;
    }
  }, []);

  // 完成任务（儿童操作）
  const completeTask = useCallback(async (taskId: number) => {
    if (!user || user.role !== 'child') {
      throw new Error('只有儿童可以完成任务');
    }

    try {
      const response = await taskApi.complete(taskId);
      if (response.success) {
        // 更新本地状态
        setTasksState(prev => ({
          ...prev,
          tasks: prev.tasks.map(task =>
            task.id === taskId ? { ...task, status: 'completed' } : task
          ),
        }));
      } else {
        throw new Error(response.error || '完成任务失败');
      }
    } catch (error) {
      throw error;
    }
  }, [user]);

  // 批准任务（家长操作）
  const approveTask = useCallback(async (taskId: number) => {
    if (!user || user.role !== 'parent') {
      throw new Error('只有家长可以批准任务');
    }

    try {
      const response = await taskApi.approve(taskId);
      if (response.success) {
        // 更新本地状态
        setTasksState(prev => ({
          ...prev,
          tasks: prev.tasks.map(task =>
            task.id === taskId ? { ...task, status: 'approved' } : task
          ),
        }));
      } else {
        throw new Error(response.error || '批准任务失败');
      }
    } catch (error) {
      throw error;
    }
  }, [user]);

  // 拒绝任务（家长操作）
  const rejectTask = useCallback(async (taskId: number, reason?: string) => {
    if (!user || user.role !== 'parent') {
      throw new Error('只有家长可以拒绝任务');
    }

    try {
      const response = await taskApi.reject(taskId, reason);
      if (response.success) {
        // 更新本地状态
        setTasksState(prev => ({
          ...prev,
          tasks: prev.tasks.map(task =>
            task.id === taskId ? { ...task, status: 'rejected' } : task
          ),
        }));
      } else {
        throw new Error(response.error || '拒绝任务失败');
      }
    } catch (error) {
      throw error;
    }
  }, [user]);

  // 获取特定任务
  const getTask = useCallback(async (taskId: number) => {
    try {
      const response = await taskApi.getById(taskId);
      if (response.success && response.data) {
        return response.data;
      } else {
        throw new Error(response.error || '获取任务失败');
      }
    } catch (error) {
      throw error;
    }
  }, []);

  // 初始加载
  useEffect(() => {
    if (isAuthenticated) {
      loadTasks(childId ? { child_id: childId } : undefined);
    }
  }, [isAuthenticated, childId, loadTasks]);

  // 清除错误
  const clearError = useCallback(() => {
    setTasksState(prev => ({ ...prev, error: null }));
  }, []);

  return {
    ...tasksState,
    loadTasks,
    createTask,
    updateTask,
    completeTask,
    approveTask,
    rejectTask,
    getTask,
    clearError,
  };
};