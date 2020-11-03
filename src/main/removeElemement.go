package main

import "fmt"

// no := 26、27
func removeNodes(nums []int, target int) []int {
	for i := 0; i < len(nums); {
		if nums[i] == target {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return nums
}

func removeNodesTest() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(removeNodes(nums, 3))
}
