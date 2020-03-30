package main

import "fmt"

func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	l := 1
	dp := make([]int, length+1)
	dp[l] = nums[0]
	for i := 1; i < length; i++ {
		if dp[l] < nums[i] {
			l++
			dp[l] = nums[i]
		} else {
			// 更新len的值 直到匹配到小于 nums[i]
			pos := mid(dp, l, nums[i])
			dp[pos] = nums[i]
		}
	}
	return l
}

func mid(d []int, l int, num int) int {
	left := 1
	right := l
	for left <= right {
		mid := (right + left) / 2
		if d[mid] < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left

}

func main() {
	fmt.Println(lengthOfLIS([]int{
		3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12,
	}))
}
