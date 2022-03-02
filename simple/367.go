package main

/*
https://leetcode-cn.com/problems/valid-perfect-square/
367. 有效的完全平方数
给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。

进阶：不要 使用任何内置的库函数，如  sqrt 。



示例 1：

输入：num = 16
输出：true
示例 2：

输入：num = 14
输出：false
*/

func isPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}

	if num < 2 {
		return true
	}

	for i := 2; i <= num/2; i++ {
		n := i * i
		if n == num {
			return true
		}

		if n > num { // 提前跳出循环
			return false
		}
	}

	return false
}
