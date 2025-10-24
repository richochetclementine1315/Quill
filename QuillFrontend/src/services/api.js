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
  timeout: 90000, // 90 second timeout for cold starts
});

// Wake up backend (ping health check - lightweight endpoint)
let backendWakePromise = null;
export const wakeBackend = () => {
  if (!backendWakePromise) {
    const healthUrl = import.meta.env.DEV 
      ? 'http://localhost:8080/api/health' 
      : 'https://quill-backend-lgxs.onrender.com/api/health';
    
    backendWakePromise = axios.get(healthUrl, { timeout: 90000 })
      .then(() => {
        console.log('Backend is awake!');
        return true;
      })
      .catch((err) => {
        console.log('Waking backend...', err.message);
        return false;
      });
  }
  return backendWakePromise;
};

// Add response interceptor for better error handling
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;
    
    // If 502/503/504 error (backend sleeping), retry with exponential backoff
    if ([502, 503, 504].includes(error.response?.status) && !originalRequest._retryCount) {
      originalRequest._retryCount = 0;
    }
    
    if (originalRequest._retryCount !== undefined && originalRequest._retryCount < 3) {
      originalRequest._retryCount++;
      const delay = Math.min(1000 * Math.pow(2, originalRequest._retryCount), 10000);
      console.log(`Backend waking up... retry ${originalRequest._retryCount}/3 in ${delay/1000}s`);
      await new Promise(resolve => setTimeout(resolve, delay));
      return api(originalRequest);
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
