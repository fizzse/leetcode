package main

/*

https://leetcode-cn.com/problems/happy-number/
202. 快乐数
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」 定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

*/

func isHappy(n int) bool {
	m := make(map[int]bool) // 如果数字重复出现 说明陷入无限循环了

	for {
		m[n] = true
		n = dealNum(n)
		if n == 1 {
			return true
		}

		if m[n] {
			return false
		}
	}
}

func dealNum(n int) int {
	tmp := n

	ret := 0
	for tmp != 0 {
		n1 := tmp % 10
		ret += n1 * n1
		tmp /= 10
	}

	return ret
}
