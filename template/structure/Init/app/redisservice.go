package app

import "html/template"

var RedisSVTemplate = template.Must(template.New("").Parse(
`package service

import (
	"time"

	"gogengotest/app/repository"
)

type RedisService struct {
	RedisRepository repository.RedisRepository
}

func (us *RedisService) InsertToken(key string, value interface{}, expires time.Duration) error {
	return us.RedisRepository.InsertRedis(key, value, expires)
}

func (us *RedisService) FetchToken(key string) (res string, err error) {
	return us.RedisRepository.GettRedis(key)
}

`))