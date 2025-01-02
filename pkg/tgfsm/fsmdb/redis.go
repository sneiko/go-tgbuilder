package fsmdb

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"tg_star_miner/pkg/tgfsm/fsmmodel"
)

func keyUser(id uuid.UUID) string { return fmt.Sprintf("user:%s", id.String()) }

// Redis db
type Redis struct {
	client *redis.Client
}

func NewRedis(opts redis.Options) *Redis {
	return &Redis{
		client: redis.NewClient(&opts),
	}
}

func (r *Redis) GetUser(ctx context.Context, id uuid.UUID) (*fsmmodel.User, error) {
	var user *fsmmodel.User

	res := r.client.Get(ctx, keyUser(id))
	if err := res.Scan(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Redis) SaveKeyFrame(ctx context.Context, id uuid.UUID, keyFrame string) error {
	return r.client.Set(ctx, keyUser(id), fsmmodel.User{
		ID:       id,
		KeyFrame: keyFrame,
	}, 0).Err()
}
