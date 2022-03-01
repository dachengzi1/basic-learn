package main

import (
	"fmt"
)

func worker(i int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", i, j)
		results <- j * j
		//time.Sleep(time.Second)
	}
}
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//开启三个go

	//for w := 0; w < 3; w++ {
	//	go worker(w, jobs, results)
	//}
	//
	//for j := 0; j < 100; j++ {
	//	jobs <- j
	//}
	//close(jobs)
	//
	////接受者没有可以接受数据，所以导致堵塞，从而主线程死锁
	////for r := range results {
	////	fmt.Println(r)
	////}
	//
	//go func() {
	//	for r := range results {
	//		fmt.Println(r)
	//	}
	//}()
	//time.Sleep(time.Second)

	go func() {
		for i := 0; i < 100; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	go func() {
		for {
			i, ok := <-jobs
			if !ok {
				break
			}
			results <- i
		}
		close(results)
	}()

	for r := range jobs {
		fmt.Println(r)
	}

}
