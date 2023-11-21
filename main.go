package main

import (
	"redis_stream/redis"
	"redis_stream/router"
)

const (
	streamName = "mystream"
)

func main() {
	// 啟用 redis
	err := redis.InitRedis()
	if err != nil {
		panic(err)
	}

	// 設置 gin
	err = router.InitGin()
	if err != nil {
		panic(err)
	}

	err = router.GinRouter.Run(":80")
	if err != nil {
		panic(err)
	}
}
