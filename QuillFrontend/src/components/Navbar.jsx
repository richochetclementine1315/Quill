import { Link } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import QuillLogo from './QuillLogo';

export default function Navbar() {
  const { user, logout } = useAuth();

  return (
    <nav className="bg-white dark:bg-gray-800 shadow-lg">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
          <div className="flex items-center">
            <Link to="/" className="flex items-center space-x-2">
              <QuillLogo className="w-8 h-8 text-primary" />
              <span className="text-2xl font-bold text-primary">Quill</span>
            </Link>
            <div className="hidden md:flex ml-10 space-x-8">
              <Link to="/" className="text-gray-700 dark:text-gray-300 hover:text-primary px-3 py-2">
                Home
              </Link>
              {user && (
                <>
                  <Link to="/my-posts" className="text-gray-700 dark:text-gray-300 hover:text-primary px-3 py-2">
                    My Posts
                  </Link>
                  <Link to="/create" className="text-gray-700 dark:text-gray-300 hover:text-primary px-3 py-2">
                    Create Post
                  </Link>
                </>
              )}
            </div>
          </div>
          
          <div className="flex items-center space-x-4">
            {user ? (
              <button
                onClick={logout}
                className="bg-red-500 hover:bg-red-600 text-white px-4 py-2 rounded-md transition"
              >
                Logout
              </button>
            ) : (
              <>
                <Link
                  to="/login"
                  className="text-gray-700 dark:text-gray-300 hover:text-primary px-4 py-2"
                >
                  Login
                </Link>
                <Link
                  to="/register"
                  className="bg-primary hover:bg-blue-600 text-white px-4 py-2 rounded-md transition"
                >
                  Register
                </Link>
              </>
            )}
          </div>
        </div>
      </div>
    </nav>
  );
}
