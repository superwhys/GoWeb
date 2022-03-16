package redisHash

import (
	"fmt"
	"github.com/superwhys/GoWeb/goRedis/redisInit"
)

func HashGetAll(key string) (map[string]string, error) {
	value, err := redisInit.Rdb.HGetAll(key).Result()
	if err != nil {
		fmt.Println("hash get all error: ", err.Error())
		return nil, err
	}
	return value, nil
}

func HashGet(key, field string) (string, error) {
	value, err := redisInit.Rdb.HGet(key, field).Result()
	if err != nil {
		fmt.Println("hash get error: ", err.Error())
		return "", err
	}
	return value, nil
}
