// API 服务配置
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1';

// 禁用模拟模式
const USE_MOCK = false;

// 用于调试的日志
const API_DEBUG = true;

// 临时解决方案参数
const MAX_NONCE_RETRIES = 3;       // 获取nonce最大重试次数
const NONCE_RETRY_DELAY = 1000;    // 重试间隔(毫秒)

// API 响应类型
interface ApiResponse<T> {
  success: boolean;
  data?: T;
  message?: string;
  error?: string;
  status?: number;
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
  assigned_child_id?: number;
  created_by: string;
  contract_task_id?: number;
  due_date?: string;
  created_at: string;
  updated_at: string;
}

// 奖品相关类型
interface Reward {
  id: number;
  family_id: number;
  name: string;
  description: string;
  image_url: string;
  token_price: number;
  created_by: number;
  active: boolean;
  stock: number;
  created_at: string;
  updated_at: string;
}

// 兑换记录相关类型
interface Exchange {
  id: number;
  reward_id: number;
  child_id: number;
  token_amount: number;
  status: 'pending' | 'completed' | 'cancelled';
  exchange_date: string;
  completed_date?: string;
  notes?: string;
  created_at: string;
  updated_at: string;
  reward_name?: string;
  reward_image?: string;
  child_name?: string;
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
    const headers: Record<string, string> = {
      'Content-Type': 'application/json',
      ...(options.headers as Record<string, string> || {})
    };

    // 每次请求时都从localStorage获取最新的token
    const currentToken = this.token || localStorage.getItem('auth_token');
    if (currentToken) {
      headers.Authorization = `Bearer ${currentToken}`;
      if (API_DEBUG) console.log('使用token进行请求:', currentToken.substring(0, 20) + '...');
    } else {
      if (API_DEBUG) console.log('没有token，发送未认证请求');
    }

    // 详细日志请求信息
    if (API_DEBUG) {
      console.log(`API请求: ${options.method || 'GET'} ${url}`);
      if (options.body) {
        try {
          console.log('请求数据:', JSON.parse(options.body as string));
        } catch (e) {
          console.log('请求数据:', options.body);
        }
      }
    }

    try {
      // 创建一个带超时的fetch Promise
      const fetchPromise = fetch(url, {
        ...options,
        headers,
      });

      // 创建一个超时Promise
      const timeoutPromise = new Promise<Response>((_, reject) => {
        setTimeout(() => reject(new Error('请求超时，请稍后再试')), 15000); // 15秒超时
      });

      // 使用Promise.race竞争
      const response = await Promise.race([fetchPromise, timeoutPromise]) as Response;
      
      // 打印详细响应状态
      if (API_DEBUG) {
        console.log(`API响应: ${response.status} ${response.statusText} - ${url}`);
        console.log(`响应头: ${JSON.stringify([...response.headers.entries()].reduce((obj, [key, val]) => {
          obj[key] = val;
          return obj;
        }, {} as Record<string, string>))}`);
      }
      
      // 尝试解析响应体
      let data;
      const contentType = response.headers.get('content-type');
      if (contentType && contentType.includes('application/json')) {
        try {
          data = await response.json();
          if (API_DEBUG) console.log('响应数据:', data);
        } catch (e) {
          console.error('解析响应JSON出错:', e);
          return {
            success: false,
            error: '服务器返回的数据格式无效',
            status: response.status
          };
        }
      } else {
        // 非JSON响应
        try {
          const textData = await response.text();
          if (API_DEBUG) console.log('响应文本:', textData);
          try {
            // 尝试将文本解析为JSON
            data = JSON.parse(textData);
          } catch (e) {
            data = { message: textData };
          }
        } catch (e) {
          console.error('读取响应出错:', e);
          return {
            success: false,
            error: '无法读取服务器响应',
            status: response.status
          };
        }
      }

      if (!response.ok) {
        // 特殊处理500错误
        if (response.status === 500) {
          console.error('服务器内部错误:', data);
          return {
            success: false,
            error: data.error || '服务器内部错误，请稍后再试',
            status: 500
          };
        }
        
        return {
          success: false,
          error: data.error || data.message || `HTTP ${response.status}`,
          status: response.status
        };
      }

      return {
        success: true,
        data: data.data || data,
        message: data.message,
      };
    } catch (error) {
      console.error('网络请求出错:', error);
      return {
        success: false,
        error: error instanceof Error ? error.message : '网络错误，请检查网络连接',
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

  // 添加避免缓存的参数
  addCacheBusters(url: string): string {
    const separator = url.includes('?') ? '&' : '?';
    return `${url}${separator}_t=${Date.now()}&_r=${Math.random().toString(36).substring(2)}`;
  }

  // 清除认证令牌
  clearToken() {
    this.token = null;
    localStorage.removeItem('auth_token');
    console.log('Token已清除');
  }

  // GET 请求
  async get<T>(endpoint: string): Promise<ApiResponse<T>> {
    // 添加cache buster参数以避免缓存
    const updatedEndpoint = this.addCacheBusters(endpoint);
    return this.request<T>(updatedEndpoint, { method: 'GET' });
  }

  // POST 请求
  async post<T>(endpoint: string, data?: any): Promise<ApiResponse<T>> {
    // 添加避免缓存的属性
    const updatedData = data ? {
      ...data,
      _cache_buster: Date.now()
    } : undefined;
    
    return this.request<T>(endpoint, {
      method: 'POST',
      body: updatedData ? JSON.stringify(updatedData) : undefined,
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

// 模拟数据服务
const mockService = {
  getNonce: (walletAddress: string): ApiResponse<{ nonce: string }> => {
    console.log('使用模拟nonce数据');
    // 生成一个随机nonce
    const randomNonce = Math.random().toString(36).substring(2, 15);
    return {
      success: true,
      data: { nonce: randomNonce }
    };
  },
  
  login: (walletAddress: string, signature: string, role: 'parent' | 'child' = 'parent'): ApiResponse<{ token: string; user: User }> => {
    console.log('使用模拟登录数据');
    // 生成一个模拟token
    const mockToken = 'mock_' + Math.random().toString(36).substring(2, 15);
    return {
      success: true,
      data: {
        token: mockToken,
        user: {
          id: 1,
          wallet_address: walletAddress,
          role: role,
          created_at: new Date().toISOString(),
          updated_at: new Date().toISOString()
        }
      }
    };
  },
  
  register: (walletAddress: string, role: 'parent' | 'child'): ApiResponse<{ user: User }> => {
    console.log('使用模拟注册数据');
    return {
      success: true,
      data: {
        user: {
          id: new Date().getTime(),
          wallet_address: walletAddress,
          role: role,
          created_at: new Date().toISOString(),
          updated_at: new Date().toISOString()
        }
      }
    };
  },
  
  logout: (): ApiResponse<null> => {
    console.log('使用模拟登出数据');
    return {
      success: true
    };
  },

  // 模拟服务其他方法
  refresh: (): ApiResponse<{ token: string }> => {
    console.log('使用模拟刷新令牌数据');
    return {
      success: true,
      data: {
        token: 'mock_refresh_token_' + Math.random().toString(36).substring(2, 15)
      }
    };
  },
};

// 临时修复工具
const tempFixes = {
  // 生成随机nonce，作为临时解决方案
  generateRandomNonce: (): string => {
    const timestamp = new Date().getTime();
    const randomPart = Math.random().toString(36).substring(2, 10);
    return `${timestamp}-${randomPart}`;
  },
  
  // 格式化钱包地址
  formatAddress: (walletAddress: string): string => {
    if (!walletAddress) return '';
    return walletAddress.toLowerCase();
  },
  
  // 直接通过正则来判断是否是nonce相关错误
  isNonceError: (error?: string): boolean => {
    if (!error) return false;
    return /failed to update nonce/i.test(error) || 
           /nonce.*error/i.test(error) ||
           /nonce.*failed/i.test(error);
  },
  
  // 等待函数
  delay: (ms: number) => new Promise(resolve => setTimeout(resolve, ms))
};

// 认证相关 API
export const authApi = {
  // 获取 nonce
  getNonce: (walletAddress: string) => {
    // 保证钱包地址格式正确
    const address = walletAddress.toLowerCase();
    
    // 输出调试信息
    if (API_DEBUG) {
      console.log(`请求nonce - 钱包地址: ${address}`);
    }
    
    // 使用API客户端的缓存避免机制
    return apiClient.get<{ nonce: string }>(`/auth/nonce/${address}`);
  },

  // 钱包登录
  login: (walletAddress: string, signature: string, role?: 'parent' | 'child') => {
    // 规范化地址
    const address = walletAddress.toLowerCase();
    
    if (API_DEBUG) {
      console.log(`登录请求 - 钱包地址: ${address}, 角色: ${role || 'default'}`);
      console.log(`签名: ${signature.substring(0, 20)}...`);
    }
    
    return apiClient.post<{ token: string; user: User }>('/auth/login', {
      wallet_address: address,
      signature,
      // 添加自定义标记和时间戳，避免服务端缓存
      _t: Date.now(),
      _r: Math.random().toString(36).substring(2, 10),
      ...(role && { role }),
    });
  },

  // 注册
  register: (walletAddress: string, role: 'parent' | 'child') => {
    // 规范化地址
    const address = walletAddress.toLowerCase();
    
    if (API_DEBUG) {
      console.log(`注册请求 - 钱包地址: ${address}, 角色: ${role}`);
    }
    
    return apiClient.post<{ user: User }>('/auth/register', {
      wallet_address: address,
      role,
      // 添加自定义标记和时间戳，避免服务端缓存
      _t: Date.now(),
      _r: Math.random().toString(36).substring(2, 10),
    });
  },

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

  // 删除儿童
  delete: (id: number) => apiClient.delete(`/children/${id}`),
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
  update: (id: number, data: any) => {
    console.log(`[API] 更新任务 ${id}，数据:`, data);
    return apiClient.put<Task>(`/tasks/${id}`, data);
  },

  // 完成任务
  complete: (id: number, submissionNote?: string) => {
    console.log(`[API] 完成任务 ${id}，备注:`, submissionNote);
    // 使用空对象作为默认请求体，确保即使没有备注也发送一个有效的JSON
    const payload = {};
    return apiClient.post<Task>(`/tasks/${id}/complete`, payload);
  },

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
export type { ApiResponse, User, Family, Child, Task, Reward, Exchange };

// 奖品相关 API
export const rewardApi = {
  // 创建奖品
  create: (familyId: number, rewardData: {
    name: string;
    description: string;
    image_url: string;
    token_price: number;
    stock: number; // 兼容旧接口，但前端固定为1
  }) => apiClient.post<Reward>(`/rewards/family/${familyId}`, rewardData),

  // 获取家庭奖品列表
  getAll: async (familyId: number, activeOnly: boolean = false) => {
    const query = activeOnly ? '?active_only=true' : '?active_only=false';
    // 添加日志记录，便于调试
    console.log(`[API] 获取家庭奖品列表, familyId: ${familyId}, activeOnly: ${activeOnly}, URL: /rewards/family/${familyId}${query}`);
    
    // 确保familyId是有效的数字
    if (!familyId || isNaN(Number(familyId))) {
      console.error(`[API] 无效的familyId: ${familyId}`);
      return { 
        success: false, 
        error: '无效的家庭ID', 
        status: 400 
      };
    }
    
    try {
      // 添加额外的URL参数，避免缓存问题
      const timestamp = Date.now();
      const modifiedQuery = `${query}&_t=${timestamp}`;
      const result = await apiClient.get<Reward[]>(`/rewards/family/${familyId}${modifiedQuery}`);
      console.log(`[API] 奖品列表响应:`, result);
      return result;
    } catch (error) {
      console.error(`[API] 获取奖品列表出错:`, error);
      return { 
        success: false, 
        error: '获取奖品列表出错', 
        status: 500 
      };
    }
  },

  // 获取奖品详情
  getById: (id: number) => apiClient.get<Reward>(`/rewards/${id}`),

  // 更新奖品信息
  update: (id: number, data: Partial<Reward>) =>
    apiClient.put<Reward>(`/rewards/${id}`, data),

  // 删除奖品
  delete: (id: number) => apiClient.delete(`/rewards/${id}`),
};

// 兑换相关 API
export const exchangeApi = {
  // 兑换奖品
  create: (data: { reward_id: number; child_id?: number; notes?: string }) =>
    apiClient.post<Exchange>('/exchanges', data),

  // 获取孩子的兑换记录
  getByChild: () => apiClient.get<Exchange[]>('/exchanges/my'),

  // 获取家庭的兑换记录
  getByFamily: (familyId: number) =>
    apiClient.get<Exchange[]>(`/families/${familyId}/exchanges`),

  // 获取兑换详情
  getById: (id: number) => apiClient.get<Exchange>(`/exchanges/${id}`),

  // 更新兑换状态
  updateStatus: (id: number, status: 'pending' | 'completed' | 'cancelled', notes?: string) =>
    apiClient.put<Exchange>(`/exchanges/${id}`, { status, notes }),

  // 批准兑换
  approve: (id: number, notes?: string) =>
    apiClient.put<Exchange>(`/exchanges/${id}`, { status: 'completed', notes }),

  // 取消兑换
  cancel: (id: number, notes?: string) =>
    apiClient.put<Exchange>(`/exchanges/${id}`, { status: 'cancelled', notes }),
};