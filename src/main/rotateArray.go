package main

import "fmt"

// no := 189
func rotate(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}

func reverse(arr []int) {
	for i := 0; i < len(arr)>>1; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

func rotateTest() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate(nums, 3)
	fmt.Println(nums)
}
