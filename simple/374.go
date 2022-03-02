package main

import "fmt"

/*
https://leetcode-cn.com/problems/guess-number-higher-or-lower/
74. 猜数字大小
猜数字游戏的规则如下：

每轮游戏，我都会从 1 到 n 随机选择一个数字。 请你猜选出的是哪个数字。
如果你猜错了，我会告诉你，你猜测的数字比我选出的数字是大了还是小了。
你可以通过调用一个预先定义好的接口 int guess(int num) 来获取猜测结果，返回值一共有 3 种可能的情况（-1，1 或 0）：

-1：我选出的数字比你猜的数字小 pick < num
1：我选出的数字比你猜的数字大 pick > num
0：我选出的数字和你猜的数字一样。恭喜！你猜对了！pick == num
返回我选出的数字。



示例 1：

输入：n = 10, pick = 6
输出：6
示例 2：

输入：n = 1, pick = 1
输出：1
示例 3：

输入：n = 2, pick = 1
输出：1
示例 4：

输入：n = 2, pick = 2
输出：2
*/

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int {
	k := 6
	if num == k {
		return 0
	}

	if num > k {
		return -1
	}

	return 1
}

func guessNumber(n int) int {
	start := 1
	end := n
	for {
		if start == end-1 {
			if guess(start) == 0 {
				return start
			}

			return end
		}

		num := (start + end) / 2
		ret := guess(num)
		if ret == 0 {
			return num
		}

		if ret < 0 {
			end = num
			continue
		}

		start = num
	}
}

func main() {
	o := guessNumber(10)
	fmt.Println(o)
}
