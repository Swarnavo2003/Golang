package main

import "fmt"

// const age = 21

// name := "golang" // not allowed to declare like this outside the main function
// var name string = "golang" // allowed

func main() {
	// const name string = "golang"
	// const name = "golang"

	// const age int = 30
	// const age = 30

	// fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}