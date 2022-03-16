package redisPipeline

import (
	"fmt"
	"github.com/superwhys/GoWeb/goRedis/redisInit"
	"time"
)

func PipelineDemo() {
	pipe := redisInit.Rdb.Pipeline()

	incr := pipe.Incr("pipeline_counter")
	pipe.Expire("pipeline_counter", time.Hour)

	_, err := pipe.Exec()
	fmt.Println(incr.Val(), err)
}

// TxPipelineDemo 事务， 能够确保两个语句之间的命令之间没有其他客户端正在执行命令。
func TxPipelineDemo() {
	pipe := redisInit.Rdb.TxPipeline()

	incr := pipe.Incr("tx_pipeline_counter")
	pipe.Expire("tx_pipeline_counter", time.Hour)

	_, err := pipe.Exec()
	fmt.Println(incr.Val(), err)
}
