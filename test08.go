package main

// 添加互斥锁
// 互斥锁是用来保护共享资源的，避免多个 goroutine 同时访问同一资源导致数据不一致的问题

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

)

var wait sync.WaitGroup
var count  = 0
var mutex sync.Mutex

func Test08() {
	wait.Add(10)
	for i := 0; i < 10; i++ {
		go func(data *int) {
			// 模拟访问耗时
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			// 加锁
			mutex.Lock()
			defer mutex.Unlock()

			// 访问数据
			temp := *data
			// 模拟计算耗时
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			ans := 1
			// 修改数据
			*data = temp + ans
			fmt.Println(*data)
			wait.Done()
		}(&count)
	}
	wait.Wait()
	fmt.Println("最终结果", count)
}
  