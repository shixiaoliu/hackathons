import React, { useState } from 'react';
import { ConnectButton } from '@rainbow-me/rainbowkit';
import { useAuthContext } from '../context/AuthContext';

interface LoginModalProps {
  isOpen: boolean;
  onClose: () => void;
}

export const LoginModal: React.FC<LoginModalProps> = ({ isOpen, onClose }) => {
  const { login, isLoading, error, clearError, isWalletConnected } = useAuthContext();
  const [selectedRole, setSelectedRole] = useState<'parent' | 'child'>('parent');
  const [connectionMode, setConnectionMode] = useState<'wallet' | 'manual'>('wallet');
  const [manualAddress, setManualAddress] = useState('');

  const handleLogin = async () => {
    if (connectionMode === 'manual') {
      // 验证手动输入的地址格式
      if (!manualAddress || !/^0x[a-fA-F0-9]{40}$/.test(manualAddress)) {
        alert('请输入有效的以太坊地址格式');
        return;
      }
      // 使用手动输入的地址登录
      const success = await login(selectedRole, manualAddress);
      if (success) {
        onClose();
      }
    } else {
      // 使用钱包连接登录
      const success = await login(selectedRole);
      if (success) {
        onClose();
      }
    }
  };

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-white rounded-lg p-6 w-full max-w-md mx-4">
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-xl font-bold">登录到 FamilyChain</h2>
          <button
            onClick={onClose}
            className="text-gray-500 hover:text-gray-700"
          >
            ✕
          </button>
        </div>

        {error && (
          <div className="mb-4 p-3 bg-red-100 border border-red-400 text-red-700 rounded">
            {error}
            <button
              onClick={clearError}
              className="ml-2 text-red-500 hover:text-red-700"
            >
              ✕
            </button>
          </div>
        )}

        <div className="space-y-4">
          {/* 角色选择 */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              选择您的角色
            </label>
            <div className="flex space-x-4">
              <label className="flex items-center">
                <input
                  type="radio"
                  value="parent"
                  checked={selectedRole === 'parent'}
                  onChange={(e) => setSelectedRole(e.target.value as 'parent' | 'child')}
                  className="mr-2"
                />
                家长
              </label>
              <label className="flex items-center">
                <input
                  type="radio"
                  value="child"
                  checked={selectedRole === 'child'}
                  onChange={(e) => setSelectedRole(e.target.value as 'parent' | 'child')}
                  className="mr-2"
                />
                儿童
              </label>
            </div>
          </div>

          {/* 连接方式选择 */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              连接方式
            </label>
            <div className="flex space-x-4 mb-3">
              <label className="flex items-center">
                <input
                  type="radio"
                  value="wallet"
                  checked={connectionMode === 'wallet'}
                  onChange={(e) => setConnectionMode(e.target.value as 'wallet' | 'manual')}
                  className="mr-2"
                />
                钱包连接
              </label>
              <label className="flex items-center">
                <input
                  type="radio"
                  value="manual"
                  checked={connectionMode === 'manual'}
                  onChange={(e) => setConnectionMode(e.target.value as 'wallet' | 'manual')}
                  className="mr-2"
                />
                手动输入
              </label>
            </div>
            
            {connectionMode === 'wallet' ? (
              <ConnectButton />
            ) : (
              <div>
                <input
                  type="text"
                  placeholder="请输入钱包地址 (0x...)"
                  value={manualAddress}
                  onChange={(e) => setManualAddress(e.target.value)}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                />
                {manualAddress && !/^0x[a-fA-F0-9]{40}$/.test(manualAddress) && (
                  <p className="text-red-500 text-xs mt-1">请输入有效的以太坊地址格式</p>
                )}
              </div>
            )}
          </div>

          {/* 登录按钮 */}
          {((connectionMode === 'wallet' && isWalletConnected) || 
            (connectionMode === 'manual' && manualAddress && /^0x[a-fA-F0-9]{40}$/.test(manualAddress))) && (
            <button
              onClick={handleLogin}
              disabled={isLoading}
              className="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed"
            >
              {isLoading ? '登录中...' : '登录'}
            </button>
          )}

          {connectionMode === 'wallet' && !isWalletConnected && (
            <p className="text-sm text-gray-600 text-center">
              请先连接您的钱包，然后点击登录
            </p>
          )}
          
          {connectionMode === 'manual' && (!manualAddress || !/^0x[a-fA-F0-9]{40}$/.test(manualAddress)) && (
            <p className="text-sm text-gray-600 text-center">
              请输入有效的钱包地址，然后点击登录
            </p>
          )}
        </div>
      </div>
    </div>
  );
};