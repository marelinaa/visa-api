package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Repo manages the database operations related to carts.
type Repo struct {
	redisClient *redis.Client
}

// NewRepository creates a new Repository with the given database connection.
func NewRepository(redisClient *redis.Client) *Repo {
	repo := &Repo{
		redisClient: redisClient,
	}

	return repo
}

func (repo *Repo) Set(key string, value interface{}, exp time.Duration) error {
	return repo.redisClient.Set(key, value, exp).Err()
}

func (repo *Repo) Delete(key string) error {
	err := repo.redisClient.Del(key).Err()
	if err != nil {
		return fmt.Errorf("failed to from redis: %w", err)
	}

	return nil
}

// redis-cli --scan --pattern '*'
// brew services stop redis
