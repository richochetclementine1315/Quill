# Quill Backend API

A RESTful blog backend API built with Go, Fiber, GORM, and MySQL. Features JWT-based authentication, user management, and complete blog post CRUD operations.

## 📋 Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Installation & Setup](#installation--setup)
- [Environment Variables](#environment-variables)
- [Database Setup](#database-setup)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
  - [Authentication Endpoints](#authentication-endpoints)
  - [Blog Post Endpoints](#blog-post-endpoints)
- [Authentication & Authorization](#authentication--authorization)
- [Project Structure](#project-structure)
- [Error Handling](#error-handling)
- [Security Features](#security-features)

## ✨ Features

- ✅ User registration and login with JWT authentication
- ✅ Password hashing with bcrypt
- ✅ Protected routes with middleware
- ✅ Complete CRUD operations for blog posts
- ✅ Pagination support for posts
- ✅ User-specific post management
- ✅ Foreign key relationships between users and posts
- ✅ Email validation
- ✅ HTTPOnly cookies for secure token storage

## 🛠 Tech Stack

- **Framework**: [Fiber v2](https://gofiber.io/) - Fast HTTP web framework
- **ORM**: [GORM](https://gorm.io/) - Go Object Relational Mapping
- **Database**: MySQL 8.0
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt
- **Environment**: godotenv

## 📦 Prerequisites

Before running this application, ensure you have:

- Go 1.24.0 or higher
- MySQL 8.0 or higher (or Docker with MySQL image)
- Git

## 🚀 Installation & Setup

### 1. Clone the repository

```bash
git clone https://github.com/richochetclementine1315/Quill.git
cd Quill/QuillBackend
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Set up environment variables

Create a `.env` file in the `QuillBackend` directory:

```env
PORT=YOUR_PORT
DSN=root:password@tcp(127.0.0.1:3306)/quilldb?charset=utf8mb4&parseTime=True&loc=Local
```

Replace `root:password` with your MySQL credentials.

## 🗄 Database Setup

You have three options to set up MySQL for this project:

### Option 1: Using XAMPP (Easiest for Windows)

**This project uses XAMPP for MySQL.**

1. **Download and install XAMPP** from [https://www.apachefriends.org/](https://www.apachefriends.org/)

2. **Start MySQL** from XAMPP Control Panel:
   - Open XAMPP Control Panel
   - Click "Start" button next to MySQL
   - MySQL will run on `localhost:3306` by default

3. **Create the database**:
   - Open phpMyAdmin: `http://localhost/phpmyadmin`
   - Click "New" in the left sidebar
   - Database name: `quilldb`
   - Collation: `utf8mb4_unicode_ci`
   - Click "Create"

4. **Update your `.env` file**:
   ```env
   PORT=8080
   DSN=root:@tcp(127.0.0.1:3306)/quilldb?charset=utf8mb4&parseTime=True&loc=Local
   ```
   Note: Default XAMPP MySQL has no password for root user (empty password after `root:`)

**XAMPP Advantages**:
- ✅ Easy GUI-based setup
- ✅ Includes phpMyAdmin for database management
- ✅ No command-line required
- ✅ All-in-one package (Apache, MySQL, PHP)
- ✅ Perfect for Windows development

### Option 2: Using Docker (Recommended for cross-platform)

Create a `docker-compose.yml` file:

```yaml
version: "3.8"
services:
  db:
    image: mysql:8.0
    container_name: quill-db
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: quilldb
      MYSQL_USER: quilluser
      MYSQL_PASSWORD: quillpass
    ports:
      - "3306:3306"
    volumes:
      - db_data:/var/lib/mysql

volumes:
  db_data:
```

Start the database:

```bash
docker compose up -d
```

### Option 3: Manual MySQL Setup

1. Install MySQL locally
2. Create a database:

```sql
CREATE DATABASE quilldb CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### Auto-Migration

The application uses GORM's AutoMigrate feature to automatically create tables on startup. Tables created:
- `users` - Stores user accounts
- `blogs` - Stores blog posts with foreign key to users

## 🏃 Running the Application

```bash
go run main.go
```

The server will start on `http://localhost:8080` (or the port specified in `.env`)

## 📚 API Documentation

Base URL: `http://localhost:8080`

### Authentication Endpoints

#### 1. Register User

Create a new user account.

**Endpoint**: `POST /api/register`

**Access**: Public (no authentication required)

**Request Body**:
```json
{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com",
  "phone": "+1234567890",
  "password": "SecurePass123"
}
```

**Validation Rules**:
- `password`: Must be longer than 6 characters
- `email`: Must be valid email format
- Email must not already exist in database

**Success Response** (200 OK):
```json
{
  "user": {
    "id": 1,
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com",
    "phone": "+1234567890"
  },
  "message": "Registration successful"
}
```

**Error Responses**:
- `400 Bad Request`: Password too short, invalid email, or email already registered
- `500 Internal Server Error`: Database error

---

#### 2. Login

Authenticate user and receive JWT token.

**Endpoint**: `POST /api/login`

**Access**: Public

**Request Body**:
```json
{
  "email": "john.doe@example.com",
  "password": "SecurePass123"
}
```

**Success Response** (200 OK):
```json
{
  "message": "Login successful"
}
```

**Cookie Set**: 
- Name: `jwt`
- HttpOnly: `true`
- Expires: 24 hours
- Contains: JWT token with user ID

**Error Responses**:
- `404 Not Found`: User not found
- `400 Bad Request`: Incorrect password
- `500 Internal Server Error`: Token generation failed

---

### Blog Post Endpoints

All blog endpoints require authentication (JWT cookie).

#### 3. Create Post

Create a new blog post.

**Endpoint**: `POST /api/post`

**Access**: Protected (requires authentication)

**Headers**: 
- Cookie: `jwt=<token>` (automatically sent by browser after login)

**Request Body**:
```json
{
  "title": "My First Blog Post",
  "desc": "This is the content of my blog post...",
  "image": "https://example.com/image.jpg"
}
```

**Success Response** (200 OK):
```json
{
  "message": "Blog Post Created Successfully!",
  "post": {
    "id": 1,
    "title": "My First Blog Post",
    "desc": "This is the content of my blog post...",
    "image": "https://example.com/image.jpg",
    "userid": 1,
    "user": {
      "id": 1,
      "first_name": "John",
      "last_name": "Doe",
      "email": "john.doe@example.com",
      "phone": "+1234567890"
    }
  }
}
```

**Error Responses**:
- `401 Unauthorized`: Missing or invalid JWT token
- `400 Bad Request`: Invalid payload or foreign key constraint violation

---

#### 4. Get All Posts (Paginated)

Retrieve all blog posts with pagination.

**Endpoint**: `GET /api/allposts?page=1`

**Access**: Protected

**Query Parameters**:
- `page` (optional, default: 1): Page number for pagination

**Success Response** (200 OK):
```json
{
  "data": [
    {
      "id": 1,
      "title": "My First Blog Post",
      "desc": "Content...",
      "image": "https://example.com/image.jpg",
      "userid": 1,
      "user": {
        "id": 1,
        "first_name": "John",
        "last_name": "Doe",
        "email": "john.doe@example.com",
        "phone": "+1234567890"
      }
    }
  ],
  "meta": {
    "page": 1,
    "total": 15,
    "last_page": 3
  }
}
```

**Pagination Details**:
- Items per page: 5
- Includes user information via preload
- Returns total count and last page number

**Error Responses**:
- `401 Unauthorized`: Missing or invalid JWT token

---

#### 5. Get Single Post Detail

Retrieve details of a specific blog post by ID.

**Endpoint**: `GET /api/allposts/:id`

**Access**: Protected

**URL Parameters**:
- `id`: Post ID (integer)

**Example**: `GET /api/allposts/1`

**Success Response** (200 OK):
```json
{
  "data": {
    "id": 1,
    "title": "My First Blog Post",
    "desc": "Content...",
    "image": "https://example.com/image.jpg",
    "userid": 1,
    "user": {
      "id": 1,
      "first_name": "John",
      "last_name": "Doe",
      "email": "john.doe@example.com",
      "phone": "+1234567890"
    }
  }
}
```

**Error Responses**:
- `401 Unauthorized`: Missing or invalid JWT token
- Returns empty object if post not found

---

#### 6. Update Post

Update an existing blog post.

**Endpoint**: `PUT /api/updatepost/:id`

**Access**: Protected

**URL Parameters**:
- `id`: Post ID to update

**Request Body** (all fields optional):
```json
{
  "title": "Updated Title",
  "desc": "Updated content...",
  "image": "https://example.com/new-image.jpg"
}
```

**Success Response** (200 OK):
```json
{
  "id": 1,
  "title": "Updated Title",
  "desc": "Updated content...",
  "image": "https://example.com/new-image.jpg",
  "userid": 1
}
```

**Error Responses**:
- `401 Unauthorized`: Missing or invalid JWT token
- `400 Bad Request`: Invalid payload

---

#### 7. Get User's Own Posts

Retrieve all posts created by the currently authenticated user.

**Endpoint**: `GET /api/uniquepost`

**Access**: Protected

**Success Response** (200 OK):
```json
[
  {
    "id": 1,
    "title": "My First Post",
    "desc": "Content...",
    "image": "https://example.com/image.jpg",
    "userid": 1,
    "user": {
      "id": 1,
      "first_name": "John",
      "last_name": "Doe",
      "email": "john.doe@example.com",
      "phone": "+1234567890"
    }
  },
  {
    "id": 5,
    "title": "My Second Post",
    "desc": "More content...",
    "image": "https://example.com/image2.jpg",
    "userid": 1,
    "user": { /* ... */ }
  }
]
```

**How it works**:
- Extracts user ID from JWT token
- Filters posts by `user_id`
- Preloads user information

**Error Responses**:
- `401 Unauthorized`: Missing or invalid JWT token

---

#### 8. Delete Post

Delete a blog post by ID.

**Endpoint**: `DELETE /api/deletepost/:id`

**Access**: Protected

**URL Parameters**:
- `id`: Post ID to delete

**Example**: `DELETE /api/deletepost/1`

**Success Response** (200 OK):
```json
{
  "message": "Post deleted successfully"
}
```

**Error Responses**:
- `401 Unauthorized`: Missing or invalid JWT token
- `400 Bad Request`: Record not found

---

## 🔐 Authentication & Authorization

### How Authentication Works

1. **Registration**: 
   - User submits credentials
   - Password is hashed using bcrypt (cost: 14)
   - User stored in database

2. **Login**:
   - User submits email and password
   - Backend verifies credentials
   - JWT token generated with user ID as issuer
   - Token stored in HttpOnly cookie (expires in 24 hours)

3. **Protected Routes**:
   - Client automatically sends JWT cookie with each request
   - Middleware intercepts request
   - Validates JWT token
   - Extracts user ID from token
   - Allows or denies access

### JWT Token Structure

```javascript
{
  "iss": "1",              // User ID (issuer)
  "exp": 1729756800        // Expiration timestamp
}
```

**Secret Key**: Stored in `utils/helper.go` (should be moved to environment variable for production)

### Middleware Flow

```
Request → IsAuthenticate Middleware → Parse JWT Cookie → Validate Token → Next() or 401
```

### Security Best Practices Implemented

✅ **Password Security**:
- Bcrypt hashing with cost 14
- Passwords never returned in API responses (`json:"-"`)

✅ **Token Security**:
- HttpOnly cookies (prevents XSS attacks)
- 24-hour expiration
- Secure token validation

✅ **Input Validation**:
- Email format validation (regex)
- Password length requirement (>6 characters)
- Duplicate email check

✅ **Database Security**:
- Foreign key constraints
- GORM parameterized queries (prevents SQL injection)

## 📁 Project Structure

```
QuillBackend/
├── main.go                 # Application entry point
├── .env                    # Environment variables
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
│
├── controller/             # Request handlers
│   ├── authcontroller.go  # Register & Login handlers
│   └── postcontroller.go  # Blog post CRUD handlers
│
├── database/               # Database connection
│   └── connect.go         # GORM MySQL connection & AutoMigrate
│
├── middleware/             # HTTP middleware
│   └── middleware.go      # JWT authentication middleware
│
├── models/                 # Database models
│   ├── user.go            # User model & password methods
│   └── blog.go            # Blog post model
│
├── routes/                 # Route definitions
│   └── route.go           # All API route mappings
│
└── utils/                  # Helper functions
    └── helper.go          # JWT generation & parsing
```

## ⚠️ Error Handling

### Common Error Responses

**401 Unauthorized**:
```json
{
  "message": "unauthenticated"
}
```
Cause: Missing or invalid JWT token

**400 Bad Request**:
```json
{
  "message": "Password must be at least 6 characters long"
}
```
Cause: Validation failure

**404 Not Found**:
```json
{
  "message": "User not found"
}
```
Cause: Resource doesn't exist

**500 Internal Server Error**:
```json
{
  "message": "Failed to generate token"
}
```
Cause: Server-side error

## 🔒 Security Features

### Implemented Security Measures

1. **Password Hashing**: bcrypt with cost factor 14
2. **JWT Authentication**: Stateless token-based auth
3. **HttpOnly Cookies**: Prevents client-side script access
4. **Input Validation**: Email format and password length checks
5. **SQL Injection Prevention**: GORM parameterized queries
6. **Foreign Key Constraints**: Data integrity enforcement
7. **Middleware Protection**: Route-level authorization

### Production Security Recommendations

⚠️ **Before deploying to production**:

1. Move `SecretKey` from code to environment variable:
   ```go
   // In utils/helper.go
   SecretKey := os.Getenv("JWT_SECRET_KEY")
   ```

2. Enable CORS with specific origins (not `*`)

3. Add rate limiting for login/register endpoints

4. Implement refresh token mechanism

5. Use HTTPS in production

6. Add request logging and monitoring

7. Implement account lockout after failed login attempts

8. Add email verification for registration

9. Use secure cookie settings:
   ```go
   cookie := fiber.Cookie{
       Secure: true,    // Only send over HTTPS
       SameSite: "Strict",
   }
   ```

## 🧪 Testing the API

### Using cURL

**Register**:
```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"first_name":"John","last_name":"Doe","email":"john@example.com","phone":"1234567890","password":"SecurePass123"}'
```

**Login**:
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -c cookies.txt \
  -d '{"email":"john@example.com","password":"SecurePass123"}'
```

**Create Post** (using saved cookies):
```bash
curl -X POST http://localhost:8080/api/post \
  -H "Content-Type: application/json" \
  -b cookies.txt \
  -d '{"title":"My Post","desc":"Content here","image":"https://example.com/img.jpg"}'
```

**Get All Posts**:
```bash
curl -X GET "http://localhost:8080/api/allposts?page=1" \
  -b cookies.txt
```

### Using Postman or Thunder Client

1. Send POST to `/api/login` with credentials
2. Postman automatically saves the JWT cookie
3. Subsequent requests will include the cookie automatically
4. Make sure "Automatically follow redirects" is enabled in Settings

## 🤝 Contributing

Contributions are welcome! Please follow these guidelines:

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Commit your changes: `git commit -m 'Add amazing feature'`
4. Push to the branch: `git push origin feature/amazing-feature`
5. Open a Pull Request

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 👤 Author

**Mrinmoy**
- GitHub: [@richochetclementine1315](https://github.com/richochetclementine1315)

## 🙏 Acknowledgments

- [Fiber Framework](https://gofiber.io/)
- [GORM](https://gorm.io/)
- Inspired by modern REST API best practices

---

**Need Help?** Open an issue on GitHub or check the [Fiber documentation](https://docs.gofiber.io/).
