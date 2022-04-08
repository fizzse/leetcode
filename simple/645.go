package main

/*
https://leetcode-cn.com/problems/set-mismatch/
645. 错误的集合
集合 s 包含从 1 到 n 的整数。不幸的是，因为数据错误，导致集合里面某一个数字复制了成了集合里面的另外一个数字的值，导致集合 丢失了一个数字 并且 有一个数字重复 。

给定一个数组 nums 代表了集合 S 发生错误后的结果。

请你找出重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。

示例 1：

输入：nums = [1,2,2,4]
输出：[2,3]
示例 2：

输入：nums = [1,1]
输出：[1,2]
*/

func findErrorNums(nums []int) []int {
	m := make(map[int]int)

	for _, n := range nums {
		m[n]++
	}

	empty := 0
	dump := 0
	for i := 1; i < len(nums)+1; i++ {
		switch m[i] {
		case 0:
			empty = i

		case 1:
			if dump != 0 && empty != 0 {
				return []int{dump, empty}
			}

		case 2:
			dump = i
		}
	}

	return []int{dump, empty}
}
