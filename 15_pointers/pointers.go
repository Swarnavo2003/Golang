package main

import "fmt"

// pass by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum",num)
// }

// pass by reference
// func changeNum(num *int) {
// 	*num = 5
// 	fmt.Println("In changeNum", *num)
// }

func changeSlice(nums *[]int) {
	*nums = append(*nums, 4)
}

func main() {
	// num := 1
	// changeNum(num)

	// fmt.Println("Memory address", &num)
	// changeNum(&num)
	// fmt.Println("After changeNum",num)

	var nums = []int{1,2,3}
	fmt.Println("Before change", nums)
	changeSlice(&nums)
	fmt.Println("After change", nums)
}