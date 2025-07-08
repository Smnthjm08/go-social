package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Users interface {
		Create(context.Context, *User) error
	}
	Posts interface {
		Create(context.Context, *Post) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UserStore{db},
	}
}
