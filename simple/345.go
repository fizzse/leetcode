package main

import "fmt"

/*
https://leetcode-cn.com/problems/reverse-vowels-of-a-string/submissions/
345. 反转字符串中的元音字母
给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。

元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。



示例 1：

输入：s = "hello"
输出："holle"
示例 2：

输入：s = "leetcode"
输出："leotcede"
*/

func reverseVowels(s string) string {
	size := len(s)
	if size < 2 {
		return s
	}

	buf := make([]byte, size)
	m := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	i := 0
	j := size - 1
	for i <= j {
		c1 := s[i]
		c2 := s[j]
		if !m[c1] {
			buf[i] = c1
			i++
			continue
		}

		if !m[c2] {
			buf[j] = c2
			j--
			continue
		}

		buf[i] = s[j]
		buf[j] = s[i]
		i++
		j--
	}

	return string(buf)
}

func main() {
	in := "leetcode"
	out := reverseVowels(in)
	fmt.Println(out)
}
