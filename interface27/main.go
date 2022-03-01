package main

import "fmt"

type dog struct{}

type cat struct{}

type sayer interface {
	say()
}

func (d dog) say() {
	fmt.Println("wang wang")
}

func (d cat) say() {
	fmt.Println("miao miao")
}

func do(s sayer) {
	s.say()
}
func main() {

	var s sayer
	s = dog{}
	do(s)
	s = cat{}
	do(s)

	var x interface{}
	x = make([]int, 2)
	v, ok := x.(string)
	if ok {
		fmt.Printf("string: %s", v)
	} else {
		fmt.Println("not string")
	}

	switch v := x.(type) {
	case []int:
		fmt.Printf("[]int: %v", v)
	case int:
		fmt.Printf("int: %d", v)
	case bool:
		fmt.Printf("bool: %b", v)

	default:
		fmt.Printf("not type")
	}

}
