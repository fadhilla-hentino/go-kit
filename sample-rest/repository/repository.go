package repository

import "context"

type UserRepository interface {
	Store(ctx context.Context, key, value string) error
	Load(ctx context.Context, key string) (string, error)
}
