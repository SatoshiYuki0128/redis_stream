package main

const (
	streamName = "mystream"
)

func main() {
	err := InitRedis()
	if err != nil {
		panic(err)
	}

	err = InitGin()
	if err != nil {
		panic(err)
	}

	err = GinRouter.Run(":80")
	if err != nil {
		panic(err)
	}
}
