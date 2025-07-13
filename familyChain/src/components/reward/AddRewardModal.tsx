import React from 'react';
import { X } from 'lucide-react';
import RewardForm from './RewardForm';
import { RewardCreateRequest } from '../../types/reward';

interface AddRewardModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSubmit: (data: RewardCreateRequest) => void;
  isLoading?: boolean;
}

const AddRewardModal: React.FC<AddRewardModalProps> = ({
  isOpen,
  onClose,
  onSubmit,
  isLoading = false
}) => {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 overflow-y-auto">
      <div 
        className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0"
        onClick={onClose} // Close when clicking background
      >
        {/* 背景遮罩 */}
        <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
        
        {/* 模态框 */}
        <div 
          className="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg"
          onClick={(e) => e.stopPropagation()} // Prevent closing when clicking content area
        >
          {/* 标题栏 */}
          <div className="bg-gray-50 px-4 py-3 flex justify-between items-center">
            <h3 className="text-lg font-medium leading-6 text-gray-900">Add Reward</h3>
            <button 
              type="button" 
              className="text-gray-400 hover:text-gray-500"
              onClick={onClose}
            >
              <X className="h-5 w-5" />
            </button>
          </div>
          
          {/* 表单内容 */}
          <div className="px-4 pt-5 pb-4 sm:p-6">
            <RewardForm 
              onSubmit={onSubmit}
              onCancel={onClose}
              isLoading={isLoading}
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default AddRewardModal; 