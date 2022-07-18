package main

/*
https://leetcode.cn/problems/4sum/
18. 四数之和
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：

0 <= a, b, c, d < n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。



示例 1：

输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
示例 2：

输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]
*/

/*
 思路同三数之和，简化思路
*/

import (
	"log"
	"sort"
)

func towSum1(nums []int, target int, pos int, numMap map[int]int) [][4]int {
	var rets [][4]int
	for i, n := range nums {
		index, ok := numMap[target-n]
		if !ok || index <= pos+i {
			continue
		}

		// 按照 大小排序
		ret := [4]int{0, 0, n, target - n}
		rets = append(rets, ret)
	}
	return rets
}

func threeSum2(nums []int, target int, pos int, numMap map[int]int) [][4]int {
	sort.Ints(nums)

	size := len(nums)
	var retList [][4]int

	for i := 0; i < size-2; i++ {
		target2 := target - nums[i]
		rets := towSum1(nums[i+1:], target2, pos+i+1, numMap)
		for _, ret := range rets {
			ret[1] = nums[i]
			retList = append(retList, ret)
		}
	}
	return retList
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	size := len(nums)
	filter := make(map[[4]int]bool)
	var retList [][]int

	numMap := make(map[int]int)
	for i, num := range nums {
		numMap[num] = i
	}

	for i := 0; i < size-3; i++ {
		target2 := target - nums[i]
		rets := threeSum2(nums[i+1:], target2, i+1, numMap)
		for _, ret := range rets {
			ret[0] = nums[i]
			if filter[ret] {
				continue
			}

			ret := ret
			filter[ret] = true
			retList = append(retList, ret[:4])
		}
	}

	return retList
}

func main() {
	in := []int{1, -2, -5, -4, -3, 3, 3, 5}
	target := -11
	rets := fourSum(in, target)
	for _, ret := range rets {
		log.Println(ret)
	}
}
