package main

import (
	"fmt"
)

// slice -> dynamic
// most used construct in go
// + useful methods
func main() {
	// uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)
	// capacity -> maximum number of elements can fit
	// fmt.Println(cap(nums))
	
	// fmt.Println(nums)
	// fmt.Println(nums == nil)

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	// nums := []int{}
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	// var nums = make([]int, 2, 5)
	// nums[0] = 3
	// nums[1] = 5
	// fmt.Println(nums);
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	// copy function
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	// slice operator
	// var nums = []int{1, 2, 3, 4, 5}
	// fmt.Println(nums[0:1])
	// fmt.Println(nums[:2])
	// fmt.Println(nums[1:])

	// slices package
	// var nums1 = []int{1,2,3}
	// var nums2 = []int{1,2,4}
	// fmt.Println(slices.Equal(nums1, nums2))

	// 2D slices
	var nums = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(nums)
}