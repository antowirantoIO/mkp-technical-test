# MKP Boarding Test - Maritime Management System

## 🚀 Features

- **Clean Architecture**: Separation of concerns with clear dependency boundaries
- **PostgreSQL Database**: Robust relational database with GORM ORM
- **JWT Authentication**: Secure token-based authentication with refresh tokens
- **Role-Based Access Control**: Flexible permission system with user roles
- **Swagger Documentation**: Interactive API documentation with complete endpoint coverage
- **Docker Support**: Containerized deployment with docker-compose
- **Database Migrations**: Version-controlled schema changes with migration tools
- **Input Validation**: Comprehensive request validation using go-playground/validator
- **Structured Logging**: Logrus-based logging with proper error handling
- **Maritime Domain**: Specialized for shipping industry operations

## 📋 Domain Models

- **Users**: User management with authentication, email verification, and profile management
- **Roles & Permissions**: Complete RBAC system with assignable permissions
- **Operators**: Maritime operator/company management with detailed business information
- **Ships**: Comprehensive ship registry with technical specifications and tracking
- **Harbors**: Harbor information system with facilities and operational details

## 🛠️ Prerequisites

- Go 1.21 or higher
- PostgreSQL 15+
- Docker & Docker Compose (optional)
- Make (optional, for convenience commands)

## 🚀 Quick Start

### Using Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone git@github.com:antowirantoIO/mkp-technical-test.git
   cd mkp-boarding-test
   ```

2. **Start all services**
   ```bash
   docker-compose up -d
   ```

3. **Run database migrations**
   ```bash
   make db-migrate-up
   ```

4. **Access the application**
   - API: http://localhost:3000
   - Swagger UI: http://localhost:3000/swagger/

### Manual Setup

1. **Install dependencies**
   ```bash
   go mod download
   ```

2. **Setup PostgreSQL database**
   ```bash
   createdb mkp_technicaltest
   ```

3. **Configure environment**
   Update `config.json` with your database credentials:
   ```json
   {
     "database": {
       "host": "localhost",
       "port": 5432,
       "username": "postgres",
       "password": "postgres",
       "name": "mkp_technicaltest"
     }
   }
   ```

4. **Run database migrations**
   ```bash
   make db-migrate-up
   ```

5. **Start the application**
   ```bash
   make dev
   ```

## 📝 API Usage Examples

### Authentication Flow

1. **Register a new user**
   ```bash
   curl -X POST http://localhost:3000/api/users/register \
     -H "Content-Type: application/json" \
     -d '{
       "username": "john_doe",
       "email": "john@example.com",
       "password": "SecurePass123!",
       "full_name": "John Doe"
     }'
   ```

2. **Login to get JWT token**
   ```bash
   curl -X POST http://localhost:3000/api/users/login \
     -H "Content-Type: application/json" \
     -d '{
       "username": "john_doe",
       "password": "SecurePass123!"
     }'
   ```

3. **Use the token for authenticated requests**
   ```bash
   curl -X GET http://localhost:3000/api/users/_current \
     -H "Authorization: Bearer YOUR_JWT_TOKEN"
   ```

### Maritime Operations Examples

1. **Create a maritime operator**
   ```bash
   curl -X POST http://localhost:3000/api/operators \
     -H "Authorization: Bearer YOUR_JWT_TOKEN" \
     -H "Content-Type: application/json" \
     -d '{
       "company_name": "Ocean Shipping Ltd",
       "operator_type": "shipping_line",
       "registration_number": "OSL-2024-001",
       "country": "Singapore",
       "city": "Singapore",
       "email": "contact@oceanshipping.com",
       "website": "https://www.oceanshipping.com"
     }'
   ```

2. **Register a new ship**
   ```bash
   curl -X POST http://localhost:3000/api/ships \
     -H "Authorization: Bearer YOUR_JWT_TOKEN" \
     -H "Content-Type: application/json" \
     -d '{
       "operator_id": "operator-uuid-here",
       "ship_name": "MV Ocean Pioneer",
       "imo_number": "1234567",
       "ship_type": "container",
       "flag_state": "Singapore",
       "port_of_registry": "Singapore",
       "length": 300.5,
       "beam": 48.2,
       "gross_tonnage": 150000
     }'
   ```

3. **Create a harbor record**
   ```bash
   curl -X POST http://localhost:3000/api/harbors \
     -H "Authorization: Bearer YOUR_JWT_TOKEN" \
     -H "Content-Type: application/json" \
     -d '{
       "harbor_code": "SGSIN",
       "harbor_name": "Port of Singapore",
       "un_locode": "SGSIN",
       "country": "Singapore",
       "province": "Singapore",
       "city": "Singapore",
       "latitude": 1.2966,
       "longitude": 103.8006,
       "has_container": true,
       "has_customs": true,
       "has_pilotage": true
     }'
   ```

## 📚 API Documentation

Once the application is running, visit:
- **Swagger UI**: http://localhost:3000/swagger/
- **API Base URL**: http://localhost:3000/api

### Authentication

The API uses JWT Bearer tokens for authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

### Available Endpoints

#### Authentication (Public Endpoints)
- `POST /api/users/register` - User registration with email verification
- `POST /api/users/login` - User login with JWT token generation

#### User Management (Protected)
- `GET /api/users/_current` - Get current authenticated user profile
- `PATCH /api/users/_current` - Update current user profile
- `DELETE /api/users` - User logout (token invalidation)

#### Role Management (Protected)
- `GET /api/roles` - List all roles with pagination
- `POST /api/roles` - Create new role
- `GET /api/roles/{roleId}` - Get role details by ID
- `PUT /api/roles/{roleId}` - Update role information
- `DELETE /api/roles/{roleId}` - Delete role
- `POST /api/roles/{roleId}/permissions` - Assign permissions to role
- `DELETE /api/roles/{roleId}/permissions` - Remove permissions from role

#### Permission Management (Protected)
- `GET /api/permissions` - List all permissions with filtering
- `POST /api/permissions` - Create new permission
- `GET /api/permissions/{permissionId}` - Get permission details
- `PUT /api/permissions/{permissionId}` - Update permission
- `DELETE /api/permissions/{permissionId}` - Delete permission

#### Operator Management (Protected)
- `GET /api/operators` - List operators with advanced filtering (company name, type, status, location)
- `POST /api/operators` - Create new maritime operator
- `GET /api/operators/{operatorId}` - Get operator details
- `PUT /api/operators/{operatorId}` - Update operator information
- `DELETE /api/operators/{operatorId}` - Delete operator

#### Ship Management (Protected)
- `GET /api/ships` - List ships with filtering (operator, name, flag state, type, status)
- `POST /api/ships` - Register new ship with technical specifications
- `GET /api/ships/{shipId}` - Get detailed ship information
- `PUT /api/ships/{shipId}` - Update ship details and specifications
- `DELETE /api/ships/{shipId}` - Remove ship from registry

#### Harbor Management (Protected)
- `GET /api/harbors` - List harbors with location and facility filtering
- `POST /api/harbors` - Create new harbor with comprehensive facility details
- `GET /api/harbors/{harborId}` - Get harbor information and facilities
- `PUT /api/harbors/{harborId}` - Update harbor details and capabilities
- `DELETE /api/harbors/{harborId}` - Delete harbor record

## 🧪 Testing

### Run Tests
```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage
```

### Test Structure
- Unit tests: Test individual components in isolation
- Integration tests: Test complete workflows with database

## 🔧 Development

### Available Make Commands

```bash
# Development
make dev              # Run in development mode
make build            # Build the application
make test             # Run tests
make test-coverage    # Run tests with coverage

# Dependencies
make deps             # Download dependencies
make tidy             # Tidy dependencies

# Database
make db-migrate-up    # Run migrations
make db-migrate-down  # Rollback migrations
make db-create        # Create database
make db-drop          # Drop database

# Docker
make docker-up        # Start all services
make docker-down      # Stop all services
make docker-rebuild   # Rebuild and restart

# Swagger
make swagger-gen      # Generate Swagger docs

# Code Quality
make fmt              # Format code
make vet              # Run go vet
make lint             # Run linter
```

### Project Structure Guidelines

1. **Entity Layer**: Core business entities with no external dependencies
2. **Repository Layer**: Data access interfaces and implementations
3. **Use Case Layer**: Business logic and orchestration
4. **Delivery Layer**: HTTP handlers, middleware, and routing
5. **Gateway Layer**: External service integrations

## 🐳 Docker Services

- **app**: Main Go application (port 3000)
- **postgres**: PostgreSQL database (port 5432)
- **adminer**: Database management UI (port 8080)

## 📊 Monitoring & Observability

- **Logging**: Structured logging with Logrus
- **Request Validation**: Comprehensive input validation with detailed error messages
- **Error Handling**: Consistent error responses with proper HTTP status codes
- **Swagger Documentation**: Complete API documentation with request/response examples

## 🔒 Security

- **JWT Authentication**: Secure token-based authentication with configurable expiration
- **Role-Based Access Control**: Granular permission system for maritime operations
- **Input Validation**: Comprehensive request validation with detailed error messages
- **SQL Injection Prevention**: GORM ORM with parameterized queries
- **Password Security**: Bcrypt hashing with salt
- **Email Verification**: User account verification system
- **Authorization Middleware**: Protected endpoints with proper access control

## 🌊 Maritime Domain Features

### Ship Management

### Harbor Operations

### Operator Management

## 🚀 Deployment

### Production Build
```bash
make prod-build
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Run the test suite
6. Submit a pull request

## 📄 License

This project is licensed under the Apache 2.0 License - see the LICENSE file for details.

## 🛠️ Technology Stack

- **Language**: Go 1.25
- **Web Framework**: Fiber v2 (Express-inspired)
- **Database**: PostgreSQL with GORM ORM
- **Authentication**: JWT with golang-jwt/jwt
- **Validation**: go-playground/validator
- **Documentation**: Swagger/OpenAPI 3.0
- **Logging**: Logrus structured logging
- **Configuration**: Viper configuration management
- **Containerization**: Docker & Docker Compose
- **Database Migrations**: golang-migrate

## 📞 Support

For support and questions:
- Create an issue in the repository
- Review the Swagger documentation at `/docs/`
- Check the logs for detailed error information

---

**Built with ❤️ for the Maritime Industry using Go and Clean Architecture principles**