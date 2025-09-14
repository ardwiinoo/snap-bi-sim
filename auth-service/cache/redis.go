package cache

import (
	"context"
	"fmt"
	"log"

	"github.com/ardwiinoo/snap-bi-sim/auth-service/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis(cfg *config.AppConfig) *redis.Client {
	addr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.RedisPass,
		DB:       0,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("failed to connect Redis: %v", err)
	}

	log.Println("Connected to Redis:", addr)
	return rdb
}
