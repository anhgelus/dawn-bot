package redis

import "github.com/redis/go-redis/v9"

type RedisOptions struct {
	Host     string `toml:"host"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

var client *redis.Client

// ConnectClient connect to the Redis Client
func ConnectClient(o RedisOptions) {
	options := o.toOptions()
	client = redis.NewClient(&options)
}

func (o *RedisOptions) toOptions() redis.Options {
	return redis.Options{
		Addr:     o.Host,
		Password: o.Password,
		DB:       o.DB,
	}
}

// Connect to the RedisClient
func (o *RedisOptions) Connect() {
	ConnectClient(*o)
}
