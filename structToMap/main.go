package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

//json 必须首字母大写，不然输出{}
type Persion struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func toMap(data interface{}) (map[string]interface{}, error) {
	var mp map[string]interface{}

	if b, err := json.Marshal(data); err != nil {
		return nil, err
	} else {
		d := json.NewDecoder(bytes.NewReader(b))
		d.UseNumber() //设置将float64转为一个number
		if err = d.Decode(&mp); err != nil {
			return nil, err
		}

		for k, v := range mp {
			mp[k] = v
		}
	}
	return mp, nil
}
func reflecttoMap(data interface{}) map[string]interface{} {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	var mp2 = make(map[string]interface{}) // map：= map[string]interface{} ， map == nil 无法赋值

	for i := 0; i < t.NumField(); i++ {
		mp2[t.Field(i).Name] = v.Field(i).Interface()
	}
	return mp2
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("toMap err: ", err.(string))
		}
	}()
	p := Persion{Name: "xiixix", Age: 20}
	m, err := toMap(p)
	if err != nil {
		//fmt.Println("toMap:", err)
		panic(err.Error())
	}
	fmt.Println("toMap m:", m)

	m2 := reflecttoMap(p)
	fmt.Println("reflect toMap m2:", m2)
}
