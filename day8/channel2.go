package main

import (
	"fmt"
	"time"
)

func f(c chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Println(1)
	}
	c <- 10
	fmt.Println("this")
	<-c
	fmt.Println("this2")

}

func main() {
	c := make(chan int, 1)
	//c <- 10
	go f(c)
	fmt.Println("NO deadlock")
	time.Sleep(1000 * time.Millisecond)
}
