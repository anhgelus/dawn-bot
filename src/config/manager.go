package config

import (
	"dawn-bot/src/db/postgres"
	"dawn-bot/src/db/redis"
	"dawn-bot/src/utils"
	"github.com/pelletier/go-toml"
	"os"
)

type DatabasesConfig struct {
	Postgres postgres.PostgresOptions
	Redis    redis.RedisOptions
}

func LoadAndImportDatabaseConfig() DatabasesConfig {
	c, err := os.ReadFile(ConfigPath + "databases.toml")
	utils.PanicError(err)
	var config DatabasesConfig
	err = toml.Unmarshal(c, &config)
	utils.PanicError(err)

	return config
}
