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
  showTakeButton?: boolean;
}

const TaskCard = ({ task, onClick, onTakeTask, showTakeButton = false }: TaskCardProps) => {
  const isOverdue = new Date(task.deadline) < new Date() && task.status === 'open';
  
  const statusColors = {
    'open': 'bg-blue-100 text-blue-800',
    'in-progress': 'bg-yellow-100 text-yellow-800',
    'completed': 'bg-orange-100 text-orange-800',
    'approved': 'bg-green-100 text-green-800',
    'rejected': 'bg-red-100 text-red-800',
  };
  
  const difficultyColors = {
    'easy': 'bg-green-100 text-green-800',
    'medium': 'bg-yellow-100 text-yellow-800',
    'hard': 'bg-red-100 text-red-800',
  };

  return (
    <Card 
      hoverable 
      onClick={onClick} 
      className={`h-full transition-all duration-300 ${isOverdue ? 'border-red-300 border' : ''}`}
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
          <span className={isOverdue ? 'text-red-500 font-medium' : ''}>
            {isOverdue 
              ? 'Overdue by ' + formatDistanceToNow(new Date(task.deadline))
              : formatDistanceToNow(new Date(task.deadline)) + ' left'}
          </span>
        </div>
        
        <div className="flex items-center text-sm text-gray-500">
          <Award className="h-4 w-4 mr-1" />
          <span className="font-medium text-primary-600">{task.reward} ETH</span>
        </div>
      </CardBody>
      
      <CardFooter className="flex justify-between items-center py-3 bg-gray-50">
        <span className={`text-xs font-medium px-2 py-1 rounded-full ${statusColors[task.status]}`}>
          {task.status.charAt(0).toUpperCase() + task.status.slice(1)}
        </span>
        
        <div className="flex items-center gap-2">
          {showTakeButton && task.status === 'open' && onTakeTask && (
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
          
          {isOverdue && (
            <div className="flex items-center text-xs text-red-500">
              <AlertTriangle className="h-3 w-3 mr-1" />
              Overdue
            </div>
          )}
        </div>
      </CardFooter>
    </Card>
  );
};

export default TaskCard;