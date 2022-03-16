package redisSet

import (
	"fmt"
	"github.com/superwhys/GoWeb/goRedis/redisInit"
	"time"
)

func RedisSet(key string, value interface{}, timeOut time.Duration) error {
	err := redisInit.Rdb.Set(key, value, timeOut).Err()
	if err != nil {
		fmt.Println("set error: ", err.Error())
		return err
	}
	return nil
}
