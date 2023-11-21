package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func InitRedis() error {
	options := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       9,
	}

	Client = redis.NewClient(options)

	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	return nil
}

func AddStream(streamName string, message map[string]interface{}) error {
	xAddArgs := &redis.XAddArgs{
		Stream: streamName,
		Values: message,
	}

	xAddResult, err := Client.XAdd(context.Background(), xAddArgs).Result()
	if err != nil {
		fmt.Println("Failed to add message to Stream:", err)
		return err
	}

	fmt.Println("Message added to Stream:", xAddResult)

	return nil
}
