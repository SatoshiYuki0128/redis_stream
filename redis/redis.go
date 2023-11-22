package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	_const "redis_stream/const"
)

var Client *redis.Client

func InitRedis() error {
	options := &redis.Options{
		Addr:     "redis:6379",
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

func DeleteStream(ctx context.Context, id string) error {
	_, err := Client.XDel(ctx, _const.StreamName, id).Result()
	if err != nil {
		return err
	}

	return nil
}

func GetStreamList(ctx context.Context) ([]redis.XStream, error) {
	messages, err := Client.XRead(ctx, &redis.XReadArgs{
		Streams: []string{_const.StreamName, "0"},
		Count:   10,
		Block:   0, // 使用 0 表示非阻塞
	}).Result()

	if err != nil {
		return nil, err
	}

	return messages, nil
}
