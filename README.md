# Golang Clean Architecture API

A production-ready REST API built with Go using Clean Architecture principles, featuring PostgreSQL database, Swagger documentation, and Docker containerization.

## 🚀 Features

- **Clean Architecture**: Separation of concerns with clear dependency boundaries
- **PostgreSQL Database**: Robust relational database with GORM ORM
- **JWT Authentication**: Secure token-based authentication
- **Role-Based Access Control**: Flexible permission system
- **Swagger Documentation**: Interactive API documentation
- **Docker Support**: Containerized deployment with docker-compose
- **Comprehensive Testing**: Unit and integration tests
- **Kafka Integration**: Event-driven messaging
- **Redis Caching**: Optional caching layer
- **Database Migrations**: Version-controlled schema changes

## 📋 Domain Models

- **Users**: User management with authentication
- **Roles & Permissions**: RBAC system
- **Contacts**: Contact information management
- **Addresses**: Address management linked to contacts
- **Operators**: Operator management
- **Ships**: Ship registry and management
- **Harbors**: Harbor information system

## 🏗️ Architecture

```
cmd/
├── web/           # HTTP server entry point
└── worker/        # Background worker entry point

internal/
├── config/        # Configuration and dependency injection
├── delivery/      # Delivery layer (HTTP handlers, middleware)
├── entity/        # Domain entities
├── gateway/       # External service gateways
├── model/         # DTOs and request/response models
├── repository/    # Data access layer
└── usecase/       # Business logic layer

db/
└── migrations/    # Database migration files

docs/              # Swagger documentation
test/              # Test files
```

## 🛠️ Prerequisites

- Go 1.21 or higher
- PostgreSQL 15+
- Docker & Docker Compose (optional)
- Make (optional, for convenience commands)

## 🚀 Quick Start

### Using Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd mkp-boarding-test
   ```

2. **Start all services**
   ```bash
   docker-compose up -d
   ```

3. **Access the application**
   - API: http://localhost:3000
   - Swagger UI: http://localhost:3000/swagger/
   - Database Admin (Adminer): http://localhost:8080

### Manual Setup

1. **Install dependencies**
   ```bash
   go mod download
   ```

2. **Setup PostgreSQL database**
   ```bash
   createdb golang_clean_architecture
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
       "name": "golang_clean_architecture"
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

#### Authentication
- `POST /api/users/login` - User login
- `POST /api/users/register` - User registration
- `GET /api/users/current` - Get current user

#### User Management
- `GET /api/users` - List users
- `GET /api/users/{id}` - Get user by ID
- `PUT /api/users/{id}` - Update user
- `DELETE /api/users/{id}` - Delete user

#### Role & Permission Management
- `GET /api/roles` - List roles
- `POST /api/roles` - Create role
- `PUT /api/roles/{id}` - Update role
- `DELETE /api/roles/{id}` - Delete role
- `POST /api/roles/{id}/permissions` - Assign permissions
- `DELETE /api/roles/{id}/permissions` - Remove permissions


#### Operator Management
- `GET /api/operators` - List operators
- `POST /api/operators` - Create operator
- `GET /api/operators/{id}` - Get operator
- `PUT /api/operators/{id}` - Update operator
- `DELETE /api/operators/{id}` - Delete operator

#### Ship Management
- `GET /api/ships` - List ships
- `POST /api/ships` - Create ship
- `GET /api/ships/{id}` - Get ship
- `PUT /api/ships/{id}` - Update ship
- `DELETE /api/ships/{id}` - Delete ship

#### Harbor Management
- `GET /api/harbors` - List harbors
- `POST /api/harbors` - Create harbor
- `GET /api/harbors/{id}` - Get harbor
- `PUT /api/harbors/{id}` - Update harbor
- `DELETE /api/harbors/{id}` - Delete harbor

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

### Adding New Features

1. Define entity in `internal/entity/`
2. Create repository interface and implementation
3. Implement use case with business logic
4. Create DTOs in `internal/model/`
5. Add HTTP controller in `internal/delivery/http/`
6. Register routes in `internal/delivery/http/route/`
7. Add to dependency injection in `internal/config/app.go`
8. Write tests

## 🐳 Docker Services

- **app**: Main Go application (port 3000)
- **postgres**: PostgreSQL database (port 5432)
- **redis**: Redis cache (port 6379)
- **adminer**: Database management UI (port 8080)

## 📊 Monitoring & Observability

- **Logging**: Structured logging with Logrus
- **Health Checks**: Built-in health check endpoints
- **Metrics**: Application metrics collection
- **Tracing**: Distributed tracing support

## 🔒 Security

- JWT-based authentication
- Role-based access control (RBAC)
- Input validation and sanitization
- SQL injection prevention with GORM
- CORS configuration
- Rate limiting middleware

## 🚀 Deployment

### Production Build
```bash
make prod-build
```

### Docker Production
```bash
docker-compose -f docker-compose.prod.yml up -d
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

## 📞 Support

For support and questions:
- Create an issue in the repository
- Contact: support@example.com

---

**Built with ❤️ using Go and Clean Architecture principles**