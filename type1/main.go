package main

import "fmt"

func main()  {

	var a interface{}
	a = []int{1,2,3}
	value, ok := a.([]int)
	if !ok {
		fmt.Println("It's not ok for type string")
		return
	}
	fmt.Println("The value is ", value)
}
