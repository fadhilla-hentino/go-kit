package datasource

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(host, port, password string) *redis.Client {
	redisOptions := &redis.Options{
		Addr:            fmt.Sprintf("%s:%s", host, port),
		Password:        password,
		DB:              0,
		MaxRetries:      10,
		MaxRetryBackoff: time.Second,
		PoolSize:        10,
	}

	return redis.NewClient(redisOptions)
}
