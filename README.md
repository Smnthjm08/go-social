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

smnthjm08:go-social% migrate -path=./cmd/migrate/migrations -database="postgres://postgres:postgres@localhost/social?sslmode=disable" force 1

smnthjm08:go-social% migrate -path=./cmd/migrate/migrations -database="postgres://postgres:postgres@localhost/social?sslmode=disable" drop -f

smnthjm08:go-social% migrate -path=./cmd/migrate/migrations -database="postgres://postgres:postgres@localhost/social?sslmode=disable" up

Fixed the migration state (force version to 1)

Reset the database (drop everything)

Re-applied migrations cleanly
