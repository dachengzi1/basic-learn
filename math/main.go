package main

import (
	"fmt"
	"math/rand"
)

func sumN(n int) {
	r_num := rand.Intn(n)
	fmt.Printf("rand int: %d\n", r_num)
	sum := 0
	for r_num != 0 {
		temp := r_num % 10
		sum += temp
		r_num /= 10
	}
	fmt.Printf("sum: %d\n", sum)
}

func sumRange(n, min int) {
	r_num := rand.Intn(n)
	fmt.Printf("rand int: %d\n", r_num)
	sum := 0
	for r_num != 0 {
		temp := r_num % 10
		temp += min
		sum += temp
		r_num /= 10
	}
	fmt.Printf("sum: %d", sum)
}

// 数字位数相加
func main() {
	sumN(100)
	sumRange(20, 10)
}
