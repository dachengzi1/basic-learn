package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"reflect"
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

func limit(str string, interval time.Duration, limitNum int64) bool {
	tick := time.Now().Unix() / int64(interval.Seconds())

	fmt.Printf("interval Seconds: %v, tick: %d, time now: %d\n", interval.Seconds(), tick, time.Now().Unix())

	key := fmt.Sprint("%s:%d:%d", str, interval, limitNum, tick)

	_, err := client.SetNX(key, 0, interval).Result()
	if err != nil {
		panic(err)
	}
	quantum, err := client.Incr(key).Result()
	if err != nil {
		panic(err)
	}
	if quantum > limitNum {
		return false
	}
	return true
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err, reflect.TypeOf(err))
		}
	}()
	test2()

	var ch = make(chan int64) //阻塞
	<-ch

}

func test2() {
	fillInteval := time.Minute

	var limitNum int64 = 5

	waitTime := 30

	fmt.Printf("time range from 0 to %d\n", waitTime)

	time.Sleep(time.Duration(waitTime) * time.Second)

	for i := 0; i < 10; i++ {

		fmt.Printf("time range from %d to %d\n", i*10+waitTime, (i+1)*10+waitTime)

		rs := limit("test2", fillInteval, limitNum)

		fmt.Println("result is:", rs)

		time.Sleep(10 * time.Second)
	}
}

func test() {
	for i := 0; i < 100; i++ {
		go func() {
			rs := limit("test1", 1*time.Second, 5)
			fmt.Println("result is:", rs)
		}()
	}
}
