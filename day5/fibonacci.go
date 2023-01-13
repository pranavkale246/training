package main

import "fmt"

func fib() func() {
	f1 := 0
	f2 := 1
	f3 := 0

	return func() {
		fmt.Println(f1)
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
}

func main() {
	f := fib()

	for i := 0; i < 10; i++ {
		f()
	}
}
