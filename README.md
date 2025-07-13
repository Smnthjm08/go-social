# go-social

A simple social media app using React (frontend) and Go (backend).

Currently in development.

**Tech stack:**

- Go
- Docker
- Postgres (Docker)
- Swagger (API docs)
- Golang Migrate (DB migrations)

**Folders:**

- `bin`: compiled app
- `cmd`: backend executables
  - `api`: runs the API server
  - `migrate`: handles database migrations
- `internal`: private Go packages
- `docs`: Swagger docs
- `scripts`: setup scripts
- `web`: frontend

**Database Migration Commands**

- **Fix the migration state (force version to 1):**

  ```bash
  migrate -path=./cmd/migrate/migrations -database="postgres://postgres:postgres@localhost/social?sslmode=disable" force 1
  ```

- **Reset the database (drop everything):**

  ```bash
  migrate -path=./cmd/migrate/migrations -database="postgres://postgres:postgres@localhost/social?sslmode=disable" drop -f
  ```

- **Re-apply migrations cleanly:**

  ```bash
  migrate -path=./cmd/migrate/migrations -database="postgres://postgres:postgres@localhost/social?sslmode=disable" up
  ```
