import React, { useState, useEffect } from 'react';
import Card, { CardBody } from '../common/Card';
import { Reward } from '../../types/reward';
import { Edit, Trash2, AlertCircle, Image as ImageIcon } from 'lucide-react';
import Button from '../common/Button';

interface RewardCardProps {
  reward: Reward;
  onEdit?: () => void;
  onDelete?: () => void;
}

// 安全的Base64编码函数，支持Unicode字符
const safeBase64Encode = (str: string): string => {
  try {
    // 对于现代浏览器，使用内置的编码API
    return btoa(encodeURIComponent(str).replace(/%([0-9A-F]{2})/g, (_, p1) => 
      String.fromCharCode(parseInt(p1, 16))
    ));
  } catch (e) {
    console.error('编码失败:', e);
    return '';
  }
};

const RewardCard: React.FC<RewardCardProps> = ({ reward, onEdit, onDelete }) => {
  // 使用内联SVG数据URL作为默认图片，避免外部依赖
  const defaultImage = `data:image/svg+xml;base64,${safeBase64Encode('<svg xmlns="http://www.w3.org/2000/svg" width="300" height="200" viewBox="0 0 300 200"><rect width="300" height="200" fill="#f0f0f0"/><text x="150" y="100" font-family="Arial" font-size="24" text-anchor="middle" fill="#888888">' + (reward.name || '奖品') + '</text></svg>')}`;
  
  // 图片状态
  const [imageUrl, setImageUrl] = useState<string>(reward.image_url || defaultImage);
  const [imageError, setImageError] = useState<boolean>(false);
  
  // 当reward变化时更新图片URL
  useEffect(() => {
    if (reward.image_url) {
      setImageUrl(reward.image_url);
      setImageError(false);
    }
  }, [reward.image_url]);
  
  // 处理图片URL
  const processImageUrl = (url: string): string => {
    // 如果是base64编码的图片，直接使用
    if (url.startsWith('data:image')) {
      return url;
    }
    
    // 不再使用placeholder.com的URL
    return url;
  };
  
  // 处理图片加载错误
  const handleImageError = () => {
    console.log('图片加载失败:', imageUrl);
    setImageError(true);
    setImageUrl(defaultImage);
  };

  return (
    <Card 
      className="h-full flex flex-col transition-all duration-200 hover:shadow-lg"
      hoverable={false}
    >
      {/* 奖品图片 */}
      <div className="relative w-full h-48 bg-gray-100 overflow-hidden flex items-center justify-center">
        {imageError ? (
          <div className="flex flex-col items-center justify-center text-gray-400">
            <ImageIcon className="h-12 w-12 mb-2" />
            <span className="text-sm">{reward.name}</span>
          </div>
        ) : (
          <img 
            src={processImageUrl(imageUrl)} 
            alt={reward.name} 
            className="w-full h-full object-contain"
            onError={handleImageError}
            loading="lazy"
          />
        )}
        
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
            {Math.floor(reward.token_price)} FCT
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