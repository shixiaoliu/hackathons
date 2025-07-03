import { useState, useEffect } from 'react';
import { Routes, Route, Navigate } from 'react-router-dom';
import { useAccount } from 'wagmi';
import Layout from './components/common/Layout';
import Home from './pages/Home';
import ParentDashboard from './pages/ParentDashboard';
import ChildDashboard from './pages/ChildDashboard';
import TaskDetail from './pages/TaskDetail';
import CreateTask from './pages/CreateTask';
import SubmitTask from './pages/SubmitTask';
import RewardManagement from './pages/RewardManagement';
import RewardStore from './pages/RewardStore';
import { LoginModal } from './components/LoginModal';
import { useAuthContext } from './context/AuthContext';
import { UserRoleProvider } from './context/UserRoleContext';
import { FamilyProvider } from './context/FamilyContext';
import { TaskProvider } from './context/TaskContext';
import { RewardProvider } from './context/RewardContext';
import TaskList from './pages/TaskList';

function App() {
  const { isConnected } = useAccount();
  const { isAuthenticated, user, isLoading } = useAuthContext();
  const [showLoginModal, setShowLoginModal] = useState(false);
  const [authChecked, setAuthChecked] = useState(false);

  // 检查是否需要显示登录模态框
  const shouldShowLogin = isConnected && !isAuthenticated;

  // 当钱包连接但未认证时，自动显示登录模态框
  useEffect(() => {
    if (shouldShowLogin) {
      setShowLoginModal(true);
    }
  }, [shouldShowLogin]);

  // 确保认证状态已经检查完毕
  useEffect(() => {
    if (!isLoading) {
      setAuthChecked(true);
    }
  }, [isLoading]);

  // 根据用户角色决定重定向
  const getDefaultRoute = () => {
    if (!isAuthenticated || !user) return '/';
    return user.role === 'parent' ? '/parent' : '/child';
  };

  // 如果认证状态仍在加载中，显示加载状态
  if (!authChecked) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary-600"></div>
      </div>
    );
  }

  return (
    <UserRoleProvider>
      <FamilyProvider>
        <TaskProvider>
          <RewardProvider>
            <Layout>
              <Routes>
                <Route 
                  path="/" 
                  element={ 
                      <Home onLoginClick={() => setShowLoginModal(true)} />
                  } 
                />
                <Route 
                  path="/parent/*" 
                  element={
                    isAuthenticated && user?.role === 'parent' ? 
                      <ParentDashboard /> : 
                      <Navigate to="/" />
                  } 
                />
                <Route 
                  path="/child/*" 
                  element={
                    isAuthenticated && user?.role === 'child' ? 
                      <ChildDashboard /> : 
                      <Navigate to="/" />
                  } 
                />
                <Route 
                  path="/tasks" 
                  element={
                    isAuthenticated ? 
                      <TaskList /> : 
                      <Navigate to="/" />
                  } 
                />
                <Route 
                  path="/task/:id" 
                  element={isAuthenticated ? <TaskDetail /> : <Navigate to="/" />} 
                />
                <Route 
                  path="/create-task" 
                  element={
                    isAuthenticated && user?.role === 'parent' ? 
                      <CreateTask /> : 
                      <Navigate to="/" />
                  } 
                />
                <Route 
                  path="/submit-task/:id" 
                  element={
                    isAuthenticated && user?.role === 'child' ? 
                      <SubmitTask /> : 
                      <Navigate to="/" />
                  } 
                />
                <Route 
                  path="/rewards" 
                  element={
                    isAuthenticated && user?.role === 'parent' ? 
                      <RewardManagement /> : 
                      <Navigate to="/" />
                  } 
                />
                <Route 
                  path="/reward-store" 
                  element={
                    isAuthenticated && user?.role === 'child' ? 
                      <RewardStore /> : 
                      <Navigate to="/" />
                  } 
                />
              </Routes>
              
              {/* 登录模态框 */}
              <LoginModal 
                isOpen={shouldShowLogin || showLoginModal}
                onClose={() => setShowLoginModal(false)}
              />
            </Layout>
          </RewardProvider>
        </TaskProvider>
      </FamilyProvider>
    </UserRoleProvider>
  );
}

export default App;