package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

//json 必须首字母大写，不然输出{}
type Persion struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func JSONMethod(content interface{}) map[string]interface{} {
	var name map[string]interface{}
	if marshalContent, err := json.Marshal(content); err != nil {
		fmt.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber() // 设置将float64转为一个number
		if err := d.Decode(&name); err != nil {
			fmt.Println(err)
		} else {
			for k, v := range name {
				name[k] = v
			}
		}
	}
	return name
}
func main() {
	p := Persion{"张三", 1}
	json1 := JSONMethod(p)
	fmt.Println("json method: ", json1)

	p2 := Persion{"王五", 30}
	// 编码结果暂存到 buffer
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(&p2)
	if err == nil {
		fmt.Print("json.NewEncoder 编码结果: ", b.String())
	}

	var str = `{"name": "上三", "age": 1}`
	var p3 Persion
	err = json.NewDecoder(strings.NewReader(str)).Decode(p3)
	if err == nil {
		fmt.Print("json.NewDecoder 编码结果: ", p3.Name, p3.Age)
	}

	str1:= JSONMethod(p)
	fmt.Print("json.NewDecoder 编码结果: ",str1)



}
