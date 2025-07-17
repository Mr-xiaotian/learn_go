package main

import (
	"fmt"
	"sync"
)

func chapter_10() {
	// 并发打印数字
	fmt.Println("\n并发打印数字:")
	var wg sync.WaitGroup
	for i := 1; i < 6; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()

	// 通道通信
	fmt.Println("\n通道通信:")
	intChan := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range intChan {
			fmt.Println(i)

		}
	}()
	for i := 6; i < 16; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Wait()

	// select 多路监听
	fmt.Println("\nselect 多路监听:")
	in1 := make(chan int)
	in2 := make(chan int)
	go func() {
		for i := 16; i < 21; i++ {
			in1 <- i
		}
		close(in1)
	}()
	go func() {
		for i := 20; i < 26; i++ {
			in2 <- i
		}
		close(in2)
	}()
	for in1 != nil || in2 != nil {
		select {
		case v, ok := <-in1:
			if !ok {
				in1 = nil
			} else {
				fmt.Println(v)
			}
		case v, ok := <-in2:
			if !ok {
				in2 = nil
			} else {
				fmt.Println(v)
			}
		}
	}

	// 互斥锁保证并发安全
	fmt.Println("\n互斥锁保证并发安全:")
	var l sync.RWMutex
	shareNum := 0

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			l.Lock()
			defer wg.Done()
			defer l.Unlock()
			shareNum += 1
		}()
	}
	wg.Wait()
	fmt.Println(shareNum)

}
