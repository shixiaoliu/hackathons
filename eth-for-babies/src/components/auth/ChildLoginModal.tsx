import React, { useState } from 'react';
import { X, User, LogIn } from 'lucide-react';
import { useFamily } from '../../context/FamilyContext';
import Button from '../common/Button';
import Card, { CardBody, CardHeader } from '../common/Card';
import { Child } from '../../types/child';

interface ChildLoginModalProps {
  isOpen: boolean;
  onClose: () => void;
  onChildSelected: (child: Child) => void;
  availableChildren: Child[];
}

const ChildLoginModal: React.FC<ChildLoginModalProps> = ({ 
  isOpen, 
  onClose, 
  onChildSelected, 
  availableChildren 
}) => {
  const [selectedChildId, setSelectedChildId] = useState<string>('');

  const handleLogin = () => {
    const selectedChild = availableChildren.find(child => child.id === selectedChildId);
    if (selectedChild) {
      onChildSelected(selectedChild);
      onClose();
    }
  };

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-white rounded-lg shadow-xl max-w-md w-full mx-4">
        <Card>
          <CardHeader>
            <div className="flex justify-between items-center">
              <h3 className="text-lg font-semibold text-gray-900 flex items-center">
                <LogIn className="h-5 w-5 mr-2 text-primary-600" />
                Select Your Profile
              </h3>
              <button
                onClick={onClose}
                className="text-gray-400 hover:text-gray-600 transition-colors"
              >
                <X className="h-5 w-5" />
              </button>
            </div>
          </CardHeader>
          
          <CardBody>
            <p className="text-gray-600 mb-4">
              Multiple child profiles are associated with this wallet. Please select your profile:
            </p>
            
            <div className="space-y-3 mb-6">
              {availableChildren.map((child) => (
                <label
                  key={child.id}
                  className={`flex items-center p-3 border rounded-lg cursor-pointer transition-colors ${
                    selectedChildId === child.id
                      ? 'border-primary-500 bg-primary-50'
                      : 'border-gray-200 hover:border-gray-300'
                  }`}
                >
                  <input
                    type="radio"
                    name="child"
                    value={child.id}
                    checked={selectedChildId === child.id}
                    onChange={(e) => setSelectedChildId(e.target.value)}
                    className="sr-only"
                  />
                  <div className="flex items-center space-x-3 flex-1">
                    <div className="w-10 h-10 rounded-full bg-primary-100 flex items-center justify-center">
                      {child.avatar ? (
                        <img src={child.avatar} alt={child.name} className="w-10 h-10 rounded-full" />
                      ) : (
                        <User className="h-5 w-5 text-primary-600" />
                      )}
                    </div>
                    <div>
                      <p className="font-medium text-gray-900">{child.name}</p>
                      <p className="text-sm text-gray-500">Age {child.age}</p>
                    </div>
                  </div>
                  <div className={`w-4 h-4 rounded-full border-2 ${
                    selectedChildId === child.id
                      ? 'border-primary-500 bg-primary-500'
                      : 'border-gray-300'
                  }`}>
                    {selectedChildId === child.id && (
                      <div className="w-2 h-2 bg-white rounded-full mx-auto mt-0.5"></div>
                    )}
                  </div>
                </label>
              ))}
            </div>

            <div className="flex gap-3">
              <Button
                type="button"
                onClick={onClose}
                className="flex-1 bg-gray-200 text-gray-800 hover:bg-gray-300"
              >
                Cancel
              </Button>
              <Button
                type="button"
                onClick={handleLogin}
                disabled={!selectedChildId}
                className="flex-1"
              >
                <LogIn className="h-4 w-4 mr-2" />
                Login
              </Button>
            </div>
          </CardBody>
        </Card>
      </div>
    </div>
  );
};

export default ChildLoginModal;