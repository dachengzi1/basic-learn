package main

import "fmt"

func _for() {

	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			if i == 4 && j == 10 { //跳出当前代码块
				break
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}

	}
	fmt.Println("for end...")
}
func main() {

	_for()
}
