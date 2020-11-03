package main

import "fmt"

// leet code no := 122
func bestTicket(nums []int) int {

	if len(nums) < 2 {
		return 0
	}
	j := 0
	for i := 0; i <= len(nums)-1; i++ {
		if i == len(nums)-1 {
			break
		}
		if nums[i+1] > nums[i] {
			j += nums[i+1] - nums[i]
		}
	}
	return j
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bestTicketTest() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 1, 3, 1, 4, 9, 5, 6, 1}
	fmt.Println(maxProfit(nums))
}
