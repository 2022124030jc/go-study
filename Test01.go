package main

import (
	"fmt"
	"sync"
)

func Test01() {
	// map
	mp := make(map[string]int, 10)
	var wg sync.WaitGroup
	var mu sync.Mutex // 用于保护 map 的并发访问

	for i := 0; i < 10; i++ {
		wg.Add(1) // 增加计数器
		go func() {
			defer wg.Done() // 协程完成时减少计数器

			// 写操作
			for i := 0; i < 100; i++ {
				mu.Lock()
				mp["helloworld"] = 1
				mu.Unlock()
			}

			// 读操作
			for i := 0; i < 10; i++ {
				mu.Lock()
				fmt.Println(mp["helloworld"])
				mu.Unlock()
			}
		}()
	}

	wg.Wait() // 等待所有协程完成
}
