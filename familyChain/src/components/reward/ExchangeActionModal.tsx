import React, { useState } from 'react';
import { X } from 'lucide-react';
import Button from '../common/Button';
import { Exchange, ExchangeStatus } from '../../types/reward';

interface ExchangeActionModalProps {
  isOpen: boolean;
  onClose: () => void;
  onConfirm: (notes?: string) => void;
  exchange: Exchange;
  actionType: 'approve' | 'cancel';
  isLoading?: boolean;
}

const ExchangeActionModal: React.FC<ExchangeActionModalProps> = ({
  isOpen,
  onClose,
  onConfirm,
  exchange,
  actionType,
  isLoading = false
}) => {
  const [notes, setNotes] = useState<string>('');
  
  if (!isOpen) return null;
  
  const title = actionType === 'approve' ? 'Approve Exchange Request' : 'Reject Exchange Request';
  const confirmText = actionType === 'approve' ? 'Approve' : 'Reject';
  const confirmColor = actionType === 'approve' ? 'bg-green-600 hover:bg-green-700' : 'bg-red-600 hover:bg-red-700';

  // 获取状态描述
  const getStatusDescription = () => {
    if (actionType === 'approve') {
      return `确认此兑换请求将会：
1. 减少 ${exchange.child_name || '孩子'} 的代币余额 ${exchange.token_amount} FCT
2. 将奖品标记为已兑换
3. 从可用奖品列表中移除此奖品`;
    } else {
      return `拒绝此兑换请求将会：
1. 保留 ${exchange.child_name || '孩子'} 的代币余额
2. 将奖品保留在可用奖品列表中
3. 将此请求标记为已取消`;
    }
  };

  return (
    <div className="fixed inset-0 z-50 overflow-y-auto">
      <div 
        className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0"
        onClick={onClose} // 点击背景关闭
      >
        {/* 背景遮罩 */}
        <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
        
        {/* 模态框 */}
        <div 
          className="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg"
          onClick={(e) => e.stopPropagation()} // 防止点击内容区域关闭
        >
          {/* 标题栏 */}
          <div className="bg-gray-50 px-4 py-3 flex justify-between items-center">
            <h3 className="text-lg font-medium leading-6 text-gray-900">{title}</h3>
            <button 
              type="button" 
              className="text-gray-400 hover:text-gray-500"
              onClick={onClose}
            >
              <X className="h-5 w-5" />
            </button>
          </div>
          
          {/* 内容 */}
          <div className="px-4 pt-5 pb-4 sm:p-6">
            <div className="mb-4">
              <p className="text-gray-700 whitespace-pre-line">
                {getStatusDescription()}
              </p>
            </div>
            
            {/* 兑换信息 */}
            <div className="bg-gray-50 p-3 rounded-md mb-4">
              <p><strong>Reward:</strong> {exchange.reward_name || 'Unnamed Reward'}</p>
              <p><strong>Exchanger:</strong> {exchange.child_name || 'Unknown'}</p>
              <p><strong>Tokens:</strong> {exchange.token_amount}</p>
              {exchange.notes && <p><strong>Notes:</strong> {exchange.notes}</p>}
            </div>
            
            {/* 添加备注 */}
            <div className="mb-4">
              <label htmlFor="notes" className="block text-sm font-medium text-gray-700 mb-1">
                备注 (可选)
              </label>
              <textarea
                id="notes"
                value={notes}
                onChange={(e) => setNotes(e.target.value)}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-primary-500"
                rows={3}
                placeholder={actionType === 'approve' ? 'Add notes (e.g., issuance instructions)' : 'Add rejection reason'}
                disabled={isLoading}
              ></textarea>
            </div>
            
            {/* 操作按钮 */}
            <div className="mt-4 flex justify-end gap-3">
              <Button
                type="button"
                variant="outline"
                onClick={onClose}
                disabled={isLoading}
              >
                取消
              </Button>
              <Button
                type="button"
                className={confirmColor}
                onClick={() => onConfirm(notes)}
                isLoading={isLoading}
                disabled={isLoading}
              >
                {confirmText}
              </Button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ExchangeActionModal; 