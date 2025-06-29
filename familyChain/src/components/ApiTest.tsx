import React, { useState } from 'react';
import { authApi, familyApi, childApi, taskApi, contractApi } from '../services/api';
import { useAuthContext } from '../context/AuthContext';

interface ApiTestProps {
  isOpen: boolean;
  onClose: () => void;
}

export const ApiTest: React.FC<ApiTestProps> = ({ isOpen, onClose }) => {
  const { user, isAuthenticated } = useAuthContext();
  const [testResults, setTestResults] = useState<string[]>([]);
  const [loading, setLoading] = useState(false);

  const addResult = (message: string) => {
    setTestResults(prev => [...prev, `${new Date().toLocaleTimeString()}: ${message}`]);
  };

  const clearResults = () => {
    setTestResults([]);
  };

  const testHealthCheck = async () => {
    setLoading(true);
    try {
      const response = await fetch('http://localhost:8080/api/v1/health');
      const data = await response.json();
      addResult(`✅ 健康检查成功: ${JSON.stringify(data)}`);
    } catch (error) {
      addResult(`❌ 健康检查失败: ${error}`);
    }
    setLoading(false);
  };

  const testGetNonce = async () => {
    if (!user) {
      addResult('❌ 请先登录');
      return;
    }
    
    setLoading(true);
    try {
      const response = await authApi.getNonce(user.wallet_address);
      if (response.success) {
        addResult(`✅ 获取 Nonce 成功: ${response.data?.nonce}`);
      } else {
        addResult(`❌ 获取 Nonce 失败: ${response.error}`);
      }
    } catch (error) {
      addResult(`❌ 获取 Nonce 异常: ${error}`);
    }
    setLoading(false);
  };

  const testGetFamilies = async () => {
    if (!isAuthenticated) {
      addResult('❌ 请先登录');
      return;
    }
    
    setLoading(true);
    try {
      const response = await familyApi.getAll();
      if (response.success) {
        addResult(`✅ 获取家庭列表成功: ${JSON.stringify(response.data)}`);
      } else {
        addResult(`❌ 获取家庭列表失败: ${response.error}`);
      }
    } catch (error) {
      addResult(`❌ 获取家庭列表异常: ${error}`);
    }
    setLoading(false);
  };

  const testGetChildren = async () => {
    if (!isAuthenticated) {
      addResult('❌ 请先登录');
      return;
    }
    
    setLoading(true);
    try {
      const response = await childApi.getAll();
      if (response.success) {
        addResult(`✅ 获取儿童列表成功: ${JSON.stringify(response.data)}`);
      } else {
        addResult(`❌ 获取儿童列表失败: ${response.error}`);
      }
    } catch (error) {
      addResult(`❌ 获取儿童列表异常: ${error}`);
    }
    setLoading(false);
  };

  const testGetTasks = async () => {
    if (!isAuthenticated) {
      addResult('❌ 请先登录');
      return;
    }
    
    setLoading(true);
    try {
      const response = await taskApi.getAll();
      if (response.success) {
        addResult(`✅ 获取任务列表成功: ${JSON.stringify(response.data)}`);
      } else {
        addResult(`❌ 获取任务列表失败: ${response.error}`);
      }
    } catch (error) {
      addResult(`❌ 获取任务列表异常: ${error}`);
    }
    setLoading(false);
  };

  const testCreateFamily = async () => {
    if (!isAuthenticated || user?.role !== 'parent') {
      addResult('❌ 只有家长可以创建家庭');
      return;
    }
    
    setLoading(true);
    try {
      const response = await familyApi.create('测试家庭');
      if (response.success) {
        addResult(`✅ 创建家庭成功: ${JSON.stringify(response.data)}`);
      } else {
        addResult(`❌ 创建家庭失败: ${response.error}`);
      }
    } catch (error) {
      addResult(`❌ 创建家庭异常: ${error}`);
    }
    setLoading(false);
  };

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-white rounded-lg p-6 w-full max-w-4xl mx-4 max-h-[80vh] overflow-hidden">
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-xl font-bold">API 连接测试</h2>
          <button
            onClick={onClose}
            className="text-gray-500 hover:text-gray-700"
          >
            ✕
          </button>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          {/* 测试按钮 */}
          <div className="space-y-2">
            <h3 className="font-semibold mb-2">API 测试</h3>
            
            <button
              onClick={testHealthCheck}
              disabled={loading}
              className="w-full bg-green-600 text-white py-2 px-4 rounded-md hover:bg-green-700 disabled:opacity-50"
            >
              测试健康检查
            </button>
            
            <button
              onClick={testGetNonce}
              disabled={loading || !user}
              className="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 disabled:opacity-50"
            >
              测试获取 Nonce
            </button>
            
            <button
              onClick={testGetFamilies}
              disabled={loading || !isAuthenticated}
              className="w-full bg-purple-600 text-white py-2 px-4 rounded-md hover:bg-purple-700 disabled:opacity-50"
            >
              测试获取家庭列表
            </button>
            
            <button
              onClick={testGetChildren}
              disabled={loading || !isAuthenticated}
              className="w-full bg-orange-600 text-white py-2 px-4 rounded-md hover:bg-orange-700 disabled:opacity-50"
            >
              测试获取儿童列表
            </button>
            
            <button
              onClick={testGetTasks}
              disabled={loading || !isAuthenticated}
              className="w-full bg-red-600 text-white py-2 px-4 rounded-md hover:bg-red-700 disabled:opacity-50"
            >
              测试获取任务列表
            </button>
            
            <button
              onClick={testCreateFamily}
              disabled={loading || !isAuthenticated || user?.role !== 'parent'}
              className="w-full bg-indigo-600 text-white py-2 px-4 rounded-md hover:bg-indigo-700 disabled:opacity-50"
            >
              测试创建家庭
            </button>
            
            <button
              onClick={clearResults}
              className="w-full bg-gray-600 text-white py-2 px-4 rounded-md hover:bg-gray-700"
            >
              清除结果
            </button>
          </div>

          {/* 测试结果 */}
          <div>
            <div className="flex justify-between items-center mb-2">
              <h3 className="font-semibold">测试结果</h3>
              <div className="text-sm text-gray-600">
                {isAuthenticated ? (
                  <span className="text-green-600">✅ 已认证 ({user?.role})</span>
                ) : (
                  <span className="text-red-600">❌ 未认证</span>
                )}
              </div>
            </div>
            
            <div className="bg-gray-100 p-4 rounded-md h-96 overflow-y-auto">
              {testResults.length === 0 ? (
                <p className="text-gray-500">点击测试按钮查看结果...</p>
              ) : (
                <div className="space-y-1">
                  {testResults.map((result, index) => (
                    <div key={index} className="text-sm font-mono">
                      {result}
                    </div>
                  ))}
                </div>
              )}
            </div>
          </div>
        </div>

        {/* 当前状态 */}
        <div className="mt-4 p-3 bg-blue-50 rounded-md">
          <h4 className="font-semibold mb-2">当前状态</h4>
          <div className="text-sm space-y-1">
            <div>API 地址: http://localhost:8080/api/v1</div>
            <div>认证状态: {isAuthenticated ? '已认证' : '未认证'}</div>
            {user && (
              <>
                <div>用户角色: {user.role}</div>
                <div>钱包地址: {user.wallet_address}</div>
              </>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};