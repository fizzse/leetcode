package main

import "fmt"

/*
796. 旋转字符串
给定两个字符串, s 和 goal。如果在若干次旋转操作之后，s 能变成 goal ，那么返回 true 。

s 的 旋转操作 就是将 s 最左边的字符移动到最右边。

例如, 若 s = 'abcde'，在旋转一次之后结果就是'bcdea' 。


示例 1:

输入: s = "abcde", goal = "cdeab"
输出: true
示例 2:

输入: s = "abcde", goal = "abced"
输出: false
*/

func rotateString(s string, goal string) bool {
	if s == goal {
		return true
	}

	s1 := len(s)
	s2 := len(goal)
	if s1 != s2 {
		return false
	}

	b1 := []byte(s)
	b2 := []byte(goal)
	for i := 0; i < s1; i++ {
		move(b1)
		if equal(b1, b2) {
			return true
		}
	}

	return false
}

func move(b1 []byte) {
	c := b1[0]
	size := len(b1)
	for i := 0; i < size-1; i++ {
		b1[i] = b1[i+1]
	}

	b1[size-1] = c
	return
}

func equal(b1, b2 []byte) bool {
	for i, c := range b1 {
		if b2[i] != c {
			return false
		}
	}

	return true
}

func main() {
	a := []byte{'a', 'b', 'c'}
	move(a)

	for _, c := range a {
		fmt.Printf("%c\t", c)
	}
}
