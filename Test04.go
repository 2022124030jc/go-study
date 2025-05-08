package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Test04() {
	// 创建一个定时器，每秒触发一次
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// 创建一个 10 秒后触发的计时器
	done := time.After(10 * time.Second)

	// 使用当前时间作为随机数种子
	rand.Seed(time.Now().UnixNano())

	for {
		select {
		case <-ticker.C:
			// 生成 0-100 之间的随机数
			num := rand.Intn(100)
			fmt.Printf("随机数: %d\n", num)
		case <-done:
			fmt.Println("程序结束")
			return
		}
	}
}
