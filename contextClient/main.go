package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type ResponseData struct {
	Resp *http.Response
	Err  error
}

func parseBody(by []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	fmt.Println(string(by))
	err := json.Unmarshal(by, &result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func doCall(ctx context.Context) {
	transport := http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: &transport,
	}
	respChan := make(chan *ResponseData, 1)
	req, err := http.NewRequest("GET", "http://localhost:8080/?name=张三&age=3", nil)
	if err != nil {
		fmt.Println("req err:", err)
		return
	}
	req = req.WithContext(ctx) // 使用带超时的ctx创建一个新的client request

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		resp, err := client.Do(req)
		fmt.Printf("client.do resp:%v, err:%v\n", resp, err)

		respChan <- &ResponseData{
			Resp: resp,
			Err:  err,
		}
		wg.Done()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("call api timeout")
	case result := <-respChan:
		fmt.Println("call server api success")
		if result.Err != nil {
			fmt.Println("req result Err:", result.Err)
			return
		}
		defer result.Resp.Body.Close()
		data, _ := ioutil.ReadAll(result.Resp.Body)
		json, err := parseBody(data)
		if err != nil {
			fmt.Println("un marshal Err:", err)
			return
		}
		fmt.Printf("resp:%v\n", json)

	}

}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	doCall(ctx)
}
