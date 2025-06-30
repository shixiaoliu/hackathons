import React from 'react';
import Card, { CardBody, CardFooter } from '../common/Card';
import Button from '../common/Button';
import { Exchange } from '../../types/reward';
import { CheckCircle, XCircle, Clock } from 'lucide-react';

interface ExchangeCardProps {
  exchange: Exchange;
  onApprove: () => void;
  onCancel: () => void;
}

const ExchangeCard: React.FC<ExchangeCardProps> = ({ 
  exchange,
  onApprove,
  onCancel
}) => {
  // 默认图片
  const defaultImage = 'https://via.placeholder.com/300x200?text=奖品图片';
  
  // 状态显示
  const statusDisplay = () => {
    switch (exchange.status) {
      case 'completed':
        return (
          <span className="inline-flex items-center text-green-600">
            <CheckCircle className="h-4 w-4 mr-1" />
            已批准
          </span>
        );
      case 'cancelled':
        return (
          <span className="inline-flex items-center text-red-600">
            <XCircle className="h-4 w-4 mr-1" />
            已拒绝
          </span>
        );
      case 'pending':
      default:
        return (
          <span className="inline-flex items-center text-yellow-600">
            <Clock className="h-4 w-4 mr-1" />
            待处理
          </span>
        );
    }
  };

  // 格式化日期
  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    return new Intl.DateTimeFormat('zh-CN', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit'
    }).format(date);
  };

  return (
    <Card className="transition-all duration-200 hover:shadow-lg">
      <CardBody>
        <div className="flex flex-col sm:flex-row gap-4">
          {/* 奖品图片 */}
          <div className="w-full sm:w-1/4">
            <div className="relative w-full h-32 bg-gray-100 rounded-md overflow-hidden">
              <img 
                src={exchange.reward_image || defaultImage} 
                alt={exchange.reward_name || '奖品图片'} 
                className="w-full h-full object-cover"
                onError={(e) => {
                  const target = e.target as HTMLImageElement;
                  target.src = defaultImage;
                }}
              />
            </div>
          </div>
          
          {/* 兑换信息 */}
          <div className="w-full sm:w-3/4 flex flex-col">
            <div className="flex justify-between items-start">
              <div>
                <h3 className="text-lg font-semibold text-gray-900">
                  {exchange.reward_name || '未命名奖品'}
                </h3>
                <p className="text-sm text-gray-600 mt-1">兑换者: {exchange.child_name || '未知'}</p>
              </div>
              <div className="px-3 py-1 bg-primary-100 text-primary-800 text-sm font-medium rounded-md">
                {exchange.token_amount} 代币
              </div>
            </div>
            
            {/* 兑换状态和日期 */}
            <div className="mt-4 flex flex-wrap justify-between">
              <div>
                <p className="text-sm text-gray-500">
                  兑换时间: {formatDate(exchange.exchange_date)}
                </p>
                {exchange.completed_date && (
                  <p className="text-sm text-gray-500">
                    处理时间: {formatDate(exchange.completed_date)}
                  </p>
                )}
              </div>
              <div>{statusDisplay()}</div>
            </div>
            
            {/* 备注 */}
            {exchange.notes && (
              <div className="mt-2 p-2 bg-gray-50 rounded-md">
                <p className="text-sm text-gray-600">{exchange.notes}</p>
              </div>
            )}
          </div>
        </div>
      </CardBody>
      
      {/* 操作按钮 - 只对待处理的请求显示 */}
      {exchange.status === 'pending' && (
        <CardFooter className="flex justify-end space-x-2">
          <Button 
            variant="outline" 
            className="border-red-300 text-red-600 hover:bg-red-50"
            onClick={(e) => {
              e.stopPropagation();
              onCancel();
            }}
          >
            拒绝
          </Button>
          <Button
            className="bg-green-600 hover:bg-green-700"
            onClick={(e) => {
              e.stopPropagation();
              onApprove();
            }}
          >
            批准
          </Button>
        </CardFooter>
      )}
    </Card>
  );
};

export default ExchangeCard; 