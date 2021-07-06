package cache

import (
	"sync"
	"time"

	"whats/core/config"

	"github.com/allegro/bigcache/v3"
	"github.com/go-redis/redis/v8"
)

var (
	initOnce sync.Once
	// RDB redis client
	RDB *redis.Client
	// Cache  big cache
	Cache *bigcache.BigCache
)

// Init init cache
func Init() {
	initOnce.Do(func() {
		RDB = redis.NewClient(&redis.Options{
			Addr:     config.GetEnv("REDIS_ADDR"),
			Password: config.GetEnv("REDIS_PASSWORD"),
		})

		Cache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	})
}
