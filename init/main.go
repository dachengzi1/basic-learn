package main

import (
	"fmt"
	"math"
)

var MIN = 0.00001

func isEqual(ff1, ff2 float64) bool {
	if ff1 > ff2 {
		return math.Dim(ff1, ff2) < MIN
	} else {
		return math.Dim(ff2, ff1) < MIN
	}
}
func main() {
	//var a *int = new(int)
	//var b float32 = 5.0
	//*a = b
	//fmt.Println(*a)

	//var a float32 = 0.0
	//fmt.Println(a / 0)
	//
	//var f float32 = 0.1
	//fmt.Println(f / 0)
	//
	//var f1 float32 = -0.1
	//fmt.Println(f1 / 0)
	//
	//f2 := 0.9
	//f3 := 1.0
	//
	//fmt.Println(isEqual(f2, f3))

	//var x int = 1
	//x,x = 1,2
	//
	//fmt.Println(x)

	var x = [3]int{}
	var a = 0

	a, x[a] = 1,2
	fmt.Println(x)
	fmt.Println(a)

}
