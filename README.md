# 🔐 Auth Service

A robust authentication service written in **Go (v1.24)** that leverages **JWT tokens** for secure session management. The service provides features like **user registration, login, instant logout**, and **role-based access control (RBAC)**.

## 🚀 Features

- ✅ User registration & login
- ✅ JWT-based authentication
- ✅ Refresh & access token support
- ✅ Redis-backed instant logout mechanism (logoutPin)
- ✅ Role-based access control (RBAC) with `user` and `admin` roles
- ✅ Graceful lifecycle management with Uber Fx
- ✅ PostgreSQL persistence
- ✅ Docker-compatible

## 🧱 Tech Stack

| Component     | Description                         |
|---------------|-------------------------------------|
| Go            | v1.24                               |
| Gin           | HTTP routing framework              |
| Uber Fx       | Dependency injection & lifecycle    |
| PostgreSQL    | User and token persistence          |
| Redis         | Token blacklist and cache layer     |
| JWT (go-jwt)  | Access and refresh token handling   |

## ⚙️ Environment Variables

| Key            | Description                  | Example                          |
|----------------|------------------------------|----------------------------------|
| `PORT`         | Server port                  | `8080`                           |
| `DATABASE_URL` | PostgreSQL DSN               | `postgres://user:pass@...`       |
| `REDIS_URL`    | Redis connection string      | `redis://localhost:6379/0`       |
| `JWT_SECRET`   | Secret for signing JWT tokens| `your_secret_key`                |
| `TOKEN_TTL`    | Token lifetime in minutes    | `15`                             |

## 📦 Running the Project

### Prerequisites

- Go 1.24+
- Docker & Docker Compose

### Registration
curl -X POST http://localhost:8080/api/auth/register \
-H "Content-Type: application/json" \
-d '{
"email": "sdf@example.com",
"password": "sdf"
}'

### Login
curl -X POST http://localhost:8080/api/auth/login \
-H "Content-Type: application/json" \
-d '{
"email": "sdf@example.com",
"password": "sdf"
}'


## PLANS FOR FURTHER DEVELOPMENT
-- Add user verification via email
