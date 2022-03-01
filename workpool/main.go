package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Job struct {
	Id   int
	Rnum int
}

type Result struct {
	Job *Job
	Sum int
}

func createPool(n int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < n; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			for job := range jobChan {
				num := job.Rnum
				var sum = 0
				for num != 0 {
					temp := num % 10
					sum += temp
					num /= 10
				}
				resultChan <- &Result{
					Job: job,
					Sum: sum,
				}
			}
		}(jobChan, resultChan)
	}
}
func main() {
	job := make(chan *Job, 10)
	result := make(chan *Result, 10)

	createPool(10, job, result)

	go func(result chan *Result) {
		for r := range result {
			fmt.Printf("job id:%d, rand num: %d, sum is: %d \n", r.Job.Id,r.Job.Rnum, r.Sum)
		}
	}(result)

	for i := 0; i < 10; i++ {
		job <- &Job{
			Id:   i,
			Rnum: rand.Intn(100),
		}
	}
	time.Sleep(time.Second)

}
