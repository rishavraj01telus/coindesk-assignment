package cache

import (
	"coindesk/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() (*RedisClient, error) {

	redisConfig := config.GetRedisConfig()

	redisClient := redis.NewClient(
		&redis.Options{
			Addr:     redisConfig.ClientAddress,
			Password: redisConfig.Password,
			DB:       redisConfig.Db,
		})
	return &RedisClient{
		client: redisClient,
	}, nil
}

func (rdb *RedisClient) GetValue(ctx context.Context, key string) (string, error) {
	value, err := rdb.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (rdb *RedisClient) SetValue(ctx context.Context, key string, val interface{}, expiry time.Duration) error {
	_, err := rdb.client.Set(ctx, key, val, expiry).Result()
	if err == redis.Nil {
		fmt.Println(key + " does not exist")
	} else if err != nil {
		return err
	}
	return nil
}
