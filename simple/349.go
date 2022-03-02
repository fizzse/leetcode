package main

/*
https://leetcode-cn.com/problems/intersection-of-two-arrays/
349. 两个数组的交集
给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。



示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
示例 2：

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
解释：[4,9] 也是可通过的
*/

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	for _, num := range nums1 {
		m1[num]++
	}

	for _, num := range nums2 {
		if m1[num] > 0 {
			m2[num]++
		}
	}

	ret := make([]int, 0, len(m2))
	for k := range m2 {
		ret = append(ret, k)
	}

	return ret
}
