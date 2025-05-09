package main

// 无缓冲chan与容量为1的chan的区别：无缓冲chan会双向阻塞，发送方与接收方必须一一配对

import (
	"fmt"
	"time"

)

func Test13() {
	finish := make(chan struct{},1)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2秒后结束")
		finish <- struct{}{}
		// 如果是无缓冲的话这里会有一个阻塞，但是容量为1的chan不会阻塞
		fmt.Println("协程结束")
	}()
	time.Sleep(4 * time.Second)
	<-finish
	time.Sleep(2 * time.Second)
	fmt.Println("主协程结束")
}
