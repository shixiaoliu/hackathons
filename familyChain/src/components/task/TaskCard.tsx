import React from 'react';
import { Clock, Award, AlertTriangle } from 'lucide-react';
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

  return (
    <Card 
      hoverable={!isExpired} 
      onClick={isExpired ? undefined : onClick} 
      className={`h-full transition-all duration-300 ${isExpired ? 'border-red-300 border opacity-70' : ''}`}
    >
      {task.imageUrl && (
        <div className="h-40 overflow-hidden">
          <img 
            src={task.imageUrl} 
            alt={task.title} 
            className="w-full h-full object-cover"
          />
        </div>
      )}
      
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