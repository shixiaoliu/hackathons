import { useState } from 'react';
import { PlusCircle, Edit, Trash2, User } from 'lucide-react';
import { useFamily } from '../../context/FamilyContext';
import Button from '../common/Button';
import Card, { CardBody, CardHeader } from '../common/Card';
import AddChildModal from './AddChildModal';

const ChildrenManager = () => {
  const { children, selectChild, selectedChild, removeChild, isParent } = useFamily();
  const [showAddModal, setShowAddModal] = useState(false);
  const [childToDelete, setChildToDelete] = useState<string | null>(null);

  // 添加调试信息
  console.log('[ChildrenManager] children:', children);
  console.log('[ChildrenManager] isParent:', isParent);
  console.log('[ChildrenManager] children length:', children.length);

  const handleDeleteChild = (childId: string, e: React.MouseEvent) => {
    e.stopPropagation(); // 防止触发卡片的点击事件
    setChildToDelete(childId);
  };

  const confirmDelete = () => {
    if (childToDelete) {
      removeChild(childToDelete);
      setChildToDelete(null);
    }
  };

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h2 className="text-2xl font-bold text-gray-900">My Children</h2>
        <Button
          onClick={() => setShowAddModal(true)}
          leftIcon={<PlusCircle className="h-5 w-5" />}
        >
          Add Child
        </Button>
      </div>

      {children.length === 0 ? (
        <div className="text-center py-12">
          <User className="mx-auto h-12 w-12 text-gray-400" />
          <h3 className="mt-2 text-sm font-medium text-gray-900">No children added yet</h3>
          <p className="mt-1 text-sm text-gray-500">
            Get started by adding your first child to the family.
          </p>
          <div className="mt-6">
            <Button
              onClick={() => setShowAddModal(true)}
              leftIcon={<PlusCircle className="h-5 w-5" />}
            >
              Add Your First Child
            </Button>
          </div>
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {children.map((child) => (
            <Card 
              key={child.id} 
              className={`cursor-pointer transition-all ${
                selectedChild?.id === child.id 
                  ? 'ring-2 ring-primary-500 bg-primary-50' 
                  : 'hover:shadow-md'
              }`}
              onClick={() => selectChild(child.id)}
            >
              <CardHeader>
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-3">
                    <div className="w-12 h-12 rounded-full bg-primary-100 flex items-center justify-center">
                      {child.avatar ? (
                        <img src={child.avatar} alt={child.name} className="w-12 h-12 rounded-full" />
                      ) : (
                        <User className="h-6 w-6 text-primary-600" />
                      )}
                    </div>
                    <div>
                      <h3 className="font-semibold text-gray-900">{child.name}</h3>
                      <p className="text-sm text-gray-500">Age {child.age}</p>
                    </div>
                  </div>
                  <Button
                    variant="ghost"
                    size="sm"
                    onClick={(e) => handleDeleteChild(child.id, e)}
                    className="text-red-600 hover:text-red-700 hover:bg-red-50"
                  >
                    <Trash2 className="h-4 w-4" />
                  </Button>
                </div>
              </CardHeader>
              <CardBody>
                <div className="space-y-2">
                  <div className="flex justify-between text-sm">
                    <span className="text-gray-500">Tasks Completed:</span>
                    <span className="font-medium">{child.totalTasksCompleted}</span>
                  </div>
                  <div className="flex justify-between text-sm">
                    <span className="text-gray-500">Total Earned:</span>
                    <span className="font-medium">{child.totalRewardsEarned} ETH</span>
                  </div>
                  <div className="text-xs text-gray-400 truncate">
                    {child.walletAddress}
                  </div>
                </div>
              </CardBody>
            </Card>
          ))}
        </div>
      )}

      {showAddModal && (
        <AddChildModal
          isOpen={showAddModal}
          onClose={() => setShowAddModal(false)}
        />
      )}

      {/* Delete Confirmation Modal */}
      {childToDelete && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 max-w-sm w-full mx-4">
            <h3 className="text-lg font-semibold text-gray-900 mb-4">Delete Child</h3>
            <p className="text-gray-600 mb-6">
              Are you sure you want to delete this child? This action cannot be undone.
            </p>
            <div className="flex justify-end space-x-3">
              <Button
                variant="secondary"
                onClick={() => setChildToDelete(null)}
              >
                Cancel
              </Button>
              <Button
                variant="danger"
                onClick={confirmDelete}
              >
                Delete
              </Button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default ChildrenManager;