package tgfsmdb

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/sneiko/go-tgbuilder/pkg/tgfsm"
)

func keyUser(id string) string { return fmt.Sprintf("user:%s", id) }

// Redis db
type Redis struct {
	client *redis.Client
}

func NewRedis(opts redis.Options) *Redis {
	return &Redis{
		client: redis.NewClient(&opts),
	}
}

func (r *Redis) GetUser(ctx context.Context, id string) (*tgfsm.User, error) {
	var user *tgfsm.User

	res := r.client.Get(ctx, keyUser(id))
	if err := res.Scan(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Redis) SaveKeyFrame(ctx context.Context, id string, keyFrame string) error {
	return r.client.Set(ctx, keyUser(id), tgfsm.User{
		ID:       id,
		KeyFrame: keyFrame,
	}, 0).Err()
}
