import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { postAPI } from '../services/api';

export default function Home() {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [page, setPage] = useState(1);
  const [meta, setMeta] = useState({});

  useEffect(() => {
    fetchPosts();
  }, [page]);

  const fetchPosts = async () => {
    try {
      setError(null);
      const response = await postAPI.getAllPosts(page);
      setPosts(response.data.data);
      setMeta(response.data.meta);
    } catch (error) {
      console.error('Error fetching posts:', error);
      if (error.response?.status === 502) {
        setError('Backend is waking up from sleep (free tier). Please wait 30 seconds and refresh...');
      } else {
        setError('Failed to load posts. Please try again.');
      }
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return (
      <div className="flex flex-col justify-center items-center min-h-screen">
        <div className="text-xl text-gray-600 dark:text-gray-400 mb-2">Loading posts...</div>
        <div className="text-sm text-gray-500 dark:text-gray-500">
          (First load may take 30-60 seconds on free tier)
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex flex-col justify-center items-center min-h-screen">
        <div className="text-xl text-red-600 dark:text-red-400 mb-4">{error}</div>
        <button
          onClick={() => {
            setLoading(true);
            fetchPosts();
          }}
          className="bg-primary hover:bg-primary-dark text-white font-medium py-2 px-6 rounded-md transition"
        >
          Retry
        </button>
      </div>
    );
  }

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
      <div className="text-center mb-12">
        <h1 className="text-4xl font-bold text-gray-900 dark:text-white mb-4">
          Welcome to Quill
        </h1>
        <p className="text-xl text-gray-600 dark:text-gray-300">
          Discover amazing stories and share your own
        </p>
      </div>

      {posts.length === 0 ? (
        <div className="text-center text-gray-600 dark:text-gray-400">
          <p className="text-xl">No posts yet. Be the first to create one!</p>
        </div>
      ) : (
        <>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            {posts.map((post) => (
              <div key={post.id} className="bg-white dark:bg-gray-800 rounded-lg shadow-md overflow-hidden hover:shadow-lg transition">
                {post.image && (
                  <img
                    src={post.image}
                    alt={post.title}
                    className="w-full h-48 object-cover"
                  />
                )}
                <div className="p-6">
                  <h2 className="text-2xl font-bold text-gray-900 dark:text-white mb-2">
                    {post.title}
                  </h2>
                  <p className="text-gray-600 dark:text-gray-300 mb-4 line-clamp-3">
                    {post.desc}
                  </p>
                  <div className="flex items-center justify-between">
                    <div className="text-sm text-gray-500 dark:text-gray-400">
                      By {post.user?.first_name} {post.user?.last_name}
                    </div>
                    <Link
                      to={`/post/${post.id}`}
                      className="text-primary hover:text-blue-600 font-medium"
                    >
                      Read More →
                    </Link>
                  </div>
                </div>
              </div>
            ))}
          </div>

          {/* Pagination */}
          {meta.last_page > 1 && (
            <div className="mt-12 flex justify-center space-x-2">
              <button
                onClick={() => setPage(page - 1)}
                disabled={page === 1}
                className="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-md text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-700 disabled:opacity-50 disabled:cursor-not-allowed"
              >
                Previous
              </button>
              <span className="px-4 py-2 text-gray-700 dark:text-gray-300">
                Page {page} of {meta.last_page}
              </span>
              <button
                onClick={() => setPage(page + 1)}
                disabled={page >= meta.last_page}
                className="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-md text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-700 disabled:opacity-50 disabled:cursor-not-allowed"
              >
                Next
              </button>
            </div>
          )}
        </>
      )}
    </div>
  );
}
