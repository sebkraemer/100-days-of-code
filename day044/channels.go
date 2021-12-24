package main

import (
	"fmt"
	"time"
)

func runWorkers(workers []func() int) int {
	done := make(chan struct{})
	result := make(chan int)

	if len(workers) < 1 {
		return 0 // func will deadlock without workers if we don't return here
	}

	for i, worker := range workers {
		fmt.Println("spawning worker", i)
		go func(worker func() int) {
			select {
			case result <- worker():
			case <-done:
			}
		}(worker)
	}
	res := <-result
	close(done) // todo: what is the difference here to leaving it out?

	return res
}

func main() {
	workers := make([]func() int, 10)

	for i := range workers {
		i := i // beware of unwanted closure capturing of i!
		fmt.Println(i)
		workers[i] = func() int {
			duration := time.Duration((1+i)*300) * time.Millisecond
			fmt.Println("worker with duration", duration)
			time.Sleep(duration)
			return i
		}
	}

	start := time.Now()
	res := runWorkers(workers)
	stop := time.Now()

	fmt.Println("elapsed:", (stop.Sub(start)))

	// 0 will be printed since the function will always return after the shorted worker with index 0
	fmt.Println(res)
}
