package main

import (
	"fmt"
	"time"
)


func write(ch chan string) {
	time.Sleep(time.Second * 10)
	ch <- "C"
	fmt.Println("内部调用")
}

func read(ch chan string) {
	fmt.Println("读取channel开始")
	 str:= <- ch
	 fmt.Println("读取channel:", str)
}
func main() {
	ch:= make(chan string, 1)
	fmt.Println("外部调用")
	go write(ch)
	read(ch)

	fmt.Println("外部调用")

}
