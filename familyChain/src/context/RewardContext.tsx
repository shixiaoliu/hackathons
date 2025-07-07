import React, { createContext, useState, useEffect, ReactNode } from 'react';
import { rewardApi, exchangeApi, familyApi } from '../services/api';
import { Reward, Exchange } from '../types/reward';
import { useAuthContext } from './AuthContext';
import { useFamily } from './FamilyContext';

// 定义上下文类型
interface RewardContextType {
  rewards: Reward[];
  exchanges: Exchange[];
  loading: boolean;
  error: string | null;
  fetchRewards: () => Promise<void>;
  setRewards: React.Dispatch<React.SetStateAction<Reward[]>>;
  createReward: (rewardData: {
    name: string;
    description: string;
    image_url: string;
    token_price: number;
    stock: number;
    contract_reward_id?: number;
  }) => Promise<Reward | null>;
  updateReward: (id: number, data: Partial<Reward>) => Promise<boolean>;
  deleteReward: (id: number) => Promise<boolean>;
  fetchExchanges: () => Promise<void>;
  approveExchange: (id: number, notes?: string) => Promise<boolean>;
  cancelExchange: (id: number, notes?: string) => Promise<boolean>;
}

// 创建上下文
export const RewardContext = createContext<RewardContextType | undefined>(undefined);

// 提供者组件
export const RewardProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [rewards, setRewards] = useState<Reward[]>([]);
  const [exchanges, setExchanges] = useState<Exchange[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  
  const { user } = useAuthContext();
  const { selectedFamily, currentChild } = useFamily();

  // 获取家庭奖品列表
  const fetchRewards = async () => {
    let familyId: number | null = null;
    
    // 优先使用selectedFamily
    if (selectedFamily && selectedFamily.id) {
      familyId = parseInt(selectedFamily.id);
      console.log('从selectedFamily获取家庭ID:', familyId);
    } 
    // 如果是孩子用户但没有selectedFamily，尝试从API获取家庭信息
    else if (user && user.role === 'child') {
      console.log('当前是孩子用户，尝试获取所有家庭');
      try {
        const familiesResponse = await familyApi.getAll();
        if (familiesResponse.success && familiesResponse.data && familiesResponse.data.length > 0) {
          // 获取第一个家庭的ID
          familyId = familiesResponse.data[0].id;
          console.log('获取到的第一个家庭ID:', familyId);
        } else {
          console.log('获取家庭列表失败或列表为空:', familiesResponse);
          // 尝试使用默认ID 1
          familyId = 1;
          console.log('使用默认家庭ID: 1');
        }
      } catch (error) {
        console.error('获取家庭列表出错:', error);
        // 出错时使用默认ID 1
        familyId = 1;
        console.log('出错后使用默认家庭ID: 1');
      }
    }
    
    if (!familyId) {
      console.log('没有找到有效的家庭ID，无法获取奖品列表');
      return;
    }

    setLoading(true);
    setError(null);
    console.log('正在获取家庭奖品列表，家庭ID:', familyId);
    try {
      // 首先尝试获取激活的奖品，不考虑库存
      let response = await rewardApi.getAll(familyId, true);
      console.log('获取激活奖品响应:', response);
      
      // 如果没有找到激活的奖品，尝试获取所有奖品
      if ((!response.success || !response.data || response.data.length === 0) && response.status === 404) {
        console.log('未找到激活的奖品，尝试获取所有奖品');
        response = await rewardApi.getAll(familyId, false);
      }
      console.log('获取奖品列表响应:', response);
      if (response.success) {
        // 确保即使返回的数据为null或undefined也将rewards设置为空数组
        console.log('获取到的奖品列表:', response.data);
        
        // 过滤掉库存为0的奖品
        const filteredRewards = (response.data || []).filter(reward => 
          reward.stock > 0 && reward.active
        );
        console.log('过滤后的奖品列表 (仅显示有库存的):', filteredRewards);
        
        setRewards(filteredRewards);
      } else {
        // 当出现404错误时，表示没有奖品记录，将rewards设置为空数组
        if (response.status === 404) {
          console.log('奖品列表为空');
          setRewards([]);
        } else {
          console.error('获取奖品列表返回错误:', response.error);
          setError(response.error || '获取奖品列表失败');
        }
      }
    } catch (err: any) {
      console.error('获取奖品列表出错:', err);
      // 当出现异常时，如果是404错误，表示没有奖品记录，将rewards设置为空数组
      if (err.response && err.response.status === 404) {
        console.log('奖品列表为空 (404)');
        setRewards([]);
      } else {
        setError('获取奖品列表时发生错误');
      }
    } finally {
      setLoading(false);
    }
  };

  // 创建新奖品
  const createReward = async (rewardData: {
    name: string;
    description: string;
    image_url: string;
    token_price: number;
    stock: number;
    contract_reward_id?: number;
  }): Promise<Reward | null> => {
    if (!selectedFamily || !selectedFamily.id) {
      setError('没有选择家庭，无法创建奖品');
      return null;
    }

    setLoading(true);
    setError(null);
    try {
      const response = await rewardApi.create(parseInt(selectedFamily.id), rewardData);
      if (response.success && response.data) {
        // 添加到列表
        setRewards(prev => [...prev, response.data as Reward]);
        return response.data;
      } else {
        setError(response.error || '创建奖品失败');
        return null;
      }
    } catch (err) {
      console.error('创建奖品出错:', err);
      setError('创建奖品时发生错误');
      return null;
    } finally {
      setLoading(false);
    }
  };

  // 更新奖品
  const updateReward = async (id: number, data: Partial<Reward>): Promise<boolean> => {
    setLoading(true);
    setError(null);
    try {
      const response = await rewardApi.update(id, data);
      if (response.success && response.data) {
        // 更新列表
        setRewards(prev => prev.map(reward => 
          reward.id === id ? { ...reward, ...response.data as Reward } : reward
        ));
        return true;
      } else {
        setError(response.error || '更新奖品失败');
        return false;
      }
    } catch (err) {
      console.error('更新奖品出错:', err);
      setError('更新奖品时发生错误');
      return false;
    } finally {
      setLoading(false);
    }
  };

  // 删除奖品
  const deleteReward = async (id: number): Promise<boolean> => {
    setLoading(true);
    setError(null);
    try {
      const response = await rewardApi.delete(id);
      if (response.success) {
        // 从列表中移除
        setRewards(prev => prev.filter(reward => reward.id !== id));
        return true;
      } else {
        setError(response.error || '删除奖品失败');
        return false;
      }
    } catch (err) {
      console.error('删除奖品出错:', err);
      setError('删除奖品时发生错误');
      return false;
    } finally {
      setLoading(false);
    }
  };

  // 获取兑换请求列表
  const fetchExchanges = async () => {
    if (!selectedFamily || !selectedFamily.id) {
      console.log('没有选择家庭，无法获取兑换请求');
      return;
    }

    setLoading(true);
    setError(null);
    try {
      const response = await exchangeApi.getByFamily(parseInt(selectedFamily.id));
      if (response.success) {
        // 确保即使返回的数据为null或undefined也将exchanges设置为空数组
        console.log('获取到家庭兑换记录:', response.data);
        setExchanges(response.data || []);
      } else {
        // 当出现404错误时，表示没有兑换记录，将exchanges设置为空数组
        if (response.status === 404) {
          console.log('兑换请求列表为空');
          setExchanges([]);
        } else {
          console.error('获取兑换请求列表返回错误:', response.error);
          setError(response.error || '获取兑换请求列表失败');
        }
      }
    } catch (err: any) {
      console.error('获取兑换请求列表出错:', err);
      // 当出现异常时，如果是404错误，表示没有兑换记录，将exchanges设置为空数组
      if (err.response && err.response.status === 404) {
        console.log('兑换请求列表为空 (404)');
        setExchanges([]);
      } else {
        setError('获取兑换请求列表时发生错误');
      }
    } finally {
      setLoading(false);
    }
  };

  // 批准兑换请求
  const approveExchange = async (id: number, notes?: string): Promise<boolean> => {
    setLoading(true);
    setError(null);
    try {
      const response = await exchangeApi.approve(id, notes);
      if (response.success && response.data) {
        // 更新列表
        setExchanges(prev => prev.map(exchange => 
          exchange.id === id ? { ...exchange, status: 'completed', notes: notes || exchange.notes } : exchange
        ));
        return true;
      } else {
        setError(response.error || '批准兑换请求失败');
        return false;
      }
    } catch (err) {
      console.error('批准兑换请求出错:', err);
      setError('批准兑换请求时发生错误');
      return false;
    } finally {
      setLoading(false);
    }
  };

  // 取消兑换请求
  const cancelExchange = async (id: number, notes?: string): Promise<boolean> => {
    setLoading(true);
    setError(null);
    try {
      const response = await exchangeApi.cancel(id, notes);
      if (response.success && response.data) {
        // 更新列表
        setExchanges(prev => prev.map(exchange => 
          exchange.id === id ? { ...exchange, status: 'cancelled', notes: notes || exchange.notes } : exchange
        ));
        return true;
      } else {
        setError(response.error || '取消兑换请求失败');
        return false;
      }
    } catch (err) {
      console.error('取消兑换请求出错:', err);
      setError('取消兑换请求时发生错误');
      return false;
    } finally {
      setLoading(false);
    }
  };

  // 当家庭变更时，获取奖品和兑换列表
  useEffect(() => {
    if (selectedFamily?.id && user) {
      fetchRewards();
      fetchExchanges();
    }
  }, [selectedFamily?.id, user]);

  const contextValue: RewardContextType = {
    rewards,
    exchanges,
    loading,
    error,
    fetchRewards,
    setRewards,
    createReward,
    updateReward,
    deleteReward,
    fetchExchanges,
    approveExchange,
    cancelExchange
  };

  return (
    <RewardContext.Provider value={contextValue}>
      {children}
    </RewardContext.Provider>
  );
};

// 自定义Hook，方便组件使用此上下文
export const useReward = () => {
  const context = React.useContext(RewardContext);
  if (context === undefined) {
    throw new Error('useReward must be used within a RewardProvider');
  }
  return context;
}; 