package main

import "fmt"

func fibi(x int) int {
	if x == 0 || x == 1 {
		//fmt.Println(x)
		return x
	} else {
		return fibi(x-1) + fibi(x-2)

	}
}
func main() {
	x := 10
	for i := 0; i < x; i++ {
		fmt.Println(fibi(i))
		fmt.Println("Done")
	}
}
