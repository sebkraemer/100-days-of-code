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

	// https://stackoverflow.com/a/36613932/4312669
	// it's not strictly necessary to close a channel, it can be garbage collected either way if it is not referenced any more
	// depending on the situation, other parts of the program might depend on the close() behavior
	close(done)

	return res
}

func main() {
	workers := make([]func() int, 10)

	for i := range workers {
		i := i // beware of unwanted closure capturing of i!
		workers[i] = func() int {
			duration := time.Duration((1+i)*300) * time.Millisecond
			fmt.Println("worker with duration", duration)
			start := time.Now()
			for {
				stop := time.Now()
				d := stop.Sub(start)
				fmt.Println(i, d)
				if d >= duration {
					break
				}
				time.Sleep(time.Duration(100) * time.Millisecond)
			}
			fmt.Println("exiting", i)
			return i
		}
	}

	start := time.Now()
	res := runWorkers(workers)
	stop := time.Now()

	fmt.Println("elapsed:", (stop.Sub(start)))

	// 0 will be printed since the function will always return after the shorted worker with index 0
	fmt.Println(res)

	// even though runWorkers returned the goroutines spawned from it are still running;
	// without the sleep, they were just not visible and 'killed' without apparent
	// problem when quitting the program
	// I assume that e.g. a cancellation context would be needed to, well, cancel the workers prematurely.
	fmt.Println("waiting for other routines to finish (?)")
	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println("finished")
}
