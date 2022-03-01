package main

import "fmt"

//func rec(ch chan int) {
//	ret := <-ch
//	fmt.Println("rec:", ret)
//}

//1.关闭的通道不能再接受数据，否则panic
//2,一个关闭并且没有值的通道，读出的数据是该通道类型的零值
//3，一个关闭的通道进行取值，会一直读完
//4。对关闭的通道关闭会panic
func main() {
	//var ch = make(chan int)//无缓存通道，必须有接收方
	//go rec(ch)
	//ch <- 2
	//time.Sleep(time.Second)

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		fmt.Println("ch1 close...")
		close(ch2)
	}()

	for c := range ch2 {
		fmt.Println(c)
	}
}
