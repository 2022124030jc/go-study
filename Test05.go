package main

import (
	"fmt"
	"sync"

)

func Test05() {
	var wait = sync.WaitGroup{}
	wait.Add(1)
	go func() {
		defer wait.Done()
		fmt.Println("hello world")
	}()
	wait.Wait()
	fmt.Println("main end")
}