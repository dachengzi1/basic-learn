package main

import (
	"fmt"
)

func subPrint(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
}
func print(args ...int) {
	subPrint(args...)
}

func main() {
	print(1, 2, 34, 5, 6)
}
