import React, { useState } from 'react';
import { X, ImagePlus } from 'lucide-react';
import Button from '../common/Button';
import { Reward } from '../../types/reward';

interface RewardFormProps {
  initialData?: Reward;
  onSubmit: (data: {
    name: string;
    description: string;
    image_url: string;
    token_price: number;
    stock: number;
  }) => void;
  onCancel: () => void;
  isLoading?: boolean;
}

const RewardForm: React.FC<RewardFormProps> = ({
  initialData,
  onSubmit,
  onCancel,
  isLoading = false
}) => {
  const [name, setName] = useState(initialData?.name || '');
  const [description, setDescription] = useState(initialData?.description || '');
  const [imageUrl, setImageUrl] = useState(initialData?.image_url || '');
  const [tokenPrice, setTokenPrice] = useState(initialData?.token_price || 10);
  const [stock, setStock] = useState(initialData?.stock || 1);
  const [previewImage, setPreviewImage] = useState<string | null>(initialData?.image_url || null);
  const [errors, setErrors] = useState<Record<string, string>>({});

      // 表单验证
  const validate = (): boolean => {
    const newErrors: Record<string, string> = {};
    
    if (!name.trim()) {
      newErrors.name = '名称不能为空';
    }
    
    if (!tokenPrice || tokenPrice <= 0) {
      newErrors.tokenPrice = '代币价格必须大于0';
    }
    
    // 库存固定为1，移除验证
        
    // 恢复图片URL验证
    if (!imageUrl.trim()) {
      newErrors.imageUrl = '请上传图片或提供图片URL';
    }
    
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  // 提交表单
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!validate()) {
      return;
    }
    
    // 使用用户选择的图片或默认图片
    const finalImageUrl = imageUrl || 'https://via.placeholder.com/400x300?text=%E5%A5%96%E5%93%81';
    
    onSubmit({
      name,
      description,
      image_url: finalImageUrl,
      token_price: tokenPrice,
      stock: 1  // 库存固定为1
    });
  };

  // 处理图片上传
  const handleImageUpload = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;
    
    // 确保文件是图片
    if (!file.type.match('image.*')) {
      setErrors({...errors, imageUrl: '请上传图片文件'});
      return;
    }
    
    // 文件大小限制（2MB）
    if (file.size > 2 * 1024 * 1024) {
      setErrors({...errors, imageUrl: '图片大小不能超过2MB'});
      return;
    }
    
    // 创建预览URL
    const reader = new FileReader();
    reader.onload = (e) => {
      const result = e.target?.result as string;
      setPreviewImage(result);
      setImageUrl(result); // Base64编码的图片数据
    };
    reader.readAsDataURL(file);
    
    // 清除错误
    if (errors.imageUrl) {
      const { imageUrl, ...restErrors } = errors;
      setErrors(restErrors);
    }
  };

  // 处理外部图片URL
  const handleExternalImageUrl = (url: string) => {
    setImageUrl(url);
    setPreviewImage(url);
    
    // 清除错误
    if (errors.imageUrl) {
      const { imageUrl, ...restErrors } = errors;
      setErrors(restErrors);
    }
  };

  // 清除图片
  const clearImage = () => {
    setImageUrl('');
    setPreviewImage(null);
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-6">
      {/* 奖品名称 */}
      <div>
        <label htmlFor="name" className="block text-sm font-medium text-gray-700 mb-1">
          奖品名称<span className="text-red-500">*</span>
        </label>
        <input
          type="text"
          id="name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          className={`w-full px-3 py-2 border rounded-md ${
            errors.name ? 'border-red-500' : 'border-gray-300'
          } focus:outline-none focus:ring-1 focus:ring-primary-500`}
          placeholder="输入奖品名称"
          disabled={isLoading}
        />
        {errors.name && <p className="mt-1 text-sm text-red-500">{errors.name}</p>}
      </div>
      
      {/* 奖品描述 */}
      <div>
        <label htmlFor="description" className="block text-sm font-medium text-gray-700 mb-1">
          奖品描述
        </label>
        <textarea
          id="description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          rows={3}
          className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-primary-500"
          placeholder="输入奖品描述"
          disabled={isLoading}
        ></textarea>
      </div>
      
      {/* 奖品图片 */}
      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          奖品图片<span className="text-red-500">*</span>
        </label>
        
        <div className="mt-1 flex flex-col items-center space-y-2">
          {/* 预览区域 */}
          {previewImage ? (
            <div className="relative w-full h-48 bg-gray-100 rounded-md overflow-hidden flex items-center justify-center">
              <img
                src={previewImage}
                alt="奖品预览"
                className="w-full h-full object-contain"
              />
              <button
                type="button"
                onClick={clearImage}
                className="absolute top-2 right-2 p-1 bg-white rounded-full shadow-md hover:bg-gray-100"
                disabled={isLoading}
              >
                <X className="h-5 w-5 text-gray-500" />
              </button>
            </div>
          ) : (
            <div className="w-full h-48 border-2 border-dashed border-gray-300 rounded-md flex flex-col items-center justify-center bg-gray-50">
              <ImagePlus className="h-12 w-12 text-gray-400" />
              <p className="mt-2 text-sm text-gray-500">点击上传图片或输入图片URL</p>
            </div>
          )}
          
          {/* 上传控件 */}
          <div className="flex flex-col sm:flex-row w-full gap-2">
            <div className="flex-1">
              <input
                type="file"
                id="image-upload"
                accept="image/*"
                onChange={handleImageUpload}
                className="hidden"
                disabled={isLoading}
              />
              <label
                htmlFor="image-upload"
                className="w-full flex justify-center py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 cursor-pointer"
              >
                选择图片
              </label>
            </div>
            <div className="flex-1">
              <input
                type="text"
                value={imageUrl}
                onChange={(e) => handleExternalImageUrl(e.target.value)}
                placeholder="或输入图片URL"
                className={`w-full px-3 py-2 border ${
                  errors.imageUrl ? 'border-red-500' : 'border-gray-300'
                } rounded-md focus:outline-none focus:ring-1 focus:ring-primary-500`}
                disabled={isLoading}
              />
            </div>
          </div>
          
          {errors.imageUrl && (
            <p className="text-sm text-red-500 self-start">{errors.imageUrl}</p>
          )}
        </div>
      </div>
      
      {/* 代币价格 */}
      <div>
        <label htmlFor="token-price" className="block text-sm font-medium text-gray-700 mb-1">
          代币价格<span className="text-red-500">*</span>
        </label>
        <div className="relative mt-1 rounded-md shadow-sm">
          <input
            type="number"
            id="token-price"
            value={tokenPrice}
            onChange={(e) => setTokenPrice(parseInt(e.target.value) || 0)}
            className={`w-full px-3 py-2 border ${
              errors.tokenPrice ? 'border-red-500' : 'border-gray-300'
            } rounded-md focus:outline-none focus:ring-1 focus:ring-primary-500`}
            min="1"
            step="1"
            disabled={isLoading}
          />
          <div className="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
            <span className="text-gray-500">代币</span>
          </div>
        </div>
        {errors.tokenPrice && (
          <p className="mt-1 text-sm text-red-500">{errors.tokenPrice}</p>
        )}
      </div>
      
      {/* 库存固定为1，不再显示输入框 */}
      
      {/* 操作按钮 */}
      <div className="flex justify-end space-x-3 pt-4">
        <Button
          type="button"
          variant="outline"
          onClick={onCancel}
          disabled={isLoading}
        >
          取消
        </Button>
        <Button
          type="submit"
          isLoading={isLoading}
          disabled={isLoading}
        >
          {initialData ? '更新奖品' : '创建奖品'}
        </Button>
      </div>
    </form>
  );
};

export default RewardForm; 