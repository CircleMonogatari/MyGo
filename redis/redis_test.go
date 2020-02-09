package redis

import (
	"testing"

	"github.com/go-redis/redis"
)

func InitRedisTesting(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})

	defer client.Close()

	// //ping

	// pong, err := client.Ping().Result()
	// if err != nil {

	// }
}
