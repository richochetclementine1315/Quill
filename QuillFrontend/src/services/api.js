import axios from 'axios';

// Use environment variable in production, localhost in development
const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_BASE_URL,
  withCredentials: true, // Important: sends cookies with requests
  headers: {
    'Content-Type': 'application/json',
  }
});

// Auth API calls
export const authAPI = {
  register: (data) => api.post('/register', data),
  login: (data) => api.post('/login', data),
};

// Post API calls
export const postAPI = {
  getAllPosts: (page = 1) => api.get(`/allposts?page=${page}`),
  getPostById: (id) => api.get(`/allposts/${id}`),
  createPost: (data) => api.post('/post', data),
  updatePost: (id, data) => api.put(`/updatepost/${id}`, data),
  deletePost: (id) => api.delete(`/deletepost/${id}`),
  getUserPosts: () => api.get('/uniquepost'),
};

// Image upload API
export const uploadImage = (file) => {
  const formData = new FormData();
  formData.append('image', file);
  
  return axios.post(`${API_BASE_URL}/upload-image`, formData, {
    withCredentials: true,
    headers: {
      'Content-Type': 'multipart/form-data',
    }
  });
};

export default api;
