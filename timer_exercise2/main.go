package main

import (
	"fmt"
	"time"
)

type Response struct {
	Id  int
	Res chan string
}

func worker(res *Response) {

	time.Sleep(2 * time.Second)
	fmt.Println("res channel...")
	res.Res <- "success"

}
func main() {

	res := &Response{
		Id:  0,
		Res: make(chan string, 1),
	}
	go worker(res)

	t := time.NewTicker(3 * time.Second)
	defer func() {
		close(res.Res)
		t.Stop()
	}()
	for {
		select {
		case <-t.C:
			fmt.Println("time out...")
			return
		case res, ok := <-res.Res:
			if !ok {
				fmt.Println("channel close...")
				break
			}
			return
			fmt.Println("req response:", res)

		}
	}
	fmt.Println("end ...")
	time.Sleep(10 * time.Second)
}
