import { Heart, Github } from 'lucide-react';

const Footer = () => {
  const currentYear = new Date().getFullYear();

  return (
    <footer className="bg-white border-t border-gray-200">
      <div className="container mx-auto px-4 py-6">
        <div className="flex flex-col md:flex-row justify-between items-center">
          <div className="mb-4 md:mb-0">
            <p className="text-gray-600 text-sm">
              &copy; {currentYear} FamilyChain. All rights reserved.
            </p>
          </div>
          
          <div className="flex items-center">
            <span className="text-gray-600 text-sm flex items-center">
              Made with <Heart className="h-4 w-4 text-error-500 mx-1" /> for families
            </span>
            <a 
              href="#" 
              className="ml-4 text-gray-600 hover:text-primary-600 transition-colors"
              aria-label="GitHub"
            >
              <Github className="h-5 w-5" />
            </a>
          </div>
        </div>
      </div>
    </footer>
  );
};

export default Footer;