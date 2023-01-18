package main

import "fmt"

type regular struct {
	string
	int
}

type employee struct {
	name     string
	salary   int
	currency string
}

type intern struct {
	name   string
	salary int
}

func display(e *employee) {
	fmt.Printf("Function called , Name is:%v\n", e.name)
}
func (r regular) display() {
	fmt.Printf("Name: %v\n", r.string)
}
func (i intern) display() {
	fmt.Printf("Intern Name: %v\n", i.name)
}

func (e employee) display() {
	e.name = "Pranav Kale"
	fmt.Printf("After: Name is %v, Salary = %v%v\n ", e.name, e.currency, e.salary)
}

func main() {
	e := employee{
		"pranav",
		300000,
		"$",
	}

	i := intern{
		"Pranav Kale",
		30000,
	}

	r := regular{
		"Pranav Kale",
		4500000,
	}

	fmt.Printf("Before: Name is %v, Salary = %v%v\n ", e.name, e.currency, e.salary)
	e.display()
	fmt.Printf("After After: Name is %v, Salary = %v%v\n ", e.name, e.currency, e.salary)
	i.display()
	r.display()

	display(&e)
}
