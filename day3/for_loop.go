package main

import "fmt"

func result(i, j, k int) (sum int) {
	sum = i + j + k
	return
}

func try_switch() {

	for i := 1; i < 4; i++ {
		switch i {
		case 1:
			fmt.Println("This is case 1")
		case 2:
			fmt.Println("This is case 2")
		default:
			fmt.Println("This is default case")
		}
	}

}

func try_defer() {
	name := "Smart"
	name2 := "Zop"

	defer fmt.Println(name)
	fmt.Println(name2)
}

func main() {
	//for x := 1; x < 10; x++ {
	//	fmt.Println(x)
	//}

	//var i int = 1
	//for i < 5 {
	//
	//	fmt.Println(i)
	//	i += i
	//

	//fmt.Println(result(1, 2, 3))

	try_switch()

	try_defer()
}
