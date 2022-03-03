package main

/*
https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/
453. 最小操作次数使数组元素相等
给你一个长度为 n 的整数数组，每次操作将会使 n - 1 个元素增加 1 。返回让数组所有元素相等的最小操作次数。



示例 1：

输入：nums = [1,2,3]
输出：3
解释：
只需要3次操作（注意每次操作会增加两个元素的值）：
[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
示例 2：

输入：nums = [1,1,1]
输出：0
*/

/*
题解:
	每一次 n+1数字+1 等于有一个数字-1 直到所有数字等于最小数字
*/

func minMoves(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	cnt := 0
	for _, num := range nums {
		cnt += num - min
	}

	return cnt
}
