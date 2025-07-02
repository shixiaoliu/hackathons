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

const RewardStore = () => {
  const navigate = useNavigate();
  const { rewards, loading, error, fetchRewards } = useReward();
  const { currentChild, selectedFamily } = useFamily();
  const { user } = useAuthContext();
  const [activeTab, setActiveTab] = useState<'available' | 'exchanged'>('available');
  const [exchanges, setExchanges] = useState<Exchange[]>([]);
  const [balance, setBalance] = useState<string>('0');
  const [loadingExchanges, setLoadingExchanges] = useState<boolean>(false);
  const [loadingBalance, setLoadingBalance] = useState<boolean>(false);
  const [exchangeInProgress, setExchangeInProgress] = useState<boolean>(false);
  
  // 获取代币余额
  const fetchBalance = async () => {
    if (!currentChild?.walletAddress) return;
    
    setLoadingBalance(true);
    try {
      const response = await contractApi.getBalance(currentChild.walletAddress);
      if (response.success && response.data) {
        setBalance(response.data.balance);
      }
    } catch (error) {
      console.error('获取余额失败:', error);
    } finally {
      setLoadingBalance(false);
    }
  };
  
  // 获取兑换记录
  const fetchExchanges = async () => {
    if (!currentChild) return;
    
    setLoadingExchanges(true);
    try {
      const response = await exchangeApi.getByChild();
      if (response.success && response.data) {
        setExchanges(response.data);
      }
    } catch (error) {
      console.error('获取兑换记录失败:', error);
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
      
      const response = await exchangeApi.create({
        reward_id: reward.id,
        notes: `由${currentChild.name}兑换`
      });
      
      if (response.success) {
        alert('兑换申请已提交');
        // 刷新数据
        fetchBalance();
        fetchExchanges();
        fetchRewards();
      } else {
        console.error('兑换失败详情:', response);
        alert(`兑换失败: ${response.error || '未知错误'}`);
      }
    } catch (error: any) {
      console.error('兑换过程出错:', error);
      const errorMessage = error.response?.data?.error || error.message || '未知错误';
      alert(`兑换过程发生错误: ${errorMessage}`);
    } finally {
      setExchangeInProgress(false);
    }
  };
  
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
            <div>
              <p className="text-sm text-gray-500">Token Balance</p>
              <p className="text-2xl font-bold text-gray-900">
                {loadingBalance ? (
                  <span className="animate-pulse">Loading...</span>
                ) : (
                  `${parseFloat(balance).toFixed(4)} tokens`
                )}
              </p>
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
                        {reward.token_price} tokens
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
          ) : exchanges.length === 0 ? (
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
                        <p className="text-2xl font-bold text-primary-600">{totalSpent.toFixed(4)} tokens</p>
                      </div>
                    </div>
                  </CardBody>
                </Card>
              </div>
              
              <div className="space-y-4">
                {exchanges.map(exchange => (
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