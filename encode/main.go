package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

//编码 将 struct format json string
//json.Marshal
//json.NewEncoder

//解码  将 json string format struct

//json.Unmarshal
//json.NewDecoder

type Persion struct {
	Name string
	Age  int
}

func main() {

	// json.Marshal
	p := Persion{Name: "aaa", Age: 2}
	b, err := json.MarshalIndent(&p, "", "  ")

	if err != nil {
		fmt.Println("json.Marshal err:", err)
	}
	fmt.Println("json.Marshal:", string(b), reflect.TypeOf(string(b)))

	// json.NewEncoder
	p2 := Persion{Name: "bbb", Age: 3}
	var b2 bytes.Buffer
	err = json.NewEncoder(&b2).Encode(&p2)

	if err != nil {
		fmt.Println("json.Marshal err:", err)
	}
	fmt.Println("json.Marshal:", string(b2.Bytes()), reflect.TypeOf(string(b2.Bytes())))

	// json.Unmarshal
	b3 := []byte(`{"Name":"ccc","Age":3}`)
	var p4 Persion
	err = json.Unmarshal(b3, &p4)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
	}
	fmt.Println("json.Unmarshal:", p4, reflect.TypeOf(p4))

	// json.NewEncoder
	var p5 Persion
	var str = `{"name":"赵六","age":28}`
	err = json.NewDecoder(strings.NewReader(str)).Decode(&p5)

	if err != nil {
		fmt.Println("json.Marshal err:", err)
	}
	fmt.Println("json.Marshal:", p5, reflect.TypeOf(p5))
}
