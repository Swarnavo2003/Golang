package main

import (
	"fmt"
)

func main() {
	// simple switch
	// i := 3
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	// multiple condition switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's weekend")
	// default:
	// 	fmt.Println("It's workday")
	// }

	// type switch
	whoAmI := func (i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("its and integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("other")
		}
	}

	whoAmI("golang")
	whoAmI(50)
	whoAmI(true)	
	whoAmI(55.90)	
}