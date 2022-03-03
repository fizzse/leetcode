package main

import "fmt"

/*
434. 字符串中的单词数
统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。

请注意，你可以假定字符串里不包括任何不可打印的字符。

示例:

输入: "Hello, my name is John"
输出: 5
解释: 这里的单词是指连续的不是空格的字符，所以 "Hello," 算作 1 个单词。
*/

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}

	m := byte(' ')
	s = s + " " // 末尾加一个空格 减少判断条件

	size := len(s)
	cnt := 0

	for i := 0; i < size-1; i++ {
		if s[i] != m && s[i+1] == m {
			cnt++
		}
	}

	return cnt
}

func main() {
	in := ", , , ,        a, eaefa"
	fmt.Println(countSegments(in))
}
