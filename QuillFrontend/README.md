# Quill Frontend

A modern React frontend for the Quill blog platform built with Vite, React Router, and TailwindCSS.

## 🚀 Features

- ✅ User authentication (Register/Login)
- ✅ JWT-based session management with HttpOnly cookies
- ✅ Create, read, update, and delete blog posts
- ✅ Image upload functionality
- ✅ Responsive design with TailwindCSS
- ✅ Pagination for blog posts
- ✅ Protected routes for authenticated users
- ✅ User-specific post management

## 📦 Tech Stack

- **React 18** - UI library
- **Vite** - Build tool and dev server
- **React Router v6** - Client-side routing
- **Axios** - HTTP client
- **TailwindCSS** - Utility-first CSS framework

## 🛠 Prerequisites

- Node.js 16+ and npm
- Backend server running on `http://localhost:8080`

## 📥 Installation

### 1. Install dependencies

```cmd
cd QuillFrontend
npm install
```

### 2. Start development server

```cmd
npm run dev
```

The app will run on `http://localhost:3000`

## 🗂 Project Structure

```
QuillFrontend/
├── index.html
├── package.json
├── vite.config.js
├── tailwind.config.js
├── postcss.config.js
│
└── src/
    ├── main.jsx              # App entry point
    ├── App.jsx               # Main app component with routing
    ├── index.css             # Global styles with Tailwind
    │
    ├── components/           # Reusable components
    │   ├── Navbar.jsx       # Navigation bar
    │   └── ProtectedRoute.jsx # Auth guard for protected routes
    │
    ├── pages/                # Page components
    │   ├── Home.jsx         # Homepage with all posts
    │   ├── Login.jsx        # Login page
    │   ├── Register.jsx     # Registration page
    │   ├── CreatePost.jsx   # Create new post
    │   ├── EditPost.jsx     # Edit existing post
    │   ├── PostDetail.jsx   # Single post view
    │   └── MyPosts.jsx      # User's posts dashboard
    │
    ├── context/              # React Context
    │   └── AuthContext.jsx  # Authentication state management
    │
    └── services/             # API services
        └── api.js           # Axios configuration and API calls
```

## 🔐 Authentication Flow

1. **Register**: User creates account → Backend hashes password → Success
2. **Login**: User logs in → Backend sets HttpOnly JWT cookie → Redirect to home
3. **Protected Routes**: Middleware checks for user session → Allow or redirect to login
4. **Logout**: Clear user state → Redirect to login

## 📡 API Integration

### Base Configuration

```javascript
baseURL: 'http://localhost:8080/api'
withCredentials: true  // Sends cookies with requests
```

### API Endpoints Used

**Authentication:**
- `POST /api/register` - Create new user
- `POST /api/login` - Login user

**Posts:**
- `GET /api/allposts?page=1` - Get all posts (paginated)
- `GET /api/allposts/:id` - Get single post
- `POST /api/post` - Create post (protected)
- `PUT /api/updatepost/:id` - Update post (protected)
- `DELETE /api/deletepost/:id` - Delete post (protected)
- `GET /api/uniquepost` - Get current user's posts (protected)

**Images:**
- `POST /api/upload-image` - Upload image (protected)
- `GET /api/uploads/:filename` - Static file serving

## 🎨 Key Features Explained

### 1. Authentication Context

```javascript
// Provides auth state across the app
const { user, login, register, logout } = useAuth();
```

### 2. Protected Routes

```javascript
// Redirects unauthenticated users to login
<ProtectedRoute>
  <CreatePost />
</ProtectedRoute>
```

### 3. Image Upload

```javascript
// Upload file and get URL
const response = await uploadImage(file);
const imageUrl = response.data.url;
```

### 4. Pagination

```javascript
// Navigate through pages
<button onClick={() => setPage(page + 1)}>Next</button>
```

## 🔧 Configuration

### Vite Proxy Configuration

The app uses Vite's proxy feature to avoid CORS issues during development:

```javascript
server: {
  proxy: {
    '/api': {
      target: 'http://localhost:8080',
      changeOrigin: true
    }
  }
}
```

## 🚦 Running in Production

### Build for production

```cmd
npm run build
```

### Preview production build

```cmd
npm run preview
```

## 🎯 Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run preview` - Preview production build

## 🔒 Security Features

- HttpOnly cookies prevent XSS attacks
- Protected routes require authentication
- Credentials included in all API requests
- Form validation on registration
- Password confirmation check

## 🐛 Common Issues & Solutions

### Issue: "unauthenticated" error on protected routes

**Solution**: Make sure you're logged in. The backend uses HttpOnly cookies which are automatically sent.

### Issue: CORS errors

**Solution**: 
1. Backend must allow credentials: `withCredentials: true`
2. Backend CORS config should allow the frontend origin
3. Use Vite proxy in development

### Issue: Images not loading

**Solution**: Check that:
1. Backend is running on port 8080
2. `uploads/` folder exists in backend
3. Image URLs are correct (http://localhost:8080/uploads/...)

## 🎨 Customization

### Change Primary Color

Edit `tailwind.config.js`:

```javascript
theme: {
  extend: {
    colors: {
      primary: '#3B82F6',  // Change this
    }
  }
}
```

### Change Backend URL

Edit `src/services/api.js`:

```javascript
baseURL: 'https://your-backend-url.com/api'
```

## 📱 Responsive Design

The app is fully responsive with breakpoints:
- Mobile: < 768px
- Tablet: 768px - 1024px
- Desktop: > 1024px

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Commit your changes: `git commit -m 'Add amazing feature'`
4. Push to the branch: `git push origin feature/amazing-feature`
5. Open a Pull Request

## 📝 License

This project is open source and available under the MIT License.

## 👤 Author

**Mrinmoy**
- GitHub: [@richochetclementine1315](https://github.com/richochetclementine1315)

---

**Need Help?** Check the backend README or open an issue on GitHub.
