import { ReactNode } from 'react';

interface CardProps {
  children: ReactNode;
  className?: string;
  onClick?: () => void;
  hoverable?: boolean;
}

const Card = ({ 
  children, 
  className = '',
  onClick,
  hoverable = false
}: CardProps) => {
  const hoverClass = hoverable 
    ? 'transition-transform duration-200 hover:-translate-y-1 hover:shadow-lg cursor-pointer' 
    : '';
  
  return (
    <div 
      className={`bg-white rounded-lg shadow-md overflow-hidden ${hoverClass} ${className}`}
      onClick={onClick}
    >
      {children}
    </div>
  );
};

export const CardHeader = ({ 
  children, 
  className = '' 
}: { 
  children: ReactNode;
  className?: string;
}) => {
  return (
    <div className={`px-6 py-4 border-b border-gray-200 ${className}`}>
      {children}
    </div>
  );
};

export const CardBody = ({ 
  children, 
  className = '' 
}: { 
  children: ReactNode;
  className?: string;
}) => {
  return (
    <div className={`px-6 py-4 ${className}`}>
      {children}
    </div>
  );
};

export const CardFooter = ({ 
  children, 
  className = '' 
}: { 
  children: ReactNode;
  className?: string;
}) => {
  return (
    <div className={`px-6 py-4 border-t border-gray-200 ${className}`}>
      {children}
    </div>
  );
};

export default Card;