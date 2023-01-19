package main

import "fmt"

func channel3(c chan<- int) {
	c <- 10
	c <- 20

}

func main() {
	c := make(chan int)

	go channel3(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
