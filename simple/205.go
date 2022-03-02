package main

import "fmt"

/*
https://leetcode-cn.com/problems/isomorphic-strings/
205. 同构字符串
给定两个字符串 s 和 t ，判断它们是否是同构的。

如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。

每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

示例 1:

输入：s = "egg", t = "add"
输出：true
示例 2：

输入：s = "foo", t = "bar"
输出：false
示例 3：

输入：s = "paper", t = "title"
输出：true
*/

func isIsomorphic(s string, t string) bool {
	s1 := len(s)
	s2 := len(t)
	if s1 != s2 {
		return false
	}

	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)

	for i := 0; i < s1; i++ {
		c1 := s[i]
		c2 := t[i]

		if m1[c1] == 0 && m2[c2] == 0 {
			m1[c1] = c2
			m2[c2] = c1
			continue
		}
		if m1[c1] != c2 || m2[c2] != c1 {
			return false
		}
	}

	return true
}

func main() {
	s := "add"
	d := "caa"
	fmt.Println(isIsomorphic(s, d))
}
