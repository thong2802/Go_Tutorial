package main

import "fmt"

func main() {
	// Declaring variables
	var myString string = "Hello"
	var myInt int = 100
	var myFloat float64 = 45.65

	fmt.Println(myString, myInt, myFloat)

	// ================================
	// multiple declaring variable
	var (
		employeeId          int    = 1
		firstName, lastName string = "Satoshi", "Nakomoto"
	)

	fmt.Println(employeeId, firstName, lastName)

	// ==================================
	// Short variable declaring syntax
	name := "Thong"
	age, salary, isProgramming := "22", 100000.0, true

	fmt.Println(name, age, salary, isProgramming)
}
