import { FC } from 'react';
import { formatDistanceToNow } from '../../utils/dateUtils';
import { Exchange } from '../../services/api';
import { Check, XIcon, ShoppingBag, AlertTriangle, User } from 'lucide-react';
import Card, { CardBody } from '../common/Card';

interface ExchangeCardProps {
  exchange: Exchange;
  isChild?: boolean;
}

const ExchangeCard: FC<ExchangeCardProps> = ({ 
  exchange, 
  isChild = false
}) => {
  // Status badge
  const renderStatusBadge = () => {
    switch (exchange.status) {
      case 'completed':
      case 'confirmed':
        return (
          <div className="flex items-center px-2 py-1 rounded bg-green-100 text-green-800 text-xs font-medium">
            <Check className="h-3 w-3 mr-1" />
            Completed
          </div>
        );
      case 'cancelled':
        return (
          <div className="flex items-center px-2 py-1 rounded bg-red-100 text-red-800 text-xs font-medium">
            <XIcon className="h-3 w-3 mr-1" />
            Cancelled
          </div>
        );
      case 'failed':
        return (
          <div className="flex items-center px-2 py-1 rounded bg-red-100 text-red-800 text-xs font-medium">
            <AlertTriangle className="h-3 w-3 mr-1" />
            Failed
          </div>
        );
      default:
        return (
          <div className="flex items-center px-2 py-1 rounded bg-green-100 text-green-800 text-xs font-medium">
            <Check className="h-3 w-3 mr-1" />
            Completed
          </div>
        );
    }
  };
  
  // Format time
  const formatTime = (dateString: string) => {
    try {
      const date = new Date(dateString);
      return formatDistanceToNow(date);
    } catch (error) {
      console.error('Date formatting error:', error);
      return 'Unknown time';
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
                alt={exchange.reward_name || 'Reward image'}
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
                {exchange.reward_name || 'Unnamed reward'}
              </h3>
              {renderStatusBadge()}
            </div>
            
            <div className="flex items-center text-sm text-gray-500 mb-1">
              <span className="font-medium text-primary-600">{Math.floor(exchange.token_amount)} FCT</span>
              <span className="mx-2">•</span>
              <span title={new Date(exchange.exchange_date).toLocaleString()}>
                {formatTime(exchange.exchange_date)}
              </span>
              {/* Display child name */}
              {exchange.child_name && !isChild && (
                <>
                  <span className="mx-2">•</span>
                  <span className="flex items-center">
                    <User className="h-3 w-3 mr-1" />
                    {exchange.child_name}
                  </span>
                </>
              )}
            </div>
            
            {exchange.notes && (
              <p className="text-sm text-gray-600 mt-1">{exchange.notes}</p>
            )}
          </div>
        </div>
      </CardBody>
    </Card>
  );
};

export default ExchangeCard; 