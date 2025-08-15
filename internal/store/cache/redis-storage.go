package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/smnthjm08/go-social/internal/store"
)

type Strorage struct {
	Users interface {
		Get(context.Context, int64) (*store.User, error)
		Set(context.Context, *store.User) error
	}
}

func NewRedisStorage(rdb *redis.Client) Strorage {
	return Strorage{
		Users: &UserStore{rdb: rdb},
	}
}
