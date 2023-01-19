package main

import (
	"fmt"
	"time"
)

func hello(ch chan bool) {
	fmt.Println("Hello Sleep")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Hello Again")
	ch <- true
}

func main() {
	ch := make(chan bool)
	fmt.Println("Main flow")
	go hello(ch)
	<-ch
	fmt.Println("Bye")
}
