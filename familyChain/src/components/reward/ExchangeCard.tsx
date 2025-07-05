import { FC } from 'react';
import { formatDistanceToNow } from '../../utils/dateUtils';
import { Exchange } from '../../services/api';
import { Check, XIcon, Clock, ShoppingBag, AlertTriangle } from 'lucide-react';
import Card, { CardBody } from '../common/Card';

interface ExchangeCardProps {
  exchange: Exchange;
  isChild?: boolean;
  onApprove?: (exchangeId: number) => void;
  onCancel?: (exchangeId: number) => void;
}

const ExchangeCard: FC<ExchangeCardProps> = ({ 
  exchange, 
  isChild = false,
  onApprove,
  onCancel
}) => {
  // 状态徽章
  const renderStatusBadge = () => {
    switch (exchange.status) {
      case 'pending':
        return (
          <div className="flex items-center px-2 py-1 rounded bg-yellow-100 text-yellow-800 text-xs font-medium">
            <Clock className="h-3 w-3 mr-1" />
            待处理
          </div>
        );
      case 'completed':
        return (
          <div className="flex items-center px-2 py-1 rounded bg-green-100 text-green-800 text-xs font-medium">
            <Check className="h-3 w-3 mr-1" />
            已完成
          </div>
        );
      case 'cancelled':
        return (
          <div className="flex items-center px-2 py-1 rounded bg-red-100 text-red-800 text-xs font-medium">
            <XIcon className="h-3 w-3 mr-1" />
            已取消
          </div>
        );
      case 'confirmed':
        return (
          <div className="flex items-center px-2 py-1 rounded bg-blue-100 text-blue-800 text-xs font-medium">
            <Check className="h-3 w-3 mr-1" />
            已确认
          </div>
        );
      case 'failed':
        return (
          <div className="flex items-center px-2 py-1 rounded bg-red-100 text-red-800 text-xs font-medium">
            <AlertTriangle className="h-3 w-3 mr-1" />
            失败
          </div>
        );
      default:
        return null;
    }
  };
  
  // 格式化时间
  const formatTime = (dateString: string) => {
    try {
      const date = new Date(dateString);
      return formatDistanceToNow(date);
    } catch (error) {
      console.error('日期格式化错误:', error);
      return '未知时间';
    }
  };

  return (
    <Card>
      <CardBody>
        <div className="flex items-center">
          <div className="h-16 w-16 rounded-lg bg-gray-100 overflow-hidden flex-shrink-0 mr-4">
            {exchange.reward_image ? (
              <img
                src={exchange.reward_image}
                alt={exchange.reward_name || '奖品图片'}
                className="w-full h-full object-contain"
              />
            ) : (
              <div className="w-full h-full flex items-center justify-center">
                <ShoppingBag className="h-6 w-6 text-gray-400" />
              </div>
            )}
          </div>
          
          <div className="flex-grow">
            <div className="flex items-center justify-between mb-1">
              <h3 className="font-medium text-gray-900">
                {exchange.reward_name || '未命名奖品'}
              </h3>
              {renderStatusBadge()}
            </div>
            
            <div className="flex items-center text-sm text-gray-500 mb-1">
              <span className="font-medium text-primary-600">{Math.floor(exchange.token_amount)} FCT</span>
              <span className="mx-2">•</span>
              <span title={new Date(exchange.exchange_date).toLocaleString()}>
                {formatTime(exchange.exchange_date)}
              </span>
            </div>
            
            {exchange.notes && (
              <p className="text-sm text-gray-600 mt-1">{exchange.notes}</p>
            )}
            
            {!isChild && exchange.status === 'pending' && (
              <div className="flex space-x-2 mt-2">
                {onApprove && (
                  <button
                    onClick={() => onApprove(exchange.id)}
                    className="text-sm px-3 py-1 bg-green-100 text-green-700 rounded hover:bg-green-200"
                  >
                    Approve
                  </button>
                )}
                
                {onCancel && (
                  <button
                    onClick={() => onCancel(exchange.id)}
                    className="text-sm px-3 py-1 bg-red-100 text-red-700 rounded hover:bg-red-200"
                  >
                    Reject
                  </button>
                )}
              </div>
            )}
          </div>
        </div>
      </CardBody>
    </Card>
  );
};

export default ExchangeCard; 