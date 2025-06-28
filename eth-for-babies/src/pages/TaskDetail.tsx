import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { useUserRole } from '../context/UserRoleContext';
import { useTask } from '../context/TaskContext';
import { useAuthContext } from '../context/AuthContext';
import { useAccount } from 'wagmi';
import { Clock, Award, ChevronLeft, Calendar, User, CheckCircle, AlertTriangle } from 'lucide-react';
import Button from '../components/common/Button';
import Card, { CardBody, CardHeader, CardFooter } from '../components/common/Card';
import { formatDateTime, formatDate } from '../utils/dateUtils';
import { ethers } from 'ethers';
import { TaskContractABI } from '../contracts/TaskContract';

// Get contract address from environment variables
const TASK_CONTRACT_ADDRESS = import.meta.env.VITE_TASK_CONTRACT_ADDRESS || '0x123456789...'; // Replace placeholder with actual fallback address

const TaskDetail = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const { userRole } = useUserRole();
  const { tasks, assignTask, submitTask, approveTask, rejectTask } = useTask();
  const { user } = useAuthContext();
  const { address } = useAccount();
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [contractError, setContractError] = useState<string | null>(null);
  
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

  return (
    <div className="max-w-3xl mx-auto">
      <div className="flex items-center mb-6">
        <button
          onClick={() => navigate(isParent ? '/parent' : '/child')}
          className="flex items-center text-gray-600 hover:text-gray-900"
        >
          <ChevronLeft className="h-5 w-5 mr-1" />
          Back to dashboard
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
                <Button onClick={() => approveTask(task.id)}>
                  Approve
                </Button>
                <Button 
                  variant="secondary"
                  onClick={() => rejectTask(task.id)}
                >
                  Reject
                </Button>
              </div>
            )}
          </div>
        </CardFooter>
      </Card>
    </div>
  );
};

export default TaskDetail;