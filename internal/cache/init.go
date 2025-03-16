package cache

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisCache interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (interface{}, error)
}

type RedisCacheImpl struct {
	cli *redis.Client
}

func NewRedisClient(ctx context.Context) RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("could not connect to redis, err: %v", err)
	}
	return &RedisCacheImpl{
		cli: client,
	}
}

func (r *RedisCacheImpl) Set(ctx context.Context, key string, value interface{}) error {
	if res := r.cli.Set(ctx, key, value, 0); res.Err() != nil {
		log.Printf("failed to set, err: %v", res.Err())
		return res.Err()
	}
	return nil
}

func (r *RedisCacheImpl) Get(ctx context.Context, key string) (interface{}, error) {
	value, err := r.cli.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return value, nil
}
