# Go REST API

A REST API built with Go, PostgreSQL, and GORM.

## Prerequisites

Make sure you have the following installed:

- Go
- PostgreSQL/NeonDB

## Configuration

Before starting the server, create the configuration file in the root directory:

```text
config/local.yaml
```

Add the following configuration:

```yaml
env: "dev"

postgres:
  database_url: "<your_neondb_url>"

http_server:
  address: "localhost:8080"
```

Replace `<your_neondb_url>` with your PostgreSQL connection string.

For example:

```yaml
postgres:
  database_url: "postgresql://<db_username>:<db_password>@<neondb_host_server>/neondb?sslmode=require&channel_binding=require"
```

## Install the required modules

```bash
go mod tidy
```

## Database migration for first time (Important)

Uncomment this section from the `cmd/rest-api/main.go` file.

```go
// other code
import (
  // other modules
  "github.com/Subham-Das-98/go-rest-api/internal/models"
)

func main() {
  // other code

  // migration
  err = db.AutoMigrate(&models.User{})
  if err != nil {
    slog.Error("database migration failed",
      slog.String("error", err.Error()),
    )
    os.Exit(1)
  }

  // other code
}

```

## Run the Server

Start the API using:

```bash
go run cmd/rest-api/main.go -config config/local.yaml
```

The server will be available at:

```text
http://localhost:8080
```

## Project Structure

```text
.
├── cmd/
│   └── rest-api/
│       └── main.go
├── config/
│   └── local.yaml
├── internal/
│   ├── config/
│   ├── handler/
│   ├── middleware/
│   ├── models/
│   ├── repository/
│   ├── router/
│   ├── service/
│   └── ...
├── go.mod
├── go.sum
└── README.md
```

## Environment

The local configuration uses:

```yaml
env: "dev"
```

You can create separate configuration files for other environments, such as staging or production.

## Database

The application uses NeonDB PostgreSQL as its database. Configure the connection through:

```yaml
postgres:
  database_url: "<your_neondb_url>"
```

## Development

Run the application directly with Go:

```bash
go run cmd/rest-api/main.go -config config/local.yaml
```

To format the Go source code:

```bash
gofmt -w .
```

To run tests:

```bash
go test ./...
```
