import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { PlusCircle, Package, Gift, RefreshCw, AlertCircle } from 'lucide-react';
import Button from '../components/common/Button';
import { useFamily } from '../context/FamilyContext';
import { useReward } from '../context/RewardContext';
import { useAuthContext } from '../context/AuthContext';
import RewardCard from '../components/reward/RewardCard';
import ExchangeCard from '../components/reward/ExchangeCard';
import AddRewardModal from '../components/reward/AddRewardModal';
import EditRewardModal from '../components/reward/EditRewardModal';
import { Reward, Exchange, RewardCreateRequest, RewardUpdateRequest } from '../types/reward';
import { createRewardWithContract } from '../services/tokenService';

const RewardManagement = () => {
  const navigate = useNavigate();
  const { user } = useAuthContext();
  const { selectedFamily } = useFamily();
  const { 
    rewards, 
    exchanges, 
    loading, 
    error, 
    fetchRewards, 
    createReward, 
    updateReward, 
    deleteReward, 
    fetchExchanges
  } = useReward();
  
  // 组件状态
  const [activeTab, setActiveTab] = useState('available'); // 'available' 或 'exchanged'
  const [filter, setFilter] = useState('all'); // 'all', 'active', 'inactive'
  const [isRefreshing, setIsRefreshing] = useState(false);
  const [createLoading, setCreateLoading] = useState(false); // 添加创建奖品的loading状态
  const [editLoading, setEditLoading] = useState(false); // 添加编辑奖品的loading状态
  
  // 模态框状态
  const [addModalOpen, setAddModalOpen] = useState(false);
  const [editModalOpen, setEditModalOpen] = useState(false);
  const [selectedReward, setSelectedReward] = useState<Reward | null>(null);
  
  // 处理刷新数据
  const handleRefresh = async () => {
    setIsRefreshing(true);
    if (activeTab === 'available') {
      console.log('手动刷新奖品数据');
      await fetchRewards();
    } else {
      console.log('手动刷新兑换记录数据');
      await fetchExchanges();
    }
    setTimeout(() => setIsRefreshing(false), 500); // 提供视觉反馈
  };
  
  // 处理添加奖品
  const handleAddReward = async (data: RewardCreateRequest) => {
    setCreateLoading(true);
    // 检查图片URL是否为placeholder.com的URL
    if (data.image_url && data.image_url.includes('placeholder.com')) {
      // 使用安全的Base64编码函数
      const safeBase64Encode = (str: string): string => {
        try {
          return btoa(encodeURIComponent(str).replace(/%([0-9A-F]{2})/g, (_, p1) => 
            String.fromCharCode(Number.parseInt(p1, 16))
          ));
        } catch (e) {
          console.error('编码失败:', e);
          return '';
        }
      };
      
      // 替换为内联SVG
      const svg = `<svg xmlns="http://www.w3.org/2000/svg" width="400" height="300" viewBox="0 0 400 300">
        <rect width="400" height="300" fill="#f0f0f0"/>
        <text x="200" y="150" font-family="Arial" font-size="24" text-anchor="middle" fill="#888888">${data.name || '奖品'}</text>
      </svg>`;
      
      data.image_url = `data:image/svg+xml;base64,${safeBase64Encode(svg)}`;
    }
    
    try {
      console.log('开始创建奖品，数据:', data);
      
      // 检查是否需要在区块链上创建奖品
      if (data.create_on_blockchain) {
        console.log('在区块链上创建奖品');
        
        if (!selectedFamily || !selectedFamily.id) {
          console.error('未选择家庭，无法在区块链上创建奖品');
          return;
        }
        
        // 调用合约创建奖品
        const familyId = parseInt(selectedFamily.id);
        const contractRewardId = await createRewardWithContract(
          familyId,
          data.name,
          data.description || '',
          data.image_url,
          data.token_price,
          data.stock || 1
        );
        
        if (contractRewardId > 0) {
          console.log('区块链上创建奖品成功，奖品ID:', contractRewardId);
          // 在后端创建奖品记录，并关联区块链ID
          const result = await createReward({
            ...data,
            contract_reward_id: contractRewardId
          });
          
          if (result) {
            console.log('创建奖品成功:', result);
            // 显示成功提示
            alert(`奖品 "${data.name}" 创建成功！`);
            // 刷新奖品列表
            fetchRewards();
            // 关闭模态框
            setAddModalOpen(false);
          }
        } else {
          console.error('区块链上创建奖品失败');
        }
      } else {
        // 仅在后端创建奖品
        const result = await createReward(data);
        if (result) {
          console.log('创建奖品成功:', result);
          // 显示成功提示
          alert(`奖品 "${data.name}" 创建成功！`);
          // 刷新奖品列表
          fetchRewards();
          // 关闭模态框
          setAddModalOpen(false);
        }
      }
    } catch (error) {
      console.error('创建奖品失败:', error);
      // 显示错误提示
      alert(`创建奖品失败: ${error instanceof Error ? error.message : '未知错误'}`);
    } finally {
      setCreateLoading(false);
    }
  };
  
  // 处理编辑奖品
  const handleEditReward = async (data: RewardUpdateRequest) => {
    setEditLoading(true);
    // 检查图片URL是否为placeholder.com的URL
    if (data.image_url && data.image_url.includes('placeholder.com')) {
      // 使用安全的Base64编码函数
      const safeBase64Encode = (str: string): string => {
        try {
          return btoa(encodeURIComponent(str).replace(/%([0-9A-F]{2})/g, (_, p1) => 
            String.fromCharCode(Number.parseInt(p1, 16))
          ));
        } catch (e) {
          console.error('编码失败:', e);
          return '';
        }
      };
      
      // 替换为内联SVG
      const svg = `<svg xmlns="http://www.w3.org/2000/svg" width="400" height="300" viewBox="0 0 400 300">
        <rect width="400" height="300" fill="#f0f0f0"/>
        <text x="200" y="150" font-family="Arial" font-size="24" text-anchor="middle" fill="#888888">${data.name || '奖品'}</text>
      </svg>`;
      
      data.image_url = `data:image/svg+xml;base64,${safeBase64Encode(svg)}`;
    }
    
    if (!selectedReward) {
      setEditLoading(false);
      return;
    }
    
    try {
      console.log('开始更新奖品，数据:', data);
      const result = await updateReward(selectedReward.id, data);
      if (result) {
        console.log('更新奖品成功');
        // 显示成功提示
        alert(`奖品 "${data.name}" 更新成功！`);
        // 刷新奖品列表
        fetchRewards();
        // 关闭模态框
        setEditModalOpen(false);
      }
    } catch (error) {
      console.error('更新奖品失败:', error);
      // 显示错误提示
      alert(`更新奖品失败: ${error instanceof Error ? error.message : '未知错误'}`);
    } finally {
      setEditLoading(false);
    }
  };
  
  // 处理删除奖品
  const handleDeleteReward = async (reward: Reward) => {
    if (window.confirm(`确定要删除奖品 "${reward.name}" 吗？`)) {
      try {
        const result = await deleteReward(reward.id);
        if (result) {
          alert(`奖品 "${reward.name}" 已成功删除！`);
        }
      } catch (error) {
        console.error('删除奖品失败:', error);
        alert(`删除奖品失败: ${error instanceof Error ? error.message : '未知错误'}`);
      }
    }
  };
  
  // 根据过滤条件筛选奖品
  const filteredRewards = rewards.filter(reward => {
    if (filter === 'all') return true;
    if (filter === 'active') return reward.active;
    if (filter === 'inactive') return !reward.active;
    return true;
  });
  
  // 根据状态过滤兑换请求
  const filteredExchanges = exchanges.filter(exchange => {
    if (filter === 'all') return true;
    if (filter === 'completed') return exchange.status === 'completed';
    if (filter === 'cancelled') return exchange.status === 'cancelled';
    return true;
  });
  
  // 监听标签页切换
  useEffect(() => {
    console.log('标签页切换:', activeTab);
    if (selectedFamily) {
      console.log('当前选择的家庭:', selectedFamily);
      if (activeTab === 'available') {
        console.log('获取奖品列表');
        fetchRewards();
      } else {
        console.log('获取兑换记录');
        fetchExchanges();
      }
    } else {
      console.log('未选择家庭，无法获取数据');
    }
  }, [activeTab]);
  
  // 监听家庭变更
  useEffect(() => {
    console.log('家庭变更:', selectedFamily);
    if (selectedFamily) {
      if (activeTab === 'available') {
        console.log('获取奖品列表');
        fetchRewards();
      } else {
        console.log('获取兑换记录');
        fetchExchanges();
      }
    }
  }, [selectedFamily]);
  
  // 组件加载时打印信息
  useEffect(() => {
    console.log('RewardManagement 组件加载');
    console.log('当前用户:', user);
    console.log('当前选择的家庭:', selectedFamily);
    console.log('当前认证令牌:', localStorage.getItem('auth_token'));
  }, []);

  return (
    <div className="max-w-6xl mx-auto">
      <div className="flex flex-col md:flex-row md:items-center md:justify-between mb-8">
        <div>
          <h1 className="text-3xl font-bold text-gray-900 mb-2">Reward Management</h1>
          <p className="text-gray-600 mb-4">
            Manage your rewards and exchanges here.
            {activeTab === 'exchanged' && exchanges.length === 0 && (
              <button 
                className="ml-2 text-primary-600 hover:underline"
                onClick={() => {
                  console.log('手动刷新兑换记录');
                  if (selectedFamily) {
                    console.log('当前选择的家庭:', selectedFamily);
                    fetchExchanges();
                  } else {
                    console.log('未选择家庭');
                  }
                }}
              >
                点击刷新兑换记录
              </button>
            )}
          </p>
        </div>
        
        <div className="mt-4 md:mt-0 flex flex-col sm:flex-row gap-3">
          <Button 
            onClick={handleRefresh}
            variant="secondary"
            leftIcon={<RefreshCw className={`h-5 w-5 ${isRefreshing ? 'animate-spin' : ''}`} />}
          >
            Refresh Data
          </Button>
          
          {activeTab === 'available' && (
            <Button 
              onClick={() => setAddModalOpen(true)}
              leftIcon={<PlusCircle className="h-5 w-5" />}
            >
              Add Reward
            </Button>
          )}
        </div>
      </div>
      
      {/* 标签页导航 */}
      <div className="mb-8">
        <div className="flex space-x-2 border-b border-gray-200">
          <button
            className={`px-4 py-2 text-sm font-medium ${
              activeTab === 'available'
                ? 'text-primary-600 border-b-2 border-primary-600'
                : 'text-gray-500 hover:text-gray-700'
            }`}
            onClick={() => setActiveTab('available')}
          >
            Available Rewards
          </button>
          <button
            className={`px-4 py-2 text-sm font-medium ${
              activeTab === 'exchanged'
                ? 'text-primary-600 border-b-2 border-primary-600'
                : 'text-gray-500 hover:text-gray-700'
            }`}
            onClick={() => setActiveTab('exchanged')}
          >
            My Exchange Records
          </button>
        </div>
      </div>
      
      {/* 内容区域 */}
      <div>
        {loading && (
          <div className="text-center py-12">
            <div className="inline-block animate-spin rounded-full h-8 w-8 border-2 border-gray-300 border-t-primary-600"></div>
            <p className="mt-2 text-gray-500">Loading...</p>
          </div>
        )}
        
        {error && (
          <div className="bg-red-50 border-l-4 border-red-400 p-4 mb-8 rounded-md">
            <div className="flex">
              <div className="flex-shrink-0">
                <AlertCircle className="h-5 w-5 text-red-400" />
              </div>
              <div className="ml-3">
                <p className="text-sm text-red-700">
                  Loading failed: {error}
                </p>
              </div>
            </div>
          </div>
        )}
        
        {!loading && !error && (
          <>
            {activeTab === 'available' ? (
              // 奖品列表
              <>
                {filteredRewards.length > 0 ? (
                  <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                    {filteredRewards.map((reward) => (
                      <RewardCard 
                        key={reward.id}
                        reward={reward}
                        onEdit={() => {
                          setSelectedReward(reward);
                          setEditModalOpen(true);
                        }}
                        onDelete={() => handleDeleteReward(reward)}
                      />
                    ))}
                  </div>
                ) : (
                  <div className="text-center py-12">
                    <Package className="mx-auto h-12 w-12 text-gray-400" />
                    <p className="mt-2 text-gray-500 text-lg">No rewards available</p>
                    {filter !== 'inactive' && (
                      <Button 
                        onClick={() => setAddModalOpen(true)}
                        className="mt-4"
                      >
                        Add First Reward
                      </Button>
                    )}
                  </div>
                )}
              </>
            ) : (
              // 兑换请求列表
              <>
                {filteredExchanges.length > 0 ? (
                  <div className="space-y-4">
                    {filteredExchanges.map((exchange) => (
                      <ExchangeCard 
                        key={exchange.id}
                        exchange={exchange}
                      />
                    ))}
                  </div>
                ) : (
                  <div className="text-center py-12">
                    <Package className="mx-auto h-12 w-12 text-gray-400" />
                    <p className="mt-2 text-gray-500 text-lg">No exchange requests</p>
                    <p className="mt-1 text-sm text-gray-500">
                      {selectedFamily ? 
                        `Family "${selectedFamily.name}" has no exchange records` :
                        'Please select a family to view exchange records'
                      }
                    </p>
                    <div className="mt-4 flex flex-col items-center gap-2">
                      <button 
                        className="text-primary-600 hover:underline"
                        onClick={() => {
                          console.log('手动刷新兑换记录');
                          if (selectedFamily) {
                            console.log('当前选择的家庭:', selectedFamily);
                            fetchExchanges();
                          } else {
                            console.log('未选择家庭');
                          }
                        }}
                      >
                        Click to Refresh Exchange Records
                      </button>
                    </div>
                  </div>
                )}
              </>
            )}
          </>
        )}
      </div>
      
      {/* 添加奖品模态框 */}
      <AddRewardModal 
        isOpen={addModalOpen}
        onClose={() => setAddModalOpen(false)}
        onSubmit={handleAddReward}
        isLoading={createLoading}
      />
      
      {/* 编辑奖品模态框 */}
      {selectedReward && (
        <EditRewardModal 
          isOpen={editModalOpen}
          onClose={() => {
            setEditModalOpen(false);
            setSelectedReward(null);
          }}
          onSubmit={handleEditReward}
          reward={selectedReward}
          isLoading={editLoading}
        />
      )}
    </div>
  );
};

export default RewardManagement; 