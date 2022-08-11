package user

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type repo struct {
	redisClient *redis.Client
}

func NewUserRepo(redisClient *redis.Client) *repo {
	return &repo{
		redisClient: redisClient,
	}
}

func (repo *repo) Store(ctx context.Context, key, value string) error {
	err := repo.redisClient.SetEX(ctx, key, value, 1*time.Hour).Err()
	return err
}

func (repo *repo) Load(ctx context.Context, key string) (string, error) {
	return repo.redisClient.Get(ctx, key).Result()
}
