package cache

import (
	"sync"

	"whats/core/config"

	"github.com/go-redis/redis/v8"
)

var initOnce sync.Once

// RDB redis client
var RDB *redis.Client

// Init init cache
func Init() {
	initOnce.Do(func() {
		RDB = redis.NewClient(&redis.Options{
			Addr:     config.GetEnv("REDIS_ADDR"),
			Password: config.GetEnv("REDIS_PASSWORD"),
		})
	})
}
