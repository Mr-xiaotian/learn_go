package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("[生产者] 正在生产:", i)
		ch <- i
		time.Sleep(500 * time.Millisecond) // 模拟生产延迟
	}
	close(ch) // 生产完成，关闭通道
}

func consumer(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // goroutine 结束时通知 wg

	for val := range ch {
		fmt.Printf("[消费者 %d] 消费了: %d\n", id, val)
		time.Sleep(300 * time.Millisecond) // 模拟消费延迟
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	// 启动生产者
	go producer(ch)

	// 启动多个消费者（这里是 2 个）
	consumerCount := 2
	wg.Add(consumerCount)
	for i := 1; i <= consumerCount; i++ {
		go consumer(i, ch, &wg)
	}

	wg.Wait() // 等待所有消费者完成
	fmt.Println("主程序结束")
}
