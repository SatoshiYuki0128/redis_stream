package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() error {
	options := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       9,
	}

	RedisClient := redis.NewClient(options)

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	return nil
}

func AddStream(client *redis.Client, stream string, message map[string]interface{}) (string, error) {
	xAddArgs := &redis.XAddArgs{
		Stream: streamName,
		Values: message,
	}

	xAddResult, err := client.XAdd(context.Background(), xAddArgs).Result()
	if err != nil {
		return "", err
	}

	return xAddResult, nil
}
