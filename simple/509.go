package main

import "fmt"

/*
https://leetcode-cn.com/problems/fibonacci-number/
509. 斐波那契数
斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

F(0) = 0，F(1) = 1
F(n) = F(n - 1) + F(n - 2)，其中 n > 1
给定 n ，请计算 F(n) 。

示例 1：

输入：n = 2
输出：1
解释：F(2) = F(1) + F(0) = 1 + 0 = 1
示例 2：

输入：n = 3
输出：2
解释：F(3) = F(2) + F(1) = 1 + 1 = 2
示例 3：

输入：n = 4
输出：3
解释：F(4) = F(3) + F(2) = 2 + 1 = 3
*/

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func fibByLoop(n int) int {
	if n < 2 {
		return n
	}

	n1 := 0
	n2 := 1
	for i := 1; i < n; i++ {
		n2 += n1
		n1 = n2 - n1
	}

	return n2
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fib(i), fibByLoop(i))
	}
}
