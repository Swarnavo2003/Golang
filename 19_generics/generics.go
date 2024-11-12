package main

import "fmt"

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// generics
// func printSlice[T int | string| bool](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// comparable -> all types
func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

// LIFO
type stack[T any] struct {
	elements []T
}

func main() {
	// nums := []int{1,2,3}
	// printSlice(nums)

	// names := []string{"golang","typescript"}
	// printSlice(names)

	vals := []bool{true, false, true}
	printSlice(vals, "john")

	// myStack := stack[string]{
	// 	elements: []string{"golang","typescript"},
	// }
	// fmt.Println(myStack)


}