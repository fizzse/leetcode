package main

import (
	"log"
	"sort"
)

/*
https://leetcode.cn/problems/3sum/
15. 三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

输入：nums = []
输出：[]
示例 3：

输入：nums = [0]
输出：[]
*/

/*
 思路：	1.将三数之和 转换成两数之和
		2.先排序，后面方便去重
*/

func towSum(nums []int, target int, pos int, numMap map[int]int, filter map[[3]int]bool) [][]int {
	var rets [][]int
	for i, n := range nums {
		index, ok := numMap[target-n]
		if !ok || index <= pos+i {
			continue
		}

		// 按照 大小排序
		ret := [3]int{0 - target, n, target - n}
		if filter[ret] {
			continue
		}

		filter[ret] = true
		rets = append(rets, ret[:])
	}

	return rets
}

func threeSum1(nums []int, target int) [][]int {
	sort.Ints(nums)

	size := len(nums)
	filter := make(map[[3]int]bool)
	var retList [][]int
	var rets [][]int

	numMap := make(map[int]int)
	for i, num := range nums {
		numMap[num] = i
	}

	for i := 0; i < size-2; i++ {
		if nums[i] > target {
			break
		}

		target := target - nums[i]
		rets = towSum(nums[i+1:], target, i+1, numMap, filter)
		retList = append(retList, rets...)
	}
	return retList
}

func threeSum(nums []int) [][]int {
	return threeSum1(nums, 0)
}

func main() {
	in := []int{-1, 0, 1, 2, -1, -4}

	rets := threeSum(in)
	for _, ret := range rets {
		log.Println(ret)
	}
}
