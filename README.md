
# Book Backend

## Table of Contents

1. [Introduction](#introduction)
2. [Prerequisites](#prerequisites)
3. [Installation](#installation)
4. [Project Structure](#project-structure)
5. [Configuration](#configuration)
6. [Usage](#usage)
7. [Testing](#testing)

## Introduction

Project to understand clean architecture with Golang by assuming that we are building a book backend service.

## Prerequisites

List the necessary prerequisites for setting up the project. Include versions where applicable.

- [Go] version 1.22.4
- [Docker] version 20.10.8

## Installation

### Clone the Repository

\`\`\`bash
git clone https://github.com/a-zit/book-backend.git
cd book-backend
\`\`\`

### Install Dependencies

Ensure you have Go installed and set up correctly. Then, install the required dependencies:

\`\`\`bash
go mod tidy
\`\`\`

## Project Structure

Explain the structure of the project directory. Highlight important directories and files.

\`\`\`plaintext
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
\`\`\`

## Configuration

Describe how to configure the project. Include any environment variables or configuration files needed.

### Environment Variables

| Variable       | Description                   | Default Value |
|----------------|-------------------------------|---------------|
| \`PORT\`         | Port to run the application on| \`9000\`        |
| \`DATABASE_URL\` | Database connection URL       | \`user:password@tcp(localhost:3306)/mydatabase\`          |

## Usage

Explain how to run and use the project. Provide examples of common tasks.

### Running the Application

#### Docker Compose

The project includes a \`docker-compose.yml\` file for running the application with a MariaDB database.

### Running the Application

\`\`\`bash
docker-compose up -d

export DATABASE_URL="user:password@tcp(localhost:3306)/mydatabase"
export PORT=9000

go run cmd/backend-bo-api/main.go
\`\`\`

#### VS Code

If you have VS Code, you can run the application by using the VS Code launch.json file.

### Example Usage

Provide examples of how to use the application. This could include API endpoints, CLI commands, or other interfaces.

\`\`\`bash
curl http://localhost:9000/health
\`\`\`

## Testing

Describe how to run tests for the project.

### Running Tests

\`\`\`bash
go test ./...

# Run tests with coverage
make test coverage
\`\`\`
