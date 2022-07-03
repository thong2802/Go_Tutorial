package main

import "fmt"

func main() {
	//Type Inference
	var name = " Thong Nguyen" // type Inference is Optional here
	fmt.Printf("Variable 'name' is of type:  %T \n", name)

	// =============================
	// Multiple variable declaring with Inference Types
	var firstName, lastName, age, salary = "Thong", "Nguyen", 22, 10.000
	fmt.Printf("Variable 'firstName is of type: %T \n, Variable 'lastName is of type %T \n, Variable 'age is of type %T \n,  Variable 'salary is of type %T \n |",
		firstName, lastName, age, salary)

}
