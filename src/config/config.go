package config

type Configurations struct {
	Redis  RedisConfiguration
	Server ServerConfigurations
}

type RedisConfiguration struct {
	RedisAddr     string
	RedisPassword string
}

type ServerConfigurations struct {
	Port int
}
