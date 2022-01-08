package main

import (
	"fmt"
	"time"
)

func scz(ch chan<- int) {
	for i := 0; i <= 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
		fmt.Printf("proudcing data: %d\n", i)
	}
	defer close(ch)
}

func xfz(ch <-chan int) {
	for k := range ch {
		fmt.Printf("get data: %d\n", k)
	}
}

func main() {
	ch := make(chan int, 10)
	go scz(ch)
	xfz(ch)
}
