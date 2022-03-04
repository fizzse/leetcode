package main

import "fmt"

/*
485. 最大连续 1 的个数
给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。



示例 1：

输入：nums = [1,1,0,1,1,1]
输出：3
解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.
示例 2:

输入：nums = [1,0,1,1,0,1]
输出：2
*/

func findMaxConsecutiveOnes(nums []int) int {
	maxCnt := 0

	sum := 0
	for _, v := range nums {
		if v == 1 {
			sum++

			if sum > maxCnt {
				maxCnt = sum
			}
		} else {
			sum = 0
		}
	}

	return maxCnt
}

func main() {
	src := []int{1, 0, 1, 1, 0, 0}
	fmt.Println(findMaxConsecutiveOnes(src))
}
