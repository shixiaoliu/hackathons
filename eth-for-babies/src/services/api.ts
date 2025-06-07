// API 服务配置
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1';

// API 响应类型
interface ApiResponse<T> {
  success: boolean;
  data?: T;
  message?: string;
  error?: string;
}

// 用户相关类型
interface User {
  id: number;
  wallet_address: string;
  role: 'parent' | 'child';
  created_at: string;
  updated_at: string;
}

// 家庭相关类型
interface Family {
  id: number;
  name: string;
  parent_address: string;
  created_at: string;
  updated_at: string;
  children?: Child[];
}

// 儿童相关类型
interface Child {
  id: number;
  name: string;
  wallet_address: string;
  age: number;
  avatar?: string;
  parent_address: string;
  total_tasks_completed: number;
  total_rewards_earned: string;
  created_at: string;
  updated_at: string;
}

// 任务相关类型
interface Task {
  id: number;
  title: string;
  description: string;
  reward_amount: string;
  difficulty: 'easy' | 'medium' | 'hard';
  status: 'pending' | 'in_progress' | 'completed' | 'approved' | 'rejected';
  assigned_child_id: number;
  created_by: string;
  due_date?: string;
  created_at: string;
  updated_at: string;
}

// HTTP 请求工具函数
class ApiClient {
  private baseURL: string;
  private token: string | null = null;

  constructor(baseURL: string) {
    this.baseURL = baseURL;
    this.token = localStorage.getItem('auth_token');
  }

  private async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<ApiResponse<T>> {
    const url = `${this.baseURL}${endpoint}`;
    const headers: HeadersInit = {
      'Content-Type': 'application/json',
      ...options.headers,
    };

    // 每次请求时都从localStorage获取最新的token
    const currentToken = this.token || localStorage.getItem('auth_token');
    if (currentToken) {
      headers.Authorization = `Bearer ${currentToken}`;
      console.log('使用token进行请求:', currentToken.substring(0, 20) + '...');
    } else {
      console.log('没有token，发送未认证请求');
    }

    try {
      const response = await fetch(url, {
        ...options,
        headers,
      });

      const data = await response.json();

      if (!response.ok) {
        return {
          success: false,
          error: data.message || `HTTP ${response.status}`,
        };
      }

      return {
        success: true,
        data: data.data || data,
        message: data.message,
      };
    } catch (error) {
      return {
        success: false,
        error: error instanceof Error ? error.message : 'Network error',
      };
    }
  }

  // 设置认证令牌
  setToken(token: string) {
    this.token = token;
    localStorage.setItem('auth_token', token);
    console.log('Token已设置:', token.substring(0, 20) + '...');
  }

  // 获取当前token
  getToken(): string | null {
    return this.token || localStorage.getItem('auth_token');
  }

  // 清除认证令牌
  clearToken() {
    this.token = null;
    localStorage.removeItem('auth_token');
    console.log('Token已清除');
  }

  // GET 请求
  async get<T>(endpoint: string): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, { method: 'GET' });
  }

  // POST 请求
  async post<T>(endpoint: string, data?: any): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, {
      method: 'POST',
      body: data ? JSON.stringify(data) : undefined,
    });
  }

  // PUT 请求
  async put<T>(endpoint: string, data?: any): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, {
      method: 'PUT',
      body: data ? JSON.stringify(data) : undefined,
    });
  }

  // DELETE 请求
  async delete<T>(endpoint: string): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, { method: 'DELETE' });
  }
}

// 创建 API 客户端实例
const apiClient = new ApiClient(API_BASE_URL);

// 认证相关 API
export const authApi = {
  // 获取 nonce
  getNonce: (walletAddress: string) =>
    apiClient.get<{ nonce: string }>(`/auth/nonce/${walletAddress}`),

  // 钱包登录
  login: (walletAddress: string, signature: string, role?: 'parent' | 'child') =>
    apiClient.post<{ token: string; user: User }>('/auth/login', {
      wallet_address: walletAddress,
      signature,
      ...(role && { role }),
    }),

  // 注册
  register: (walletAddress: string, role: 'parent' | 'child') =>
    apiClient.post<{ user: User }>('/auth/register', {
      wallet_address: walletAddress,
      role,
    }),

  // 刷新令牌
  refresh: () => apiClient.post<{ token: string }>('/auth/refresh'),

  // 登出
  logout: () => apiClient.post('/auth/logout'),
};

// 家庭相关 API
export const familyApi = {
  // 创建家庭
  create: (name: string) =>
    apiClient.post<Family>('/families', { name }),

  // 获取家庭列表
  getAll: () => apiClient.get<Family[]>('/families'),

  // 获取家庭详情
  getById: (id: number) => apiClient.get<Family>(`/families/${id}`),

  // 更新家庭信息
  update: (id: number, data: Partial<Family>) =>
    apiClient.put<Family>(`/families/${id}`, data),

  // 添加家庭成员
  addMember: (id: number, walletAddress: string) =>
    apiClient.post(`/families/${id}/members`, { wallet_address: walletAddress }),
};

// 儿童相关 API
export const childApi = {
  // 添加儿童
  create: (childData: Omit<Child, 'id' | 'created_at' | 'updated_at'>) =>
    apiClient.post<Child>('/children', childData),

  // 获取儿童列表
  getAll: () => apiClient.get<Child[]>('/children/my'),

  // 获取儿童详情
  getById: (id: number) => apiClient.get<Child>(`/children/${id}`),

  // 更新儿童信息
  update: (id: number, data: Partial<Child>) =>
    apiClient.put<Child>(`/children/${id}`, data),

  // 获取儿童进度
  getProgress: (id: number) =>
    apiClient.get<any>(`/children/${id}/progress`),
};

// 任务相关 API
export const taskApi = {
  // 创建任务
  create: (taskData: Omit<Task, 'id' | 'created_at' | 'updated_at'>) =>
    apiClient.post<Task>('/tasks', taskData),

  // 获取任务列表
  getAll: (params?: { child_id?: number; status?: string }) => {
    const queryParams = new URLSearchParams();
    if (params?.child_id) queryParams.append('child_id', params.child_id.toString());
    if (params?.status) queryParams.append('status', params.status);
    const query = queryParams.toString();
    return apiClient.get<Task[]>(`/tasks${query ? `?${query}` : ''}`);
  },

  // 获取任务详情
  getById: (id: number) => apiClient.get<Task>(`/tasks/${id}`),

  // 更新任务
  update: (id: number, data: Partial<Task>) =>
    apiClient.put<Task>(`/tasks/${id}`, data),

  // 完成任务
  complete: (id: number) => apiClient.post(`/tasks/${id}/complete`),

  // 批准任务
  approve: (id: number) => apiClient.post(`/tasks/${id}/approve`),

  // 拒绝任务
  reject: (id: number, reason?: string) =>
    apiClient.post(`/tasks/${id}/reject`, { reason }),
};

// 合约相关 API
export const contractApi = {
  // 获取代币余额
  getBalance: (address: string) =>
    apiClient.get<{ balance: string }>(`/contracts/balance/${address}`),

  // 转移代币
  transfer: (to: string, amount: string) =>
    apiClient.post<{ transaction_hash: string }>('/contracts/transfer', {
      to,
      amount,
    }),

  // 获取交易状态
  getTransactionStatus: (hash: string) =>
    apiClient.get<any>(`/contracts/transactions/${hash}`),
};

// 导出 API 客户端
export { apiClient };
export type { ApiResponse, User, Family, Child, Task };