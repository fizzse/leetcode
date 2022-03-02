package main

/*
3. 无重复字符的最长子串
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。



示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

import (
	"fmt"
)

func lengthOfLongestSubstring1(s string) int {
	size := len(s)
	if size < 2 {
		return size
	}

	max := 0

	for i := 0; i < size; i++ {
		if size-i < max {
			return max
		}
		m := make(map[byte]int)
		for j := i; j < size; j++ {
			char := s[j]
			m[char]++
			if m[char] > 1 {
				break
			}
		}

		subSize := len(m)
		if subSize > max {
			max = subSize
		}
	}

	return max
}

func lengthOfLongestSubstring(s string) int {
	min := func(n1, n2 int) int {
		if n1 < n2 {
			return n1
		}
		return n2
	}

	max := func(n1, n2 int) int {
		if n1 < n2 {
			return n2
		}
		return n1
	}

	size := len(s)
	if size < 2 {
		return size
	}

	m := make(map[byte]int)
	m[s[0]] = -1 // 默认值0

	rlt := 1
	cnt := 1
	for i := 1; i < size; i++ {
		cnt++
		char := s[i]
		if m[char] != len(m)-1 { // 之前没有出现过
			cnt = min(cnt, i-m[char]+1)
		}
		rlt = max(cnt, rlt)
		m[char] = i - 1
	}

	return rlt
}

func main() {
	l := lengthOfLongestSubstring("au")
	fmt.Println(l)
}
