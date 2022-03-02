package main

/*
https://leetcode-cn.com/problems/longest-palindrome/
409. 最长回文串
给定一个包含大写字母和小写字母的字符串 s ，返回 通过这些字母构造成的 最长的回文串 。

在构造过程中，请注意 区分大小写 。比如 "Aa" 不能当做一个回文字符串。
示例 1:

输入:s = "abccccdd"
输出:7
解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
示例 2:

输入:s = "a"
输入:1
示例 3:

输入:s = "bb"
输入: 2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestPalindrome(s string) int {
	m := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	sum := 0
	for _, v := range m {
		sum += v - (v & 1) // 都用偶数的
	}

	l := len(s)
	if sum < l { // 说明有字符出现了单数次
		return sum + 1 // 将单数放中间
	}

	return l
}
