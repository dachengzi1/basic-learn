package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exit bool // 定义全局变量跳出 子协程

func work() {
	fmt.Println("work...")
	for exit {
		break
	}
	wg.Done()
}

func main() {
	wg.Add(1)

	go work()

	time.Sleep(time.Second * 3)

	exit = true

	wg.Wait()

	fmt.Println("exit...")
}
