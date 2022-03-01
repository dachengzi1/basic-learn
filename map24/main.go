package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {

	//初始化map

	obj := make(map[string]int, 20)
	obj["张三"] = 20
	obj["张二"] = 28
	fmt.Println("obj:", obj)

	obj2 := map[string]int{
		"张三": 30,
		"里斯": 29,
	}
	fmt.Println("obj2:", obj2)

	//删除
	//delete(obj2, "张三")
	//fmt.Println("obj2:", obj2)

	//判断某个键是否存在

	v, ok := obj2["张三"]
	if ok {
		fmt.Println("find name 张三", v)
	} else {
		fmt.Println("not found 张三")
	}

	//遍历
	for key, value := range obj2 {
		fmt.Printf("name: %s, value:%d\n", key, value)
	}

	//key := rand.Intn(100)//  生成0~99的随机整数
	//fmt.Println(key)

	obj3 := make(map[string]int, 100)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%d", i)
		value := rand.Intn(100)
		obj3[key] = value
	}

	var keys = make([]string, 200)
	for key, _ := range obj3 {
		//fmt.Printf("obj3, name: %s, value:%d\n", key, value)
		keys = append(keys, key)
	}

	sort.Strings(keys) // 排序
	for _, v := range keys {
		fmt.Printf("name: %s, value:%d\n", v, obj3[v])
	}

	//元素为map类型的切片
	arr := make([]map[string]int, 1)

	obj4 := map[string]int{
		"张三": 1,
	}
	arr = append(arr, obj4)
	fmt.Println(arr)

	//值为切片类型的map

	obj5 := make(map[string][]int, 10)
	arr2 := make([]int, 10)
	arr2[0] = 11
	obj5["里斯"] = arr2
	fmt.Println("obj5:", obj5)

	//写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。

	str := "how do you do"

	obj6 := make(map[string]int, 10)

	arr3 := strings.Split(str, " ")

	for _, v := range arr3 {
		fmt.Println("string:", v)
		_, ok := obj6[v]
		if ok {
			obj6[v]++
		} else {
			obj6[v] = 1
		}
	}

	for key, value := range obj6 {
		fmt.Printf("%s:%d\n", key, value)

	}

	//观察下面代码，写出最终的打印结果。
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])




}
