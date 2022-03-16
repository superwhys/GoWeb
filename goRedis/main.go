package main

import (
	"fmt"
	"github.com/superwhys/GoWeb/goRedis/redisGet"
	"github.com/superwhys/GoWeb/goRedis/redisHash"
	"github.com/superwhys/GoWeb/goRedis/redisInit"
	"github.com/superwhys/GoWeb/goRedis/redisPipeline"
	"github.com/superwhys/GoWeb/goRedis/redisSet"
	"github.com/superwhys/GoWeb/goRedis/redisWatch"
)

func main() {
	err := redisInit.InitRedisClient()
	if err != nil {
		fmt.Println("redis init error: ", err.Error())
		return
	}
	defer redisInit.Rdb.Close()

	fmt.Println("redis inint success! ")

	err = redisSet.RedisSet("name", "superyong", 0)
	if err != nil {
		return
	}
	value, err := redisGet.RedisGet("name")
	if err != nil {
		return
	}
	fmt.Println("value is: ", value)

	hashValue, err := redisHash.HashGetAll("user")
	if err != nil {
		return
	}
	fmt.Println(hashValue)

	value, err = redisHash.HashGet("user", "name")
	if err != nil {
		return
	}
	fmt.Println(value)

	redisPipeline.PipelineDemo()
	value, err = redisGet.RedisGet("pipeline_counter")
	if err != nil {
		return
	}
	fmt.Println("value is: ", value)

	redisPipeline.TxPipelineDemo()
	err = redisWatch.WatchDemo()
	if err != nil {
		return
	}
}
