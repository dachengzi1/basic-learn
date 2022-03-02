package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"reflect"
	"strconv"
	"time"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:       "localhost:6379",
		Password:   "",
		DB:         0,
		PoolSize:   3,
		MaxRetries: 3,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis client success!")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err, reflect.TypeOf(err))
		}
	}()
	test2()
	ch := make(chan int)
	<-ch
}

func test2() {
	var (
		intervalTime = time.Minute
		egmentNum    = int64(6)
		limitNum     = int64(5)
		waitTime     = 30
	)
	time.Sleep(time.Duration(waitTime) * time.Second)
	for i := 0; i < 10; i++ {
		//fmt.Printf("time range from %d to %d\n", i*10+waitTime, (i+1)*10+waitTime)
		for j := 0; j < 8; j++ {
			res := limit("test2", intervalTime, egmentNum, limitNum)
			fmt.Printf("time range from %d to %d, res: %v\n", i*10+waitTime, (i+1)*10+waitTime, res)

		}
		time.Sleep(10 * time.Second)
	}
}

func limit(str string, interval time.Duration, segmentNum int64, limitNum int64) bool {

	segmentNumInterval := interval.Seconds() / float64(segmentNum)

	tick := float64(time.Now().Unix()) / segmentNumInterval

	key := fmt.Sprintf("%s_%d_%d_%d_%f", str, interval, segmentNum, limitNum, tick)

	fmt.Printf("segmentNumInterval: %v, tick: %v, time now: %d, key: %v\n", segmentNumInterval, tick, time.Now().Unix(), key)

	_, err := client.SetNX(key, 0, interval).Result()
	if err != nil {
		panic(err)
	}
	quantum, err := client.Incr(key).Result()
	if err != nil {
		panic(err)
	}

	for i := segmentNumInterval; i < interval.Seconds(); i += segmentNumInterval {
		//tick -= 1
		//preKey := fmt.Sprintf("%s_%d_%d_%d_%f", str, interval, segmentNum, limitNum, tick)
		//
		//fmt.Printf("preKey: %v\n", preKey)
		//
		//val, err := client.Get(preKey).Result()
		//if err != nil {
		//	val = "0"
		//}
		//num, err := strconv.ParseInt(val, 0, 64)
		//if err != nil {
		//	panic(err.Error())
		//}
		//num = num + quantum
		//
		//if num > limitNum {
		//	return false
		//}
		tick = tick - 1
		preKey := fmt.Sprintf("%s_%d_%d_%d_%f", str, interval, segmentNum, limitNum, tick)
		fmt.Printf("preKey: %v\n", preKey)
		val, err := client.Get(preKey).Result()
		if err != nil {
			val = "0"
		}
		num, err := strconv.ParseInt(val, 0, 64)
		quantum = quantum + num
		if quantum > limitNum {
			client.Decr(key).Result()
			return false
		}
	}
	return true

}
