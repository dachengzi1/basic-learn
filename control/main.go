package main

import (
	"fmt"
	"time"
)

func run() {

	for {
		var i = 0
		for {
			if i > 5 {
				fmt.Println("i == 5")
				continue
			}
			fmt.Println("i:", i)
			i++
			//select {
			//case :
			//
			//}
		}
	}
}
func main() {
	run()
	time.Sleep(time.Second * 100)
}
