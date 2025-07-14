# Database Migration Commands

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
