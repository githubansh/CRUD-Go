# User Service

A simple production-style Go backend project using Gin.

## Project Structure

```text
user-service/
|-- cmd/
|   `-- server/
|       `-- main.go
|-- internal/
|   |-- handler/
|   |   `-- user_handler.go
|   |-- service/
|   |   `-- user_service.go
|   |-- repository/
|   |   `-- user_repository.go
|   `-- model/
|       `-- user.go
|-- go.mod
`-- README.md
```

## Requirements

- Go installed
- Internet access for the first `go mod tidy`

## Setup

```bash
go mod tidy
```

## Run the Server

```bash
go run ./cmd/server
```

The server starts on:

```text
http://localhost:8080
```

## APIs

### Create User

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Ansh\",\"age\":22,\"email\":\"ansh@example.com\"}"
```

Example response:

```json
{
  "id": "1",
  "name": "Ansh",
  "age": 22,
  "email": "ansh@example.com"
}
```

### Get User By ID

```bash
curl http://localhost:8080/users/1
```

Example response:

```json
{
  "id": "1",
  "name": "Ansh",
  "age": 22,
  "email": "ansh@example.com"
}
```

## Notes

- Data is stored in memory, so users are lost when the server stops.
- The request flow is `handler -> service -> repository`.
