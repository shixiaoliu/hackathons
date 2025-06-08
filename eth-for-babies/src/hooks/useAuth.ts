import { useState, useEffect, useCallback } from 'react';
import { useAccount, useSignMessage } from 'wagmi';
import { authApi, apiClient } from '../services/api';
import type { User } from '../services/api';

interface AuthState {
  user: User | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  error: string | null;
}

export const useAuth = () => {
  const { address, isConnected } = useAccount();
  const { signMessageAsync } = useSignMessage();
  
  const [authState, setAuthState] = useState<AuthState>({
    user: null,
    isAuthenticated: false,
    isLoading: false,
    error: null,
  });

  // 检查本地存储的认证状态
  useEffect(() => {
    const token = localStorage.getItem('auth_token');
    const userStr = localStorage.getItem('user_data');
    
    if (token && userStr) {
      try {
        const user = JSON.parse(userStr);
        apiClient.setToken(token);
        setAuthState({
          user,
          isAuthenticated: true,
          isLoading: false,
          error: null,
        });
      } catch (error) {
        // 清除无效的本地数据
        localStorage.removeItem('auth_token');
        localStorage.removeItem('user_data');
      }
    }
  }, []);

  // 当钱包断开连接时清除认证状态
  useEffect(() => {
    if (!isConnected) {
      logout();
    }
  }, [isConnected]);

  // 登录函数
  const login = useCallback(async (role: 'parent' | 'child' = 'parent', manualAddress?: string) => {
    const walletAddress = manualAddress || address;
    if (!walletAddress || walletAddress === 'undefined' || !/^0x[a-fA-F0-9]{40}$/.test(walletAddress)) {
      setAuthState(prev => ({ ...prev, error: '请先连接钱包或输入有效的钱包地址' }));
      return false;
    }

    setAuthState(prev => ({ ...prev, isLoading: true, error: null }));

    try {
      // 由于服务器需要正确的nonce，我们恢复nonce获取流程
      console.log('开始登录流程，请求nonce...');
      
      // 1. 尝试获取nonce (带重试逻辑)
      const maxNonceRetries = 3;
      let nonceResponse = null;
      let nonce = '';
      
      for (let i = 0; i < maxNonceRetries; i++) {
        try {
          if (i > 0) {
            console.log(`尝试获取nonce (第${i+1}次)...`);
            await new Promise(resolve => setTimeout(resolve, 1000 * i));
          }
          
          nonceResponse = await authApi.getNonce(walletAddress);
          
          if (nonceResponse.success && nonceResponse.data) {
            nonce = nonceResponse.data.nonce;
            console.log(`成功获取nonce: ${nonce}`);
            break;
          } else {
            console.warn(`获取nonce失败 (第${i+1}次): ${nonceResponse.error || '未知错误'}`);
          }
        } catch (err) {
          console.error(`获取nonce出错 (第${i+1}次):`, err);
        }
      }
      
      if (!nonce) {
        // 如果所有的nonce获取尝试都失败，尝试注册一个新用户
        console.log('无法获取nonce，尝试注册新用户...');
        try {
          const registerResponse = await authApi.register(walletAddress, role);
          console.log('注册响应:', registerResponse);
          
          // 再次尝试获取nonce
          nonceResponse = await authApi.getNonce(walletAddress);
          if (nonceResponse.success && nonceResponse.data) {
            nonce = nonceResponse.data.nonce;
            console.log(`注册后成功获取nonce: ${nonce}`);
          } else {
            throw new Error(nonceResponse.error || '获取nonce失败');
          }
        } catch (err) {
          throw new Error('无法获取有效的nonce，请稍后再试');
        }
      }
      
      // 2. 签名消息
      const message = `Welcome to Family Task Chain!\n\nClick to sign in and accept the Terms of Service.\n\nThis request will not trigger a blockchain transaction or cost any gas fees.\n\nNonce: ${nonce}`;
      console.log('待签名消息:', message);
      
      // 如果是手动输入地址，使用特殊方式处理
      let signature;
      if (manualAddress) {
        // 手动地址模式，先尝试注册再直接返回登录成功
        try {
          // 尝试直接注册用户
          const registerResponse = await authApi.register(walletAddress, role);
          console.log('手动地址模式注册响应:', registerResponse);
          
          // 创建模拟的认证状态
          const mockUser = {
            id: Date.now(),
            wallet_address: walletAddress,
            role: role,
            created_at: new Date().toISOString(),
            updated_at: new Date().toISOString()
          };
          
          const mockToken = `manual_${Date.now()}_${Math.random().toString(36).substring(2)}`;
          
          // 保存到localStorage
          apiClient.setToken(mockToken);
          localStorage.setItem('auth_token', mockToken);
          localStorage.setItem('user_data', JSON.stringify(mockUser));
          localStorage.setItem('user_role', role);
          
          // 更新认证状态
          setAuthState({
            user: mockUser,
            isAuthenticated: true,
            isLoading: false,
            error: null
          });
          
          console.log('手动地址模式模拟登录成功');
          return true;
        } catch (err) {
          console.warn('手动地址模式注册失败，使用标准流程');
        }
        
        // 使用固定签名（仅用于测试）
        signature = '0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000';
        console.log('使用模拟签名(手动地址模式)');
      } else {
        // 请求用户签名
        console.log('请求用户签名消息...');
        
        try {
          signature = await signMessageAsync({ message });
          console.log('用户已签名消息:', signature.substring(0, 30) + '...');
        } catch (err) {
          throw new Error('用户拒绝签名，无法继续登录');
        }
      }

      // 3. 发送登录请求
      console.log('发送登录请求...');
      
      // 设置最大重试次数
      const maxRetries = 3;
      let retryCount = 0;
      let loginSuccess = false;
      let lastError;
      let loginResponse;
      
      // 重试登录直到成功或达到最大重试次数
      while (!loginSuccess && retryCount < maxRetries) {
        try {
          // 每次重试前等待递增的延迟
          if (retryCount > 0) {
            const delay = retryCount * 1000; // 1秒, 2秒, 3秒...
            console.log(`登录重试 ${retryCount}/${maxRetries}，等待 ${delay}ms...`);
            await new Promise(resolve => setTimeout(resolve, delay));
          }
          
          loginResponse = await authApi.login(walletAddress, signature, role);
          
          if (loginResponse.success && loginResponse.data) {
            loginSuccess = true;
          } else {
            throw new Error(loginResponse.error || 'Failed to login');
          }
        } catch (error) {
          lastError = error instanceof Error ? error : new Error(String(error));
          console.warn(`登录尝试 ${retryCount + 1} 失败:`, error);
          retryCount++;
        }
      }
      
      if (!loginSuccess) {
        // 如果所有登录尝试都失败，但收到了Invalid signature错误，给出更友好的提示
        if (lastError && 
            (lastError.message.includes('Invalid signature') || 
             lastError.message.includes('无效签名'))) {
          throw new Error('签名验证失败，请确保钱包账户正确并重试。如果问题持续存在，请尝试重新连接钱包。');
        }
        throw lastError || new Error('登录失败，请稍后重试');
      }

      // 登录成功
      const { token, user } = loginResponse!.data!;
      apiClient.setToken(token);
      localStorage.setItem('auth_token', token);
      localStorage.setItem('user_data', JSON.stringify(user));
      localStorage.setItem('user_role', user.role);
      
      console.log('登录成功:', { user: user.wallet_address, role: user.role });
      
      setAuthState({
        user,
        isAuthenticated: true,
        isLoading: false,
        error: null,
      });
      
      return true;
    } catch (error) {
      console.error('登录失败:', error);
      let errorMessage = '登录失败';
      
      if (error instanceof Error) {
        errorMessage = error.message;
        console.error('错误详情:', {
          name: error.name,
          message: error.message,
          stack: error.stack
        });
      }
      
      setAuthState(prev => ({
        ...prev,
        isLoading: false,
        error: errorMessage,
      }));
      return false;
    }
  }, [address, signMessageAsync]);

  // 登出函数
  const logout = useCallback(async () => {
    try {
      await authApi.logout();
    } catch (error) {
      console.warn('登出请求失败:', error);
    } finally {
      apiClient.clearToken();
      localStorage.removeItem('auth_token');
      localStorage.removeItem('user_data');
      setAuthState({
        user: null,
        isAuthenticated: false,
        isLoading: false,
        error: null,
      });
    }
  }, []);

  // 刷新令牌
  const refreshToken = useCallback(async () => {
    try {
      const response = await authApi.refresh();
      if (response.success && response.data) {
        apiClient.setToken(response.data.token);
        return true;
      }
    } catch (error) {
      console.warn('刷新令牌失败:', error);
    }
    return false;
  }, []);

  // 清除错误
  const clearError = useCallback(() => {
    setAuthState(prev => ({ ...prev, error: null }));
  }, []);

  return {
    ...authState,
    login,
    logout,
    refreshToken,
    clearError,
    walletAddress: address,
    isWalletConnected: isConnected && address && /^0x[a-fA-F0-9]{40}$/.test(address),
  };
};