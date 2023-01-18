package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type People struct {
	s []Person
}

func (ppl *People) sortPeople() {
	sort.Slice(ppl.s, func(p, q int) bool {
		//return ppl[p]. < ppl[q].age
		return ppl.s[p].age < ppl.s[q].age
	})
}

func main() {
	p := []Person{}
	p = append(p, Person{"Pranav", 21}, Person{"Kale", 31}, Person{"uerier", 25})
	ppl := People{p}

	fmt.Println(ppl)
	ppl.sortPeople()
	fmt.Println(ppl)
}
