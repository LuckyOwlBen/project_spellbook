package definitions

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ProbeRedis(ipAddress string) {
	fmt.Println("Probing Redis on", ipAddress)
	rdb := redis.NewClient(&redis.Options{
		Addr:     ipAddress, // use the appropriate address and port
		Password: "",        // no password set
		DB:       0,         // use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	fmt.Println("Redis connected:", pong)
	keys := rdb.Keys(ctx, "*")
	fmt.Println("Keys in Redis:")
	for _, key := range keys.Val() {
		fmt.Println(key)
		if key == "flag" {
			fmt.Println("Flag found in Redis")
			fmt.Println(rdb.Get(ctx, key))
		}
	}
}
