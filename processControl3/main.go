package main

import (
	"fmt"
	"time"
)

// for select ,无法感知channel关闭
//使用 _,ok := <-channel, 当ok = false （channel关闭） 时，直接return
//
//某个通道关闭时，不再处理该通道，而继续处理其他通道， 可以使用 channel = nil（select 不会再nil的通道上等待）

func main() {
	inChan := make(chan int, 10)

	t := time.NewTicker(time.Millisecond * 500)
	go func() {
		for {
			select {
			case x, ok := <-inChan:
				if !ok {
					fmt.Println("exit .")
					return
				}
				fmt.Println(x)
			case <-t.C:
				fmt.Println("Working .")
			}

		}
	}()

	for i := 0; i < 10; i++ {
		inChan <- i
	}
	close(inChan)
	time.Sleep(time.Second * 2)
}
