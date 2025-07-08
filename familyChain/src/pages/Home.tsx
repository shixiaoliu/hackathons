import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAccount } from 'wagmi';
import { ConnectButton } from '@rainbow-me/rainbowkit';
import { CheckCircle, Shield, Award, ChevronRight, User, Settings } from 'lucide-react';
import Button from '../components/common/Button';
import { useUserRole } from '../context/UserRoleContext';
import { useFamily } from '../context/FamilyContext';
import { useAuthContext } from '../context/AuthContext';
import ChildLoginModal from '../components/auth/ChildLoginModal';
import { ApiTest } from '../components/ApiTest';
import { LoginModal } from '../components/LoginModal';

interface HomeProps {
  onLoginClick?: () => void;
}

const Home: React.FC<HomeProps> = ({ onLoginClick }) => {
  const { isConnected, address } = useAccount();
  const { user, isAuthenticated } = useAuthContext();
  const navigate = useNavigate();
  const { setUserRole } = useUserRole();
  const { getAllChildren, loginAsChild } = useFamily();
  const [showChildLoginModal, setShowChildLoginModal] = useState(false);
  const [showApiTest, setShowApiTest] = useState(false);
  const [showLoginModal, setShowLoginModal] = useState(false);

  // 获取当前用户的钱包地址（优先使用user.wallet_address，其次使用wagmi的address）
  const currentWalletAddress = user?.wallet_address || address;
  
  // 当用户成功登录后，自动关闭登录模态窗口
  useEffect(() => {
    if (isAuthenticated) {
      setShowLoginModal(false);
    }
  }, [isAuthenticated]);

  const handleRoleSelect = (role: 'parent' | 'child') => {
    if (role === 'child' && currentWalletAddress) {
      // 检查是否有多个child账户关联到这个钱包
      const allChildren = getAllChildren();
      
      const availableChildren = allChildren.filter(child => 
        child.walletAddress.toLowerCase() === currentWalletAddress.toLowerCase()
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
            onClick={() => setShowLoginModal(true)}
            className="flex items-center gap-2 px-4 py-2 text-sm bg-gray-100 hover:bg-gray-200 rounded-md transition-colors mr-2"
          > 
            <User className="w-4 h-4" />
            Login
          </button>
          <button
            onClick={() => setShowApiTest(true)}
            className="flex items-center gap-2 px-4 py-2 text-sm bg-gray-100 hover:bg-gray-200 rounded-md transition-colors"
          >
            <Settings className="w-4 h-4" />
            API Test
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
          <div className="flex flex-col sm:flex-row justify-center gap-6 mb-10">
            <Button 
              size="lg" 
              onClick={() => handleRoleSelect('parent')}
              rightIcon={<ChevronRight />}
              className="rounded-full shadow-lg px-8 py-4 text-lg font-semibold transition-transform duration-150 hover:scale-105 hover:shadow-xl bg-gradient-to-r from-primary-500 to-primary-400 text-white"
            >
              I'm a Parent
            </Button>
            <Button 
              size="lg" 
              variant="secondary" 
              onClick={() => handleRoleSelect('child')}
              rightIcon={<ChevronRight />}
              className="rounded-full shadow-lg px-8 py-4 text-lg font-semibold transition-transform duration-150 hover:scale-105 hover:shadow-xl bg-gradient-to-r from-blue-400 to-blue-500 text-white"
            >
              I'm a Child
            </Button>
          </div>
        )}
        
        <div className="mt-10 flex flex-col items-center">
          <div className="relative w-full max-w-4xl h-96 rounded-3xl overflow-hidden shadow-2xl border border-gray-200 bg-white">
            {/* Main image */}
            <img 
              src="https://images.pexels.com/photos/4262010/pexels-photo-4262010.jpeg" 
              alt="Family working together" 
              className="w-full h-full object-cover object-center"
            />
            {/* Overlay and text */}
            <div className="absolute inset-0 bg-gradient-to-b from-black/40 via-black/10 to-white/70 z-10 flex flex-col justify-between">
              <div className="p-8">
                <div className="text-4xl md:text-5xl font-extrabold text-white drop-shadow-lg tracking-wide mb-2">FamilyChain</div>
                <div className="text-lg md:text-2xl text-white/90 font-medium drop-shadow-md">A New Way for Family Tasks · Blockchain Rewards for Growth</div>
              </div>
              <div className="p-8 pb-6 text-lg text-gray-800 font-semibold drop-shadow-sm">
                <span className="bg-white/70 rounded-xl px-4 py-2">Make family collaboration fun and rewards transparent!</span>
              </div>
            </div>
          </div>
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

      {/* Login Modal */}
      <LoginModal
        isOpen={showLoginModal}
        onClose={() => setShowLoginModal(false)}
      />
    </div>
  );
};

export default Home;