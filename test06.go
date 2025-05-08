package main

// waitGroup 是一个计数信号量，用于等待一组 goroutine 完成
import (
	"fmt"
	"sync"
)

func Test06() {
	var wait = sync.WaitGroup{}
	var mainWait = sync.WaitGroup{}
	mainWait.Add(10)

	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func(i int) {
			defer wait.Done()
			defer mainWait.Done()
			fmt.Println("hello world", i)
		}(i)
		wait.Wait() // 等待所有协程完成，使得协程可以按照顺序执行，但是会导致运行速率变慢
	}

	mainWait.Wait()
	fmt.Println("main end")
}
