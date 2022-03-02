package main

/*

https://leetcode-cn.com/problems/contains-duplicate-ii/
219. 存在重复元素 II
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。



示例 1：

输入：nums = [1,2,3,1], k = 3
输出：true
示例 2：

输入：nums = [1,0,1,1], k = 1
输出：true
示例 3：

输入：nums = [1,2,3,1,2,3], k = 2
输出：false
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	m1 := make(map[int]int)

	for i, n := range nums {
		if i0, ok := m1[n]; ok { // 比较下标的差值
			if i-i0 <= k {
				return true
			}
		}
		m1[n] = i // 如果不存在 设置下标。或者间隔超过了k 设置为后一个相同数据的下标
	}

	return false
}
