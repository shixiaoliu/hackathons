import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { ShoppingBag, Check, History, Wallet, BarChart4, RefreshCw } from 'lucide-react';
import Card, { CardBody, CardHeader } from '../components/common/Card';
import Button from '../components/common/Button';
import { useReward } from '../context/RewardContext';
import { useFamily } from '../context/FamilyContext';
import { useAuthContext } from '../context/AuthContext';
import { exchangeApi, contractApi, Reward, Exchange, apiClient } from '../services/api';
import ExchangeCard from '../components/reward/ExchangeCard';
import { getTokenBalance, updateLocalBalance, clearBalanceCache, exchangeRewardWithContract, burnTokens } from '../services/tokenService';
import { ethers } from 'ethers';

const RewardStore = () => {
  const navigate = useNavigate();
  const { rewards, loading, error, fetchRewards, setRewards } = useReward();
  const { currentChild, selectedFamily } = useFamily();
  const { user } = useAuthContext();
  const [activeTab, setActiveTab] = useState<'available' | 'exchanged'>('available');
  const [exchanges, setExchanges] = useState<Exchange[]>([]);
  const [balance, setBalance] = useState<string>('0');
  const [loadingExchanges, setLoadingExchanges] = useState<boolean>(false);
  const [loadingBalance, setLoadingBalance] = useState<boolean>(false);
  const [exchangeInProgress, setExchangeInProgress] = useState<boolean>(false);
  
  // 获取代币余额 - 使用区块链直接查询
  const fetchBalance = async () => {
    if (!currentChild?.walletAddress) return;
    
    setLoadingBalance(true);
    try {
      // 强制清除缓存，确保获取最新数据
      clearBalanceCache(currentChild.walletAddress);
      
      // 使用tokenService获取余额，它会处理缓存逻辑
      console.log('[RewardStore] 获取代币余额...');
      const balance = await getTokenBalance(currentChild.walletAddress);
      
      if (balance !== '0') {
        console.log('[RewardStore] 成功获取余额:', balance);
        setBalance(balance);
      } else {
        console.log('[RewardStore] 获取余额为0，尝试通过API获取');
        // 尝试通过API获取余额
        try {
          const response = await contractApi.getBalance(currentChild.walletAddress);
          if (response.success && response.data && response.data.balance) {
            console.log('[RewardStore] API获取余额成功:', response.data.balance);
            setBalance(response.data.balance);
          } else {
            console.log('[RewardStore] API获取余额失败或为0');
          }
        } catch (apiError) {
          console.error('[RewardStore] API获取余额失败:', apiError);
        }
      }
    } catch (error) {
      console.error('[RewardStore] 获取余额失败:', error);
    } finally {
      setLoadingBalance(false);
    }
  };
  
  // 刷新代币余额
  const refreshBalance = async () => {
    if (!currentChild?.walletAddress) return;
    
    setLoadingBalance(true);
    try {
      // 清除缓存，强制获取最新数据
      clearBalanceCache(currentChild.walletAddress);
      
      console.log('[RewardStore] 强制刷新余额...');
      const blockchainBalance = await getTokenBalance(currentChild.walletAddress);
      console.log('[RewardStore] 刷新余额结果:', blockchainBalance);
      
      // 确保余额不为0
      if (blockchainBalance === '0') {
        console.log('[RewardStore] 获取到的余额为0，尝试再次获取...');
        // 尝试直接通过API获取
        const response = await contractApi.getBalance(currentChild.walletAddress);
        if (response.success && response.data && response.data.balance !== '0') {
          console.log('[RewardStore] API获取余额成功:', response.data.balance);
          setBalance(response.data.balance);
          return;
        }
      }
      
      setBalance(blockchainBalance);
    } catch (error) {
      console.error('[RewardStore] 刷新余额失败:', error);
      
      // 如果刷新失败，尝试直接通过API获取
      try {
        console.log('[RewardStore] 尝试通过API获取余额...');
        const response = await contractApi.getBalance(currentChild.walletAddress);
        if (response.success && response.data) {
          console.log('[RewardStore] API获取余额成功:', response.data.balance);
          setBalance(response.data.balance);
        }
      } catch (apiError) {
        console.error('[RewardStore] API获取余额也失败:', apiError);
      }
    } finally {
      setLoadingBalance(false);
    }
  };
  
  // 获取兑换记录
  const fetchExchanges = async () => {
    if (!currentChild) {
      console.log('没有当前孩子信息，无法获取兑换记录');
      return;
    }
    
    setLoadingExchanges(true);
    try {
      console.log('开始获取兑换记录，当前孩子ID:', currentChild.id);
      
      // 打印认证和用户信息，用于调试
      const token = localStorage.getItem('auth_token');
      const userData = localStorage.getItem('user_data');
      const userRole = localStorage.getItem('user_role');
      console.log('认证信息:', {
        token: token ? '已设置' : '未设置',
        userData: userData ? JSON.parse(userData) : null,
        userRole
      });
      
      // 添加延迟，确保后端有足够时间处理之前的请求
      await new Promise(resolve => setTimeout(resolve, 500));
      
      const response = await exchangeApi.getByChild();
      console.log('获取兑换记录响应:', JSON.stringify(response));
      
      if (response.success && Array.isArray(response.data)) {
        console.log('成功获取兑换记录，数量:', response.data.length);
        
        // 打印每条记录的详细信息
        response.data.forEach((exchange, index) => {
          console.log(`兑换记录 ${index + 1}:`, {
            id: exchange.id,
            reward_id: exchange.reward_id,
            child_id: exchange.child_id,
            reward_name: exchange.reward_name,
            status: exchange.status,
            date: exchange.exchange_date
          });
        });
        
        // 处理后端返回的数据，确保status字段正确映射
        const processedExchanges = response.data.map(exchange => ({
          ...exchange,
          // 确保status字段符合前端期望的值，将confirmed映射为completed
          status: exchange.status === 'confirmed' ? 'completed' : exchange.status
        }));
        
        setExchanges(processedExchanges);
      } else {
        console.error('获取兑换记录失败:', response.error || '未知错误', '状态码:', response.status);
        
        // 如果是404错误或没有数据，设置为空数组
        if (response.status === 404 || !response.data) {
          console.log('没有兑换记录或API返回404，设置为空数组');
          setExchanges([]);
        }
      }
    } catch (error) {
      console.error('获取兑换记录出错:', error);
      // 确保即使出错也设置为空数组
      setExchanges([]);
    } finally {
      setLoadingExchanges(false);
    }
  };
  
  // 兑换奖品
  const handleExchange = async (reward: Reward) => {
    if (!currentChild) {
      alert('请先登录');
      return;
    }
    
    if (parseFloat(balance) < reward.token_price) {
      alert('代币余额不足');
      return;
    }
    
    setExchangeInProgress(true);
    try {
      // 记录兑换请求详情，便于调试
      console.log('兑换请求数据:', {
        reward_id: reward.id,
        childInfo: {
          name: currentChild.name,
          walletAddress: currentChild.walletAddress,
        }
      });

      // 先尝试通过RewardRegistry合约兑换奖品
      let exchangeSuccess = false;
      try {
        console.log('尝试通过RewardRegistry合约兑换奖品...');
        if (!reward.contract_reward_id) {
          console.warn('奖品没有关联的区块链ID，无法使用合约兑换:', reward.id);
          throw new Error('奖品没有关联的区块链ID');
        }
        
        console.log('使用区块链奖品ID进行兑换:', {
          app_reward_id: reward.id,
          blockchain_reward_id: reward.contract_reward_id
        });
        
        exchangeSuccess = await exchangeRewardWithContract(reward.contract_reward_id);
        console.log('合约兑换结果:', exchangeSuccess ? '成功' : '失败');
        
        // 如果合约兑换成功，立即更新余额显示
        if (exchangeSuccess) {
          // 清除余额缓存
          clearBalanceCache(currentChild.walletAddress);
          
          // 获取最新余额
          const newBalance = await getTokenBalance(currentChild.walletAddress);
          console.log('合约兑换成功，更新余额:', balance, '->', newBalance);
          
          // 更新本地状态
          setBalance(newBalance);
        }
      } catch (contractError) {
        console.error('通过RewardRegistry合约兑换奖品失败:', contractError);
      }

      // 如果合约兑换失败，使用API方式兑换
      if (!exchangeSuccess) {
        // 调用后端API处理兑换
        const response = await exchangeApi.create({
          reward_id: reward.id,
          notes: `由${currentChild.name}兑换`,
          token_burned: true // 告诉后端代币未在前端销毁，需要后端处理
        });
        
        console.log('API兑换响应:', response);
        
        if (response.success) {
          alert('兑换申请已提交');
          
          // 计算新余额
          const newBalance = (parseFloat(balance) - reward.token_price).toString();
          console.log('更新余额:', balance, '->', newBalance);
          
          // 更新本地状态
          setBalance(newBalance);
          
          // 更新本地缓存的余额，确保刷新页面后余额仍然正确
          if (currentChild.walletAddress) {
            updateLocalBalance(currentChild.walletAddress, newBalance);
          }
          
          // 立即从本地状态中移除已兑换的奖品
          console.log('从本地移除奖品:', reward.id);
          const updatedRewards = rewards.filter(r => r.id !== reward.id);
          console.log('更新后的奖品列表:', updatedRewards);
          
          // 直接更新本地奖品列表状态
          setRewards(updatedRewards);
          
          // 添加到兑换记录
          if (response.data) {
            console.log('添加到兑换记录:', response.data);
            // 确保response.data是一个有效的Exchange对象
            const newExchange = response.data as Exchange;
            
            // 添加额外的字段，确保显示正确
            const enhancedExchange = {
              ...newExchange,
              reward_name: reward.name,
              reward_image: reward.image_url,
              child_name: currentChild.name,
              token_amount: reward.token_price,
              token_burned: false // 记录代币是否已经被销毁
            };
            
            console.log('增强后的兑换记录:', enhancedExchange);
            
            // 添加到兑换记录列表的开头
            setExchanges(prev => [enhancedExchange, ...prev]);
            
            // 切换到兑换记录标签
            setActiveTab('exchanged');
          }
        } else {
          alert(`兑换失败: ${response.error || '未知错误'}`);
        }
      } else {
        // 如果合约兑换成功，更新UI
        alert('兑换成功！');
        
        // 立即从本地状态中移除已兑换的奖品
        console.log('从本地移除奖品:', reward.id);
        const updatedRewards = rewards.filter(r => r.id !== reward.id);
        setRewards(updatedRewards);
        
        // 刷新兑换记录
        fetchExchanges();
        
        // 切换到兑换记录标签
        setActiveTab('exchanged');
      }
    } catch (error) {
      console.error('兑换过程出错:', error);
      alert(`兑换过程出错: ${error instanceof Error ? error.message : '未知错误'}`);
    } finally {
      setExchangeInProgress(false);
    }
  };
  
  // 当切换到兑换记录标签时，刷新数据
  useEffect(() => {
    if (activeTab === 'exchanged' && currentChild) {
      console.log('切换到兑换记录标签，刷新数据');
      fetchExchanges();
    }
  }, [activeTab]);
  
  // 加载页面数据
  useEffect(() => {
    if (currentChild) {
      console.log('当前孩子信息:', currentChild);
      // 将当前孩子信息存储到localStorage，以便API调用时可以获取
      localStorage.setItem('current_child', JSON.stringify({
        id: currentChild.id,
        name: currentChild.name,
        walletAddress: currentChild.walletAddress
      }));
      
      // 不要在useEffect内部调用hooks
      console.log('当前选择的家庭:', selectedFamily);
      
      // 打印认证信息，用于调试
      const token = localStorage.getItem('auth_token');
      const userData = localStorage.getItem('user_data');
      const userRole = localStorage.getItem('user_role');
      console.log('认证信息:', {
        token: token ? '已设置' : '未设置',
        userData: userData ? JSON.parse(userData) : null,
        userRole
      });
      
      // 立即获取数据
      fetchRewards();
      // 使用普通的fetchBalance而不是强制刷新
      fetchBalance();
      fetchExchanges();
    }
  }, [currentChild, selectedFamily]);
  
  // 如果没有孩子信息，显示提示
  if (!currentChild) {
    return (
      <div className="text-center py-12">
        <div className="bg-yellow-50 border border-yellow-200 rounded-lg p-6 max-w-lg mx-auto">
          <ShoppingBag className="h-12 w-12 text-yellow-600 mx-auto mb-4" />
          <h2 className="text-xl font-semibold text-gray-900 mb-2">Login Required</h2>
          <p className="text-gray-600 mb-4">
            Please log in with your child account to access the reward store.
          </p>
          <Button
            onClick={() => navigate('/')}
            variant="primary"
          >
            返回首页
          </Button>
        </div>
      </div>
    );
  }
  
  // 可兑换的奖品 - 放宽条件，显示所有奖品
  console.log('所有奖品数据:', rewards);
  
  // 仅使用真实数据
  const availableRewards = rewards;
  
  // 已兑换的奖品
  const exchangedRewards = exchanges.filter(exchange => exchange.status !== 'cancelled');
  
  // 统计数据
  const pendingExchanges = exchanges.filter(exchange => exchange.status === 'pending').length;
  const completedExchanges = exchanges.filter(exchange => exchange.status === 'completed').length;
  const totalSpent = exchanges
    .filter(exchange => exchange.status !== 'cancelled')
    .reduce((total, exchange) => total + exchange.token_amount, 0);
  
  // 生成默认图片的数据URL
  const generateDefaultImageDataUrl = (text: string): string => {
    try {
      const svg = `<svg xmlns="http://www.w3.org/2000/svg" width="400" height="300" viewBox="0 0 400 300">
        <rect width="400" height="300" fill="#f0f0f0"/>
        <text x="200" y="150" font-family="Arial" font-size="24" text-anchor="middle" fill="#888888">${text || '奖品'}</text>
      </svg>`;
      return `data:image/svg+xml;base64,${btoa(svg)}`;
    } catch (error) {
      console.error('生成默认图片失败:', error);
      // 提供一个极简的备用数据URL
      return 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg==';
    }
  };

  // 处理图片加载错误
  const handleImageError = (reward: Reward) => {
    console.log('图片加载失败:', reward.image_url);
    // 创建一个新的奖品对象，替换图片URL为默认图片
    const updatedReward = {
      ...reward,
      image_url: generateDefaultImageDataUrl(reward.name)
    };
    
    // 更新奖品列表中的对应项
    setRewards(prevRewards => 
      prevRewards.map(r => r.id === reward.id ? updatedReward : r)
    );
  };
  
  return (
    <div className="max-w-6xl mx-auto">
      <div className="mb-8 flex justify-between items-center">
        <div>
          <h1 className="text-3xl font-bold text-gray-900 mb-2">Reward Store</h1>
          <p className="text-gray-600">Use your tokens to redeem rewards</p>
        </div>
        <Button 
          variant="secondary"
          onClick={() => {
            console.log('手动刷新奖品列表');
            fetchRewards();
          }}
          className="flex items-center"
        >
          <RefreshCw className="h-4 w-4 mr-2" />
          Refresh Rewards
        </Button>
      </div>
      
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
        <Card className="flex-1">
          <CardBody className="flex items-center">
            <div className="flex-shrink-0 mr-4">
              <div className="w-12 h-12 rounded-full bg-green-100 flex items-center justify-center">
                <Wallet className="h-6 w-6 text-green-600" />
              </div>
            </div>
            <div className="flex-grow">
              <p className="text-sm text-gray-500">FCT Balance</p>
              <div className="flex items-center">
                <p className="text-2xl font-bold text-gray-900">
                  {loadingBalance ? (
                    <span className="animate-pulse">Loading...</span>
                  ) : (
                    `${Math.floor(parseFloat(balance) || 0)} FCT`
                  )}
                </p>
                <button 
                  onClick={async () => {
                    if (!currentChild?.walletAddress) return;
                    
                    setLoadingBalance(true);
                    try {
                      // 强制清除缓存
                      clearBalanceCache();
                      
                      // 尝试获取最新余额
                      console.log('[RewardStore] 强制刷新余额...');
                      const blockchainBalance = await getTokenBalance(currentChild.walletAddress);
                      
                      if (blockchainBalance && parseFloat(blockchainBalance) > 0) {
                        console.log('[RewardStore] 刷新余额成功:', blockchainBalance);
                        setBalance(blockchainBalance);
                      } else {
                        // 如果获取失败，尝试通过API获取
                        console.log('[RewardStore] 通过区块链获取余额失败，尝试API...');
                        const response = await contractApi.getBalance(currentChild.walletAddress);
                        if (response.success && response.data) {
                          console.log('[RewardStore] API获取余额成功:', response.data.balance);
                          setBalance(response.data.balance);
                        } else {
                          console.log('[RewardStore] 所有获取余额方式均失败');
                        }
                      }
                    } catch (error) {
                      console.error('[RewardStore] 刷新余额失败:', error);
                    } finally {
                      setLoadingBalance(false);
                    }
                  }}
                  className="ml-2 text-blue-500 hover:text-blue-700"
                  disabled={loadingBalance}
                  title="刷新余额"
                >
                  <RefreshCw className={`h-4 w-4 ${loadingBalance ? 'animate-spin' : ''}`} />
                </button>
              </div>
            </div>
          </CardBody>
        </Card>
        
        <Card>
          <CardBody className="flex items-center">
            <div className="flex-shrink-0 mr-4">
              <div className="w-12 h-12 rounded-full bg-blue-100 flex items-center justify-center">
                <ShoppingBag className="h-6 w-6 text-blue-600" />
              </div>
            </div>
            <div>
              <p className="text-sm text-gray-500">Exchanged Rewards</p>
              <p className="text-2xl font-bold text-gray-900">{exchangedRewards.length}</p>
            </div>
          </CardBody>
        </Card>
        
        <Card>
          <CardBody className="flex items-center">
            <div className="flex-shrink-0 mr-4">
              <div className="w-12 h-12 rounded-full bg-yellow-100 flex items-center justify-center">
                <History className="h-6 w-6 text-yellow-600" />
              </div>
            </div>
            <div>
              <p className="text-sm text-gray-500">Pending Requests</p>
              <p className="text-2xl font-bold text-gray-900">{pendingExchanges}</p>
            </div>
          </CardBody>
        </Card>
      </div>
      
      <div className="mb-6">
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
      
      {activeTab === 'available' && (
        <>
          {loading ? (
            <div className="flex justify-center items-center py-12">
              <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary-600"></div>
            </div>
          ) : availableRewards.length === 0 ? (
            <div className="text-center py-12 bg-gray-50 rounded-lg">
              <ShoppingBag className="h-12 w-12 text-gray-400 mx-auto mb-4" />
              <h3 className="text-lg font-medium text-gray-900 mb-2">No rewards available</h3>
              <p className="text-gray-600">Currently, there are no rewards available for redemption. Please check back later.</p>
            </div>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {availableRewards.map((reward) => (
                <div key={reward.id} className="w-full">
                  <Card 
                    className="h-full cursor-pointer hover:shadow-lg transition-shadow duration-200"
                    onClick={() => handleExchange(reward)}
                  >
                    {/* 奖品图片 */}
                    <div className="relative w-full h-48 bg-gray-100 overflow-hidden flex items-center justify-center">
                      <img 
                        src={reward.image_url || generateDefaultImageDataUrl(reward.name)} 
                        alt={reward.name} 
                        className="w-full h-full object-contain"
                        onError={() => handleImageError(reward)}
                        loading="lazy"
                      />
                      <div className="absolute top-0 left-0 bg-blue-500 text-white px-2 py-1 text-xs font-bold">
                        Limited: 1
                      </div>
                    </div>
                    
                    <CardBody>
                      <div className="mb-2 flex justify-between items-start">
                        <h3 className="text-lg font-semibold text-gray-900">{reward.name}</h3>
                        <div className="px-2 py-1 bg-primary-100 text-primary-800 text-sm font-medium rounded-md">
                          {reward.token_price} FCT
                        </div>
                      </div>
                      <p className="text-gray-600 text-sm mb-4">
                        {reward.description || '暂无描述'}
                      </p>
                      <div className="mt-auto pt-2 border-t border-gray-100 flex justify-between items-center">
                        <span className="text-sm text-gray-500">Limited: 1</span>
                        <Button 
                          variant="primary" 
                          size="sm"
                          disabled={parseFloat(balance) < reward.token_price || exchangeInProgress}
                        >
                          {exchangeInProgress ? '处理中...' : '兑换奖品'}
                        </Button>
                      </div>
                    </CardBody>
                  </Card>
                </div>
              ))}
            </div>
          )}
        </>
      )}
      
      {activeTab === 'exchanged' && (
        <>
          {loadingExchanges ? (
            <div className="flex justify-center items-center py-12">
              <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary-600"></div>
            </div>
          ) : exchangedRewards.length === 0 ? (
            <div className="text-center py-12 bg-gray-50 rounded-lg">
              <History className="h-12 w-12 text-gray-400 mx-auto mb-4" />
              <h3 className="text-lg font-medium text-gray-900 mb-2">No exchange requests</h3>
              <p className="text-gray-600">You have not exchanged any rewards yet.</p>
            </div>
          ) : (
            <>
              <div className="mb-6">
                <Card>
                  <CardHeader>
                    <h3 className="font-semibold text-lg flex items-center">
                      <BarChart4 className="h-5 w-5 mr-2 text-primary-600" />
                      交易统计
                    </h3>
                  </CardHeader>
                  <CardBody>
                    <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                      <div className="border rounded-lg p-4 bg-gray-50">
                        <p className="text-sm text-gray-500">总兑换次数</p>
                        <p className="text-2xl font-bold text-gray-900">{exchangedRewards.length}</p>
                      </div>
                      <div className="border rounded-lg p-4 bg-gray-50">
                        <p className="text-sm text-gray-500">已完成兑换</p>
                        <p className="text-2xl font-bold text-green-600">{completedExchanges}</p>
                      </div>
                      <div className="border rounded-lg p-4 bg-gray-50">
                        <p className="text-sm text-gray-500">总消费代币</p>
                        <p className="text-2xl font-bold text-primary-600">{Math.floor(totalSpent)} FCT</p>
                      </div>
                    </div>
                  </CardBody>
                </Card>
              </div>
              
              <div className="space-y-4">
                {exchangedRewards.map(exchange => (
                  <ExchangeCard
                    key={exchange.id}
                    exchange={exchange}
                    isChild={true}
                  />
                ))}
              </div>
            </>
          )}
        </>
      )}
    </div>
  );
};

export default RewardStore; 