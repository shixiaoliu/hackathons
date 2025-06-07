import React, { useState } from 'react';
import { useNavigate, Routes, Route } from 'react-router-dom';
import { useAccount } from 'wagmi';
import { PlusCircle, Filter, AlertCircle, Users } from 'lucide-react';
import Button from '../components/common/Button';
import { mockTasks } from '../data/mockTasks';
import TaskCard from '../components/task/TaskCard';
import ParentTaskReview from './ParentTaskReview';
import ChildrenManager from '../components/family/ChildrenManager';
import { useFamily } from '../context/FamilyContext';
import { useTask } from '../context/TaskContext';

const ParentDashboard = () => {
  const navigate = useNavigate();
  const { address } = useAccount();
  const [filter, setFilter] = useState('all');
  const [activeTab, setActiveTab] = useState('tasks');
  const { children, selectedChild } = useFamily();
  const { tasks, getTasksForParent, approveTask, rejectTask } = useTask();
  
  // 获取当前家长的任务
  const parentTasks = address ? getTasksForParent(address) : [];
  
  // 添加调试信息
  console.log('[ParentDashboard] address:', address);
  console.log('[ParentDashboard] parentTasks:', parentTasks);
  console.log('[ParentDashboard] selectedChild:', selectedChild);
  console.log('[ParentDashboard] filter:', filter);
  
  // 根据选中的孩子过滤任务
  const filteredTasks = parentTasks.filter(task => {
    // 如果选中了孩子，只显示该孩子的任务；如果没有选中孩子，显示所有任务
    let matchesChild = true;
    if (selectedChild) {
      matchesChild = task.assignedChildId === selectedChild.id || 
                    task.assignedTo === selectedChild.walletAddress;
    }
    
    // 根据状态过滤
    let matchesStatus = true;
    if (filter === 'pending') {
      matchesStatus = task.status === 'completed';
    } else if (filter === 'active') {
      matchesStatus = task.status === 'open' || task.status === 'in-progress';
    } else if (filter === 'completed') {
      matchesStatus = task.status === 'approved' || task.status === 'rejected';
    }
    // filter === 'all' 时 matchesStatus 保持为 true
    
    return matchesChild && matchesStatus;
  });
  
  console.log('[ParentDashboard] filteredTasks:', filteredTasks);

  // 处理任务审批
  const handleApproveTask = (taskId) => {
    approveTask(taskId);
  };

  const handleRejectTask = (taskId) => {
    rejectTask(taskId);
  };
  
  const pendingReviewCount = parentTasks.filter(task => task.status === 'completed').length;

  return (
    <Routes>
      <Route path="/" element={
        <div className="max-w-6xl mx-auto">
          <div className="flex flex-col md:flex-row md:items-center md:justify-between mb-8">
            <div>
              <h1 className="text-3xl font-bold text-gray-900 mb-2">Parent Dashboard</h1>
              <p className="text-gray-600">
                {selectedChild 
                  ? `Managing tasks for ${selectedChild.name}`
                  : 'Manage children and tasks'
                }
              </p>
            </div>
            
            <div className="mt-4 md:mt-0 flex flex-col sm:flex-row gap-3">
              <Button 
                onClick={() => navigate('/create-task')}
                leftIcon={<PlusCircle className="h-5 w-5" />}
                disabled={!selectedChild}
              >
                Create Task
              </Button>
            </div>
          </div>

          {/* 标签页导航 */}
          <div className="mb-8">
            <div className="flex space-x-2 border-b border-gray-200">
              <button
                className={`px-4 py-2 text-sm font-medium flex items-center space-x-2 ${
                  activeTab === 'children'
                    ? 'text-primary-600 border-b-2 border-primary-600'
                    : 'text-gray-500 hover:text-gray-700'
                }`}
                onClick={() => setActiveTab('children')}
              >
                <Users className="h-4 w-4" />
                <span>Children ({children.length})</span>
              </button>
              <button
                className={`px-4 py-2 text-sm font-medium ${
                  activeTab === 'tasks'
                    ? 'text-primary-600 border-b-2 border-primary-600'
                    : 'text-gray-500 hover:text-gray-700'
                }`}
                onClick={() => setActiveTab('tasks')}
              >
                Tasks
                {selectedChild && ` for ${selectedChild.name}`}
              </button>
            </div>
          </div>

          {/* 内容区域 */}
          {activeTab === 'children' ? (
            <ChildrenManager />
          ) : (
            <div>
              {!selectedChild && (
                <div className="bg-blue-50 border-l-4 border-blue-400 p-4 mb-8 rounded-md">
                  <div className="flex">
                    <div className="flex-shrink-0">
                      <Users className="h-5 w-5 text-blue-400" />
                    </div>
                    <div className="ml-3">
                      <p className="text-sm text-blue-700">
                        Please select a child from the Children tab to create and manage tasks.
                      </p>
                    </div>
                  </div>
                </div>
              )}
              
              {selectedChild && (
                <>
                  {pendingReviewCount > 0 && (
                    <div className="bg-yellow-50 border-l-4 border-yellow-400 p-4 mb-8 rounded-md">
                      <div className="flex">
                        <div className="flex-shrink-0">
                          <AlertCircle className="h-5 w-5 text-yellow-400" />
                        </div>
                        <div className="ml-3">
                          <p className="text-sm text-yellow-700">
                            You have {pendingReviewCount} task{pendingReviewCount > 1 ? 's' : ''} waiting for review.
                          </p>
                        </div>
                      </div>
                    </div>
                  )}
                  
                  <div className="mb-8">
                    <div className="flex space-x-2 border-b border-gray-200">
                      <button
                        className={`px-4 py-2 text-sm font-medium ${
                          filter === 'all'
                            ? 'text-primary-600 border-b-2 border-primary-600'
                            : 'text-gray-500 hover:text-gray-700'
                        }`}
                        onClick={() => setFilter('all')}
                      >
                        All Tasks
                      </button>
                      <button
                        className={`px-4 py-2 text-sm font-medium ${
                          filter === 'active'
                            ? 'text-primary-600 border-b-2 border-primary-600'
                            : 'text-gray-500 hover:text-gray-700'
                        }`}
                        onClick={() => setFilter('active')}
                      >
                        Active
                      </button>
                      <button
                        className={`px-4 py-2 text-sm font-medium ${
                          filter === 'pending'
                            ? 'text-primary-600 border-b-2 border-primary-600'
                            : 'text-gray-500 hover:text-gray-700'
                        }`}
                        onClick={() => setFilter('pending')}
                      >
                        Pending Review
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
                  
                  <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                    {filteredTasks.map((task) => (
                      <div key={task.id} className="relative">
                        <TaskCard task={task} onClick={() => navigate(`/task/${task.id}`)} />
                        {task.status === 'completed' && (
                          <div className="absolute top-2 right-2 flex gap-1">
                            <Button 
                              size="sm" 
                              onClick={(e) => {
                                e.stopPropagation();
                                handleApproveTask(task.id);
                              }}
                              className="bg-green-600 hover:bg-green-700"
                            >
                              Approve
                            </Button>
                            <Button 
                              size="sm" 
                              variant="secondary"
                              onClick={(e) => {
                                e.stopPropagation();
                                handleRejectTask(task.id);
                              }}
                              className="bg-red-600 hover:bg-red-700 text-white"
                            >
                              Reject
                            </Button>
                          </div>
                        )}
                      </div>
                    ))}
                  </div>
                  
                  {filteredTasks.length === 0 && (
                    <div className="text-center py-12">
                      <p className="text-gray-500 text-lg mb-4">No tasks found</p>
                      <Button onClick={() => navigate('/create-task')}>
                        Create First Task
                      </Button>
                    </div>
                  )}
                </>
              )}
            </div>
          )}
        </div>
      } />
      <Route path="/review/:taskId" element={<ParentTaskReview />} />
    </Routes>
  );
};

export default ParentDashboard;