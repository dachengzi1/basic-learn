package main

import (
	"fmt"
	"reflect"
)

type aaa int32

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println(t.Name(), t.Kind())
}

func reflectValue(x interface{}) {
	t := reflect.ValueOf(x)
	fmt.Println(t.Interface())
}

type persion struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func main() {
	var i int32
	var n rune
	var f float64 = 399
	var r aaa
	//int32 int32
	//int32 int32
	//float64 float64
	//aaa int32

	reflectType(i)
	reflectType(n)
	reflectType(f)
	reflectType(r)

	reflectValue(i)

	p := persion{
		Name: "zhang",
		Age:  2,
	}

	t := reflect.TypeOf(p)

	fmt.Println(t.Name(), t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		//fmt.Println(field.Name, field.Type, field.Tag.Get("json"))
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))

	}

	if info, ok := t.FieldByName("Age"); ok {
		fmt.Println(info)

	}

}
