package main

import "sort"

/*
https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
350. 两个数组的交集 II
给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。



示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
示例 2:

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]


提示：

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000

*/

func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	for _, num := range nums1 {
		m1[num]++
	}
	for _, num := range nums2 {
		m2[num]++
	}

	var ret []int
	for k, v := range m1 {
		v2 := m2[k]
		if v2 == 0 {
			continue
		}

		if v > v2 {
			v = v2
		}

		for i := 0; i < v; i++ {
			ret = append(ret, k)
		}
	}

	return ret
}

func intersectWithSort(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2) // 如果已经有序的 优化方案 双指针

	s1 := len(nums1)
	s2 := len(nums2)
	i := 0
	j := 0

	var ret []int
	for i < s1 && j < s2 {
		if nums1[i] == nums2[j] {
			ret = append(ret, nums1[i])
			i++
			j++
			continue
		}

		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return ret
}
