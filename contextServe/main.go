package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)


type User struct {
	Name string
	Age  string
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(2)
	query := r.URL.Query()
	user := &User{
		Name: query.Get("name"),
		Age:  query.Get("age"),
	}
	json, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json Marshal err:", err)
		return
	}
	fmt.Printf("query", string(json))

	if num == 0 {
		time.Sleep(time.Microsecond * 10) // 模拟慢请求10毫秒
		//fmt.Fprintf(w, "slow response", string(json))
		w.Write(json)
		return
	}
	//fmt.Fprintf(w, "quick response", string(json))
	w.Write(json)

}
func main() {
	http.HandleFunc("/", handleFunc)
	fmt.Println("serve running")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("serve err:", r)
		}
	}()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
