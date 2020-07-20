package core

import (
	"context"
	"gin_in_action/configs"
	"github.com/go-redis/redis/v8"
	"time"
)

var Redis *redis.Client
var RedisCtx = context.Background()

func connect() {
	if Redis != nil {
		return
	}

	Redis = redis.NewClient(&redis.Options{
		Addr:     configs.RedisAddr,
		Password: configs.RedisPassword,
		DB:       configs.RedisDB,
	})

	_, err := Redis.Ping(RedisCtx).Result()
	if err != nil {
		panic(err)
	}
}

func RedisSet(key string, value interface{}, exp time.Duration) {
	connect()
	err := Redis.Set(RedisCtx, key, value, exp).Err()
	if err != nil {
		panic(err)
	}
}

func RedisGet(key string) string {
	connect()
	value, err := Redis.Get(RedisCtx, key).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		panic(err)
	} else {
		return value
	}
}

func RedisDel(keys ...string) {
	connect()
	Redis.Del(RedisCtx, keys...)
}
