package app

import "html/template"

var RedisTemplate = template.Must(template.New("").Parse(
	`package infrastructure

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

// RedisInit - Initialize the redis client
func RedisInit() ( *redis.Client) {
	dsn := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	pw := os.Getenv("REDIS_PASSWORD")
	redisClient  := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: pw, // no password set to ""
		DB:       0,  // use default DB
	})
	
	pong, err := redisClient.Ping(context.Background()).Result()
	log.Println(pong, err)

	return redisClient
}

`))
