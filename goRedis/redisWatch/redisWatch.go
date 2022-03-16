package redisWatch

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/superwhys/GoWeb/goRedis/redisInit"
	"time"
)

func WatchDemo() (err error) {
	// 监视watch_count的值，并在值不变的前提下将其值+1
	key := "watch_count"
	err = redisInit.Rdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		time.Sleep(time.Second * 5)
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			pipe.Set(key, n+1, 0)
			return nil
		})
		return err
	}, key)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return err
	}
	return nil
}
