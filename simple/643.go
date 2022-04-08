package main

import "fmt"

/*
https://leetcode-cn.com/problems/maximum-average-subarray-i/
643. 子数组最大平均数 I
给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。

请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。

任何误差小于 10-5 的答案都将被视为正确答案。

示例 1：

输入：nums = [1,12,-5,-6,50,3], k = 4
输出：12.75
解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
示例 2：

输入：nums = [5], k = 1
输出：5.00000
*/

// 暴力破解
func findMaxAverage(nums []int, k int) float64 {
	size := len(nums)

	//start := 0
	max := 0
	for i := 0; i < k; i++ {
		max += nums[i]
	}

	tmp := max
	for i := 1; i < size-k+1; i++ {
		tmp += nums[i+k-1]
		tmp -= nums[i-1] // 避免重复计算 只需要加一个数和减一个数就行

		if tmp > max {
			//start = i
			max = tmp
		}
	}

	//_ = start
	return float64(max) / float64(k)
}

// 8 6
// 8+0 6+5
//
func main() {
	in := []int{8, 0, 1, 7, 8, 6, 5, 5, 6, 7}

	o := findMaxAverage(in, 2)
	fmt.Println(o)
}
