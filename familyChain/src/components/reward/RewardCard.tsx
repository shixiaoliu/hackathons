import React from 'react';
import Card, { CardBody } from '../common/Card';
import { Reward } from '../../types/reward';
import { Edit, Trash2, AlertCircle } from 'lucide-react';
import Button from '../common/Button';

interface RewardCardProps {
  reward: Reward;
  onEdit?: () => void;
  onDelete?: () => void;
}

const RewardCard: React.FC<RewardCardProps> = ({ reward, onEdit, onDelete }) => {
  // 默认图片
  const defaultImage = 'https://via.placeholder.com/300x200?text=奖品图片';

  return (
    <Card 
      className="h-full flex flex-col transition-all duration-200 hover:shadow-lg"
      hoverable={false}
    >
      {/* 奖品图片 */}
      <div className="relative w-full h-48 bg-gray-100 overflow-hidden flex items-center justify-center">
        <img 
          src={reward.image_url || defaultImage} 
          alt={reward.name} 
          className="w-full h-full object-contain"
          onError={(e) => {
            const target = e.target as HTMLImageElement;
            target.src = defaultImage;
          }}
        />
        
        {/* 唯一性标签 - 移到左上角 */}
        <div className="absolute top-0 left-0 bg-blue-500 text-white px-2 py-1 text-xs font-bold">
          Limited: 1
        </div>
        
        {/* 编辑和删除按钮 */}
        {(onEdit || onDelete) && (
          <div className="absolute top-2 right-2 flex space-x-2">
            {onEdit && (
              <button 
                onClick={(e) => {
                  e.stopPropagation();
                  onEdit();
                }}
                className="p-1 bg-white rounded-full shadow-md hover:bg-gray-100"
              >
                <Edit className="h-4 w-4 text-gray-600" />
              </button>
            )}
            
            {onDelete && (
              <button 
                onClick={(e) => {
                  e.stopPropagation();
                  onDelete();
                }}
                className="p-1 bg-white rounded-full shadow-md hover:bg-gray-100"
              >
                <Trash2 className="h-4 w-4 text-red-500" />
              </button>
            )}
          </div>
        )}
      </div>
      
      <CardBody className="flex-1 flex flex-col">
        {/* 标题和价格 */}
        <div className="mb-2 flex justify-between items-start">
          <h3 className="text-lg font-semibold text-gray-900">{reward.name}</h3>
          <div className="px-2 py-1 bg-primary-100 text-primary-800 text-sm font-medium rounded-md">
            {reward.token_price} tokens
          </div>
        </div>
        
        {/* 描述 */}
        <p className="text-gray-600 text-sm mb-4 flex-1">
          {reward.description || '暂无描述'}
        </p>
        
        {/* 底部信息栏 */}
        <div className="flex justify-between items-center text-sm text-gray-500 mt-2">
          <div className="flex items-center">
            <span>Limited: 1</span>
          </div>
          <div>
            {!reward.active && (
              <span className="inline-flex items-center text-red-600">
                <AlertCircle className="h-4 w-4 mr-1" />
                已停用
              </span>
            )}
          </div>
        </div>
      </CardBody>
    </Card>
  );
};

export default RewardCard; 