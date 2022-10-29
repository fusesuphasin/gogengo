package repository

import (
	"context"
	"fmt"
	"time"

	"gogengotest/app/infrastructure"
)

type RedisRepository struct {
	Ctx context.Context
}

func (rr *RedisRepository) InsertRedis(key string, value interface{}, expires time.Duration) error {
	redisClient := infrastructure.RedisInit()
	set := redisClient.Set(rr.Ctx, key, value, expires).Err()
	defer redisClient.Close()
	return set
}

func (rr *RedisRepository) GettRedis(key string) (res string, err error) {
	redisClient := infrastructure.RedisInit()
	
	get, err := redisClient.Get(rr.Ctx, key).Result()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer redisClient.Close()
	return get, err 
}