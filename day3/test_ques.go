package main

import "fmt"

func calculate(units float64) float64 {
	var bill float64
	if units >= 0 && units <= 50 {
		bill += bill * 0.5
	} else if bill > 50 && bill <= 100 {
		bill += 50 * 0.5
		units -= 50
		bill += bill * 0.75
	} else if bill > 100 && bill <= 150 {
		bill += 50*0.5 + 50*0.75
		units -= 100
		bill += units * 1.
	}

	bill = bill * 1.2
	return bill
}

func main() {
	fmt.Println(calculate(150))
}
