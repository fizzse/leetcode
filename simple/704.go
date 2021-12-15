package main

/*
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。


示例 1:

输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例 2:

输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
https://leetcode-cn.com/problems/binary-search/
*/

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	mid := 0

	if end == -1 {
		return -1
	}

	for start <= end {
		mid = (start + end) / 2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			start = mid + 1
		case nums[mid] > target:
			end = mid - 1
		}

	}

	return -1
}
