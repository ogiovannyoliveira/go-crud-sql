# go-crud-sql

A simple REST API built with Go that performs CRUD operations on users persisted in a PostgreSQL database.

## How it works

Users are stored in a PostgreSQL database. Each user has `first_name`, `last_name`, and `biography` fields. IDs are UUIDs generated on creation. Routing is handled by [chi](https://github.com/go-chi/chi).

## Database

The app connects to PostgreSQL using the following default connection string (defined in `internal/store/store.go`):

```
user=giovanny dbname=go-crud-sql password=secret sslmode=disable
```

### Start the database with Docker

A `compose.yml` is provided under the `docker/` directory. It starts a PostgreSQL container and runs the table creation script automatically.

```bash
docker compose -f docker/compose.yml up -d
```

This will:
- Start a PostgreSQL instance on port **5432**
- Create the `users` table via `docker/init-db-scripts/create-table-users.sql`

To stop it:

```bash
docker compose -f docker/compose.yml down
```

## How to run

With the database running, start the server:

```bash
go run cmd/api/main.go
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
