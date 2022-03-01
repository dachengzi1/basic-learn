package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//json 必须首字母大写，不然输出{}
type Persion struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}



func main() {
	p := Persion{"张三", 1}

	//结构体转json 普通

	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("struct format json:", err)
		return
	}
	fmt.Println("struct format json:", string(b))

	//结构体转json  格式化输出
	b1, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		fmt.Println("format json err:", err)
		return
	}
	fmt.Println("format json:", string(b1))

	//json字符串 转 struct

	b3 := []byte(`{"name":"张思","age":2}`)
	var p4 Persion
	err = json.Unmarshal(b3, &p4)

	if err != nil {
		fmt.Println("json str format struct err: ", err)
	}
	fmt.Println("json str format struct: ", p4, reflect.TypeOf(p4))

	//json字符串 转 map

	b2 := []byte(`{"name":"张思","age":2}`)
	var m map[string]interface{}
	err = json.Unmarshal(b2, &m)

	if err != nil {
		fmt.Println(" json str format map err:", err)
	}
	fmt.Println(" json str format map :", m, reflect.TypeOf(m))

	//struct format map
	b4 := []byte(`{"name":"张思","age":2}`)
	var m2 map[string]interface{}
	err = json.Unmarshal(b4, &m2)

	if err != nil {
		fmt.Println(" struct format map err:", err)
	}
	fmt.Println(" struct format map :", m2, reflect.TypeOf(m2))



	//
	//m := i.(map[string]interface{})
	//fmt.Printf("m si :%T\n", m)
	//
	//for k, v := range m {
	//	switch vv := v.(type) {
	//	case string:
	//		fmt.Printf("field type: %s, value: %s\n", vv, m[k])
	//	default:
	//		fmt.Printf("field other:")
	//
	//	}
	//}
}
