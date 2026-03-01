# go-crud-in-memory

A simple REST API built with Go that performs CRUD operations on users stored in memory (no database). Data is lost on restart.

## How it works

The app stores users in a `map[ID]User` in memory. Each user has `first_name`, `last_name`, and `biography` fields. IDs are UUIDs generated on creation. Routing is handled by [chi](https://github.com/go-chi/chi).

## How to run

```bash
go run main.go
```

The server starts on port **8080**.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/health` | Health check |
| `POST` | `/api/users` | Create a user |
| `GET` | `/api/users` | List all users |
| `GET` | `/api/users/{id}` | Get user by ID |
| `PUT` | `/api/users/{id}` | Update user by ID |
| `DELETE` | `/api/users/{id}` | Delete user by ID |

## User payload

```json
{
  "first_name": "John",
  "last_name": "Doe",
  "biography": "A short bio of at least 20 characters."
}
```

**Validation rules:** `first_name` and `last_name` must be 2–20 chars; `biography` must be 20–450 chars.
