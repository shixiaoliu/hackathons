import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useTask } from '../context/TaskContext';
import { useFamily } from '../context/FamilyContext';
import TaskCard from '../components/task/TaskCard';
import { Task } from '../types/task';
import Button from '../components/common/Button';
import { RefreshCw, Plus } from 'lucide-react';
import Modal from '../components/common/Modal';

const TaskList = () => {
  const navigate = useNavigate();
  const { getAllTasks, assignTask } = useTask();
  const { getAllChildren } = useFamily();
  const [tasks, setTasks] = useState<Task[]>([]);
  const [selectedTask, setSelectedTask] = useState<Task | null>(null);
  const [showAssignModal, setShowAssignModal] = useState(false);
  const [selectedChildId, setSelectedChildId] = useState<string>('');
  const [children, setChildren] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [isRefreshing, setIsRefreshing] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  const fetchTasks = async () => {
    try {
      setLoading(true);
      setError(null);
      const allTasks = await getAllTasks();
      console.log('Retrieved tasks:', allTasks);
      
      // 确保 allTasks 符合 Task 类型
      const typedTasks: Task[] = allTasks.map((task: any) => ({
        id: task.id?.toString() || '',
        title: task.title || '',
        description: task.description || '',
        reward: task.reward || '0',
        deadline: task.deadline || new Date().toISOString(),
        difficulty: task.difficulty || 'medium',
        status: task.status || 'open',
        createdBy: task.createdBy || '',
        createdAt: task.createdAt || new Date().toISOString(),
        updatedAt: task.updatedAt || new Date().toISOString(),
        completionCriteria: task.completionCriteria || '',
        assignedTo: task.assignedTo,
        assignedChildId: task.assignedChildId,
        imageUrl: task.imageUrl,
        contractTaskId: task.contractTaskId
      }));
      
      setTasks(typedTasks);
      
      // 获取所有子账户
      const childrenData = getAllChildren();
      setChildren(childrenData);
    } catch (err) {
      console.error('Failed to fetch tasks:', err);
      setError('Failed to fetch tasks. Please try again later.');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchTasks();
  }, []);

  const handleRefresh = async () => {
    setIsRefreshing(true);
    await fetchTasks();
    setIsRefreshing(false);
  };

  const handleCreateTask = () => {
    navigate('/create-task');
  };

  const handleAssignClick = (task: Task) => {
    setSelectedTask(task);
    setShowAssignModal(true);
  };

  const handleAssignTask = async () => {
    if (!selectedTask || !selectedChildId) return;
    
    setIsLoading(true);
    try {
      const selectedChild = children.find(child => child.id === selectedChildId);
      if (!selectedChild) return;

      console.log('Assigning task - Task ID:', selectedTask.id, 'Child ID:', selectedChildId, 'Child wallet address:', selectedChild.walletAddress);
      console.log('Child details:', selectedChild);

      // 调用TaskContext中的assignTask方法，它会处理区块链调用和后端API更新
      await assignTask(selectedTask.id, selectedChild.walletAddress, selectedChildId);
      
      // 更新任务列表UI（虽然refreshTasks已经在assignTask内部调用，但这里额外更新本地UI以提供即时反馈）
      setTasks(prevTasks => {
        return prevTasks.map(task => {
          if (task.id === selectedTask.id) {
            return {
              ...task,
              assignedChildId: selectedChildId,
              assignedTo: selectedChild.walletAddress,
              status: 'in-progress' as const
            };
          }
          return task;
        });
      });
      
      setShowAssignModal(false);
      setSelectedChildId('');
      
      console.log('Task assigned successfully');
    } catch (error) {
      console.error('Failed to assign task:', error);
      alert(`Failed to assign task: ${error instanceof Error ? error.message : 'Unknown error'}`);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="container mx-auto p-6">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold">All Tasks</h1>
        <div className="flex space-x-2">
          <Button 
            variant="secondary" 
            leftIcon={<RefreshCw className={`h-5 w-5 ${isRefreshing ? 'animate-spin' : ''}`} />}
            onClick={handleRefresh}
            disabled={isRefreshing}
          >
            Refresh
          </Button>
          <Button 
            variant="primary" 
            leftIcon={<Plus className="h-5 w-5" />}
            onClick={handleCreateTask}
          >
            Create Task
          </Button>
        </div>
      </div>

      {loading && (
        <div className="flex justify-center my-8">
          <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary-600"></div>
        </div>
      )}
      
      {error && (
        <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
          {error}
        </div>
      )}

      {!loading && !error && tasks.length === 0 && (
        <div className="text-center py-8 bg-gray-50 rounded-lg">
          <p className="text-gray-600">No tasks found.</p>
        </div>
      )}

      {!loading && !error && tasks.length > 0 && (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {tasks.map(task => (
            <div key={task.id} className="relative">
              <TaskCard 
                task={task}
                onClick={() => navigate(`/task/${task.id}`, { state: { from: 'tasks' } })} 
                isParentDashboard={true}
                actionButtons={
                  !task.assignedChildId && !(new Date(task.deadline) < new Date()) ? (
                    <button 
                      className="px-3 py-1 bg-primary-600 text-white rounded-md text-sm hover:bg-primary-700"
                      onClick={(e) => {
                        e.stopPropagation();
                        handleAssignClick(task);
                      }}
                    >
                      Assign
                    </button>
                  ) : null
                }
              />
            </div>
          ))}
        </div>
      )}

      {/* Assign Task Modal */}
      <Modal
        isOpen={showAssignModal}
        onClose={() => setShowAssignModal(false)}
        title="Assign Task"
      >
        <div className="p-4">
          <p className="mb-4 font-medium">{selectedTask?.title}</p>
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Select Child
            </label>
            <select
              value={selectedChildId}
              onChange={(e) => setSelectedChildId(e.target.value)}
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-primary-500 focus:border-primary-500"
              disabled={isLoading}
            >
              <option value="">-- Please Select --</option>
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
              disabled={isLoading}
            >
              Cancel
            </button>
            <button 
              className="px-4 py-2 bg-primary-600 text-white rounded-md hover:bg-primary-700 disabled:bg-gray-400"
              onClick={handleAssignTask}
              disabled={!selectedChildId || isLoading}
            >
              {isLoading ? 'Processing...' : 'Confirm Assignment'}
            </button>
          </div>
        </div>
      </Modal>
    </div>
  );
};

export default TaskList;