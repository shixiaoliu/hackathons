import React, { useState, useEffect } from 'react';
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
  const [loginAttempts, setLoginAttempts] = useState(0);
  const [showRetry, setShowRetry] = useState(false);

  // 重置登录状态
  const handleReset = () => {
    clearError();
    setLoginAttempts(0);
    setShowRetry(false);
  };
  
  const handleLogin = async () => {
    setShowRetry(false);
    
    try {
      if (connectionMode === 'manual') {
        // 验证手动输入的地址格式
        if (!manualAddress || !/^0x[a-fA-F0-9]{40}$/.test(manualAddress)) {
          alert('Please enter a valid Ethereum address format');
          return;
        }
        // 使用手动输入的地址登录
        const success = await login(selectedRole, manualAddress);
        if (success) {
          onClose();
        } else {
          setLoginAttempts(prev => prev + 1);
          // 如果失败3次或以上，显示重试按钮
          if (loginAttempts >= 2) {
            setShowRetry(true);
          }
        }
      } else {
        // 使用钱包连接登录
        const success = await login(selectedRole);
        if (success) {
          onClose();
        } else {
          setLoginAttempts(prev => prev + 1);
          // 如果失败3次或以上，显示重试按钮
          if (loginAttempts >= 2) {
            setShowRetry(true);
          }
        }
      }
    } catch (err) {
      console.error('登录过程中发生错误:', err);
      setLoginAttempts(prev => prev + 1);
      if (loginAttempts >= 2) {
        setShowRetry(true);
      }
    }
  };

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-white rounded-lg p-6 w-full max-w-md mx-4">
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-lg font-bold">Login to FamilyChain</h2>
          <button
            onClick={onClose}
            className="text-gray-500 hover:text-gray-700"
          >
            ✕
          </button>
        </div>

        <p className="text-gray-600 mb-4 text-sm">
          Please select your role
        </p>

        {error && (
          <div className="mb-4 p-3 bg-red-100 border border-red-400 text-red-700 rounded">
            <div className="flex justify-between items-center">
              <div>
                <p className="font-medium mb-1">登录失败</p>
                <p className="text-sm">{error}</p>
                {error.includes('Invalid signature') && (
                  <p className="text-sm mt-1">
                    签名验证失败。请确保您使用的是正确的钱包账户，并尝试重新连接钱包。
                  </p>
                )}
                {error.includes('签名验证失败') && (
                  <p className="text-sm mt-1">
                    您可以尝试以下解决方法：
                    <br/>1. 重新连接您的钱包
                    <br/>2. 确认您选择了正确的账户
                    <br/>3. 如果问题持续存在，请尝试清除浏览器缓存并重试
                  </p>
                )}
                {error.includes('用户拒绝签名') && (
                  <p className="text-sm mt-1">
                    您拒绝了签名请求。要继续登录，您需要签名以验证钱包所有权。
                    <br/>
                    <button 
                      onClick={handleLogin}
                      className="mt-2 px-3 py-1 bg-blue-100 text-blue-700 rounded hover:bg-blue-200"
                    >
                      重新尝试登录
                    </button>
                    <button 
                      onClick={() => {
                        clearError();
                        setConnectionMode('manual');
                      }}
                      className="mt-2 ml-2 px-3 py-1 bg-gray-100 text-gray-700 rounded hover:bg-gray-200"
                    >
                      使用手动输入方式
                    </button>
                  </p>
                )}
                {error.includes('Failed to update nonce') && (
                  <p className="text-sm mt-1">
                    服务器更新nonce失败，我们正在尝试使用一种替代方法登录。请重试登录。
                  </p>
                )}
                {error.includes('服务器内部错误') && (
                  <p className="text-sm mt-1">
                    服务器暂时不可用，请稍后再试。如果问题持续存在，请联系管理员。
                  </p>
                )}
              </div>
              <button
                onClick={clearError}
                className="text-red-500 hover:text-red-700"
              >
                ✕
              </button>
            </div>
            {showRetry && (
              <div className="mt-2 flex justify-end">
                <button 
                  onClick={handleReset}
                  className="text-sm px-2 py-1 rounded bg-blue-100 text-blue-700 hover:bg-blue-200"
                >
                  重置并重试
                </button>
              </div>
            )}
          </div>
        )}

        <div className="space-y-4">
          {/* 角色选择 */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">Select your role</label>
            <div className="flex space-x-4">
              <label className="flex items-center">
                <input
                  type="radio"
                  value="parent"
                  checked={selectedRole === 'parent'}
                  onChange={(e) => setSelectedRole(e.target.value as 'parent' | 'child')}
                  className="mr-2"
                />
                Parent
              </label>
              <label className="flex items-center">
                <input
                  type="radio"
                  value="child"
                  checked={selectedRole === 'child'}
                  onChange={(e) => setSelectedRole(e.target.value as 'parent' | 'child')}
                  className="mr-2"
                />
                Child
              </label>
            </div>
          </div>

          {/* 连接方式选择 */}
          <div>
            <p className="text-gray-600 mb-4 text-sm">Connection method</p>
            <div className="flex space-x-4 mb-3">
              <label className="flex items-center">
                <input
                  type="radio"
                  value="wallet"
                  checked={connectionMode === 'wallet'}
                  onChange={(e) => setConnectionMode(e.target.value as 'wallet' | 'manual')}
                  className="mr-2"
                />
                Wallet Connection
              </label>
              <label className="flex items-center">
                <input
                  type="radio"
                  value="manual"
                  checked={connectionMode === 'manual'}
                  onChange={(e) => setConnectionMode(e.target.value as 'wallet' | 'manual')}
                  className="mr-2"
                />
                Manual Input
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
            <>
              <button
                onClick={handleLogin}
                disabled={isLoading}
                className="w-full py-3 px-4 rounded-md flex justify-center items-center font-medium text-white transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                style={{
                  backgroundColor: isLoading ? '#4B5563' : '#4F46E5',
                  boxShadow: '0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06)'
                }}
              >
                {isLoading ? (
                  <div className="flex items-center">
                    <svg className="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                      <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                      <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    登录中...
                  </div>
                ) : 'Login'}
              </button>
            </>
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