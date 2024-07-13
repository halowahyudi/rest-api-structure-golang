# Simple Rest Api Golang Structure.

This is a Go project using Gin for REST API with route definitions directly in `route.go` and MySQL integration.

## Prerequisites

- Go 1.16 or later
- MySQL database

## Packages

The project uses the following Go packages:

- `github.com/gin-gonic/gin`: Web framework for building REST APIs
- `github.com/go-sql-driver/mysql`: MySQL driver for Go's database/sql package
- `gopkg.in/yaml.v2`: YAML parser and emitter for Go

## Usage

To run the project:

1. Clone the repository.
   ```sh
   git clone https://github.com/halowahyudi/rest-api-structure-golang.git
   ```
   ```sh
   cd rest-api-structure-golang
   ```
2. Set up your `config/config.yaml` file with your database and server details.
   ```yaml
   database:
     user: your_db_user
     password: your_db_password
     host: localhost
     name: your_db_name
   server:
     port: "8080"
   ```
3. Install dependencies using `go mod tidy`.
   ```sh
   go mod tidy
   ```
4. Run the application using `go run cmd/app/main.go`.
   ```sh
   go run cmd/app/main.go
   ```

## Routes

- `/ping`: GET endpoint to test if the server is running.
- `/dbtest`: GET endpoint to test if the database connection is working.
