package main

import (
	"math/rand"
	"time"
)

func main() {
	//jobChan := make(chan int, 126)
	resultChan := make(chan int, 126)
	defer close(resultChan)

	//go func(result chan int) {
	//for r := range resultChan {
	//	fmt.Println("result:", r)
	//}
	//}(resultChan)

	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		resultChan <- num
	}
	time.Sleep(time.Second * 3)
}
