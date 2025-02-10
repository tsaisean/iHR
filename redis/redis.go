package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"iHR/config"
)

var (
	RedisClient *redis.Client
)

func Connect(redisCfg config.Redis) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", redisCfg.Host, redisCfg.Port),
	})
}
