package main

import "fmt"

func main() {
	var dayOfWeek = 2
	switch dayOfWeek {
	case 2:
		if dayOfWeek == 2 {
			goto handleCaseNumberEqual2
		}
	handleCaseNumberEqual2:
		fmt.Println("handle case 2")
		fmt.Println("Monday")
		//	fallthrough | continue only using for loop
	case 3:
		fmt.Println("Tuesday")
		//	fallthrough
		break
	case 4:
		fmt.Println("Wednesday")
		//	fallthrough
		break
	case 5:
		fmt.Println("Thursday")
		//	fallthrough
		break
	case 6:
		fmt.Println("Friday")
		//	fallthrough
		break
	case 7:
		fmt.Println("Saturday")
		fmt.Println("Weekend. Yaay!")
		//	fallthrough
		break
	case 8:
		fmt.Println("Sunday")
		fmt.Println("Weekend. Yaay!")
		//	fallthrough
		break
	default:
		fmt.Println("Invalid Day")
	}
}
