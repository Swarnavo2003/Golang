package main

import "fmt"

// iterating over data structures
func main() {
	// nums := []int{6,7,8}

	// for i := 0; i<len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// sum := 0
	// for _, num := range nums {
	// 	sum += num
	// }
	// fmt.Println(sum)

	// for i,num := range nums {
	// 	fmt.Println(num, i)
	// }

	// m := map[string]string{"fname":"john","lname":"doe"}
	// for k,v := range m {
	// 	fmt.Println(k,v)
	// }

	// c -> unicode code point rune
	// i -> starting byte of rune
	for i,c := range "golang" {
		fmt.Println(i,string(c))
	}
}