package main

import (
	"fmt"
	"strings"
)

/*
https://leetcode-cn.com/problems/repeated-substring-pattern/
459. 重复的子字符串
给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。



示例 1:

输入: s = "abab"
输出: true
解释: 可由子串 "ab" 重复两次构成。
示例 2:

输入: s = "aba"
输出: false
示例 3:

输入: s = "abcabcabcabc"
输出: true
解释: 可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)
*/

// s = n(sub) // n >= 2
// 2s = 2n(sub)
// tmp =2n(sub)[1:-1] (掐头去尾破坏两个sub)
// tmp = 2(n-1)sub
// n >=2  tmp container s

func repeatedSubstringPattern(s string) bool {
	tmp := s + s
	tmp = tmp[1 : len(tmp)-1]

	return strings.Contains(tmp, s)
}

func main() {
	a := "abcabcabc"
	fmt.Println(repeatedSubstringPattern(a))
}
