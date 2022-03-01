package main

import "fmt"

func main() {

	//1，for range  能够感知通道的关闭

	// for select

	ch := make(chan int, 10)

	for i := 0; i < 10; i++{
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:

		}
	}

}
