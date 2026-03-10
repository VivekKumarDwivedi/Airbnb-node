# AuthInGo - Authentication Service

A comprehensive authentication and authorization service built with Go, implementing JWT-based authentication and Role-Based Access Control (RBAC).

## 🏗️ Project Architecture

### Technology Stack
- **Language**: Go 1.25.4
- **Web Framework**: Chi v5.2.3
- **Database**: MySQL
- **Authentication**: JWT (golang-jwt/jwt/v5)
- **Password Hashing**: bcrypt (golang.org/x/crypto)
- **Migration Tool**: Goose
- **Validation**: go-playground/validator/v10

### Project Structure
```
AuthInGo/
├── app/                    # Application initialization and configuration
├── config/                 # Environment and database configuration
├── controllers/            # HTTP request handlers
├── db/                     # Database layer
│   ├── migrations/         # Database schema migrations
│   └── repositories/       # Data access layer
├── dto/                    # Data Transfer Objects
├── middlewares/            # HTTP middlewares (auth, validation, rate limiting)
├── models/                 # Domain models (User, Role, Permission)
├── router/                 # Route definitions
├── services/               # Business logic layer
└── utils/                  # Utility functions
```

## 🔄 Project Workflow

### 1. Application Startup Flow
1. **main.go** initializes the application
2. Load environment configuration from `.env`
3. Create application configuration with server settings
4. Initialize application instance
5. **Application.Run()**:
   - Setup database connection
   - Initialize repositories (User, Role, RolePermission, UserRole)
   - Initialize services (UserService, RoleService)
   - Initialize controllers (UserController, RoleController)
   - Setup routers with middleware chain
   - Start HTTP server on configured port

### 2. Request Processing Flow
```
HTTP Request → Middleware Chain → Controller → Service → Repository → Database
```

#### Middleware Chain (in order):
1. **Logger**: Request logging
2. **RateLimitMiddleware**: Rate limiting protection
3. **JWTAuthMiddleware**: JWT token validation (for protected routes)
4. **Role-based middleware**: Role authorization checks
5. **Request validators**: Input validation and body parsing

### 3. Authentication Flow

#### User Registration
1. Client sends `POST /signup` with user data
2. **CreateUserRequestValidator** validates request body
3. **UserController.CreateUser** processes request
4. **UserService.CreateUser**:
   - Hashes password using bcrypt
   - Calls repository to create user
5. Repository saves user to database
6. Returns created user data (without password)

#### User Login
1. Client sends `POST /login` with email and password
2. **UserLoginRequestValidator** validates request
3. **UserController.LoginUser** processes request
4. **UserService.LoginUser**:
   - Retrieves user by email from repository
   - Validates password using bcrypt comparison
   - Generates JWT token with user claims (email, id)
5. Returns JWT token to client

#### Protected Route Access
1. Client includes JWT token in Authorization header
2. **JWTAuthMiddleware** validates token
3. Extracts user claims and adds to request context
4. **Role-based middleware** checks user permissions
5. Request proceeds to controller with authenticated user context

### 4. Authorization Flow (RBAC)

#### Role-Based Access Control Implementation
- **Roles**: admin, user, moderator (pre-defined)
- **Permissions**: Granular permissions for resources and actions
- **Role-Permission Mapping**: Many-to-many relationship
- **User-Role Mapping**: Many-to-many relationship

#### Authorization Check Process
1. User authenticates and receives JWT
2. JWT contains user ID and email
3. Middleware extracts user roles from database
4. Role-based middleware checks if user has required role
5. Access granted/denied based on role permissions

## 🗄️ Database Workflow

### Database Schema

#### Core Tables

1. **users**
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```

2. **roles**
```sql
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```

3. **permissions**
```sql
CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    resource VARCHAR(255) NOT NULL,
    action VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```

4. **role_permissions** (Many-to-Many)
```sql
CREATE TABLE role_permissions (
    id SERIAL PRIMARY KEY,
    role_id INT NOT NULL,
    permission_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES roles(id),
    FOREIGN KEY (permission_id) REFERENCES permissions(id)
);
```

5. **user_roles** (Many-to-Many)
```sql
CREATE TABLE user_roles (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (role_id) REFERENCES roles(id)
);
```

### Migration Workflow

#### Migration Management
- **Tool**: Goose migration tool
- **Location**: `db/migrations/`
- **Naming**: `YYYYMMDDHHMMSS_description.sql`

#### Migration Commands (via Makefile)
```bash
# Create new migration
make migrate-create name="create_new_table"

# Run all pending migrations
make migrate-up

# Rollback last migration
make migrate-down

# Reset database (rollback all migrations)
make migrate-reset

# Check migration status
make migrate-status

# Redo last migration
make migrate-redo
```

#### Migration Files
1. `20251201175156_create_user_table.sql` - Users table
2. `20251217145802_create_role_table.sql` - Roles table with seed data
3. `20251217145834_create_permissions_table.sql` - Permissions table
4. `20251217145904_create_role_permissions_table.sql` - Role-Permission mapping
5. `20251217145929_create_user_roles_table.sql` - User-Role mapping

### Database Connection Workflow
1. Load database configuration from environment variables
2. Establish MySQL connection using `go-sql-driver/mysql`
3. Connection pooling and configuration
4. Repository pattern for data access abstraction
5. Prepared statements for SQL injection prevention

## 🚀 API Endpoints

### User Endpoints
- `POST /signup` - User registration
- `POST /login` - User authentication
- `GET /profile` - Get user profile (protected)
- `DELETE /delete` - Delete user account
- `GET /users` - Get all users

### Role Endpoints
- Role management endpoints (implemented in role controller)

### Utility Endpoints
- `GET /ping` - Health check endpoint

## 🔧 Configuration

### Environment Variables
- `PORT` - Server port (default: :8080)
- `JWT_SECRET` - JWT signing secret
- Database connection string (configured in Makefile)

### Setup Instructions
1. Clone the repository
2. Copy `.env.example` to `.env` and configure
3. Run database migrations: `make migrate-up`
4. Start the server: `go run main.go`

## 🛡️ Security Features

- **Password Hashing**: bcrypt for secure password storage
- **JWT Authentication**: Stateless authentication with configurable expiration
- **Rate Limiting**: Request rate limiting middleware
- **Input Validation**: Comprehensive request validation
- **CORS**: Cross-Origin Resource Sharing support
- **SQL Injection Prevention**: Parameterized queries and prepared statements

## 📝 Development Notes

- Follows clean architecture principles with clear separation of concerns
- Repository pattern for data access abstraction
- Service layer for business logic
- Middleware for cross-cutting concerns
- Comprehensive error handling and logging
- RESTful API design principles