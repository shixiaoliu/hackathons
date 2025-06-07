import { useParams, useNavigate } from 'react-router-dom';
import { useUserRole } from '../context/UserRoleContext';
import { Clock, Award, ChevronLeft, Calendar, User, CheckCircle, AlertTriangle } from 'lucide-react';
import Button from '../components/common/Button';
import Card, { CardBody, CardHeader, CardFooter } from '../components/common/Card';
import { mockTasks } from '../data/mockTasks';
import { formatDateTime, formatDate } from '../utils/dateUtils';

const TaskDetail = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const { userRole } = useUserRole();
  
  const task = mockTasks.find(task => task.id === id);
  
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
  const isChild = userRole === 'child';
  const isOverdue = new Date(task.deadline) < new Date() && task.status === 'open';
  
  const canAcceptTask = isChild && task.status === 'open';
  const canSubmitTask = isChild && task.status === 'in-progress' && task.assignedTo;

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
        {task.imageUrl && (
          <div className="h-64 overflow-hidden">
            <img 
              src={task.imageUrl} 
              alt={task.title} 
              className="w-full h-full object-cover"
            />
          </div>
        )}
        
        <CardHeader>
          <div className="flex flex-col md:flex-row md:items-center md:justify-between">
            <h1 className="text-2xl font-bold text-gray-900">{task.title}</h1>
            
            <div className="mt-2 md:mt-0 flex items-center">
              <span className={`px-3 py-1 rounded-full text-sm font-medium
                ${task.status === 'open' ? 'bg-blue-100 text-blue-800' : ''}
                ${task.status === 'in-progress' ? 'bg-yellow-100 text-yellow-800' : ''}
                ${task.status === 'completed' ? 'bg-orange-100 text-orange-800' : ''}
                ${task.status === 'approved' ? 'bg-green-100 text-green-800' : ''}
                ${task.status === 'rejected' ? 'bg-red-100 text-red-800' : ''}
              `}>
                {task.status.charAt(0).toUpperCase() + task.status.slice(1)}
              </span>
              
              <span className="ml-2 px-3 py-1 rounded-full text-sm font-medium
                ${task.difficulty === 'easy' ? 'bg-green-100 text-green-800' : ''}
                ${task.difficulty === 'medium' ? 'bg-yellow-100 text-yellow-800' : ''}
                ${task.difficulty === 'hard' ? 'bg-red-100 text-red-800' : ''}
              ">
                {task.difficulty.charAt(0).toUpperCase() + task.difficulty.slice(1)}
              </span>
            </div>
          </div>
        </CardHeader>
        
        <CardBody>
          <div className="space-y-6">
            <div>
              <h2 className="text-lg font-semibold text-gray-900 mb-2">Description</h2>
              <p className="text-gray-700">{task.description}</p>
            </div>
            
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div className="flex items-center">
                <Calendar className="h-5 w-5 text-gray-500 mr-2" />
                <div>
                  <p className="text-sm text-gray-500">Deadline</p>
                  <p className={`text-sm font-medium ${isOverdue ? 'text-red-600' : 'text-gray-900'}`}>
                    {formatDateTime(task.deadline)}
                    {isOverdue && ' (Overdue)'}
                  </p>
                </div>
              </div>
              
              <div className="flex items-center">
                <Award className="h-5 w-5 text-gray-500 mr-2" />
                <div>
                  <p className="text-sm text-gray-500">Reward</p>
                  <p className="text-sm font-medium text-gray-900">{task.reward} ETH</p>
                </div>
              </div>
              
              <div className="flex items-center">
                <User className="h-5 w-5 text-gray-500 mr-2" />
                <div>
                  <p className="text-sm text-gray-500">Created By</p>
                  <p className="text-sm font-medium text-gray-900">
                    {task.createdBy.substring(0, 6)}...{task.createdBy.substring(task.createdBy.length - 4)}
                  </p>
                </div>
              </div>
              
              <div className="flex items-center">
                <Clock className="h-5 w-5 text-gray-500 mr-2" />
                <div>
                  <p className="text-sm text-gray-500">Created On</p>
                  <p className="text-sm font-medium text-gray-900">{formatDate(task.createdAt)}</p>
                </div>
              </div>
            </div>
            
            <div>
              <h2 className="text-lg font-semibold text-gray-900 mb-2">Completion Criteria</h2>
              <div className="bg-gray-50 p-4 rounded-md">
                <p className="text-gray-700">{task.completionCriteria}</p>
              </div>
            </div>
            
            {task.status === 'completed' && task.proof && (
              <div>
                <h2 className="text-lg font-semibold text-gray-900 mb-2">Submission Proof</h2>
                <div className="bg-gray-50 p-4 rounded-md">
                  <p className="text-gray-700 mb-4">{task.proof.description}</p>
                  {task.proof.images.length > 0 && (
                    <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                      {task.proof.images.map((image, index) => (
                        <img 
                          key={index} 
                          src={image} 
                          alt={`Proof ${index + 1}`} 
                          className="rounded-md w-full h-48 object-cover"
                        />
                      ))}
                    </div>
                  )}
                  <p className="text-sm text-gray-500 mt-4">
                    Submitted on {formatDateTime(task.proof.submittedAt)}
                  </p>
                </div>
              </div>
            )}
            
            {(task.status === 'approved' || task.status === 'rejected') && (
              <div>
                <h2 className="text-lg font-semibold text-gray-900 mb-2">
                  {task.status === 'approved' ? 'Approval' : 'Rejection'} Details
                </h2>
                <div className={`p-4 rounded-md ${
                  task.status === 'approved' ? 'bg-green-50' : 'bg-red-50'
                }`}>
                  <div className="flex items-start">
                    {task.status === 'approved' ? (
                      <CheckCircle className="h-5 w-5 text-green-500 mr-2 mt-0.5" />
                    ) : (
                      <AlertTriangle className="h-5 w-5 text-red-500 mr-2 mt-0.5" />
                    )}
                    <div>
                      <p className={`font-medium ${
                        task.status === 'approved' ? 'text-green-800' : 'text-red-800'
                      }`}>
                        {task.status === 'approved' 
                          ? 'Task approved and reward sent!' 
                          : 'Task was rejected'}
                      </p>
                      <p className="text-sm mt-1">
                        {task.status === 'approved'
                          ? `Reward of ${task.reward} ETH has been transferred to your wallet.`
                          : 'The submission did not meet the required criteria. Please try again with improvements.'}
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            )}
          </div>
        </CardBody>
        
        <CardFooter className="flex justify-between">
          <Button 
            variant="outline" 
            onClick={() => navigate(isParent ? '/parent' : '/child')}
          >
            Back
          </Button>
          
          <div className="flex space-x-3">
            {canAcceptTask && (
              <Button onClick={() => {
                // In a real app, this would assign the task to the user
                console.log('Accepting task:', task.id);
                alert('Task accepted! You can now submit it when completed.');
              }}>
                Accept Task
              </Button>
            )}
            
            {canSubmitTask && (
              <Button onClick={() => navigate(`/submit-task/${task.id}`)}>
                Submit Proof
              </Button>
            )}
            
            {isParent && task.status === 'completed' && (
              <Button onClick={() => navigate(`/parent/review/${task.id}`)}>
                Review Submission
              </Button>
            )}
          </div>
        </CardFooter>
      </Card>
    </div>
  );
};

export default TaskDetail;