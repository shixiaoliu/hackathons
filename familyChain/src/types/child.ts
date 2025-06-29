export interface Child {
  id: string;
  name: string;
  walletAddress: string;
  age: number;
  avatar?: string;
  parentAddress: string;
  createdAt: string;
  totalTasksCompleted: number;
  totalRewardsEarned: string;
}

export interface Family {
  id: string;
  parentAddress: string;
  children: Child[];
  createdAt: string;
}