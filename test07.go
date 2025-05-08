package main

// 此程序展示了select的使用方法，使用select来监听多个channel的消息
// 这个程序的主要功能是使用select语句来监听多个channel的消息，并在接收到消息时打印出来
import (
	"fmt"
	"time"
)

func Test07() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)
	defer func() {
		close(chA)
		close(chB)
		close(chC)
	}()

	l := make(chan struct{})
	// 这是一个优雅控制协程退出的方式，使用一个信号量来控制主协程的退出
	go Send(chA)
	go Send(chB)
	go Send(chC)

	go func() {
	Loop:
		for {
			select {
			case n, ok := <-chA:
				fmt.Println("A", n, ok)
			case n, ok := <-chB:
				fmt.Println("B", n, ok)
			case n, ok := <-chC:
				fmt.Println("C", n, ok)
			case <-time.After(time.Second): // 设置1秒的超时时间
				break Loop // 退出循环
			}
		}
		l <- struct{}{} // 告诉主协程可以退出了
	}()

	<-l
}

func Send(ch chan<- int) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond)
		ch <- i
	}
}
