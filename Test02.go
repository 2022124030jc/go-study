package main

import (
	"fmt"
	"time"
)

func PrintHello() {
	time.Sleep(2 * time.Second)
	fmt.Println("hello")
}

func Solve() {
	fmt.Println("start")
	go PrintHello()
	time.Sleep(1 * time.Second)
	fmt.Println("end")
}

func Test02() {
	Solve()
	time.Sleep(3 * time.Second) // 等待 goroutine 完成
	fmt.Println("main")
}
