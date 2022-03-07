package main

import "fmt"

func main() {

	var a []int
	if a == nil {
		fmt.Println("a == nil")
	} else {
		fmt.Println(a, len(a), cap(a))
	}

	var a2 = []int{}

	if a2 == nil {
		fmt.Println("a2 == nil")
	} else {
		fmt.Println(a2, len(a2), cap(a2))
	}

	var a3 = make([]int, 5, 10)

	if a3 == nil {
		fmt.Println("a3 == nil")
	} else {
		fmt.Println(a3, len(a3), cap(a3))

	}

	for i := 0; i < 10; i++ {
		a3 = append(a3, i)
	}
	a3[0] = 11
	fmt.Println(a3)

	a4 := []int{1, 2, 3}

	n := copy(a4, a3)

	fmt.Println("copy n:", n, a4)
}
