import React, { useState } from 'react';
import { Clock, Award, AlertTriangle, Image as ImageIcon } from 'lucide-react';
import { Task } from '../../types/task';
import Card, { CardBody, CardFooter } from '../common/Card';
import Button from '../common/Button';
import { formatDistanceToNow } from '../../utils/dateUtils';

interface TaskCardProps {
  task: Task;
  onClick?: () => void;
  onTakeTask?: () => void;
  onCompleteTask?: () => void;
  showTakeButton?: boolean;
  actionButtons?: React.ReactNode;
  isParentDashboard?: boolean;
}

const TaskCard = ({ task, onClick, onTakeTask, onCompleteTask, showTakeButton = false, actionButtons, isParentDashboard = false }: TaskCardProps) => {
  const [imageError, setImageError] = useState(false);
  
  // Check if the task is past its deadline regardless of status
  const isOverdue = new Date(task.deadline) < new Date();
  
  // Only show as expired if it's open or in-progress and past deadline
  const isExpired = isOverdue && (task.status === 'open' || task.status === 'in-progress');
  
  // Check if task is in a final state (completed, approved, rejected)
  const isFinalized = ['completed', 'approved', 'rejected'].includes(task.status);
  
  const statusColors = {
    'open': 'bg-blue-100 text-blue-800',
    'in-progress': 'bg-yellow-100 text-yellow-800',
    'completed': 'bg-orange-100 text-orange-800',
    'approved': 'bg-green-100 text-green-800',
    'rejected': 'bg-red-100 text-red-800',
    'expired': 'bg-gray-100 text-gray-800',
  };
  
  const difficultyColors = {
    'easy': 'bg-green-100 text-green-800',
    'medium': 'bg-yellow-100 text-yellow-800',
    'hard': 'bg-red-100 text-red-800',
  };

  // Determine the effective status to display
  const displayStatus = isExpired ? 'expired' : task.status;
  
  // Handle image loading error
  const handleImageError = () => {
    setImageError(true);
  };

  // 默认任务图片（使用Data URI直接嵌入SVG）
  const defaultTaskImages = {
    'easy': 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="400" height="200" viewBox="0 0 400 200"%3E%3Crect width="400" height="200" fill="%23e6ffec"/%3E%3Ctext x="50%25" y="50%25" dominant-baseline="middle" text-anchor="middle" font-family="Arial" font-size="24" fill="%2322c55e"%3EEasy Task%3C/text%3E%3C/svg%3E',
    'medium': 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="400" height="200" viewBox="0 0 400 200"%3E%3Crect width="400" height="200" fill="%23fff7e6"/%3E%3Ctext x="50%25" y="50%25" dominant-baseline="middle" text-anchor="middle" font-family="Arial" font-size="24" fill="%23f59e0b"%3EMedium Task%3C/text%3E%3C/svg%3E',
    'hard': 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="400" height="200" viewBox="0 0 400 200"%3E%3Crect width="400" height="200" fill="%23fee2e2"/%3E%3Ctext x="50%25" y="50%25" dominant-baseline="middle" text-anchor="middle" font-family="Arial" font-size="24" fill="%23ef4444"%3EHard Task%3C/text%3E%3C/svg%3E',
  };

  // 根据任务难度获取默认图片
  const getDefaultImage = () => {
    return defaultTaskImages[task.difficulty] || 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="400" height="200" viewBox="0 0 400 200"%3E%3Crect width="400" height="200" fill="%23f3f4f6"/%3E%3Ctext x="50%25" y="50%25" dominant-baseline="middle" text-anchor="middle" font-family="Arial" font-size="24" fill="%239ca3af"%3EDefault Task%3C/text%3E%3C/svg%3E';
  };

  return (
    <Card 
      hoverable={!isExpired} 
      onClick={isExpired ? undefined : onClick} 
      className={`h-full transition-all duration-300 ${isExpired ? 'border-red-300 border opacity-70' : ''}`}
    >
      {/* 图片区域 - 始终显示 */}
      <div className="relative w-full h-40">
        {task.imageUrl && !imageError ? (
          <img 
            src={task.imageUrl} 
            alt={task.title} 
            className="w-full h-full object-cover"
            onError={handleImageError}
            style={{ position: 'relative', zIndex: 1 }}
          />
        ) : (
          <img
            src={getDefaultImage()}
            alt={`${task.difficulty} task`}
            className="w-full h-full object-cover"
            onError={(e) => {
              // 如果默认图片也加载失败，显示图标
              e.currentTarget.style.display = 'none';
              const iconContainer = e.currentTarget.parentElement;
              if (iconContainer) {
                const iconDiv = document.createElement('div');
                iconDiv.className = 'w-full h-full flex items-center justify-center bg-gray-100';
                iconDiv.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" class="h-10 w-10 text-gray-300" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" /></svg>';
                iconContainer.appendChild(iconDiv);
              }
            }}
          />
        )}
      </div>
      
      <CardBody className="pb-2">
        <div className="flex justify-between items-start mb-2">
          <h3 className="text-lg font-semibold text-gray-900 truncate">{task.title}</h3>
          <span className={`text-xs font-medium px-2 py-1 rounded-full ${difficultyColors[task.difficulty]}`}>
            {task.difficulty.charAt(0).toUpperCase() + task.difficulty.slice(1)}
          </span>
        </div>
        
        <p className="text-gray-600 text-sm mb-4 line-clamp-2">{task.description}</p>
        
        <div className="flex items-center text-sm text-gray-500 mb-2">
          <Clock className="h-4 w-4 mr-1" />
          <span className={isOverdue && !isFinalized ? 'text-red-500 font-medium' : ''}>
            {isOverdue && !isFinalized
              ? 'Overdue by ' + formatDistanceToNow(new Date(task.deadline))
              : formatDistanceToNow(new Date(task.deadline)) + (isFinalized ? '' : ' left')}
          </span>
        </div>
        
        <div className="flex items-center text-sm text-gray-500">
          <Award className="h-4 w-4 mr-1" />
          <span className="font-medium text-primary-600">{task.reward} ETH</span>
        </div>
      </CardBody>
      
      <CardFooter className="flex flex-col py-3 bg-gray-50">
        <div className="flex justify-between items-center w-full">
          <span className={`text-xs font-medium px-2 py-1 rounded-full ${statusColors[displayStatus]}`}>
            {displayStatus.charAt(0).toUpperCase() + displayStatus.slice(1)}
          </span>
          
          <div className="flex items-center gap-2">
            {showTakeButton && task.status === 'open' && !isExpired && onTakeTask && (
              <Button 
                size="sm" 
                onClick={(e) => {
                  e.stopPropagation();
                  onTakeTask();
                }}
              >
                Take Task
              </Button>
            )}
            
            {!isParentDashboard && task.status === 'in-progress' && !isExpired && task.assignedTo && (
              <Button 
                size="sm"
                variant="success" 
                onClick={(e) => {
                  e.stopPropagation();
                  if (typeof onCompleteTask === 'function') {
                    onCompleteTask();
                  } else if (onClick) {
                    onClick();
                  }
                }}
              >
                Complete
              </Button>
            )}
            
            {actionButtons}
          </div>
        </div>
      </CardFooter>
    </Card>
  );
};

export default TaskCard;