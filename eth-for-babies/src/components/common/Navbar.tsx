import { Link, useLocation } from 'react-router-dom';
import { ConnectButton } from '@rainbow-me/rainbowkit';
import { Menu, Home, Award, ChevronDown } from 'lucide-react';
import { useState } from 'react';
import { useUserRole } from '../../context/UserRoleContext';

const Navbar = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const { userRole } = useUserRole();
  const location = useLocation();

  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  const dashboardLink = userRole === 'parent' 
    ? '/parent' 
    : userRole === 'child' 
      ? '/child' 
      : '/';

  return (
    <nav className="bg-white shadow-md">
      <div className="container mx-auto px-4">
        <div className="flex justify-between items-center h-16">
          <div className="flex items-center">
            <Link to="/" className="flex items-center">
              <Award className="h-8 w-8 text-primary-600" />
              <span className="ml-2 text-xl font-bold text-gray-800">FamilyChain</span>
            </Link>
          </div>

          {/* Desktop Navigation */}
          <div className="hidden md:flex items-center space-x-4">
            <Link 
              to="/" 
              className={`px-3 py-2 rounded-md text-sm font-medium ${
                location.pathname === '/' 
                  ? 'text-primary-600 bg-primary-50' 
                  : 'text-gray-700 hover:bg-gray-100'
              }`}
            >
              <Home className="inline-block mr-1 h-4 w-4" />
              Home
            </Link>
            
            {userRole && (
              <Link 
                to={dashboardLink} 
                className={`px-3 py-2 rounded-md text-sm font-medium ${
                  location.pathname.includes(dashboardLink) 
                    ? 'text-primary-600 bg-primary-50' 
                    : 'text-gray-700 hover:bg-gray-100'
                }`}
              >
                Dashboard
              </Link>
            )}
            
            <div className="ml-4">
              <ConnectButton />
            </div>
          </div>

          {/* Mobile Navigation Button */}
          <div className="md:hidden flex items-center">
            <button
              onClick={toggleMenu}
              className="inline-flex items-center justify-center p-2 rounded-md text-gray-700 hover:text-primary-600 hover:bg-gray-100 focus:outline-none"
            >
              <Menu className="h-6 w-6" />
            </button>
          </div>
        </div>
      </div>

      {/* Mobile Menu */}
      {isMenuOpen && (
        <div className="md:hidden">
          <div className="px-2 pt-2 pb-3 space-y-1 sm:px-3">
            <Link 
              to="/" 
              className={`block px-3 py-2 rounded-md text-base font-medium ${
                location.pathname === '/' 
                  ? 'text-primary-600 bg-primary-50' 
                  : 'text-gray-700 hover:bg-gray-100'
              }`}
              onClick={() => setIsMenuOpen(false)}
            >
              <Home className="inline-block mr-1 h-4 w-4" />
              Home
            </Link>
            
            {userRole && (
              <Link 
                to={dashboardLink} 
                className={`block px-3 py-2 rounded-md text-base font-medium ${
                  location.pathname.includes(dashboardLink) 
                    ? 'text-primary-600 bg-primary-50' 
                    : 'text-gray-700 hover:bg-gray-100'
                }`}
                onClick={() => setIsMenuOpen(false)}
              >
                Dashboard
              </Link>
            )}
            
            <div className="mt-4 px-3">
              <ConnectButton />
            </div>
          </div>
        </div>
      )}
    </nav>
  );
};

export default Navbar;