# Personal Portfolio Website

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![TailwindCSS](https://img.shields.io/badge/Tailwind_CSS-38B2AC?style=for-the-badge&logo=tailwind-css&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=JSON%20web%20tokens&logoColor=white)

Aplikasi web portfolio personal dengan fitur admin dashboard, CRUD lengkap, dan contact form. Dibuat menggunakan Go (Golang), Chi Router, dan PostgreSQL sebagai project **Golang Intermediate Daytime Class - Lumoshive Academy Bootcamp**.

## Video Demo

[![Watch Demo](https://img.shields.io/badge/YouTube-FF0000?style=for-the-badge&logo=youtube&logoColor=white)](https://youtu.be/a7OtsdOGbyE)

**[Tonton Video Penjelasan Sistem](https://youtu.be/a7OtsdOGbyE)**

---

## Fitur Utama

- **User Authentication** - Login system dengan JWT Cookie & Bcrypt password hashing
- **Admin Dashboard** - Panel admin dengan protected routes
- **CRUD Profile** - Manajemen data profil personal
- **CRUD Experiences** - Tambah, edit, hapus pengalaman kerja
- **CRUD Skills** - Manajemen skill dengan kategori dan level
- **CRUD Projects** - Portfolio proyek dengan upload gambar
- **CRUD Publications** - Manajemen publikasi/artikel
- **Contact Form** - Form kontak dengan integrasi email (Gomail)
- **File Upload** - Upload gambar untuk profile dan project
- **Logging System** - Zap Logger dengan log rotation
- **Unit Testing** - Testing dengan mock pattern

---

## Konsep Programming yang Diimplementasikan

### 1. Clean Architecture Pattern

```
Handler (Controller) -> Service (Business Logic) -> Repository (Data Access)
```

- **Handler Layer**: HTTP request/response handling
- **Service Layer**: Business logic & validasi
- **Repository Layer**: Data access dengan PostgreSQL
- **Dependency Injection**: Interface-based design

### 2. Middleware Pattern

- **Auth Middleware**: JWT token validation & session management
- **Logging Middleware**: Request/response logging dengan Zap
- **Recovery Middleware**: Panic recovery

### 3. Database Integration

- PostgreSQL dengan driver `pgx/v5`
- Database migrations dengan SQL file
- Foreign key relationships

### 4. Testing

- Unit testing dengan `testify/mock`
- Mock repository pattern
- Table-driven tests

### 5. Security

- JWT Authentication
- Bcrypt password hashing
- Protected routes
- Input validation

---

## Struktur Project

```
project-app-portfolio-golang-alvin/
├── cmd/
│   └── hashgen/          # CLI tool untuk generate password hash
├── database/
│   ├── database.go       # Database connection
│   ├── migrations.sql    # Database schema
│   └── mock_db.go        # Mock untuk testing
├── dto/                  # Data Transfer Objects
├── handler/              # HTTP Handlers
│   ├── admin.go          # Admin dashboard handlers
│   ├── auth.go           # Authentication handlers
│   ├── contact.go        # Contact form handler
│   ├── experience.go     # Experience CRUD handlers
│   ├── portfolio.go      # Portfolio page handler
│   ├── profile.go        # Profile CRUD handlers
│   ├── project.go        # Project CRUD handlers
│   ├── publication.go    # Publication CRUD handlers
│   └── skill.go          # Skill CRUD handlers
├── middleware/           # Middleware (Auth, Logging)
├── model/                # Domain models
├── repository/           # Data access layer
├── router/               # Route definitions
├── service/              # Business logic layer
├── utils/                # Helper functions
├── views/                # HTML templates
│   ├── layouts/          # Base layouts
│   └── pages/            # Page templates
├── public/               # Static assets
│   └── assets/
│       └── uploads/      # Uploaded files
├── logs/                 # Application logs
├── main.go               # Entry point
└── go.mod                # Go modules
```

---

## Cara Menggunakan

### Prerequisites

- Go 1.21+
- PostgreSQL 13+
- Git

### Instalasi

1. **Clone repository**

   ```bash
   git clone https://github.com/Alvinnn-R/project-app-portfolio-golang-alvin.git
   cd project-app-portfolio-golang-alvin
   ```

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Setup database**

   ```bash
   # Buat database
   createdb portfolio_db

   # Import schema
   psql -U postgres -d portfolio_db -f database/migrations.sql
   ```

4. **Konfigurasi environment**

   Buat file `.env` atau edit konfigurasi di `database/database.go`:

   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=yourpassword
   DB_NAME=portfolio_db
   JWT_SECRET=your-secret-key
   ```

5. **Generate password hash** (untuk user admin)

   ```bash
   go run cmd/hashgen/main.go
   ```

6. **Jalankan aplikasi**

   ```bash
   go run .
   ```

7. **Akses aplikasi**
   - Portfolio: `http://localhost:8080`
   - Admin Login: `http://localhost:8080/login`
   - Admin Dashboard: `http://localhost:8080/admin/dashboard`

---

## API Endpoints

### Public Endpoints

| Method | Endpoint            | Description               |
| ------ | ------------------- | ------------------------- |
| GET    | `/`                 | Portfolio page            |
| GET    | `/api/v1/portfolio` | Get portfolio data (JSON) |
| POST   | `/api/v1/contact`   | Submit contact form       |

### Auth Endpoints

| Method | Endpoint  | Description   |
| ------ | --------- | ------------- |
| GET    | `/login`  | Login page    |
| POST   | `/login`  | Process login |
| GET    | `/logout` | Logout        |

### Admin Endpoints (Protected)

| Method   | Endpoint              | Description            |
| -------- | --------------------- | ---------------------- |
| GET      | `/admin/dashboard`    | Admin dashboard        |
| GET/POST | `/admin/profile`      | Profile management     |
| GET/POST | `/admin/experiences`  | Experience management  |
| GET/POST | `/admin/skills`       | Skill management       |
| GET/POST | `/admin/projects`     | Project management     |
| GET/POST | `/admin/publications` | Publication management |

### API v1 Endpoints

| Resource     | Endpoints                                                                      |
| ------------ | ------------------------------------------------------------------------------ |
| Profile      | GET, POST `/api/v1/profile`, PUT `/api/v1/profile/{id}`                        |
| Experiences  | GET, POST `/api/v1/experiences`, GET, PUT, DELETE `/api/v1/experiences/{id}`   |
| Skills       | GET, POST `/api/v1/skills`, GET, PUT, DELETE `/api/v1/skills/{id}`             |
| Projects     | GET, POST `/api/v1/projects`, GET, PUT, DELETE `/api/v1/projects/{id}`         |
| Publications | GET, POST `/api/v1/publications`, GET, PUT, DELETE `/api/v1/publications/{id}` |

---

## Database Schema

### Tables

- **users** - Admin user accounts
- **profiles** - Personal profile information
- **experiences** - Work experience entries
- **skills** - Skills with category and level
- **projects** - Portfolio projects
- **publications** - Articles/publications

### ERD

Lihat file `database/migrations.sql` untuk schema lengkap.

---

## Testing

### Run Unit Tests

```bash
# Test semua package
go test ./... -v

# Test dengan coverage
go test ./... -cover

# Test specific package
go test ./repository/... -v
go test ./service/... -v
```

### Test Files

- `repository/experience_test.go`
- `repository/profile_test.go`
- `repository/project_test.go`
- `repository/publication_test.go`
- `repository/skill_test.go`
- `service/portfolio_test.go`

---

## Learning Outcomes

Project ini mengajarkan:

- Clean Architecture dengan Go
- RESTful API design
- JWT Authentication & Authorization
- Database integration (PostgreSQL)
- Repository pattern & dependency injection
- Middleware pattern (Auth, Logging, Recovery)
- Unit testing dengan mock
- Server-side rendering dengan HTML templates
- File upload handling
- Input validation & error handling
- Logging best practices

---

## Author

**Alvin Rama S**

- GitHub: [@Alvinnn-R](https://github.com/Alvinnn-R)
- Bootcamp: Golang Intermediate Daytime Class - Lumoshive Academy

---

## License

This project is for educational purposes as part of Lumoshive Academy Bootcamp.

---

## Acknowledgments

- Lumoshive Academy - Golang Bootcamp
- Instructor & Mentors
- Fellow bootcamp participants
