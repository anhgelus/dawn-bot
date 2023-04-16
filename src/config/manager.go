package config

import (
	"dawn-bot/src/db/postgres"
	"dawn-bot/src/db/redis"
	"dawn-bot/src/utils"
	"errors"
	"github.com/pelletier/go-toml"
	"gorm.io/gorm"
	"os"
)

type DatabasesConfig struct {
	Postgres postgres.PostgresOptions
	Redis    redis.RedisOptions
}

func LoadAndImportDatabaseConfig() DatabasesConfig {
	c, err := os.ReadFile(Path + "databases.toml")
	utils.PanicError(err)
	var config DatabasesConfig
	err = toml.Unmarshal(c, &config)
	utils.PanicError(err)

	return config
}

// GetConfig return the config of the guild and true if it's a new config
func GetConfig(guildId string) (postgres.Config, bool) {
	db := postgres.Db
	config := postgres.Config{GuildID: guildId}
	result := postgres.Db.Limit(1).Find(&config)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.Create(&postgres.Config{
			WelcomeChannelID: "",
			GuildID:          guildId,
		})
		return config, true
	}
	return config, false
}
