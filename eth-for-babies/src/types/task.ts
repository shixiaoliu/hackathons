export interface Task {
  id: string;
  title: string;
  description: string;
  deadline: string;
  difficulty: 'easy' | 'medium' | 'hard';
  reward: string;
  status: 'open' | 'in-progress' | 'completed' | 'approved' | 'rejected';
  createdBy: string;
  assignedTo?: string;
  assignedChildId?: string; // 新增：指定分配给哪个孩子
  imageUrl?: string;
  completionCriteria: string;
  proof?: {
    images: string[];
    description: string;
    submittedAt: string;
  };
  createdAt: string;
  updatedAt: string;
}