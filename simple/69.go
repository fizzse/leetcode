package main

/*

实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。
*/

// 穷举法 谁都会

/*func mySqrt(x int) int {
	for i := 1; i < x; i++ {
		num := i * i
		if num == x {
			return i
		}

		if num > x {
			return i - 1
		}
	}

	return 0
}
*/

// 二分查找
func mySqrt(x int) int {
	if x == 1 {
		return x
	}

	min, max := 0, x
	for max-min > 1 {
		m := (max + min) / 2
		if x/m < m {
			max = m
		} else {
			min = m
		}
	}

	return min
}
