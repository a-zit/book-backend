## Prerequisites

- [Go] version 1.22.4
- [Docker] version 20.10.8

## Installation & Structure

### Clone the Repository & Install Dependencies

```
git clone https://github.com/a-zit/book-backend.git

cd book-backend

go mod tidy
```

## Project Structure

Concept of microservices & clean architecture with Golang

```
book-backend/
├── LICENSE
├── Makefile -- Makefile for common tasks
├── README.md
├── cmd
│   ├── backend-bo-api
│   │   ├── config
│   │   │   └── config.go -- Configuration file
│   │   ├── handler
│   │   │   └── book.go -- HTTP handler
│   │   ├── main.go
│   │   └── middleware
│   │       └── error.go -- Error handling middleware
│   └── backend-client-api
│       └── main.go
├── database -- Database configuration and scripts
│   ├── mariadb.go
│   └── script.sql
├── docker-compose.yml
├── domain
│   ├── book.go -- Book entity, use-case, repository interface
│   ├── common.go
│   ├── error.go
│   └── mocks
│       └── BookRepository.go -- Mock repository for testing
├── go.mod
├── go.sum
├── module
│   └── book
│       ├── book_suite
│       │   └── book_suite.go -- Test suite for book
│       ├── repository.go -- repository for book
│       ├── usecase.go -- business logic for book
│       └── usecase_test.go
├── pkg
│   ├── error
│   │   └── app_error.go -- Custom Application error
│   └── validator
│       └── validator.go -- Request validation
```

## Running the Application

### Terminal

```
docker-compose up -d

export DATABASE_URL="user:password@tcp(localhost:3306)/mydatabase"
export PORT=9000

go run cmd/backend-bo-api/main.go
```

### VS Code

If you have VS Code, you can run the application by using the VS Code launch.json file.

## Testing

Describe how to run tests for the project.

### Running Tests with coverage

```
make test coverage
```
