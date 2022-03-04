package main

import "fmt"

/*
520. 检测大写字母
我们定义，在以下情况时，单词的大写用法是正确的：

全部字母都是大写，比如 "USA" 。
单词中所有字母都不是大写，比如 "leetcode" 。
如果单词不只含有一个字母，只有首字母大写， 比如 "Google" 。
给你一个字符串 word 。如果大写用法正确，返回 true ；否则，返回 false 。



示例 1：

输入：word = "USA"
输出：true
示例 2：

输入：word = "FlaG"
输出：false
*/

// 大写字母在65 -> 90
func detectCapitalUse(word string) bool {
	size := len(word)
	cnt := 0
	index := 0
	for i := 0; i < size; i++ {
		c := word[i]
		if c <= 90 {
			index = i
			cnt++
		}
	}

	if cnt == size {
		return true
	}

	if cnt == 1 && index == 0 {
		return true
	}

	if cnt == 0 {
		return true
	}

	return false
}

func main() {
	src := "abc"
	fmt.Println(detectCapitalUse(src))
}
