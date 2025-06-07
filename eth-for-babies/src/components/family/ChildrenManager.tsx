import { useState } from 'react';
import { PlusCircle, Edit, Trash2, User } from 'lucide-react';
import { useFamily } from '../../context/FamilyContext';
import Button from '../common/Button';
import Card, { CardBody, CardHeader } from '../common/Card';
import AddChildModal from './AddChildModal';

const ChildrenManager = () => {
  const { children, selectChild, selectedChild, removeChild, isParent } = useFamily();
  const [showAddModal, setShowAddModal] = useState(false);

  // 添加调试信息
  console.log('[ChildrenManager] children:', children);
  console.log('[ChildrenManager] isParent:', isParent);
  console.log('[ChildrenManager] children length:', children.length);

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
    </div>
  );
};

export default ChildrenManager;