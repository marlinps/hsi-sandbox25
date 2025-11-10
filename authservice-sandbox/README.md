# AuthService Sandbox (authservice-sandbox)

This is a small sandbox project demonstrating a minimal authentication service written in Go using Fiber (HTTP framework) and GORM (MySQL) as the persistence layer. It implements a simple user entity, a repository/service layer, JWT generation via an adapter, and HTTP routes wired in `main.go`.

This README documents how to configure, build, and run the service locally (Windows / PowerShell), and explains the project structure and extension points.

## Project purpose
- Educational sandbox to learn how to wire together a web server, database, repository/service layers, and JWT authentication in Go.
- Minimal, easy-to-follow code aimed at learners. The code shows how to perform database migrations, structure packages, and register HTTP routes.

## Repository structure (relevant files)
- `main.go` — application entry point (configures DB, migrates, initializes repository/service/adapter, and registers routes).
- `api/adapter` — adapters such as a JWT generator used by the service layer.
- `api/routes` — route definitions and HTTP handlers wiring.
- `pkg/entities` — data model definitions (e.g., `User` entity struct).
- `pkg/user` — repository and service logic for user operations (persistence, business rules).

> Note: the code expects a `routes.UserRoutes(api, userService)` function to register API endpoints under an `/api` group (see `main.go`).

## Prerequisites
- Go 1.20+ (install from https://go.dev/doc/install)
- MySQL server (or compatible) running locally or accessible remotely
- PowerShell (examples below use Windows PowerShell)

## Configuration
The example `main.go` contains an example DSN string used to connect to MySQL:

```go
dsn := "root:@tcp(127.0.0.1:3306)/sandbox_golang?charset=utf8mb4&parseTime=True&loc=Local"
```

For a production-ready setup you should replace the hard-coded DSN with environment variables. Suggested environment variables:

- `DB_USER` — database username (default: `root`)
- `DB_PASS` — database password
- `DB_HOST` — host (default: `127.0.0.1`)
- `DB_PORT` — port (default: `3306`)
- `DB_NAME` — database name (default: `sandbox_golang`)
- `JWT_SECRET` — secret key used by the JWT generator adapter

Example (PowerShell) to export environment variables for the current session:

```powershell
$env:DB_USER = "root"
$env:DB_PASS = "mypassword"
$env:DB_HOST = "127.0.0.1"
$env:DB_PORT = "3306"
$env:DB_NAME = "sandbox_golang"
$env:JWT_SECRET = "replace-with-a-secure-random-secret"
```

If you keep the DSN hard-coded in `main.go` as in the sandbox, ensure your local MySQL accepts connections with that DSN.

## Database migration
The sandbox calls `db.AutoMigrate(&entities.User{})` in `main.go`. On startup the application will create or update the `users` table to match the `User` struct in `pkg/entities`.

Make sure the database `sandbox_golang` exists and the configured user has privileges to create/alter tables. You can create it in MySQL with:

```powershell
# connect to MySQL and create database (example using mysql client)
mysql -u root -p -e "CREATE DATABASE IF NOT EXISTS sandbox_golang CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;"
```

## Run (development)
From PowerShell in the `authservice-sandbox` folder:

```powershell
# run directly
cd "d:\FILE MARLIN\Silsilah Ilmiyyah\HSI Abdullah Roy\Sandbox 25 - Golang\hsi-sandbox25\authservice-sandbox"
go run .

# or build and run the binary
go build -o authservice.exe
./authservice.exe
```

The server listens on port `8080` by default (see `app.Listen(":8080")` in `main.go`). The API base path is `/api`.

## Expected endpoints
The project registers routes via `routes.UserRoutes(api, userService)` where `api := app.Group("/api")` in `main.go`. Exact endpoints depend on the implementation in `api/routes`. Typical endpoints a user-auth service may expose include:
- `POST /api/register` — create a new user
- `POST /api/login` — authenticate and receive JWT
- `GET /api/users/:id` — fetch user details (protected)

Open `api/routes` to see the precise route definitions and handler names.

## JWT and Authentication
The project uses an adapter `api/adapter` that provides a JWT generator (instantiated in `main.go` with `adapter.NewJWTGenerator()`). The service uses the JWT generator to create tokens for authenticated users. Ensure `JWT_SECRET` is set for predictable behavior.

## Development notes
- Keep secrets out of source code. Use environment variables or a secrets manager.
- Replace hard-coded DSN in `main.go` with a configuration loader (e.g., `github.com/spf13/viper`) for flexibility.
- Add input validation and proper error handling for production readiness.
- Consider using migration tools (e.g., `golang-migrate`) instead of `AutoMigrate` for better control in multi-env deployments.

## Testing
If unit tests exist under packages, run them with:

```powershell
cd "path\to\package"
go test ./...
```

## Extending the sandbox
- Add role-based access control and middleware to protect routes.
- Implement refresh tokens and token revocation.
- Add Dockerfile and docker-compose.yaml to spin up service + MySQL easily.

## Contact / Contributing
This is a sandbox; feel free to open issues or propose improvements via pull requests. When contributing, include a short description, the rationale for changes, and tests where appropriate.

---

If you want, I can also:
- Create a `docker-compose.yml` to run MySQL + the service locally.
- Replace the hard-coded DSN with environment variable parsing and add a sample `.env` file.
- Generate a focused `README.md` in Indonesian (Bahasa Indonesia) instead of English — tell me which.
