package cache

import (
	"fmt"
	"sync"

	"whats/core/config"

	"github.com/go-redis/redis/v8"
)

var initOnce sync.Once
var RDB *redis.Client

// Init init cache
func Init() {
	initOnce.Do(func() {
		RDB = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", config.GetEnv("REDIS_HOST"), config.GetEnv("REDIS_PORT")),
			Password: config.GetEnv("REDIS_PASSWORD"),
		})
	})
}
