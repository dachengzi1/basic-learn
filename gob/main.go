package main

import (
	"encoding/json"
	"fmt"
)

// A struct with a mix of fields, used for the GOB example.
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user := User{
		Name: "张三",
		Age:  1,
	}

	//结构体转json
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Println("struct transform json err:", err)
	}
	fmt.Printf("user json: %v, type: %T\n", string(u), string(u))

	//json转结构体
	var user2 User
	err = json.Unmarshal(u, &user2)

	if err != nil {
		fmt.Println("json transform struct err:", err)
	}
	fmt.Printf("user2 json: %v, type: %T\n", user2, user2)

	//字符串转map

	str := `{"name":"张三","age":1}`
	var user3 map[string]interface{}

	err = json.Unmarshal([]byte(str), &user3)

	if err != nil {
		fmt.Println("string transform json err:", err)
	}
	fmt.Printf("user3 json: %v, type: %T\n", user3, user3)


	//map 转 字符串



	//var b bytes.Buffer
	//enc := gob.NewEncoder(&b)
	//err = enc.Encode(user)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("Outer complexData byte: %T, %v， %v, %s\n", b, b, b.Bytes(), string(b.Bytes()))
	//
	//dec := gob.NewDecoder(&b)
	//var data User
	//err = dec.Decode(&data)
	//if err != nil {
	//	fmt.Println("Error decoding GOB data:", err)
	//	return
	//}

}
