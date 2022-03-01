package main

import (
	"fmt"
	"time"
)

//发送协程关闭通道，接受协程不关闭通道， 最佳实践：把接收方通道入参声明为只读，如果接收方关闭只读协程，编译时会报错
//主协程显式关闭通道stopCh可以主动处理通知多个work协程退出场景
func worker(stopCh <-chan int) {
	t := time.NewTicker(time.Millisecond * 500)
	go func() {
		defer fmt.Println("worker exit")
		// Using stop channel explicit exit
		for {
			select {
			case <-stopCh:
				fmt.Println("Recv stop signal:")
				return
			case <-t.C:
				fmt.Println("Working .")
			}
		}
	}()
	return
}
func main() {
	inChan := make(chan int, 10)

	worker(inChan)
	time.Sleep(time.Second * 2)

	//for i := 0; i < 10; i++ {
	//	inChan <- i
	//}
	close(inChan)
	//time.Sleep(time.Second * 10)
}
