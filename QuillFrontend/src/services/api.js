import axios from 'axios';

// Always use /api in production (Vercel proxy), direct URL in development
const API_BASE_URL = import.meta.env.DEV 
  ? 'http://localhost:8080/api'
  : '/api';

const api = axios.create({
  baseURL: API_BASE_URL,
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
  },
  timeout: 60000, // 60 second timeout for cold starts
});

// Add response interceptor for better error handling
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    // If 502 error (backend sleeping), retry once after delay
    if (error.response?.status === 502 && !error.config._retry) {
      error.config._retry = true;
      console.log('Backend is waking up... retrying in 5 seconds');
      await new Promise(resolve => setTimeout(resolve, 5000));
      return api(error.config);
    }
    return Promise.reject(error);
  }
);

// Auth API calls
export const authAPI = {
  register: (data) => api.post('/register', data),
  login: (data) => api.post('/login', data),
};

// Post API calls - matching reference backend routes
export const postAPI = {
  getAllPosts: (page = 1) => api.get(`/allpost?page=${page}`),
  getPostById: (id) => api.get(`/allpost/${id}`),
  createPost: (data) => api.post('/post', data),
  updatePost: (id, data) => api.put(`/updatepost/${id}`, data),
  deletePost: (id) => api.delete(`/deletepost/${id}`),
  getUserPosts: () => api.get('/uniquepost'),
};

// Image upload API
export const uploadImage = (file) => {
  const formData = new FormData();
  formData.append('image', file);
  
  // Use the configured api instance to ensure cookies are sent
  return api.post('/upload-image', formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    }
  });
};

export default api;
