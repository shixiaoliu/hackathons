import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { useAccount } from 'wagmi';
import { Child, Family } from '../types/child';
import { familyApi, childApi, taskApi } from '../services/api';
import { useAuthContext } from './AuthContext';

interface FamilyContextType {
  family: Family | null;
  children: Child[];
  selectedChild: Child | null;
  addChild: (child: Omit<Child, 'id' | 'createdAt' | 'totalTasksCompleted' | 'totalRewardsEarned'>) => void;
  selectChild: (childId: string) => void;
  updateChild: (childId: string, updates: Partial<Child>) => void;
  removeChild: (childId: string) => void;
  isParent: boolean;
  currentChild: Child | null;
  loginAsChild: (walletAddress: string) => Child | null;
  getAllChildren: () => Child[];
  findChildByAddress: (address: string) => Child | null;
}

const FamilyContext = createContext<FamilyContextType | undefined>(undefined);

export const FamilyProvider = ({ children }: { children: ReactNode }) => {
  const { address } = useAccount();
  const { user, isAuthenticated } = useAuthContext();
  const [family, setFamily] = useState<Family | null>(null);
  const [selectedChild, setSelectedChild] = useState<Child | null>(null);
  const [isParent, setIsParent] = useState(false);
  const [currentChild, setCurrentChild] = useState<Child | null>(null);
  const [loading, setLoading] = useState(false);
  const [allChildren, setAllChildren] = useState<Child[]>([]);

  // 加载用户数据
  useEffect(() => {
    if (isAuthenticated && user) {
      setIsParent(user.role === 'parent');
      loadUserData();
    } else {
      // 清除状态
      setFamily(null);
      setCurrentChild(null);
      setSelectedChild(null);
      setIsParent(false);
      // 未登录时使用模拟数据
      setAllChildren(getAllChildrenFromStorage());
    }
  }, [isAuthenticated, user]);

  // 加载用户相关数据
  const loadUserData = async () => {
    console.log('[FamilyContext] loadUserData - user:', user);
    if (!user) return;
    
    setLoading(true);
    try {
      if (user.role === 'parent') {
        // 加载家庭数据
        const familyResponse = await familyApi.getAll();
        let currentFamily = null;
        
        if (familyResponse.success && familyResponse.data && familyResponse.data.length > 0) {
          currentFamily = familyResponse.data[0];
        } else {
          // 如果没有家庭，自动创建一个默认家庭
          console.log('没有找到家庭，正在创建默认家庭...');
          const createFamilyResponse = await familyApi.create('我的家庭');
          if (createFamilyResponse.success && createFamilyResponse.data) {
            console.log('默认家庭创建成功:', createFamilyResponse.data);
            currentFamily = createFamilyResponse.data;
          } else {
            console.error('创建默认家庭失败:', createFamilyResponse.error);
          }
        }
        
        // 加载children数据
        if (currentFamily) {
          console.log('[FamilyContext] 开始加载children数据...');
          const childrenResponse = await childApi.getAll();
          console.log('[FamilyContext] childrenResponse:', childrenResponse);
          
          if (childrenResponse.success && childrenResponse.data) {
            console.log('[FamilyContext] 原始children数据:', childrenResponse.data);
            // 将API返回的children数据转换为前端格式并关联到family
            const children = childrenResponse.data
              .filter(child => child.parent_address === user.wallet_address)
              .map(child => ({
                id: child.id.toString(),
                name: child.name,
                walletAddress: child.wallet_address,
                age: child.age,
                avatar: child.avatar || undefined,
                parentAddress: child.parent_address,
                createdAt: child.created_at,
                totalTasksCompleted: child.total_tasks_completed,
                totalRewardsEarned: child.total_rewards_earned,
              }));
            
            console.log('[FamilyContext] 过滤后的children数据:', children);
            
            setFamily({
              id: currentFamily.id.toString(),
              parentAddress: currentFamily.parent_address,
              createdAt: currentFamily.created_at,
              children
            });
            setAllChildren(children);
          } else {
            console.error('[FamilyContext] 加载children失败，使用模拟数据:', childrenResponse.error);
            // 如果加载children失败，使用模拟数据
            const mockChildren = getAllChildrenFromStorage();
            setFamily({
              id: currentFamily.id.toString(),
              parentAddress: currentFamily.parent_address,
              createdAt: currentFamily.created_at,
              children: mockChildren
            });
            setAllChildren(mockChildren);
          }
        }
      } else {
        // 加载儿童数据
        const childResponse = await childApi.getAll();
        if (childResponse.success && childResponse.data) {
          const child = childResponse.data.find(c => c.wallet_address.toLowerCase() === user.wallet_address.toLowerCase());
          if (child) {
            const childData = {
              id: child.id.toString(),
              name: child.name,
              walletAddress: child.wallet_address,
              age: child.age,
              avatar: child.avatar || undefined,
              parentAddress: child.parent_address,
              createdAt: child.created_at,
              totalTasksCompleted: child.total_tasks_completed,
              totalRewardsEarned: child.total_rewards_earned,
            };
            setCurrentChild(childData);
            // 为child用户设置所有children数据
            const allChildrenData = childResponse.data.map(c => ({
              id: c.id.toString(),
              name: c.name,
              walletAddress: c.wallet_address,
              age: c.age,
              avatar: c.avatar || undefined,
              parentAddress: c.parent_address,
              createdAt: c.created_at,
              totalTasksCompleted: c.total_tasks_completed,
              totalRewardsEarned: c.total_rewards_earned,
            }));
            setAllChildren(allChildrenData);
          } else {
            // 如果API中没有找到child，设置为空
            setCurrentChild(null);
            setAllChildren([]);
          }
        } else {
          // API调用失败，设置为空
          setCurrentChild(null);
          setAllChildren([]);
        }
      }
    } catch (error) {
      console.error('加载用户数据失败:', error);
      // 发生错误时设置为空
      setAllChildren([]);
      if (user && user.role === 'child') {
        setCurrentChild(null);
      }
    } finally {
      setLoading(false);
    }
  };

  const addChild = async (childData: Omit<Child, 'id' | 'createdAt' | 'totalTasksCompleted' | 'totalRewardsEarned'>): Promise<void> => {
    console.log('[FamilyContext] addChild - user:', user);
    if (!user || user.role !== 'parent') {
      console.error('只有父母角色可以添加子女');
      throw new Error('只有父母角色可以添加子女');
    }
    
    try {
      console.log('正在添加子女:', childData);
      console.log('当前用户:', user);
      console.log('当前token:', localStorage.getItem('auth_token'));
      
      const response = await childApi.create({
        name: childData.name,
        wallet_address: childData.walletAddress,
        age: childData.age,
        avatar: childData.avatar,
        parent_address: user.wallet_address,
        total_tasks_completed: 0,
        total_rewards_earned: '0',
      });
      
      console.log('API响应:', response);
      
      if (response.success && response.data) {
        const newChild: Child = {
          id: response.data.id.toString(),
          name: response.data.name,
          walletAddress: response.data.wallet_address,
          age: response.data.age,
          avatar: response.data.avatar || undefined,
          parentAddress: response.data.parent_address,
          createdAt: response.data.created_at,
          totalTasksCompleted: response.data.total_tasks_completed,
          totalRewardsEarned: response.data.total_rewards_earned,
        };
        
        // 重新加载family数据以确保children列表更新
        await loadUserData();
        
        console.log('子女添加成功:', newChild);
        alert('子女添加成功!');
      } else {
        console.error('添加子女失败:', response.error);
        const errorMessage = `添加子女失败: ${response.error || '未知错误'}`;
        alert(errorMessage);
        throw new Error(errorMessage);
      }
    } catch (error) {
      console.error('添加儿童失败:', error);
      const errorMessage = `添加子女失败: ${error instanceof Error ? error.message : '网络错误'}`;
      alert(errorMessage);
      throw error;
    }
  };

  const selectChild = (childId: string) => {
    const child = family?.children.find(c => c.id === childId);
    setSelectedChild(child || null);
  };

  const updateChild = async (childId: string, updates: Partial<Child>) => {
    try {
      const response = await childApi.update(parseInt(childId), {
        name: updates.name,
        age: updates.age,
        avatar: updates.avatar,
      });
      
      if (response.success) {
        setFamily(prev => prev ? {
          ...prev,
          children: prev.children.map(child => 
            child.id === childId ? { ...child, ...updates } : child
          )
        } : null);
      }
    } catch (error) {
      console.error('更新儿童信息失败:', error);
    }
  };

  const removeChild = async (childId: string) => {
    try {
      // 调用后端API删除child
      const response = await childApi.delete(parseInt(childId));
      if (response.success) {
        setFamily(prev => prev ? {
          ...prev,
          children: prev.children.filter(child => child.id !== childId)
        } : null);
        setAllChildren(prev => prev.filter(child => child.id !== childId));
        alert('Child deleted successfully!');
      } else {
        const errorMessage = response.error || 'Failed to delete child.';
        alert(errorMessage);
        throw new Error(errorMessage);
      }
    } catch (error) {
      console.error('删除child失败:', error);
      const errorMessage = error instanceof Error ? error.message : 'Network error';
      alert('Failed to delete child: ' + errorMessage);
      throw error;
    }
  };

  const loginAsChild = (walletAddress: string): Child | null => {
    // 从allChildren中查找，而不是只从family.children中查找
    const child = allChildren.find(child => child.walletAddress.toLowerCase() === walletAddress.toLowerCase()) || null;
    if (child) {
      setCurrentChild(child);
      setIsParent(false);
      // 不要清空family，保持children数据可用
      return child;
    }
    return null;
  };

  const getAllChildren = (): Child[] => {
    // 优先返回allChildren，如果为空则返回模拟数据
    return allChildren.length > 0 ? allChildren : getAllChildrenFromStorage();
  };

  const findChildByAddress = (address: string): Child | null => {
    return getAllChildren().find(child => child.walletAddress.toLowerCase() === address.toLowerCase()) || null;
  };

  return (
    <FamilyContext.Provider value={{
      family,
      children: isParent ? (family?.children || []) : allChildren,
      selectedChild,
      addChild,
      selectChild,
      updateChild,
      removeChild,
      isParent,
      currentChild,
      loginAsChild,
      getAllChildren,
      findChildByAddress
    }}>
      {children}
    </FamilyContext.Provider>
  );
};

// 模拟数据加载函数
function loadFamilyData(address: string): Family | null {
  // 这个函数现在需要通过context来访问真实数据
  // 暂时返回null，实际使用时应该通过context获取
  return null;
}

function createNewFamily(address: string): Family {
  return {
    id: Date.now().toString(),
    parentAddress: address,
    children: [],
    createdAt: new Date().toISOString()
  };
}

function getAllChildrenFromStorage(): Child[] {
  // 不再使用模拟数据，返回空数组
  // 所有数据都应该从API获取
  return [];
}

export const useFamily = (): FamilyContextType => {
  const context = useContext(FamilyContext);
  if (context === undefined) {
    throw new Error('useFamily must be used within a FamilyProvider');
  }
  return context;
};