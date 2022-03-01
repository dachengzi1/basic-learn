package main

import (
	"fmt"
	"sync"
	"time"
)

//通过通道 跳出子协程

var wg sync.WaitGroup

func work(stopChan chan struct{}) {
LOOP:
	for {
		fmt.Println("work...")
		select {
		case <-stopChan:
			break LOOP
		default:

		}
	}
	wg.Done()
}

func main() {

	stopChan := make(chan struct{})
	wg.Add(1)

	go work(stopChan)

	time.Sleep(time.Second * 1)
	stopChan <- struct{}{}

	close(stopChan)
	wg.Wait()
	fmt.Println("exit...")

}
