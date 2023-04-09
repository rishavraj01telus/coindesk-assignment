package config

import (
	"github.com/spf13/viper"
)

type RedisConfig struct {
	ClientAddress string
	Password      string
	Db            int
}

func NewRedisConfig(clientAddress string, password string, db int) *RedisConfig {
	return &RedisConfig{
		ClientAddress: clientAddress,
		Password:      password,
		Db:            db,
	}
}

func GetRedisConfig() *RedisConfig {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	config := NewRedisConfig(
		viper.GetString("redis.address"),
		viper.GetString("redis.password"),
		0,
	)
	return config
}
