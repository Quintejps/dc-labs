package main

import (
	"fmt"
	"time"
)

func main() {
	x := make(chan time.Time)
	cola := x

	for Counter := 1; ; Counter++ {
		go readCola(cola, Counter)
		x <- time.Now()
		oldCola := cola
		cola = make(chan time.Time)
		go connect(oldCola, cola)
	}
}

func readCola(cola chan time.Time, Counter int) {
	start := <-cola
	end := time.Now()
	fmt.Println(Counter, end.Sub(start))
}

func connect(src, dst chan time.Time) {
	for t := range src {
		dst <- t
	}
}

