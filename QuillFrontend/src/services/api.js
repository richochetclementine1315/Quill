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
  }
});

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
