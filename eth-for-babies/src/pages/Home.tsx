import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAccount } from 'wagmi';
import { ConnectButton } from '@rainbow-me/rainbowkit';
import { CheckCircle, Shield, Award, ChevronRight, User, Settings } from 'lucide-react';
import Button from '../components/common/Button';
import { useUserRole } from '../context/UserRoleContext';
import { useFamily } from '../context/FamilyContext';
import ChildLoginModal from '../components/auth/ChildLoginModal';
import { ApiTest } from '../components/ApiTest';

interface HomeProps {
  onLoginClick?: () => void;
}

const Home: React.FC<HomeProps> = ({ onLoginClick }) => {
  const { isConnected, address } = useAccount();
  const navigate = useNavigate();
  const { setUserRole } = useUserRole();
  const { getAllChildren, loginAsChild } = useFamily();
  const [showChildLoginModal, setShowChildLoginModal] = useState(false);
  const [showApiTest, setShowApiTest] = useState(false);

  const handleRoleSelect = (role: 'arent' | 'child') => {
    if (role === 'child' && address) {
      // 检查是否有多个child账户关联到这个钱包
      const allChildren = getAllChildren();
      
      const availableChildren = allChildren.filter(child => 
        child.walletAddress.toLowerCase() === address.toLowerCase()
      );

      if (availableChildren.length === 0) {
        // 没有找到child账户，显示提示
        alert('No child account found for this wallet address. Please ask your parent to add you as a child first.');
        return;
      } else if (availableChildren.length === 1) {
        // 只有一个child账户，直接登录
        loginAsChild(availableChildren[0].walletAddress);
        setUserRole(role);
        navigate('/child');
      } else {
        // 多个child账户，显示选择器
        setShowChildLoginModal(true);
      }
    } else {
      setUserRole(role);
      navigate(role === 'parent' ? '/parent' : '/child');
    }
  };

  const handleChildSelected = (child: any) => {
    loginAsChild(child.walletAddress);
    setUserRole('child');
    navigate('/child');
  };

  return (
    <div className="max-w-5xl mx-auto">
      {/* Hero Section */}
      <section className="text-center py-12 md:py-20">
        <div className="flex justify-end mb-4">
          <button
            onClick={() => setShowApiTest(true)}
            className="flex items-center gap-2 px-4 py-2 text-sm bg-gray-100 hover:bg-gray-200 rounded-md transition-colors"
          >
            <Settings className="w-4 h-4" />
            API 测试
          </button>
        </div>
        <h1 className="text-4xl md:text-5xl font-bold text-gray-900 mb-4">
          Reward <span className="text-primary-600">Tasks</span> with <span className="text-primary-600">Blockchain</span>
        </h1>
        <p className="text-xl text-gray-600 mb-8 max-w-3xl mx-auto">
          A fun, secure way for families to incentivize chores and responsibilities using blockchain technology.
        </p>
        
        {!isConnected ? (
          <div className="flex justify-center mb-8">
            <ConnectButton label="Connect Wallet to Start" />
          </div>
        ) : (
          <div className="flex flex-col sm:flex-row justify-center gap-4 mb-8">
            <Button 
              size="lg" 
              onClick={() => handleRoleSelect('parent')}
              rightIcon={<ChevronRight />}
            >
              I'm a Parent
            </Button>
            <Button 
              size="lg" 
              variant="secondary" 
              onClick={() => handleRoleSelect('child')}
              rightIcon={<ChevronRight />}
            >
              I'm a Child
            </Button>
          </div>
        )}
        
        <div className="mt-10 relative">
          <div className="absolute inset-0 bg-gradient-to-b from-transparent to-gray-50 z-10"></div>
          <img 
            src="https://images.pexels.com/photos/4260477/pexels-photo-4260477.jpeg" 
            alt="Family using app together" 
            className="rounded-lg shadow-xl max-h-96 object-cover mx-auto"
          />
        </div>
      </section>

      {/* Features Section */}
      <section className="py-12 md:py-20">
        <h2 className="text-3xl font-bold text-center text-gray-900 mb-12">How It Works</h2>
        
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div className="text-center p-6 rounded-lg bg-white shadow-md">
            <div className="inline-flex items-center justify-center w-16 h-16 rounded-full bg-primary-100 text-primary-600 mb-4">
              <Shield className="h-8 w-8" />
            </div>
            <h3 className="text-xl font-semibold mb-2">Create Tasks</h3>
            <p className="text-gray-600">Parents create tasks with clear descriptions, deadlines, and ETH rewards.</p>
          </div>
          
          <div className="text-center p-6 rounded-lg bg-white shadow-md">
            <div className="inline-flex items-center justify-center w-16 h-16 rounded-full bg-secondary-100 text-secondary-600 mb-4">
              <CheckCircle className="h-8 w-8" />
            </div>
            <h3 className="text-xl font-semibold mb-2">Complete & Submit</h3>
            <p className="text-gray-600">Children complete tasks and submit proof with photos or videos.</p>
          </div>
          
          <div className="text-center p-6 rounded-lg bg-white shadow-md">
            <div className="inline-flex items-center justify-center w-16 h-16 rounded-full bg-accent-100 text-accent-400 mb-4">
              <Award className="h-8 w-8" />
            </div>
            <h3 className="text-xl font-semibold mb-2">Earn Rewards</h3>
            <p className="text-gray-600">Smart contracts automatically distribute ETH rewards upon task approval.</p>
          </div>
        </div>
      </section>

      {/* Benefits Section */}
      <section className="py-12 md:py-20 bg-primary-50 rounded-xl">
        <div className="container mx-auto px-4">
          <h2 className="text-3xl font-bold text-center text-gray-900 mb-12">Benefits</h2>
          
          <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div className="flex items-start">
              <div className="flex-shrink-0 mt-1">
                <div className="w-10 h-10 rounded-full bg-primary-600 text-white flex items-center justify-center">
                  <CheckCircle className="h-5 w-5" />
                </div>
              </div>
              <div className="ml-4">
                <h3 className="text-xl font-semibold mb-2">Teaches Financial Responsibility</h3>
                <p className="text-gray-600">Kids learn the value of work and how to manage digital currency.</p>
              </div>
            </div>
            
            <div className="flex items-start">
              <div className="flex-shrink-0 mt-1">
                <div className="w-10 h-10 rounded-full bg-primary-600 text-white flex items-center justify-center">
                  <CheckCircle className="h-5 w-5" />
                </div>
              </div>
              <div className="ml-4">
                <h3 className="text-xl font-semibold mb-2">Builds Trust Through Transparency</h3>
                <p className="text-gray-600">Blockchain provides a transparent, immutable record of all tasks and payments.</p>
              </div>
            </div>
            
            <div className="flex items-start">
              <div className="flex-shrink-0 mt-1">
                <div className="w-10 h-10 rounded-full bg-primary-600 text-white flex items-center justify-center">
                  <CheckCircle className="h-5 w-5" />
                </div>
              </div>
              <div className="ml-4">
                <h3 className="text-xl font-semibold mb-2">Gamifies Household Chores</h3>
                <p className="text-gray-600">Turns everyday responsibilities into engaging challenges with rewards.</p>
              </div>
            </div>
            
            <div className="flex items-start">
              <div className="flex-shrink-0 mt-1">
                <div className="w-10 h-10 rounded-full bg-primary-600 text-white flex items-center justify-center">
                  <CheckCircle className="h-5 w-5" />
                </div>
              </div>
              <div className="ml-4">
                <h3 className="text-xl font-semibold mb-2">Secure Smart Contracts</h3>
                <p className="text-gray-600">Automated payments ensure children receive rewards when tasks are approved.</p>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-12 md:py-20 text-center">
        <h2 className="text-3xl font-bold text-gray-900 mb-4">Ready to Transform Family Tasks?</h2>
        <p className="text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
          Connect your wallet and start creating blockchain-powered incentives for your family today.
        </p>
        
        {!isConnected ? (
          <div className="flex justify-center">
            <ConnectButton label="Get Started" />
          </div>
        ) : (
          <div className="flex flex-col sm:flex-row justify-center gap-4">
            <Button 
              size="lg" 
              onClick={() => handleRoleSelect('parent')}
            >
              Parent Dashboard
            </Button>
            <Button 
              size="lg" 
              variant="secondary" 
              onClick={() => handleRoleSelect('child')}
            >
              Child Dashboard
            </Button>
          </div>
        )}
      </section>

      {/* Child Login Modal */}
      {showChildLoginModal && (
        <ChildLoginModal
          isOpen={showChildLoginModal}
          onClose={() => setShowChildLoginModal(false)}
          onChildSelected={handleChildSelected}
          availableChildren={getAllChildren().filter(child => 
            child.walletAddress.toLowerCase() === address?.toLowerCase()
          )}
        />
      )}
      
      <ApiTest
        isOpen={showApiTest}
        onClose={() => setShowApiTest(false)}
      />
    </div>
  );
};

export default Home;