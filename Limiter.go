package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(500 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("done")

	burustyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burustyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(500 * time.Millisecond) {
			burustyLimiter <- t
		}
	}()

	burustyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burustyRequests <- i
	}
	close(burustyRequests)

	for req := range burustyRequests {
		<-burustyLimiter
		fmt.Println("request", req, time.Now())
	}

}
