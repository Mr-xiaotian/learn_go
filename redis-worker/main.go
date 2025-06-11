package main

import (
	"context"
	"learn_go/redis-worker/worker"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	worker.StartWorker(
		ctx,
		rdb,
		"RedisTransfer[_process_via_redis]:input",  // Redis 中任务输入的 List
		"RedisTransfer[_process_via_redis]:output", // Redis 中结果写入的 Hash
		worker.ParseNumberTask,
		worker.Fibonacci,
	)
}
