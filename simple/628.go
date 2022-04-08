package main

import "sort"

/*
628. 三个数的最大乘积
给你一个整型数组 nums ，在数组中找出由三个数组成的最大乘积，并输出这个乘积。

示例 1：

输入：nums = [1,2,3]
输出：6
示例 2：

输入：nums = [1,2,3,4]
输出：24
示例 3：

输入：nums = [-1,-2,-3]
输出：-6
*/

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	size := len(nums)

	// 全是正数或者全是负数 就是最大的三个数的乘积
	// 最小的两个负数*最大正数 与最大的三个正数做比较
	n1 := nums[0] * nums[1] * nums[size-1]
	n2 := nums[size-1] * nums[size-2] * nums[size-3]
	if n1 < n2 {
		return n2
	}

	return n1
}
