import { Task } from '../types/task';

export const mockTasks: Task[] = [
  {
    id: '1',
    title: 'Clean Your Room',
    description: 'Make your bed, organize your desk, and vacuum the floor',
    deadline: '2023-12-31T23:59:59Z',
    difficulty: 'easy',
    reward: '0.01',
    status: 'open',
    createdBy: '0x123456789abcdef0123456789abcdef012345678',
    completionCriteria: 'Take before and after photos of your room',
    createdAt: '2023-12-01T10:00:00Z',
    updatedAt: '2023-12-01T10:00:00Z'
  },
  {
    id: '2',
    title: 'Do Math Homework',
    description: 'Complete pages 15-20 in your math workbook',
    deadline: '2023-12-25T18:00:00Z',
    difficulty: 'medium',
    reward: '0.02',
    status: 'in-progress',
    createdBy: '0x123456789abcdef0123456789abcdef012345678',
    assignedTo: '0xabcdef0123456789abcdef0123456789abcdef01',
    assignedChildId: '1',
    completionCriteria: 'Submit photos of all completed problems with work shown',
    createdAt: '2023-12-05T14:30:00Z',
    updatedAt: '2023-12-06T09:15:00Z'
  },
  {
    id: '3',
    title: 'Read a Book',
    description: 'Read a book of at least 100 pages and write a summary',
    deadline: '2023-12-28T20:00:00Z',
    difficulty: 'hard',
    reward: '0.05',
    status: 'completed',
    createdBy: '0x123456789abcdef0123456789abcdef012345678',
    assignedTo: '0xabcdef0123456789abcdef0123456789abcdef01',
    assignedChildId: '1',
    completionCriteria: 'Submit a 1-page summary of the book and a photo of you reading it',
    proof: {
      images: [
        'https://placehold.co/400x300?text=Reading+Book',
        'https://placehold.co/400x300?text=Book+Summary'
      ],
      description: 'I read "The Little Prince" and wrote a summary about the main character\'s journey.',
      submittedAt: '2023-12-20T15:45:00Z'
    },
    createdAt: '2023-12-10T11:20:00Z',
    updatedAt: '2023-12-20T15:45:00Z'
  },
  {
    id: '4',
    title: 'Take Out the Trash',
    description: 'Collect all trash from bins and take to outdoor garbage cans',
    deadline: '2023-12-22T19:00:00Z',
    difficulty: 'easy',
    reward: '0.01',
    status: 'approved',
    createdBy: '0x123456789abcdef0123456789abcdef012345678',
    assignedTo: '0xabcdef0123456789abcdef0123456789abcdef02',
    assignedChildId: '2',
    completionCriteria: 'Take photo of empty indoor bins and filled outdoor bin',
    proof: {
      images: [
        'https://placehold.co/400x300?text=Empty+Indoor+Bins',
        'https://placehold.co/400x300?text=Filled+Outdoor+Bin'
      ],
      description: 'I collected all the trash from the house and placed it in the outdoor bin.',
      submittedAt: '2023-12-21T18:30:00Z'
    },
    createdAt: '2023-12-20T10:00:00Z',
    updatedAt: '2023-12-21T19:15:00Z'
  },
  {
    id: '5',
    title: 'Help with Dinner',
    description: 'Help prepare dinner by setting the table and assisting with cooking',
    deadline: '2023-12-23T17:30:00Z',
    difficulty: 'medium',
    reward: '0.03',
    status: 'rejected',
    createdBy: '0x123456789abcdef0123456789abcdef012345678',
    assignedTo: '0xabcdef0123456789abcdef0123456789abcdef01',
    assignedChildId: '1',
    completionCriteria: 'Take photos of the set table and yourself helping with cooking',
    proof: {
      images: [
        'https://placehold.co/400x300?text=Set+Table'
      ],
      description: 'I set the table for dinner.',
      submittedAt: '2023-12-23T17:15:00Z'
    },
    createdAt: '2023-12-22T10:00:00Z',
    updatedAt: '2023-12-23T18:00:00Z'
  }
]; 