# Go Template Project

This is a starter project for building web APIs with Go. It uses the [Gin](https://github.com/gin-gonic/gin) web framework and [Uber Fx](https://github.com/uber-go/fx) for dependency injection.

## Project Structure

```
.
├── cmd/main.go             # App entry point
├── docs/                   # Swagger generated files
├── internal/
│   ├── app/                # App setup, routes, DI
│   ├── controller/         # Presentation layer
│   ├── service/            # Business layer
│   ├── repository/         # Data access layer
│   └── infra/              # Infra utilities
```

## Getting Started

### 1. Install dependencies

```bash
go mod tidy
```

### 2. Install `swag` CLI (if you haven't)

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### 3. Run the project

Use VSCode launch config:

Open VSCode

Press F5 or go to Run & Debug sidebar

Select `Debug Gin Project` and start debugging

Server will start at:  
`http://localhost:8080`

### 4. Open Swagger docs

You can check the API docs at:  
`http://localhost:8080/swagger/index.html`