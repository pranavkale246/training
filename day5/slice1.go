package main

import "fmt"

func main() {
	//var slice []int
	//
	//slice = append(slice, 0)
	//slice = append(slice, 1)
	//slice = append(slice, 2, 3, 4, 5, 6, 7)
	//slice = append(slice, 8, 9, 20, 11)
	//slice = append(slice, 8, 9, 20, 11)
	//slice = append(slice, 8, 9, 20, 11)
	//
	//fmt.Println(slice)
	//fmt.Println("Length is ", len(slice))
	//fmt.Println("Capacity is ", cap(slice))

	var a []int = make([]int, 5)
	fmt.Println(a)
	s := []int{2, 3, 4}
	fmt.Println(s)
	var slice []int = make([]int, 1, 5)
	slice[0] = 1
	slice = append(slice, 2)
	slice = append(slice, s...)
	fmt.Println(slice)
}
