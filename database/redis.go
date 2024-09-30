package database

import (
	"context"
	"fmt"
	"time"

	redispkg "github.com/redis/go-redis/v9"
)

type Redis struct{}

var client *redispkg.Client
var RedisInstance Redis

func RedisInit() {
	client = redispkg.NewClient(&redispkg.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func (r Redis) Get(key string) (string, error) {
	ctx := context.Background()
	val, err := client.Get(ctx, fmt.Sprintf("%v", key)).Result()

	if err != nil {
		return "", err
	}
	return val, nil
}

func (r Redis) Set(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	err := client.Set(ctx, key, value, expiration).Err()

	if err != nil {
		return err
	}
	return nil
}
