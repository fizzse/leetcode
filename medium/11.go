package main

/*

11. 盛最多水的容器
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。



示例 1：



输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

*/

// 暴力破解
func maxArea1(height []int) int {
	size := len(height)
	max := 0
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			area := (j - i) * (getMin(height[i], height[j]))
			max = getMax(max, area)
		}
	}

	return max
}

// O(n)复杂度 area = 底*高 值小的指针向内移动，这样就减小了搜索空间 因为面积取决于指针的距离与值小的值乘积，
// 如果值大的值向内移动，距离一定减小，而求面积的另外一个乘数一定小于等于值小的值，因此面积一定减小，而我们要求最大的面积，因此值大的指针不动，而值小的指针向内移动遍历
func maxArea(height []int) int {
	size := len(height)
	if size < 2 {
		return -1
	}
	start := 0
	end := size - 1
	max := 0
	for start < end {
		area := (end - start) * getMin(height[start], height[end])
		max = getMax(max, area)
		if height[start] > height[end] {
			end--
		} else {
			start++
		}
	}

	return max
}

func getMin(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}

	return n2
}

func getMax(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
