package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	NumProducers = 3
	NumConsumers = 2
)

func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		data := rand.Intn(100)
		fmt.Printf("生产者 %d 生产: %d\n", id, data)
		ch <- data
		time.Sleep(time.Millisecond * 200)
	}
}

func consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("🚚 消费者 %d 消费: %d\n", id, data)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	ch := make(chan int, 10) // 带缓冲的通道

	var wgPruduct sync.WaitGroup
	var wgConsume sync.WaitGroup

	// 启动生产者
	for i := 1; i <= NumProducers; i++ {
		wgPruduct.Add(1)
		go producer(i, ch, &wgPruduct)
	}

	// 启动消费者
	for i := 1; i <= NumConsumers; i++ {
		wgConsume.Add(1)
		go consumer(i, ch, &wgConsume)
	}

	// 等待生产者完成
	wgPruduct.Wait() // 等待所有 goroutine
	wgPruduct.Wait() // 等待所有 goroutine
	close(ch)        // 关闭通道，通知消费者结束

	// 为防止 main 提前退出，这里只等消费者结束
	time.Sleep(time.Second * 5)
	fmt.Println("🎉 所有任务结束")
}
