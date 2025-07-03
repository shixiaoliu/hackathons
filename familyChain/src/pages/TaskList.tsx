import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useTask } from '../context/TaskContext';
import TaskCard from '../components/task/TaskCard';
import { Task } from '../types/task';
import Button from '../components/common/Button';
import { RefreshCw } from 'lucide-react';

const TaskList = () => {
  const navigate = useNavigate();
  const { getAllTasks } = useTask();
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [isRefreshing, setIsRefreshing] = useState(false);

  const fetchTasks = async () => {
    try {
      setLoading(true);
      setError(null);
      const allTasks = await getAllTasks();
      console.log('Retrieved tasks:', allTasks);
      setTasks(allTasks);
    } catch (err) {
      console.error('Failed to fetch tasks:', err);
      setError('Failed to fetch tasks, please try again later');
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

  return (
    <div className="container mx-auto p-6">
      <div className="flex flex-col md:flex-row md:items-center md:justify-between mb-8">
        <div>
          <h1 className="text-3xl font-bold text-gray-900 mb-2">All Tasks</h1>
          <p className="text-gray-600">View and manage all tasks</p>
        </div>
        
        <div className="mt-4 md:mt-0">
          <Button 
            onClick={handleRefresh}
            variant="secondary"
            leftIcon={<RefreshCw className={`h-5 w-5 ${isRefreshing ? 'animate-spin' : ''}`} />}
          >
            Refresh Data
          </Button>
        </div>
      </div>

      {loading ? (
        <div className="flex justify-center items-center h-64">
          <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary-600"></div>
        </div>
      ) : error ? (
        <div className="bg-red-50 border-l-4 border-red-400 p-4 mb-8 rounded-md">
          <div className="flex">
            <div className="ml-3">
              <p className="text-sm text-red-700">{error}</p>
            </div>
          </div>
        </div>
      ) : tasks.length === 0 ? (
        <div className="text-center py-12">
          <p className="text-gray-500 text-lg mb-4">No tasks found</p>
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {tasks.map((task) => (
            <div key={task.id} className="relative">
              <TaskCard 
                task={task} 
                onClick={() => navigate(`/task/${task.id}`, { state: { from: 'tasks' } })} 
                showTakeButton={true}
                isParentDashboard={true} // 设置为true，这样就不会显示Complete按钮
              />
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default TaskList;