package main

import "fmt"

func add1() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	for i := 0; i < 10; i++ {
		s := add1()
		fmt.Println(s(i))
	}
}
