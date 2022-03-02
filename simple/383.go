package main

/*
https://leetcode-cn.com/problems/ransom-note/
383. 赎金信
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。



示例 1：

输入：ransomNote = "a", magazine = "b"
输出：false
示例 2：

输入：ransomNote = "aa", magazine = "ab"
输出：false
示例 3：

输入：ransomNote = "aa", magazine = "aab"
输出：true
*/

func canConstruct(ransomNote string, magazine string) bool {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for i := 0; i < len(ransomNote); i++ {
		m1[ransomNote[i]]++
	}

	for i := 0; i < len(magazine); i++ {
		m2[magazine[i]]++
	}

	for k, v := range m1 {
		if m2[k] < v {
			return false
		}
	}

	return true
}
