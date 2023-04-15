package config

import (
	"dawn-bot/src/src/db/postgres"
	"dawn-bot/src/src/db/redis"
	"dawn-bot/src/utils"
	"github.com/pelletier/go-toml"
	"os"
)

type DatabasesConfig struct {
	Postgres postgres.PostgresOptions
	Redis    redis.RedisOptions
}

func LoadAndImportDatabaseConfig() DatabasesConfig {
	path, err := os.Executable()
	utils.PanicError(err)
	content, err := os.ReadFile(path + ConfigPath + "databases.toml")
	utils.PanicError(err)
	var config DatabasesConfig
	err = toml.Unmarshal(content, &config)
	utils.PanicError(err)

	return config
}