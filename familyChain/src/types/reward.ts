export interface Reward {
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

export interface RewardCreateRequest {
  name: string;
  description: string;
  image_url: string;
  token_price: number;
  stock: number;
}

export interface RewardUpdateRequest {
  name?: string;
  description?: string;
  image_url?: string;
  token_price?: number;
  active?: boolean;
  stock?: number;
}

export type ExchangeStatus = 'pending' | 'completed' | 'cancelled';

export interface Exchange {
  id: number;
  reward_id: number;
  child_id: number;
  token_amount: number;
  status: ExchangeStatus;
  exchange_date: string;
  completed_date?: string;
  notes?: string;
  created_at: string;
  updated_at: string;
  reward_name?: string;
  reward_image?: string;
  child_name?: string;
}

export interface ExchangeUpdateRequest {
  status: ExchangeStatus;
  notes?: string;
} 