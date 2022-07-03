package main

import "fmt"

func main() {
	switch dayOfWeek := 5; dayOfWeek {
	case 1, 2, 3, 4, 5:
		fmt.Println("on the weekday")
	case 6, 7:
		fmt.Println("on the weekend")
	default:
		fmt.Println("Invalid day")
	}

}
