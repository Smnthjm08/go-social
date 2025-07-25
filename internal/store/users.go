package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `
	INSERT INTO users (username, password, email)
	VALUES ($1, $2, $3) RETURNING id, created_at
	`

	err := s.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) GetByID(ctx context.Context, userId int64) (*User, error) {
	query := `
		SELECT id, username, password, email, created_at
		FROM users
		WHERE id = $1
	`

	user := &User{}
	err := s.db.QueryRowContext(
		ctx,
		query,
		userId).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}

	}

	return user, nil
}
