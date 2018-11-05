package redis

import "C"
import (
	"GinWeb/config"
	"fmt"
	"github.com/go-redis/redis"
)

type connect struct {
	Default *redis.Client
}

var Connect connect

func init() {
	Connect.Default = redis.NewClient(&redis.Options{
		Addr:     config.C.Redis.Default.Addr,
		Password: config.C.Redis.Default.Password, // no password set
		DB:       config.C.Redis.Default.DB,       // use default DB
	})
	pong, err := Connect.Default.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		panic(err)
	}
}
