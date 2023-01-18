package main

import (
	"fmt"
	"sort"
)

//type Person struct {
//	name string
//	age  int
//}

//type People struct {
//	Person
//}

func main() {
	ppl := []{
		Person{
			"Pranav",
			21,
		},
		{"Xavier",
			19,
		},
		{"ieurf",
			27,
		},
		{"oief",
			23,
		},
	}
	fmt.Println(ppl)
	sort.Slice(ppl, func(p, q int) bool {
		return ppl[p].age < ppl[q].age
	})

	fmt.Println(ppl)
}
