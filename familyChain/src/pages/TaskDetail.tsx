import React, { useState, useEffect } from 'react';
import { useParams, useNavigate, useLocation } from 'react-router-dom';
import { useUserRole } from '../context/UserRoleContext';
import { useTask } from '../context/TaskContext';
import { useFamily } from '../context/FamilyContext';
import { useAuthContext } from '../context/AuthContext';
import { useAccount } from 'wagmi';
import { Clock, Award, ChevronLeft, Calendar, User, CheckCircle, AlertTriangle } from 'lucide-react';
import Button from '../components/common/Button';
import Card, { CardBody, CardHeader, CardFooter } from '../components/common/Card';
import Modal from '../components/common/Modal';
import { formatDateTime, formatDate } from '../utils/dateUtils';

const TaskDetail = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const location = useLocation();
  const { userRole } = useUserRole();
  const { tasks, assignTask, submitTask, approveTask, rejectTask } = useTask();
  const { getAllChildren } = useFamily();
  const { user } = useAuthContext();
  const { address } = useAccount();
  
  // 检查是否从任务列表页面进入
  const isFromTaskList = location.state && location.state.from === 'tasks';
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [contractError, setContractError] = useState<string | null>(null);
  
  // Assign Modal 相关状态
  const [showAssignModal, setShowAssignModal] = useState(false);
  const [selectedChildId, setSelectedChildId] = useState<string>('');
  const [children, setChildren] = useState<any[]>([]);
  const [isAssigning, setIsAssigning] = useState(false);
  
  // 获取子账户列表
  useEffect(() => {
    const childrenData = getAllChildren();
    setChildren(childrenData);
  }, [getAllChildren]);
  
  // 从TaskContext中查找任务而不是mockTasks
  const task = tasks.find(task => task.id === id);
  
  // 检查任务的合约ID
  useEffect(() => {
    if (task && !task.contractTaskId) {
      setContractError('此任务没有关联的区块链合约ID，无法在区块链上完成任务。');
    } else {
      setContractError(null);
    }
  }, [task]);
  
  if (!task) {
    return (
      <div className="max-w-3xl mx-auto text-center py-12">
        <h2 className="text-2xl font-bold text-gray-900 mb-4">Task Not Found</h2>
        <p className="text-gray-600 mb-6">The task you're looking for doesn't exist or has been removed.</p>
        <Button onClick={() => navigate(userRole === 'parent' ? '/parent' : '/child')}>
          Return to Dashboard
        </Button>
      </div>
    );
  }
  
  const isParent = userRole === 'parent';
  const isChild = userRole === 'child' || user?.role === 'child';
  const isOverdue = new Date(task.deadline) < new Date() && task.status === 'open';
  
  const isAssignedToCurrentUser = task?.assignedTo && address && 
    task.assignedTo.toLowerCase() === address.toLowerCase();
  
  const canAcceptTask = isChild && task.status === 'open';
  const canSubmitTask = isChild && task.status === 'in-progress' && isAssignedToCurrentUser;

  // 添加调试信息
  console.log('[TaskDetail Debug] userRole:', userRole);
  console.log('[TaskDetail Debug] user?.role:', user?.role);
  console.log('[TaskDetail Debug] isChild:', isChild);
  console.log('[TaskDetail Debug] task.status:', task.status);
  console.log('[TaskDetail Debug] task.assignedTo:', task.assignedTo);
  console.log('[TaskDetail Debug] current address:', address);
  console.log('[TaskDetail Debug] isAssignedToCurrentUser:', isAssignedToCurrentUser);
  console.log('[TaskDetail Debug] canSubmitTask:', canSubmitTask);
  console.log('[TaskDetail Debug] contractTaskId:', task.contractTaskId);
  console.log('[TaskDetail Debug] contractTaskId type:', typeof task.contractTaskId);

  // 处理任务提交
  const handleTaskSubmit = async () => {
    if (!task.contractTaskId) {
      alert('此任务没有关联的区块链合约ID，无法在区块链上完成任务。请联系父母重新创建任务。');
      return;
    }
    
    setIsSubmitting(true);
    try {
      const proof = {
        description: '任务已完成',
        submittedAt: new Date().toISOString()
      };
      await submitTask(task.id, proof);
    } catch (error) {
      console.error('提交任务失败:', error);
    } finally {
      setIsSubmitting(false);
    }
  };
  
  // 处理分配任务按钮点击
  const handleAssignClick = () => {
    setSelectedChildId('');
    setShowAssignModal(true);
  };
  
  // 处理任务分配
  const handleAssignTask = async () => {
    if (!id || !selectedChildId) return;
    
    setIsAssigning(true);
    try {
      const selectedChild = children.find(child => child.id === selectedChildId);
      if (!selectedChild) return;

      console.log('分配任务 - 任务ID:', id, '子账户ID:', selectedChildId, '子账户钱包地址:', selectedChild.walletAddress);
      console.log('子账户详情:', selectedChild);

      // 调用TaskContext中的assignTask方法，它会处理区块链调用和后端API更新
      await assignTask(id, selectedChild.walletAddress, selectedChildId);
      
      setShowAssignModal(false);
      setSelectedChildId('');
      
      alert('任务已成功分配！');
    } catch (error) {
      console.error('分配任务失败:', error);
      alert('分配任务失败，请重试');
    } finally {
      setIsAssigning(false);
    }
  };

  return (
    <div className="max-w-3xl mx-auto">
      <div className="flex items-center mb-6">
        <button
          onClick={() => {
            // 根据来源决定返回位置
            if (isFromTaskList) {
              // 如果是从任务列表页面进入，则返回任务列表
              navigate('/tasks');
            } else {
              // 否则返回仪表盘
              navigate(isParent ? '/parent' : '/child');
            }
          }}
          className="flex items-center text-gray-600 hover:text-gray-900"
        >
          <ChevronLeft className="h-5 w-5 mr-1" />
          {isFromTaskList ? 'Return to All Tasks' : 'Return to Dashboard'}
        </button>
      </div>
      
      <Card>
        {/* Task image if available */}
        {task.imageUrl && (
          <div className="h-64 overflow-hidden">
            <img 
              src={task.imageUrl} 
              alt={task.title} 
              className="w-full h-full object-cover"
            />
          </div>
        )}
        
        {/* 合约错误警告 */}
        {contractError && canSubmitTask && (
          <div className="bg-yellow-50 border border-yellow-200 rounded-lg p-4 m-4">
            <div className="flex items-center">
              <AlertTriangle className="h-5 w-5 text-yellow-500 mr-2" />
              <span className="text-yellow-700 font-medium">{contractError}</span>
            </div>
          </div>
        )}
        
        <CardHeader>
          <div className="flex justify-between items-start">
            <div className="flex-1">
              <h1 className="text-2xl font-bold text-gray-900 mb-2">{task.title}</h1>
              <div className="flex items-center space-x-4 text-sm text-gray-500">
                <div className="flex items-center">
                  <Calendar className="h-4 w-4 mr-1" />
                  <span>Due {formatDate(task.deadline)}</span>
                </div>
                <div className="flex items-center">
                  <Award className="h-4 w-4 mr-1" />
                  <span className="font-medium text-primary-600">{task.reward} ETH</span>
                </div>
              </div>
            </div>
            
            <div className="flex flex-col items-end space-y-2">
              <span className={`px-3 py-1 rounded-full text-sm font-medium ${
                task.difficulty === 'easy' ? 'bg-green-100 text-green-800' :
                task.difficulty === 'medium' ? 'bg-yellow-100 text-yellow-800' :
                'bg-red-100 text-red-800'
              }`}>
                {task.difficulty.charAt(0).toUpperCase() + task.difficulty.slice(1)}
              </span>
              
              <span className={`px-3 py-1 rounded-full text-sm font-medium ${
                task.status === 'open' ? 'bg-blue-100 text-blue-800' :
                task.status === 'in-progress' ? 'bg-yellow-100 text-yellow-800' :
                task.status === 'completed' ? 'bg-orange-100 text-orange-800' :
                task.status === 'approved' ? 'bg-green-100 text-green-800' :
                'bg-red-100 text-red-800'
              }`}>
                {task.status === 'in-progress' ? 'In Progress' : 
                 task.status.charAt(0).toUpperCase() + task.status.slice(1)}
              </span>
            </div>
          </div>
        </CardHeader>
        
        <CardBody>
          {/* Overdue warning */}
          {isOverdue && (
            <div className="bg-red-50 border border-red-200 rounded-lg p-4 mb-6">
              <div className="flex items-center">
                <AlertTriangle className="h-5 w-5 text-red-500 mr-2" />
                <span className="text-red-700 font-medium">This task is overdue!</span>
              </div>
            </div>
          )}
          
          {/* 权限警告 */}
          {canSubmitTask === false && task.status === 'in-progress' && (
            <div className="bg-yellow-50 border border-yellow-200 rounded-lg p-4 mb-6">
              <div className="flex items-center">
                <AlertTriangle className="h-5 w-5 text-yellow-500 mr-2" />
                <span className="text-yellow-700 font-medium">
                  只有被分配的孩子 ({task.assignedTo?.slice(0, 6)}...{task.assignedTo?.slice(-4)}) 
                  可以完成这个任务。当前钱包地址: {address?.slice(0, 6)}...{address?.slice(-4)}
                </span>
              </div>
            </div>
          )}
          
          {/* Task description */}
          <div className="mb-6">
            <h3 className="text-lg font-semibold text-gray-900 mb-3">Description</h3>
            <p className="text-gray-700 leading-relaxed">{task.description}</p>
          </div>
          
          {/* Completion criteria */}
          {task.completionCriteria && (
            <div className="mb-6">
              <h3 className="text-lg font-semibold text-gray-900 mb-3">Completion Criteria</h3>
              <p className="text-gray-700 leading-relaxed">{task.completionCriteria}</p>
            </div>
          )}
          
          {/* Task details */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
            <div>
              <h4 className="font-medium text-gray-900 mb-2">Task Details</h4>
              <div className="space-y-2 text-sm">
                <div className="flex justify-between">
                  <span className="text-gray-500">Created:</span>
                  <span className="text-gray-900">{formatDateTime(task.createdAt)}</span>
                </div>
                <div className="flex justify-between">
                  <span className="text-gray-500">Deadline:</span>
                  <span className={`font-medium ${
                    isOverdue ? 'text-red-600' : 'text-gray-900'
                  }`}>
                    {formatDateTime(task.deadline)}
                  </span>
                </div>
                <div className="flex justify-between">
                  <span className="text-gray-500">Reward:</span>
                  <span className="text-primary-600 font-medium">{task.reward} ETH</span>
                </div>
                {task.contractTaskId && (
                  <div className="flex justify-between">
                    <span className="text-gray-500">Contract Task ID:</span>
                    <span className="text-gray-900">{task.contractTaskId}</span>
                  </div>
                )}
              </div>
            </div>
            
            {task.assignedTo && (
              <div>
                <h4 className="font-medium text-gray-900 mb-2">Assignment</h4>
                <div className="space-y-2 text-sm">
                  <div className="flex justify-between">
                    <span className="text-gray-500">Assigned to:</span>
                    <span className="text-gray-900 font-mono text-xs">
                      {task.assignedTo.slice(0, 6)}...{task.assignedTo.slice(-4)}
                    </span>
                  </div>
                  {task.updatedAt !== task.createdAt && (
                    <div className="flex justify-between">
                      <span className="text-gray-500">Last updated:</span>
                      <span className="text-gray-900">{formatDateTime(task.updatedAt)}</span>
                    </div>
                  )}
                </div>
              </div>
            )}
          </div>
          
          {/* Submission proof */}
          {task.submissionProof && (
            <div className="mb-6">
              <h3 className="text-lg font-semibold text-gray-900 mb-3">Submission</h3>
              <div className="bg-gray-50 rounded-lg p-4">
                <div className="flex items-center mb-3">
                  <CheckCircle className="h-5 w-5 text-green-500 mr-2" />
                  <span className="font-medium text-gray-900">Task Completed</span>
                  <span className="ml-auto text-sm text-gray-500">
                    {formatDateTime(task.submissionProof.submittedAt)}
                  </span>
                </div>
                
                {task.submissionProof.description && (
                  <p className="text-gray-700 mb-3">{task.submissionProof.description}</p>
                )}
                
                {task.submissionProof.imageUrl && (
                  <div className="mt-3">
                    <img 
                      src={task.submissionProof.imageUrl} 
                      alt="Task completion proof" 
                      className="max-w-full h-auto rounded-lg border"
                    />
                  </div>
                )}
              </div>
            </div>
          )}
        </CardBody>
        
        <CardFooter>
          <div className="flex flex-col w-full">
            <div className="flex justify-between items-center w-full">
              <div className="flex items-center space-x-2">
                {task.status === 'approved' && (
                  <div className="flex items-center text-green-600">
                    <CheckCircle className="h-5 w-5 mr-1" />
                    <span className="font-medium">Approved</span>
                  </div>
                )}
                
                {task.status === 'rejected' && (
                  <div className="flex items-center text-red-600">
                    <AlertTriangle className="h-5 w-5 mr-1" />
                    <span className="font-medium">Rejected</span>
                  </div>
                )}
              </div>
              
              <div className="flex space-x-3">
                {/* Child actions */}
                {canAcceptTask && (
                  <Button onClick={() => {
                    if (address && user?.id) {
                      assignTask(task.id, address, user.id.toString());
                    }
                  }}>
                    Accept Task
                  </Button>
                )}
                
                {/* Parent actions for open tasks */}
                {isParent && task.status === 'open' && !isOverdue && (
                  <Button onClick={handleAssignClick}>
                    Assign Task
                  </Button>
                )}
                
                {canSubmitTask && (
                  <Button 
                    onClick={handleTaskSubmit}
                    isLoading={isSubmitting}
                    disabled={!task.contractTaskId || !isAssignedToCurrentUser}
                    title={!task.contractTaskId ? "此任务没有关联的区块链合约ID" : 
                           !isAssignedToCurrentUser ? "只有被分配的孩子才能完成任务" : ""}
                  >
                    Complete Task
                  </Button>
                )}
              </div>
            </div>
            
            {/* Parent actions - 移动到右下角 */}
            {isParent && task.status === 'completed' && (
              <div className="flex justify-end mt-4 space-x-2">
                <Button onClick={() => approveTask(task.id)}
                       disabled={!task.contractTaskId}
                       title={!task.contractTaskId ? "此任务没有关联的区块链合约ID，无法在区块链上批准" : "批准任务并在区块链上转移奖励"}>
                  Approve
                </Button>
                <Button 
                  variant="secondary"
                  onClick={() => rejectTask(task.id)}
                  disabled={!task.contractTaskId}
                  title={!task.contractTaskId ? "此任务没有关联的区块链合约ID，无法在区块链上拒绝" : "拒绝任务并在区块链上退回奖励"}
                >
                  Reject
                </Button>
              </div>
            )}
          </div>
        </CardFooter>
      </Card>

      {/* 分配任务弹窗 */}
      <Modal
        isOpen={showAssignModal}
        onClose={() => setShowAssignModal(false)}
        title="分配任务"
      >
        <div className="p-4">
          <p className="mb-4 font-medium">{task.title}</p>
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-700 mb-1">
              选择孩子
            </label>
            <select
              value={selectedChildId}
              onChange={(e) => setSelectedChildId(e.target.value)}
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-primary-500 focus:border-primary-500"
              disabled={isAssigning}
            >
              <option value="">-- 请选择 --</option>
              {children.map(child => (
                <option key={child.id} value={child.id}>
                  {child.name}
                </option>
              ))}
            </select>
          </div>
          <div className="flex justify-end space-x-3">
            <button 
              className="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50"
              onClick={() => setShowAssignModal(false)}
              disabled={isAssigning}
            >
              取消
            </button>
            <button 
              className="px-4 py-2 bg-primary-600 text-white rounded-md hover:bg-primary-700 disabled:bg-gray-400"
              onClick={handleAssignTask}
              disabled={!selectedChildId || isAssigning}
            >
              {isAssigning ? '处理中...' : '确认分配'}
            </button>
          </div>
        </div>
      </Modal>
    </div>
  );
};

export default TaskDetail;