package main

import (
	"fmt"
	"reflect"
)

type Persion struct {
	Name string `json: "name"`
	Age  string `json: "age"`
}

func main() {
	var fl float64

	ty := reflect.TypeOf(fl)
	fmt.Println(ty, ty.Kind())

	value := reflect.ValueOf(fl)

	fmt.Println(value)

	p := Persion{
		Name: "ss",
		Age:  "1",
	}
	ty2 := reflect.TypeOf(p)
	val2 := reflect.ValueOf(p)
	fmt.Println(ty2, val2)

	for i := 0; i < ty2.NumField(); i++ {
		t := ty2.Field(i).Name
		fmt.Printf("field is: %s\n,", t)
		v := val2.Field(i).Interface()
		fmt.Printf("field value is: %s\n,", v)

	}

}
