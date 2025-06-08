import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAccount } from 'wagmi';
import { Filter, Wallet, History, Clock, User } from 'lucide-react';
import TaskCard from '../components/task/TaskCard';
import Card, { CardBody, CardHeader } from '../components/common/Card';
import { formatDistanceToNow } from '../utils/dateUtils';
import { useFamily } from '../context/FamilyContext';
import { useTask } from '../context/TaskContext';
import { useAuthContext } from '../context/AuthContext';

const ChildDashboard = () => {
  const navigate = useNavigate();
  const { address } = useAccount();
  const { user } = useAuthContext();
  const { currentChild, loginAsChild, getAllChildren, findChildByAddress } = useFamily();
  const { tasks, assignTask, getTasksForChild, getAvailableTasks, submitTask } = useTask();
  const [filter, setFilter] = useState('available');
  const [showChildSelector, setShowChildSelector] = useState(false);

  // 获取当前用户的钱包地址（优先使用user.wallet_address，其次使用wagmi的address）
  const currentWalletAddress = user?.wallet_address || address;

  // 添加调试信息
  console.log('[ChildDashboard] currentWalletAddress:', currentWalletAddress);
  console.log('[ChildDashboard] currentChild:', currentChild);
  console.log('[ChildDashboard] getAllChildren():', getAllChildren());
  console.log('[ChildDashboard] tasks:', tasks);

  // 如果没有当前child，显示child选择器
  if (!currentChild && currentWalletAddress) {
    const allChildren = getAllChildren();
    console.log('[ChildDashboard] allChildren:', allChildren);
    const availableChildren = allChildren.filter(child => 
      child.walletAddress.toLowerCase() === currentWalletAddress.toLowerCase()
    );
    console.log('[ChildDashboard] availableChildren:', availableChildren);

    if (availableChildren.length === 0) {
      return (
        <div className="max-w-2xl mx-auto text-center py-12">
          <div className="bg-yellow-50 border border-yellow-200 rounded-lg p-6">
            <User className="h-12 w-12 text-yellow-600 mx-auto mb-4" />
            <h2 className="text-xl font-semibold text-gray-900 mb-2">Child Account Not Found</h2>
            <p className="text-gray-600 mb-4">
              Your wallet address ({currentWalletAddress}) is not registered as a child account.
              Please ask your parent to add you as a child first.
            </p>
            <div className="mt-4 p-4 bg-gray-100 rounded text-sm text-left">
              <p className="font-semibold mb-2">Debug Info:</p>
              <p>Current Address: {currentWalletAddress}</p>
              <p>Available Children: {allChildren.length}</p>
              <div className="mt-2">
                <p>Children Addresses:</p>
                {allChildren.map(child => (
                  <p key={child.id} className="ml-2">- {child.name}: {child.walletAddress}</p>
                ))}
              </div>
            </div>
            <button
              onClick={() => navigate('/')}
              className="bg-primary-600 text-white px-4 py-2 rounded-md hover:bg-primary-700 mt-4"
            >
              Go Back to Home
            </button>
          </div>
        </div>
      );
    }

    // 自动登录找到的child
    const child = availableChildren[0];
    console.log('[ChildDashboard] Auto-login child:', child);
    loginAsChild(child.walletAddress);
  }

  if (!currentChild) {
    return (
      <div className="max-w-2xl mx-auto text-center py-12">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600 mx-auto"></div>
        <p className="mt-4 text-gray-600">Loading your profile...</p>
      </div>
    );
  }

  const userAddress = currentChild.walletAddress;
  
  // 获取任务数据
  const getFilteredTasks = () => {
    if (!currentChild) return [];
    
    switch (filter) {
      case 'available':
        return getAvailableTasks();
      case 'my-tasks':
        return getTasksForChild(currentChild.walletAddress).filter(task => 
          task.status === 'in-progress' || task.status === 'completed'
        );
      case 'completed':
        return getTasksForChild(currentChild.walletAddress).filter(task => 
          task.status === 'completed' || task.status === 'approved'
        );
      default:
        return [];
    }
  };

  const filteredTasks = getFilteredTasks();

  // 计算奖励和统计
  const childTasks = currentChild ? getTasksForChild(currentChild.walletAddress) : [];
  
  const earnedRewards = childTasks
    .filter(task => task.status === 'approved')
    .reduce((total, task) => total + parseFloat(task.reward), 0);

  const pendingRewards = childTasks
    .filter(task => task.status === 'completed')
    .reduce((total, task) => total + parseFloat(task.reward), 0);

  // 处理任务分配
  const handleTakeTask = (taskId: string) => {
    if (!currentChild) {
      alert('Please login as a child first');
      return;
    }
    
    assignTask(taskId, currentChild.walletAddress, currentChild.id);
    alert('Task assigned successfully!');
  };

  // 获取最近的交易记录
  const recentTransactions = childTasks
    .filter(task => task.status === 'approved')
    .slice(0, 3);

  return (
    <div className="max-w-6xl mx-auto">
      <div className="mb-8">
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold text-gray-900 mb-2">Hello, {currentChild.name}!</h1>
            <p className="text-gray-600">Find tasks, earn rewards</p>
          </div>
          <div className="flex items-center space-x-3">
            <div className="w-12 h-12 rounded-full bg-primary-100 flex items-center justify-center">
              {currentChild.avatar ? (
                <img src={currentChild.avatar} alt={currentChild.name} className="w-12 h-12 rounded-full" />
              ) : (
                <User className="h-6 w-6 text-primary-600" />
              )}
            </div>
            <div className="text-right">
              <p className="text-sm font-medium text-gray-900">{currentChild.name}</p>
              <p className="text-xs text-gray-500">Age {currentChild.age}</p>
            </div>
          </div>
        </div>
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
              <p className="text-sm text-gray-500">Earned Rewards</p>
              <p className="text-2xl font-bold text-gray-900">{earnedRewards.toFixed(4)} ETH</p>
            </div>
          </CardBody>
        </Card>
        
        <Card>
          <CardBody className="flex items-center">
            <div className="flex-shrink-0 mr-4">
              <div className="w-12 h-12 rounded-full bg-yellow-100 flex items-center justify-center">
                <Clock className="h-6 w-6 text-yellow-600" />
              </div>
            </div>
            <div>
              <p className="text-sm text-gray-500">Pending Rewards</p>
              <p className="text-2xl font-bold text-gray-900">{pendingRewards.toFixed(4)} ETH</p>
            </div>
          </CardBody>
        </Card>
        
        <Card>
          <CardBody className="flex items-center">
            <div className="flex-shrink-0 mr-4">
              <div className="w-12 h-12 rounded-full bg-blue-100 flex items-center justify-center">
                <History className="h-6 w-6 text-blue-600" />
              </div>
            </div>
            <div>
              <p className="text-sm text-gray-500">Completed Tasks</p>
              <p className="text-2xl font-bold text-gray-900">
                {childTasks.filter(task => 
                  task.status === 'approved'
                ).length}
              </p>
            </div>
          </CardBody>
        </Card>
      </div>
      
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6 mb-8">
        <div className="lg:col-span-2">
          <div className="mb-6">
            <div className="flex space-x-2 border-b border-gray-200">
              <button
                className={`px-4 py-2 text-sm font-medium ${
                  filter === 'available'
                    ? 'text-primary-600 border-b-2 border-primary-600'
                    : 'text-gray-500 hover:text-gray-700'
                }`}
                onClick={() => setFilter('available')}
              >
                Available Tasks
              </button>
              <button
                className={`px-4 py-2 text-sm font-medium ${
                  filter === 'my-tasks'
                    ? 'text-primary-600 border-b-2 border-primary-600'
                    : 'text-gray-500 hover:text-gray-700'
                }`}
                onClick={() => setFilter('my-tasks')}
              >
                My Tasks
              </button>
              <button
                className={`px-4 py-2 text-sm font-medium ${
                  filter === 'completed'
                    ? 'text-primary-600 border-b-2 border-primary-600'
                    : 'text-gray-500 hover:text-gray-700'
                }`}
                onClick={() => setFilter('completed')}
              >
                Completed
              </button>
            </div>
          </div>
          
          {filteredTasks.length === 0 ? (
            <div className="text-center py-12 bg-white rounded-lg shadow-md">
              <div className="inline-flex items-center justify-center w-16 h-16 rounded-full bg-gray-100 text-gray-400 mb-4">
                <Filter className="h-8 w-8" />
              </div>
              <h3 className="text-lg font-medium text-gray-900 mb-2">No tasks found</h3>
              <p className="text-gray-500 max-w-md mx-auto">
                {filter === 'available' 
                  ? "There are no available tasks right now. Check back later!"
                  : filter === 'my-tasks'
                    ? "You don't have any tasks in progress."
                    : "You haven't completed any tasks yet."}
              </p>
            </div>
          ) : (
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-6">
              {filteredTasks.map(task => (
                <TaskCard
                   key={task.id} 
                   task={task} 
                   onClick={() => navigate(`/task/${task.id}`)}
                   onTakeTask={filter === 'available' ? () => handleTakeTask(task.id) : undefined}
                   onCompleteTask={task.status === 'in-progress' ? () => {
                     // 直接在仪表盘完成任务
                     const proof = {
                       description: '任务已完成',
                       submittedAt: new Date().toISOString()
                     };
                     submitTask(task.id, proof);
                     alert('任务已标记为完成！');
                   } : undefined}
                   showTakeButton={filter === 'available'}
                 />
              ))}
            </div>
          )}
        </div>
        
        <div>
          <Card>
            <CardHeader>
              <h3 className="text-lg font-medium text-gray-900">Recent Transactions</h3>
            </CardHeader>
            <CardBody className="p-0">
              {recentTransactions.length === 0 ? (
                <div className="py-6 px-4 text-center">
                  <p className="text-gray-500 text-sm">No transactions yet</p>
                </div>
              ) : (
                <ul className="divide-y divide-gray-200">
                  {recentTransactions.map(transaction => (
                    <li key={transaction.id} className="px-6 py-4">
                      <div className="flex justify-between">
                        <div>
                          <p className="text-sm font-medium text-gray-900 truncate">
                            {transaction.title}
                          </p>
                          <p className="text-xs text-gray-500">
                            {formatDistanceToNow(new Date(transaction.updatedAt))} ago
                          </p>
                        </div>
                        <p className="text-sm font-semibold text-green-600">
                          +{transaction.reward} ETH
                        </p>
                      </div>
                    </li>
                  ))}
                </ul>
              )}
            </CardBody>
          </Card>
        </div>
      </div>
    </div>
  );
};

export default ChildDashboard;