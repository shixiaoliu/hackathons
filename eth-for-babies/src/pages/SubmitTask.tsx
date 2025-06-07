import { useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { ChevronLeft, Upload, X, Check } from 'lucide-react';
import Button from '../components/common/Button';
import Card, { CardBody, CardFooter } from '../components/common/Card';
import { mockTasks } from '../data/mockTasks';

const SubmitTask = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const task = mockTasks.find(task => task.id === id);
  
  const [description, setDescription] = useState('');
  const [images, setImages] = useState<string[]>([]);
  const [isSubmitting, setIsSubmitting] = useState(false);
  
  if (!task) {
    return (
      <div className="max-w-3xl mx-auto text-center py-12">
        <h2 className="text-2xl font-bold text-gray-900 mb-4">Task Not Found</h2>
        <p className="text-gray-600 mb-6">The task you're looking for doesn't exist or has been removed.</p>
        <Button onClick={() => navigate('/child')}>
          Return to Dashboard
        </Button>
      </div>
    );
  }
  
  const handleImageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files;
    if (files) {
      const newImages: string[] = [];
      
      for (let i = 0; i < files.length; i++) {
        const file = files[i];
        const reader = new FileReader();
        reader.onloadend = () => {
          newImages.push(reader.result as string);
          if (newImages.length === files.length) {
            setImages(prev => [...prev, ...newImages]);
          }
        };
        reader.readAsDataURL(file);
      }
    }
  };
  
  const removeImage = (index: number) => {
    setImages(prev => prev.filter((_, i) => i !== index));
  };
  
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);
    
    // Simulate submission
    setTimeout(() => {
      console.log('Submitting task completion:', {
        taskId: id,
        description,
        images,
      });
      
      setIsSubmitting(false);
      navigate('/child');
    }, 1500);
  };

  return (
    <div className="max-w-2xl mx-auto">
      <div className="flex items-center mb-6">
        <button
          onClick={() => navigate(`/task/${id}`)}
          className="flex items-center text-gray-600 hover:text-gray-900"
        >
          <ChevronLeft className="h-5 w-5 mr-1" />
          Back to task
        </button>
      </div>
      
      <div className="mb-6">
        <h1 className="text-2xl font-bold text-gray-900 mb-2">Submit Task Completion</h1>
        <p className="text-gray-600">Provide proof that you've completed: <span className="font-medium">{task.title}</span></p>
      </div>
      
      <Card>
        <form onSubmit={handleSubmit}>
          <CardBody>
            <div className="space-y-6">
              <div>
                <label htmlFor="description" className="block text-sm font-medium text-gray-700 mb-1">
                  Completion Description*
                </label>
                <textarea
                  id="description"
                  value={description}
                  onChange={(e) => setDescription(e.target.value)}
                  rows={4}
                  className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                  placeholder="Describe how you completed the task and any challenges you faced"
                  required
                ></textarea>
              </div>
              
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Proof Images*
                </label>
                
                <div className="grid grid-cols-2 sm:grid-cols-3 gap-4 mb-4">
                  {images.map((image, index) => (
                    <div key={index} className="relative rounded-md overflow-hidden h-32">
                      <img
                        src={image}
                        alt={`Proof ${index + 1}`}
                        className="w-full h-full object-cover"
                      />
                      <button
                        type="button"
                        onClick={() => removeImage(index)}
                        className="absolute top-2 right-2 p-1 bg-white rounded-full shadow-md"
                      >
                        <X className="h-4 w-4 text-gray-600" />
                      </button>
                    </div>
                  ))}
                  
                  {images.length < 3 && (
                    <div className="flex items-center justify-center h-32 border-2 border-dashed border-gray-300 rounded-md">
                      <label className="flex flex-col items-center justify-center cursor-pointer">
                        <Upload className="h-8 w-8 text-gray-400" />
                        <span className="mt-2 text-xs text-gray-500">Add Image</span>
                        <input
                          type="file"
                          className="hidden"
                          accept="image/*"
                          onChange={handleImageChange}
                        />
                      </label>
                    </div>
                  )}
                </div>
                
                {images.length === 0 && (
                  <div className="flex justify-center px-6 py-8 border-2 border-dashed border-gray-300 rounded-md">
                    <div className="text-center">
                      <Upload className="mx-auto h-10 w-10 text-gray-400" />
                      <div className="mt-2 flex text-sm text-gray-600">
                        <label
                          htmlFor="images"
                          className="relative cursor-pointer bg-white rounded-md font-medium text-primary-600 hover:text-primary-500"
                        >
                          <span>Upload images</span>
                          <input 
                            id="images" 
                            type="file" 
                            multiple
                            className="sr-only" 
                            accept="image/*"
                            onChange={handleImageChange}
                            required={images.length === 0}
                          />
                        </label>
                        <p className="pl-1">or drag and drop</p>
                      </div>
                      <p className="text-xs text-gray-500">
                        PNG, JPG, GIF up to 5MB each
                      </p>
                    </div>
                  </div>
                )}
                
                <p className="mt-2 text-xs text-gray-500">
                  Upload at least one image as proof of task completion. Clear photos help ensure quick approval.
                </p>
              </div>
              
              <div className="bg-yellow-50 border-l-4 border-yellow-400 p-4 rounded-md">
                <div className="flex">
                  <div className="flex-shrink-0">
                    <Check className="h-5 w-5 text-yellow-400" />
                  </div>
                  <div className="ml-3">
                    <h3 className="text-sm font-medium text-yellow-800">Reminder: Completion Criteria</h3>
                    <div className="mt-2 text-sm text-yellow-700">
                      <p>{task.completionCriteria}</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </CardBody>
          
          <CardFooter className="flex justify-between">
            <Button 
              type="button"
              variant="outline" 
              onClick={() => navigate(`/task/${id}`)}
            >
              Cancel
            </Button>
            <Button 
              type="submit"
              isLoading={isSubmitting}
              disabled={description.trim() === '' || images.length === 0}
            >
              Submit for Approval
            </Button>
          </CardFooter>
        </form>
      </Card>
    </div>
  );
};

export default SubmitTask;