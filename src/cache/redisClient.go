package cache

import (
	"coindesk/config"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"time"
)

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisClient() (*RedisClient, error) {

	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var configuration config.Configurations
	if err := viper.ReadInConfig(); err != nil {
		logger.Error(err.Error())
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		logger.Error(err.Error())
	}

	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.address"),
		Password: viper.GetString("redis.password"),
		//DB:       0,
	})

	return &RedisClient{
		client: client,
		//ctx:    ctx,
	}, nil
}

func (rdb *RedisClient) GetValue(key string) (string, error) {
	value, err := rdb.client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (rdb *RedisClient) SetValue(key string, val interface{}, expiry time.Duration) error {
	_, err := rdb.client.Set(key, val, expiry).Result()
	if err == redis.Nil {
		fmt.Println(key + " does not exist")
	} else if err != nil {
		return err
	}
	return nil
}
