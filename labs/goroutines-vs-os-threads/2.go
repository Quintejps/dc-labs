package main

import (
	"fmt"
	"time"
)

func main() {
	ping, pong := make(chan int), make(chan int)
	var count int64

	t := time.NewTimer(1 * time.Minute)
	done := make(chan struct{})
	shutdown := make(chan struct{})

	go func() {
	loop:
		for {
			select {
			case <-shutdown:
				break loop
			case v := <-ping:
				count++
				pong <- v
			}
		}
		done <- struct{}{}
	}()

	go func() {
	loop:
		for {
			select {
			case <-shutdown:
				break loop
			case v := <-pong:
				ping <- v
			}
		}
		done <- struct{}{}
	}()

	ping <- 1

	<-t.C
	close(shutdown)

	select {
	case <-ping:
	case <-pong:
	}

	<-done
	<-done
	t.Stop()
	fmt.Println("How many communications per second can the program sustain? ", count)
}
