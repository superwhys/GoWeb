package redisGet

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/superwhys/GoWeb/goRedis/redisInit"
)

func RedisGet(key string) (string, error) {
	value, err := redisInit.Rdb.Get(key).Result()
	if err == redis.Nil {
		fmt.Printf("key: %v not exist", key)
		return "", err
	} else if err != nil {
		fmt.Println("Get key error: ", err.Error())
		return "", err
	} else {
		return value, nil
	}
}
