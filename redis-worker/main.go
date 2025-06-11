package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Payload struct {
	ID   string `json:"id"`
	Task int    `json:"task"`
}

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // 改成你的 Redis 地址
		DB:   0,
	})

	for {
		// 阻塞读取任务
		result, err := rdb.BLPop(ctx, 0*time.Second, "RedisTransfer[_process_via_redis]:input").Result()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// result[1] 是 payload 字符串
		var payload Payload
		err = json.Unmarshal([]byte(result[1]), &payload)
		if err != nil {
			fmt.Println("JSON parse error:", err)
			continue
		}

		// 🔚 检查是否为终止信号
		if payload.ID == "TERMINATION_SIGNAL" {
			rdb.HSet(ctx, "RedisTransfer[_process_via_redis]:output", "TERMINATION_SIGNAL", "TERMINATION_SIGNAL")
			fmt.Println("Received TERMINATION_SIGNAL, exiting.")
			break
		}

		// 👇 实际处理逻辑：加 100
		processed := payload.Task + 100

		// 写入结果 hash
		outputKey := "RedisTransfer[_process_via_redis]:output"
		resultStr, _ := json.Marshal(processed)

		rdb.HSet(ctx, outputKey, payload.ID, resultStr)

		fmt.Printf("Processed task %s: %d -> %d\n", payload.ID, payload.Task, processed)
	}
}
