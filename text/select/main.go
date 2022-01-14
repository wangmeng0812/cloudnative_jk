package main

import (
	"fmt"
	"time"
)

func a(c chan<- int)  {
	for i := 0; i<10 ; i++{
		c <- i
	}
	defer close(c)
}

func b(c <-chan int,str string)  {
	b := str
	for i := range c{
		fmt.Printf("this is %v.%v\n",b ,i)
	}
}



func main() {
	timer := time.NewTimer(time.Second)
	aaa := make(chan int,10)
	bbb := make(chan int,10)
	go a(aaa)
	go a(bbb)
	for {
		select {
		case <-aaa:
			b(aaa,"aaa")
		case <-bbb:
			b(bbb,"bbb")
		case <-timer.C:
			fmt.Println("this is timer.c")
			return
		}
	}
}




