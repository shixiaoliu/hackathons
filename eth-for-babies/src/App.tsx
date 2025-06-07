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
import { LoginModal } from './components/LoginModal';
import { useAuthContext } from './context/AuthContext';
import { UserRoleProvider } from './context/UserRoleContext';
import { FamilyProvider } from './context/FamilyContext';
import { TaskProvider } from './context/TaskContext';

function App() {
  const { isConnected } = useAccount();
  const { isAuthenticated, user } = useAuthContext();
  const [showLoginModal, setShowLoginModal] = useState(false);

  // 检查是否需要显示登录模态框
  const shouldShowLogin = isConnected && !isAuthenticated;

  // 当钱包连接但未认证时，自动显示登录模态框
  useEffect(() => {
    if (shouldShowLogin) {
      setShowLoginModal(true);
    }
  }, [shouldShowLogin]);

  // 根据用户角色决定重定向
  const getDefaultRoute = () => {
    if (!isAuthenticated || !user) return '/';
    return user.role === 'parent' ? '/parent' : '/child';
  };

  return (
    <UserRoleProvider>
      <FamilyProvider>
        <TaskProvider>
          <Layout>
            <Routes>
              <Route 
                path="/" 
                element={
                  isAuthenticated ? 
                    <Navigate to={getDefaultRoute()} /> : 
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
            </Routes>
            
            {/* 登录模态框 */}
            <LoginModal 
              isOpen={shouldShowLogin || showLoginModal}
              onClose={() => setShowLoginModal(false)}
            />
          </Layout>
        </TaskProvider>
      </FamilyProvider>
    </UserRoleProvider>
  );
}

export default App;