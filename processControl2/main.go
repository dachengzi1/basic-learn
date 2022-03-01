package main

import (
	"fmt"
	"time"
)

// for range 能够感知channel 关闭
//协程处理一个通道的时候，优先使用for range
func main() {
	inChan := make(chan int, 10)
	go func() {
		for ch := range inChan {
			fmt.Println(ch)
		}
	}()

	for i := 0; i < 10; i++ {
		inChan <- i
	}
	close(inChan)
	time.Sleep(time.Second * 10)
}
