# Portfolio Website - Golang

A modern portfolio web application built with Golang, PostgreSQL, and TailwindCSS featuring a neobrutalist design.

## Features

- **RESTful API**: Full CRUD operations for portfolio data
- **Clean Architecture**: Separated layers (Handler → Service → Repository)
- **PostgreSQL Database**: Robust data storage with proper indexing
- **HTML Template Rendering**: Server-side rendering using Go's `html/template`
- **Neobrutalist Design**: Modern, bold UI with TailwindCSS
- **Zap Logging**: Structured logging with file rotation
- **Input Validation**: Comprehensive validation for all inputs
- **Unit Tests**: 50%+ test coverage for service and DTO layers

## Project Structure

```
project-app-portfolio-golang-alvin/
├── database/
│   ├── database.go          # Database connection and configuration
│   └── migrations.sql        # SQL schema and sample data
├── dto/
│   ├── portfolio.go          # Request DTOs with validation
│   └── portfolio_test.go     # DTO validation tests
├── handler/
│   ├── handler.go            # Handler initialization
│   └── portfolio.go          # HTTP handlers for all endpoints
├── logs/                     # Application logs directory
├── middleware/
│   ├── logging.go            # Request logging middleware
│   └── middleware.go         # Middleware initialization
├── model/
│   ├── model.go              # Model initialization
│   └── portfolio.go          # Data models (Profile, Experience, Skill, etc.)
├── public/
│   └── assets/               # Static files (CSS, images)
├── repository/
│   ├── repository.go         # Repository initialization
│   └── portfolio.go          # Database operations
├── router/
│   └── router.go             # Route definitions
├── service/
│   ├── service.go            # Service initialization
│   ├── portfolio.go          # Business logic
│   └── portfolio_test.go     # Service unit tests
├── utils/
│   ├── logger.go             # Zap logger configuration
│   ├── response.go           # HTTP response helpers
│   └── utils.go              # Utility functions
├── views/
│   └── index.html            # Go HTML template
├── go.mod                    # Go module definition
├── main.go                   # Application entry point
└── README.md                 # This file
```

## Requirements

- Go 1.21 or higher
- PostgreSQL 12 or higher

## Installation

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd project-app-portfolio-golang-alvin
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Create the database**

   ```bash
   # Connect to PostgreSQL and create the database
   psql -U postgres
   CREATE DATABASE portfolio;
   \q
   ```

4. **Run database migrations**

   ```bash
   psql -U postgres -d portfolio -f database/migrations.sql
   ```

5. **Configure environment variables** (optional)

   ```bash
   # Set environment variables for database connection
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=postgres
   export DB_PASSWORD=your_password
   export DB_NAME=portfolio
   export DB_SSLMODE=disable
   ```

6. **Run the application**

   ```bash
   go run main.go
   ```

7. **Access the application**
   - Website: http://localhost:8000
   - API: http://localhost:8000/api/v1

## API Endpoints

### Portfolio

- `GET /api/v1/portfolio` - Get full portfolio data

### Profile

- `GET /api/v1/profile` - Get profile
- `POST /api/v1/profile` - Create profile
- `PUT /api/v1/profile/{id}` - Update profile

### Experiences

- `GET /api/v1/experiences` - Get all experiences
- `GET /api/v1/experiences/{id}` - Get experience by ID
- `POST /api/v1/experiences` - Create experience
- `PUT /api/v1/experiences/{id}` - Update experience
- `DELETE /api/v1/experiences/{id}` - Delete experience

### Skills

- `GET /api/v1/skills` - Get all skills
- `GET /api/v1/skills?category={category}` - Get skills by category
- `GET /api/v1/skills/{id}` - Get skill by ID
- `POST /api/v1/skills` - Create skill
- `PUT /api/v1/skills/{id}` - Update skill
- `DELETE /api/v1/skills/{id}` - Delete skill

### Projects

- `GET /api/v1/projects` - Get all projects
- `GET /api/v1/projects/{id}` - Get project by ID
- `POST /api/v1/projects` - Create project
- `PUT /api/v1/projects/{id}` - Update project
- `DELETE /api/v1/projects/{id}` - Delete project

### Publications

- `GET /api/v1/publications` - Get all publications
- `GET /api/v1/publications/{id}` - Get publication by ID
- `POST /api/v1/publications` - Create publication
- `PUT /api/v1/publications/{id}` - Update publication
- `DELETE /api/v1/publications/{id}` - Delete publication

### Contact

- `POST /api/v1/contact` - Submit contact form

## Request Examples

### Create Experience

```json
POST /api/v1/experiences
{
  "title": "Software Engineer",
  "organization": "Tech Company",
  "period": "2022 - Present",
  "description": "Building scalable applications",
  "type": "work",
  "color": "cyan"
}
```

### Create Skill

```json
POST /api/v1/skills
{
  "category": "Programming Languages",
  "name": "Go",
  "level": "advanced",
  "color": "black"
}
```

### Create Project

```json
POST /api/v1/projects
{
  "title": "My Project",
  "description": "Project description",
  "image_url": "/public/assets/project.jpg",
  "project_url": "https://example.com",
  "github_url": "https://github.com/username/repo",
  "tech_stack": "Go, PostgreSQL, Docker",
  "color": "cyan"
}
```

## Running Tests

```bash
# Run all tests with coverage
go test ./... -cover

# Run tests with verbose output
go test ./dto/... ./service/... -v -cover

# Generate coverage report
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Environment Variables

| Variable    | Description       | Default   |
| ----------- | ----------------- | --------- |
| DB_HOST     | Database host     | localhost |
| DB_PORT     | Database port     | 5432      |
| DB_USER     | Database user     | postgres  |
| DB_PASSWORD | Database password | password  |
| DB_NAME     | Database name     | portfolio |
| DB_SSLMODE  | SSL mode          | disable   |

## Technologies Used

- **Go 1.25** - Programming language
- **Chi v5** - HTTP router
- **pgx v5** - PostgreSQL driver
- **Zap** - Structured logger
- **Lumberjack** - Log rotation
- **TailwindCSS** - CSS framework
- **html/template** - Go templating

## License

MIT License

## Author

Alvin Maulana - Bootcamp Golang PMM
