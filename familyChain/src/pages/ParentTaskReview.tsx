import { useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { ChevronLeft, Check, X, AlertTriangle, MessageCircle } from 'lucide-react';
import Button from '../components/common/Button';
import Card, { CardBody, CardHeader, CardFooter } from '../components/common/Card';
import { mockTasks } from '../data/mockTasks';
import { formatDateTime } from '../utils/dateUtils';

const ParentTaskReview = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const task = mockTasks.find(task => task.id === id);
  
  const [feedback, setFeedback] = useState('');
  const [isApproving, setIsApproving] = useState(false);
  const [isRejecting, setIsRejecting] = useState(false);
  
  if (!task || task.status !== 'completed' || !task.proof) {
    return (
      <div className="max-w-3xl mx-auto text-center py-12">
        <h2 className="text-2xl font-bold text-gray-900 mb-4">Submission Not Found</h2>
        <p className="text-gray-600 mb-6">The task submission you're looking for doesn't exist or is not ready for review.</p>
        <Button onClick={() => navigate('/parent')}>
          Return to Dashboard
        </Button>
      </div>
    );
  }
  
  const handleApprove = () => {
    setIsApproving(true);
    // In a real app, this would interact with the smart contract
    setTimeout(() => {
      console.log('Approving task:', task.id);
      setIsApproving(false);
      navigate('/parent');
    }, 1500);
  };
  
  const handleReject = () => {
    setIsRejecting(true);
    // In a real app, this would interact with the smart contract
    setTimeout(() => {
      console.log('Rejecting task:', task.id, 'Feedback:', feedback);
      setIsRejecting(false);
      navigate('/parent');
    }, 1500);
  };

  return (
    <div className="max-w-3xl mx-auto">
      <div className="flex items-center mb-6">
        <button
          onClick={() => navigate('/parent')}
          className="flex items-center text-gray-600 hover:text-gray-900"
        >
          <ChevronLeft className="h-5 w-5 mr-1" />
          Back to dashboard
        </button>
      </div>
      
      <div className="mb-6">
        <h1 className="text-2xl font-bold text-gray-900 mb-2">Review Task Submission</h1>
        <p className="text-gray-600">
          <span className="font-medium">{task.title}</span> - Submitted for your approval
        </p>
      </div>
      
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div className="md:col-span-2">
          <Card>
            <CardHeader>
              <h2 className="text-xl font-semibold text-gray-900">Submission Details</h2>
            </CardHeader>
            
            <CardBody>
              <div className="space-y-6">
                <div>
                  <h3 className="text-sm font-medium text-gray-500 mb-1">Description</h3>
                  <p className="text-gray-900">{task.proof.description}</p>
                </div>
                
                <div>
                  <h3 className="text-sm font-medium text-gray-500 mb-2">Proof Images</h3>
                  <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    {task.proof.images.map((image, index) => (
                      <img 
                        key={index} 
                        src={image} 
                        alt={`Proof ${index + 1}`} 
                        className="rounded-md w-full object-cover"
                      />
                    ))}
                  </div>
                </div>
                
                <div className="pt-4 border-t border-gray-200">
                  <h3 className="text-sm font-medium text-gray-500 mb-1">Submitted On</h3>
                  <p className="text-gray-900">{formatDateTime(task.proof.submittedAt)}</p>
                </div>
              </div>
            </CardBody>
          </Card>
        </div>
        
        <div>
          <Card>
            <CardHeader>
              <h2 className="text-xl font-semibold text-gray-900">Task Details</h2>
            </CardHeader>
            
            <CardBody>
              <div className="space-y-4">
                <div>
                  <h3 className="text-sm font-medium text-gray-500 mb-1">Reward</h3>
                  <p className="text-lg font-semibold text-primary-600">{task.reward} ETH</p>
                </div>
                
                <div>
                  <h3 className="text-sm font-medium text-gray-500 mb-1">Difficulty</h3>
                  <p className="text-gray-900 capitalize">{task.difficulty}</p>
                </div>
                
                <div>
                  <h3 className="text-sm font-medium text-gray-500 mb-1">Completion Criteria</h3>
                  <p className="text-gray-900 text-sm">{task.completionCriteria}</p>
                </div>
              </div>
            </CardBody>
            
            <CardFooter className="flex flex-col space-y-4">
              <div className="w-full">
                <label htmlFor="feedback" className="block text-sm font-medium text-gray-700 mb-1">
                  Feedback (optional)
                </label>
                <div className="relative">
                  <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <MessageCircle className="h-5 w-5 text-gray-400" />
                  </div>
                  <textarea
                    id="feedback"
                    value={feedback}
                    onChange={(e) => setFeedback(e.target.value)}
                    rows={3}
                    className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                    placeholder="Provide feedback about the submission"
                  ></textarea>
                </div>
              </div>
              
              <div className="flex flex-col space-y-2 w-full">
                <Button
                  variant="success"
                  onClick={handleApprove}
                  isLoading={isApproving}
                  disabled={isRejecting}
                  fullWidth
                  leftIcon={<Check className="h-5 w-5" />}
                >
                  Approve & Pay
                </Button>
                
                <Button
                  variant="error"
                  onClick={handleReject}
                  isLoading={isRejecting}
                  disabled={isApproving}
                  fullWidth
                  leftIcon={<X className="h-5 w-5" />}
                >
                  Reject Submission
                </Button>
              </div>
              
              <div className="text-xs text-gray-500 mt-2">
                <div className="flex items-start">
                  <AlertTriangle className="h-4 w-4 text-yellow-500 mr-1 mt-0.5 flex-shrink-0" />
                  <p>
                    Approving will transfer {task.reward} ETH from your wallet to the child's wallet through the smart contract.
                  </p>
                </div>
              </div>
            </CardFooter>
          </Card>
        </div>
      </div>
    </div>
  );
};

export default ParentTaskReview;