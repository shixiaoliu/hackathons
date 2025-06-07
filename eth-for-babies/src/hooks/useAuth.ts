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
      // 1. 获取 nonce
      const nonceResponse = await authApi.getNonce(walletAddress);
      if (!nonceResponse.success || !nonceResponse.data) {
        throw new Error(nonceResponse.error || '获取 nonce 失败');
      }

      const { nonce } = nonceResponse.data;

      // 2. 签名消息
      const message = `Welcome to Family Task Chain!\n\nClick to sign in and accept the Terms of Service.\n\nThis request will not trigger a blockchain transaction or cost any gas fees.\n\nNonce: ${nonce}`;
      
      // 如果是手动输入地址，跳过签名步骤（用于测试目的）
      let signature;
      if (manualAddress) {
        // 对于手动输入的地址，使用一个模拟签名
        signature = '0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000';
      } else {
        signature = await signMessageAsync({ message });
      }

      // 3. 发送登录请求
      const loginResponse = await authApi.login(walletAddress, signature, role);
      if (!loginResponse.success || !loginResponse.data) {
        // 检查错误类型，只有在用户不存在时才尝试注册
        if (loginResponse.error && loginResponse.error.includes('User not found')) {
          const registerResponse = await authApi.register(walletAddress, role);
          if (!registerResponse.success) {
            throw new Error(registerResponse.error || '注册失败');
          }

          // 注册成功后重新登录
          const retryLoginResponse = await authApi.login(walletAddress, signature, role);
          if (!retryLoginResponse.success || !retryLoginResponse.data) {
            throw new Error(retryLoginResponse.error || '登录失败');
          }

          const { token, user } = retryLoginResponse.data;
          apiClient.setToken(token);
          localStorage.setItem('auth_token', token);
          localStorage.setItem('user_data', JSON.stringify(user));
          
          console.log('注册并登录成功:', { user: user.wallet_address, role: user.role });
          
          setAuthState({
            user,
            isAuthenticated: true,
            isLoading: false,
            error: null,
          });
          
          return true;
        } else {
          // 其他登录错误（如签名验证失败、用户已存在等）
          throw new Error(loginResponse.error || '登录失败');
        }
      }

      const { token, user } = loginResponse.data;
      apiClient.setToken(token);
      localStorage.setItem('auth_token', token);
      localStorage.setItem('user_data', JSON.stringify(user));
      
      console.log('登录成功:', { user: user.wallet_address, role: user.role });
      
      setAuthState({
        user,
        isAuthenticated: true,
        isLoading: false,
        error: null,
      });
      
      return true;
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : '登录失败';
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