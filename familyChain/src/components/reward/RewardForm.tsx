import React, { useState } from 'react';
import { X, ImagePlus, Loader2 } from 'lucide-react';
import Button from '../common/Button';
import { Reward } from '../../types/reward';
import { taskApi } from '../../services/api';

interface RewardFormProps {
  initialData?: Reward;
  onSubmit: (data: {
    name: string;
    description: string;
    image_url: string;
    token_price: number;
    stock: number;
    create_on_blockchain?: boolean;
  }) => void;
  onCancel: () => void;
  isLoading?: boolean;
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
  const [isUploadingImage, setIsUploadingImage] = useState(false);

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
    
    // 图片验证 - 如果没有预览图片，则使用默认图片
    if (!previewImage && !imageUrl) {
      // 不再将缺少图片视为错误，将使用默认图片
      console.log('将使用默认图片');
    }
    
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  // 生成默认图片的数据URL
  const generateDefaultImageDataUrl = (text: string): string => {
    try {
      const svg = `<svg xmlns="http://www.w3.org/2000/svg" width="400" height="300" viewBox="0 0 400 300">
        <rect width="400" height="300" fill="#f0f0f0"/>
        <text x="200" y="150" font-family="Arial" font-size="24" text-anchor="middle" fill="#888888">${text || '奖品'}</text>
      </svg>`;
      return `data:image/svg+xml;base64,${safeBase64Encode(svg)}`;
    } catch (error) {
      console.error('生成默认图片失败:', error);
      // 提供一个极简的备用数据URL
      return 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg==';
    }
  };

  // 提交表单
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!validate()) {
      return;
    }
    
    // 使用上传到服务器的图片URL，如果没有则使用默认图片
    const finalImageUrl = imageUrl || generateDefaultImageDataUrl(name);
    
    onSubmit({
      name,
      description,
      image_url: finalImageUrl,
      token_price: tokenPrice,
      stock: 1,  // 库存固定为1
      create_on_blockchain: true // 默认在区块链上创建
    });
  };

  // 处理图片上传
  const handleImageUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;
    
    // 确保文件是图片
    if (!file.type.match('image.*')) {
      setErrors({...errors, imageUrl: '请上传图片文件'});
      return;
    }
    
    try {
      // 显示上传状态
      setIsUploadingImage(true);
      
      // 创建本地预览
      const reader = new FileReader();
      reader.onloadend = () => {
        setPreviewImage(reader.result as string);
      };
      reader.readAsDataURL(file);
      
      // 上传图片到服务器
      console.log('开始上传图片到服务器...');
      const response = await taskApi.uploadImage(file);
      
      if (response.success && response.data) {
        console.log('图片上传成功，URL:', response.data.url);
        // 保存服务器返回的URL
        setImageUrl(response.data.url);
      } else {
        console.error('图片上传失败:', response.error);
        setErrors({...errors, imageUrl: `图片上传失败: ${response.error || '未知错误'}`});
        // 不清除预览，让用户看到上传的是什么图片
      }
    } catch (error) {
      console.error('图片上传过程中出错:', error);
      setErrors({...errors, imageUrl: '图片上传失败，请重试'});
    } finally {
      setIsUploadingImage(false);
    }
    
    // 清除图片相关错误
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
          Reward Name<span className="text-red-500">*</span>
        </label>
        <input
          type="text"
          id="name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          className={`w-full px-3 py-2 border rounded-md ${
            errors.name ? 'border-red-500' : 'border-gray-300'
          } focus:outline-none focus:ring-1 focus:ring-primary-500`}
          placeholder="Enter reward name"
          disabled={isLoading}
        />
        {errors.name && <p className="mt-1 text-sm text-red-500">{errors.name}</p>}
      </div>
      
      {/* 奖品描述 */}
      <div>
        <label htmlFor="description" className="block text-sm font-medium text-gray-700 mb-1">
          Reward Description
        </label>
        <textarea
          id="description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          rows={3}
          className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-primary-500"
          placeholder="Enter reward description"
          disabled={isLoading}
        ></textarea>
      </div>
      
      {/* 奖品图片 */}
      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Reward Image<span className="text-red-500">*</span>
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
              {isUploadingImage && (
                <div className="absolute inset-0 flex items-center justify-center bg-black bg-opacity-50 rounded-md">
                  <Loader2 className="h-8 w-8 text-white animate-spin" />
                  <span className="text-white ml-2">上传中...</span>
                </div>
              )}
              <button
                type="button"
                onClick={clearImage}
                className="absolute top-2 right-2 p-1 bg-white rounded-full shadow-md hover:bg-gray-100"
                disabled={isLoading || isUploadingImage}
              >
                <X className="h-5 w-5 text-gray-500" />
              </button>
            </div>
          ) : (
            <div className="w-full h-48 border-2 border-dashed border-gray-300 rounded-md flex flex-col items-center justify-center bg-gray-50">
              <ImagePlus className="h-12 w-12 text-gray-400" />
              <p className="mt-2 text-sm text-gray-500">Click to upload image</p>
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
                disabled={isLoading || isUploadingImage}
              />
              <label
                htmlFor="image-upload"
                className={`w-full flex justify-center py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium ${isUploadingImage || isLoading ? 'bg-gray-100 text-gray-500 cursor-not-allowed' : 'bg-white text-gray-700 cursor-pointer hover:bg-gray-50'}`}
              >
                {isUploadingImage ? 'Uploading...' : 'Choose Image'}
              </label>
            </div>
          </div>
          
          {errors.imageUrl && (
            <p className="text-sm text-red-500 self-start">{errors.imageUrl}</p>
          )}
          
          {/* 添加提示信息 */}
          {imageUrl && (
            <p className="text-xs text-green-600 self-start">
              图片已成功上传到服务器
            </p>
          )}
          <p className="text-xs text-gray-500 self-start">
            Recommended image size: 400x300px, max 5MB
          </p>
        </div>
      </div>
      
      {/* 代币价格 */}
      <div>
        <label htmlFor="token-price" className="block text-sm font-medium text-gray-700 mb-1">
          Token Price<span className="text-red-500">*</span>
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
            <span className="text-gray-500">FCT</span>
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
          Cancel
        </Button>
        <Button
          type="submit"
          isLoading={isLoading}
          disabled={isLoading || isUploadingImage}
        >
          {isLoading ? (initialData ? "Updating..." : "Creating...") : (initialData ? 'Update Reward' : 'Create Reward')}
        </Button>
      </div>
    </form>
  );
};

export default RewardForm; 