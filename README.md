[![GitHub Actions Workflow Status](https://github.com/Aadityaa2606/Bank-API/actions/workflows/test.yml/badge.svg)](https://github.com/Aadityaa2606/Bank-API/actions)
[![wakatime](https://wakatime.com/badge/user/16514914-4626-4732-8ab1-9ea08b62263f/project/fce34056-eb02-4c7d-ab38-a37a0219985b.svg)](https://wakatime.com/badge/user/16514914-4626-4732-8ab1-9ea08b62263f/project/fce34056-eb02-4c7d-ab38-a37a0219985b)

[![wakatime](https://wakatime.com/badge/user/16514914-4626-4732-8ab1-9ea08b62263f/project/fce34056-eb02-4c7d-ab38-a37a0219985b.svg)](https://wakatime.com/badge/user/16514914-4626-4732-8ab1-9ea08b62263f/project/fce34056-eb02-4c7d-ab38-a37a0219985b)

# Simple Bank API

A robust banking API built with Go, featuring both HTTP/REST and gRPC endpoints. This project demonstrates modern backend development practices including secure user authentication, database migrations, and containerization.

📚 **API Documentation**: [https://bankapi.aadinagarajan.com](https://bankapi.aadinagarajan.com)

## 🌟 Features

- **Multi-protocol Support**
  - REST API using Gin framework
  - gRPC services with protocol buffers
- **User Management**
  - Account creation and authentication
  - JWT-based access and refresh tokens
  - Session management
- **Banking Operations**
  - Account management
  - Money transfers
  - Transaction history
- **Database**
  - PostgreSQL with pgx driver
  - Migrations using Goose
  - SQLC for type-safe database operations

## 🛠️ Technologies

- **Backend**
  - Go 1.24
  - Gin web framework
  - gRPC & Protocol Buffers
  - JWT authentication
- **Database**
  - PostgreSQL 17
  - SQLC
  - Goose migrations
- **DevOps**
  - Docker & Docker Compose
  - Make for automation
  - Air for live reloading

## 🚀 Getting Started

### Prerequisites

- Go 1.24 or later
- Docker and Docker Compose
- Make

### Environment Setup

1. Clone the repository
2. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```
3. Modify the `.env` file with your desired configuration

### Running the Application

Using Make commands:

```bash
# Start all services (Docker, PostgreSQL, and API)
make start

# Or start without database (for remote DB)
make start-remote

# Stop all services
make stop
```

### Development Commands

```bash
# Database operations
make startdb    # Start PostgreSQL container
make createdb   # Create the database
make migrateup  # Run migrations
make psql       # Connect to PostgreSQL

# Code generation
make sqlc       # Generate Go code from SQL
make proto      # Generate Go code from Protocol Buffers

# Testing
make test       # Run tests with coverage
```

## 🏗️ Project Structure

```
bank_api/
├── api/          # HTTP/REST API handlers
├── db/
│   ├── migration/  # Database migrations
│   ├── query/      # SQL queries
│   └── sqlc/       # Generated Go code
├── gapi/         # gRPC service implementations
├── pb/           # Protocol Buffer definitions
├── token/        # JWT token management
└── util/         # Utility functions
```

## 🔒 Authentication Flow

1. User registration (`POST /users`)
2. Login (`POST /users/login`)
   - Returns access & refresh tokens
3. Protected endpoints
   - Require valid access token
4. Token refresh (`POST /users/token/refresh`)
   - Uses refresh token to get new access token
5. Logout (`POST /users/logout`)
   - Invalidates refresh token

## 🌍 API Endpoints

### Public Endpoints
- `POST /users` - Create new user
- `POST /users/login` - User login

### Protected Endpoints
- `GET /accounts` - List accounts
- `POST /accounts` - Create account
- `GET /accounts/:id` - Get account details
- `PATCH /accounts/:id` - Update account
- `DELETE /accounts/:id` - Delete account
- `POST /transfer` - Create money transfer
- `POST /users/logout` - Logout user
- `POST /users/token/refresh` - Refresh access token
- `POST /users/revoke` - Revoke session

## 📝 License

This project is a learning exercise and is available under the MIT License.
