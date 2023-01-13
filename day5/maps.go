package main

import "fmt"

func main() {

	mp := make(map[int]int)
	mp[3] = 3
	mp[4] = 4

	var m = make(map[int]int)
	m[0] = 0
	m[1] = 1

	v, ok := m[1]
	fmt.Println(v, ok)

	fmt.Println(m[1])
	fmt.Println(mp)
}
