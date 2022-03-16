package redisInit

import "github.com/go-redis/redis"

var (
	Rdb *redis.Client
)

func InitRedisClient() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = Rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
