import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { ShoppingBag, Check, History, Wallet, BarChart4, RefreshCw } from 'lucide-react';
import Card, { CardBody, CardHeader } from '../components/common/Card';
import Button from '../components/common/Button';
import { useReward } from '../context/RewardContext';
import { useFamily } from '../context/FamilyContext';
import { useAuthContext } from '../context/AuthContext';
import { exchangeApi, contractApi, Reward, Exchange } from '../services/api';
import ExchangeCard from '../components/reward/ExchangeCard';
import { getTokenBalance, updateLocalBalance, clearBalanceCache, burnTokens } from '../services/tokenService';

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
      // 使用tokenService获取余额，它会处理缓存逻辑
      console.log('[RewardStore] 获取代币余额...');
      const balance = await getTokenBalance(currentChild.walletAddress);
      
      if (balance !== '0') {
        console.log('[RewardStore] 成功获取余额:', balance);
        setBalance(balance);
      } else {
        console.log('[RewardStore] 获取余额为0，可能需要进一步检查');
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
            reward_name: exchange.reward_name,
            status: exchange.status,
            date: exchange.exchange_date
          });
        });
        
        setExchanges(response.data);
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

      // 首先尝试销毁代币
      let tokenBurned = false;
      if (currentChild.walletAddress) {
        try {
          console.log('尝试销毁代币...');
          tokenBurned = await burnTokens(currentChild.walletAddress, reward.token_price);
          console.log('代币销毁结果:', tokenBurned ? '成功' : '失败');
        } catch (burnError) {
          console.error('销毁代币时出错:', burnError);
          // 如果销毁失败，继续处理，让后端处理代币销毁
        }
      }

      // 调用后端API处理兑换
      const response = await exchangeApi.create({
        reward_id: reward.id,
        notes: `由${currentChild.name}兑换`,
        token_burned: tokenBurned // 告诉后端代币是否已在前端销毁
      });
      
      console.log('兑换响应:', response);
      
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
            token_amount: reward.token_price
          };
          
          console.log('增强后的兑换记录:', enhancedExchange);
          
          // 添加到兑换记录列表的开头
          setExchanges(prev => [enhancedExchange, ...prev]);
          
          // 切换到兑换记录标签
          setActiveTab('exchanged');
        }
        
        // 后台刷新数据
        console.log('开始后台刷新数据...');
        setTimeout(() => {
          console.log('执行后台刷新...');
          refreshBalance();
          fetchExchanges();
          fetchRewards();
        }, 1000);
      } else {
        console.error('兑换失败详情:', response);
        
        // 检查是否是库存不足错误
        const isOutOfStock = response.error && response.error.includes('out of stock');
        
        if (isOutOfStock) {
          console.log('检测到库存不足错误，从本地移除奖品');
          // 即使兑换失败，也从本地列表中移除商品（因为后端已经认为它没有库存）
          const updatedRewards = rewards.filter(r => r.id !== reward.id);
          setRewards(updatedRewards);
          
          // 强制刷新奖品列表
          setTimeout(() => {
            fetchRewards();
          }, 500);
          
          alert('兑换失败: 该奖品已被兑换完毕');
        } else {
          alert(`兑换失败: ${response.error || '未知错误'}`);
        }
      }
    } catch (error: any) {
      console.error('兑换过程出错:', error);
      
      // 检查错误信息中是否包含库存不足
      const errorMessage = error.response?.data?.error || error.message || '未知错误';
      const isOutOfStock = errorMessage.includes('out of stock');
      
      if (isOutOfStock) {
        console.log('检测到库存不足错误，从本地移除奖品');
        // 即使兑换失败，也从本地列表中移除商品（因为后端已经认为它没有库存）
        const updatedRewards = rewards.filter(r => r.id !== reward.id);
        setRewards(updatedRewards);
        
        alert('兑换失败: 该奖品已被兑换完毕');
        
        // 强制刷新奖品列表
        setTimeout(() => {
          fetchRewards();
        }, 500);
      } else {
        alert(`兑换过程发生错误: ${errorMessage}`);
      }
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
      // 不要在useEffect内部调用hooks
      console.log('当前选择的家庭:', selectedFamily);
      
      // 立即获取数据
      fetchRewards();
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
        <Card>
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
          
          {/* 强制刷新按钮 */}
          {activeTab === 'exchanged' && (
            <button
              className="ml-auto px-3 py-1 text-xs bg-blue-100 text-blue-700 rounded hover:bg-blue-200 flex items-center"
              onClick={() => {
                console.log('强制刷新兑换记录');
                fetchExchanges();
              }}
              disabled={loadingExchanges}
            >
              <RefreshCw className={`h-3 w-3 mr-1 ${loadingExchanges ? 'animate-spin' : ''}`} />
              刷新记录
            </button>
          )}
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
              {availableRewards.map(reward => (
                <Card key={reward.id} className="overflow-hidden">
                  <div className="h-48 bg-gray-200 overflow-hidden">
                    {reward.image_url ? (
                      <img
                        src={reward.image_url}
                        alt={reward.name}
                        className="w-full h-full object-contain"
                      />
                    ) : (
                      <div className="w-full h-full flex items-center justify-center bg-gray-100">
                        <ShoppingBag className="h-12 w-12 text-gray-400" />
                      </div>
                    )}
                  </div>
                  <CardBody>
                    <h3 className="font-semibold text-lg mb-2">{reward.name}</h3>
                    <p className="text-gray-600 text-sm mb-3">{reward.description}</p>
                    <div className="flex items-center justify-between">
                      <div className="text-primary-600 font-bold">
                        {Math.floor(reward.token_price)} FCT
                      </div>
                      <div className="text-sm text-gray-500">
                      Limited: 1
                      </div>
                    </div>
                    <Button
                      className="w-full mt-4"
                      variant="primary"
                      onClick={() => handleExchange(reward)}
                      disabled={parseFloat(balance) < reward.token_price || exchangeInProgress}
                    >
                      {exchangeInProgress ? 'Processing...' : 'Exchange'}
                    </Button>
                  </CardBody>
                </Card>
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
              
              {/* 调试按钮 - 仅在开发模式下显示 */}
              {process.env.NODE_ENV === 'development' && (
                <button
                  className="mt-4 px-4 py-2 bg-gray-200 text-gray-700 rounded hover:bg-gray-300"
                  onClick={() => {
                    console.log('添加测试兑换记录');
                    
                    // 创建测试兑换记录
                    const testExchange1 = {
                      id: 1001,
                      reward_id: 1,
                      child_id: currentChild ? Number(currentChild.id) : 1,
                      token_amount: 10,
                      status: 'confirmed' as 'confirmed',
                      exchange_date: new Date().toISOString(),
                      notes: '测试兑换记录1',
                      created_at: new Date().toISOString(),
                      updated_at: new Date().toISOString(),
                      reward_name: '测试奖品1',
                      reward_image: 'https://via.placeholder.com/150',
                      child_name: currentChild ? currentChild.name : '测试用户'
                    };
                    
                    const testExchange2 = {
                      id: 1002,
                      reward_id: 2,
                      child_id: currentChild ? Number(currentChild.id) : 1,
                      token_amount: 20,
                      status: 'confirmed' as 'confirmed',
                      exchange_date: new Date().toISOString(),
                      notes: '测试兑换记录2',
                      created_at: new Date().toISOString(),
                      updated_at: new Date().toISOString(),
                      reward_name: '测试奖品2',
                      reward_image: 'https://via.placeholder.com/150',
                      child_name: currentChild ? currentChild.name : '测试用户'
                    };
                    
                    // 添加到兑换记录
                    setExchanges([testExchange1, testExchange2]);
                  }}
                >
                  添加测试数据（仅开发环境）
                </button>
              )}
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