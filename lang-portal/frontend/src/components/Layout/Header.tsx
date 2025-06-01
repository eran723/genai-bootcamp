
import React from 'react';
import { Clock, Target, Flame } from 'lucide-react';

const Header: React.FC = () => {
  return (
    <header className="bg-white border-b border-gray-200 px-6 py-4">
      <div className="flex items-center justify-between">
        <div>
          <h2 className="text-2xl font-semibold text-gray-900">
            Japanese Tutor Dashboard
          </h2>
          <p className="text-sm text-gray-600">
            {new Date().toLocaleDateString('en-US', { 
              weekday: 'long', 
              year: 'numeric', 
              month: 'long', 
              day: 'numeric' 
            })}
          </p>
        </div>
        
        <div className="flex items-center space-x-6">
          <div className="flex items-center space-x-2 text-orange-600">
            <Flame size={20} />
            <span className="font-medium">7 day streak</span>
          </div>
          <div className="flex items-center space-x-2 text-green-600">
            <Target size={20} />
            <span className="font-medium">85% accuracy</span>
          </div>
          <div className="flex items-center space-x-2 text-blue-600">
            <Clock size={20} />
            <span className="font-medium">2h 15m today</span>
          </div>
        </div>
      </div>
    </header>
  );
};

export default Header;
