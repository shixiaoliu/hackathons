import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { PlusCircle, Package, Gift, RefreshCw, Filter, AlertCircle } from 'lucide-react';
import Button from '../components/common/Button';
import { useFamily } from '../context/FamilyContext';
import { useReward } from '../context/RewardContext';
import { useAuthContext } from '../context/AuthContext';
import RewardCard from '../components/reward/RewardCard';
import ExchangeCard from '../components/reward/ExchangeCard';
import AddRewardModal from '../components/reward/AddRewardModal';
import EditRewardModal from '../components/reward/EditRewardModal';
import ExchangeActionModal from '../components/reward/ExchangeActionModal';
import { Reward, Exchange, RewardCreateRequest, RewardUpdateRequest } from '../types/reward';

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
    fetchExchanges,
    approveExchange,
    cancelExchange
  } = useReward();
  
  // 组件状态
  const [activeTab, setActiveTab] = useState('rewards'); // 'rewards' 或 'exchanges'
  const [filter, setFilter] = useState('all'); // 'all', 'active', 'inactive'
  const [isRefreshing, setIsRefreshing] = useState(false);
  
  // 模态框状态
  const [addModalOpen, setAddModalOpen] = useState(false);
  const [editModalOpen, setEditModalOpen] = useState(false);
  const [selectedReward, setSelectedReward] = useState<Reward | null>(null);
  const [exchangeActionModal, setExchangeActionModal] = useState<{
    open: boolean;
    exchange: Exchange | null;
    action: 'approve' | 'cancel';
  }>({
    open: false,
    exchange: null,
    action: 'approve'
  });
  
  // 处理刷新数据
  const handleRefresh = async () => {
    setIsRefreshing(true);
    if (activeTab === 'rewards') {
      await fetchRewards();
    } else {
      await fetchExchanges();
    }
    setTimeout(() => setIsRefreshing(false), 500); // 提供视觉反馈
  };
  
  // 处理添加奖品
  const handleAddReward = async (data: RewardCreateRequest) => {
    const result = await createReward(data);
    if (result) {
      setAddModalOpen(false);
    }
  };
  
  // 处理编辑奖品
  const handleEditReward = async (data: RewardUpdateRequest) => {
    if (!selectedReward) return;
    
    const success = await updateReward(selectedReward.id, data);
    if (success) {
      setEditModalOpen(false);
      setSelectedReward(null);
    }
  };
  
  // 处理删除奖品
  const handleDeleteReward = async (reward: Reward) => {
    if (window.confirm(`确定要删除奖品 "${reward.name}" 吗？`)) {
      await deleteReward(reward.id);
    }
  };
  
  // 处理批准兑换
  const handleApproveExchange = (exchange: Exchange) => {
    setExchangeActionModal({
      open: true,
      exchange,
      action: 'approve'
    });
  };
  
  // 处理拒绝兑换
  const handleCancelExchange = (exchange: Exchange) => {
    setExchangeActionModal({
      open: true,
      exchange,
      action: 'cancel'
    });
  };
  
  // 确认批准或拒绝兑换
  const confirmExchangeAction = async (notes?: string) => {
    if (!exchangeActionModal.exchange) return;
    
    const exchangeId = exchangeActionModal.exchange.id;
    let success = false;
    
    if (exchangeActionModal.action === 'approve') {
      success = await approveExchange(exchangeId, notes);
    } else {
      success = await cancelExchange(exchangeId, notes);
    }
    
    if (success) {
      setExchangeActionModal({
        open: false,
        exchange: null,
        action: 'approve'
      });
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
    if (filter === 'pending') return exchange.status === 'pending';
    if (filter === 'completed') return exchange.status === 'completed';
    if (filter === 'cancelled') return exchange.status === 'cancelled';
    return true;
  });
  
  // 统计待处理的兑换请求数量
  const pendingExchangesCount = exchanges.filter(e => e.status === 'pending').length;
  
  // 监听家庭变更
  useEffect(() => {
    if (selectedFamily) {
      fetchRewards();
      fetchExchanges();
    }
  }, [selectedFamily]);

  return (
    <div className="max-w-6xl mx-auto">
      <div className="flex flex-col md:flex-row md:items-center md:justify-between mb-8">
        <div>
          <h1 className="text-3xl font-bold text-gray-900 mb-2">Reward Management</h1>
          <p className="text-gray-600 mb-4">
            Manage your rewards and exchanges here.
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
          
          {activeTab === 'rewards' && (
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
      
      {/* 筛选器 */}
      <div className="mb-8">
        <div className="flex items-center space-x-2">
          <Filter className="h-4 w-4 text-gray-500" />
          <span className="text-sm font-medium text-gray-700">Filter:</span>
          
          {/* 根据当前标签页显示不同的筛选选项 */}
          {activeTab === 'rewards' ? (
            <div className="flex space-x-2">
              <button
                className={`px-3 py-1 text-sm rounded-full ${
                  filter === 'all'
                    ? 'bg-primary-100 text-primary-800'
                    : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
                }`}
                onClick={() => setFilter('all')}
              >
                All
              </button>
              <button
                className={`px-3 py-1 text-sm rounded-full ${
                  filter === 'active'
                    ? 'bg-primary-100 text-primary-800'
                    : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
                }`}
                onClick={() => setFilter('active')}
              >
                Active
              </button>
              <button
                className={`px-3 py-1 text-sm rounded-full ${
                  filter === 'inactive'
                    ? 'bg-primary-100 text-primary-800'
                    : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
                }`}
                onClick={() => setFilter('inactive')}
              >
                Inactive
              </button>
            </div>
          ) : (
            <div className="flex space-x-2">
              <button
                className={`px-3 py-1 text-sm rounded-full ${
                  filter === 'all'
                    ? 'bg-primary-100 text-primary-800'
                    : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
                }`}
                onClick={() => setFilter('all')}
              >
                All
              </button>
              <button
                className={`px-3 py-1 text-sm rounded-full ${
                  filter === 'pending'
                    ? 'bg-primary-100 text-primary-800'
                    : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
                }`}
                onClick={() => setFilter('pending')}
              >
                Pending
              </button>
              <button
                className={`px-3 py-1 text-sm rounded-full ${
                  filter === 'completed'
                    ? 'bg-primary-100 text-primary-800'
                    : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
                }`}
                onClick={() => setFilter('completed')}
              >
                Completed
              </button>
              <button
                className={`px-3 py-1 text-sm rounded-full ${
                  filter === 'cancelled'
                    ? 'bg-primary-100 text-primary-800'
                    : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
                }`}
                onClick={() => setFilter('cancelled')}
              >
                Cancelled
              </button>
            </div>
          )}
        </div>
      </div>
      
      {/* 内容区域 */}
      <div>
        {loading && (
          <div className="text-center py-12">
            <div className="inline-block animate-spin rounded-full h-8 w-8 border-2 border-gray-300 border-t-primary-600"></div>
            <p className="mt-2 text-gray-500">加载中...</p>
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
                  加载失败: {error}
                </p>
              </div>
            </div>
          </div>
        )}
        
        {!loading && !error && (
          <>
            {activeTab === 'rewards' ? (
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
                        添加第一个奖品
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
                        onApprove={() => handleApproveExchange(exchange)}
                        onCancel={() => handleCancelExchange(exchange)}
                      />
                    ))}
                  </div>
                ) : (
                  <div className="text-center py-12">
                    <Package className="mx-auto h-12 w-12 text-gray-400" />
                    <p className="mt-2 text-gray-500 text-lg">No exchange requests</p>
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
        isLoading={loading}
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
          isLoading={loading}
        />
      )}
      
      {/* 兑换操作模态框 */}
      {exchangeActionModal.exchange && (
        <ExchangeActionModal 
          isOpen={exchangeActionModal.open}
          onClose={() => setExchangeActionModal({
            open: false,
            exchange: null,
            action: 'approve'
          })}
          onConfirm={confirmExchangeAction}
          exchange={exchangeActionModal.exchange}
          actionType={exchangeActionModal.action}
          isLoading={loading}
        />
      )}
    </div>
  );
};

export default RewardManagement; 