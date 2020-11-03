package main

import (
	"fmt"
	"sort"
)

// 计算两个数的交集 no := 350
func intersect(nums1 []int, nums2 []int) []int {

	m0 := map[int]int{}
	for _, v := range nums1 {

		//遍历nums1，value作为key，重复+1
		m0[v] += 1
	}
	k := 0
	for _, v := range nums2 {

		//如果元素相同，将其存入nums2中，并将出现次数减1
		if m0[v] > 0 {
			m0[v] -= 1
			nums2[k] = v
			k++
		}
	}
	fmt.Println(nums2[0:k])
	return nums2[0:k]
}

// 归并排序
func intersect1(nums1 []int, nums2 []int) []int {
	i, j, k := 0, 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}

func getNumTest() {

	a := []int{0, 1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	intersect(a, b)

}
