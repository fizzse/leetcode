package main

import (
	"fmt"
	"sort"
)

/*
https://leetcode-cn.com/problems/array-of-doubled-pairs/
954. 二倍数对数组
给定一个长度为偶数的整数数组 arr，只有对 arr 进行重组后可以满足 “对于每个 0 <= i < len(arr) / 2，都有 arr[2 * i + 1] = 2 * arr[2 * i]” 时，返回 true；否则，返回 false。

示例 1：

输入：arr = [3,1,3,6]
输出：false
示例 2：

输入：arr = [2,1,2,6]
输出：false
示例 3：

输入：arr = [4,-2,2,-4]
输出：true
解释：可以用 [-2,-4] 和 [2,4] 这两组组成 [-2,-4,2,4] 或是 [2,4,-2,-4]
*/

// 思路 暴力破解
// 对数组排序 依次遍历 讲2*n 与 n 都-1
// 如果cnt(n)==0 说明作为倍数被抵消完了
// 遍历所有的cnt，都为0 说明符合条件
func canReorderDoubled(arr []int) bool {
	m := make(map[int]int)

	for _, n := range arr {
		m[n]++
	}

	// 遍历m
	sort.Ints(arr) // 主要是排序占用时间
	for _, n := range arr {
		n2 := n * 2
		if m[n] == 0 {
			continue
		}

		if m[n2] != 0 {
			m[n]--
			m[n2]--
		}
	}

	for n, cnt := range m {
		if cnt != 0 {
			fmt.Println(n, cnt)
			return false
		}
	}
	return true
}

func main() {
	in := []int{2, 4, 0, 0, 8, 1}
	o := canReorderDoubled(in)
	print(o)
}
