package main

import "fmt"

/*
https://leetcode-cn.com/problems/permutation-in-string/
567. 字符串的排列
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。

换句话说，s1 的排列之一是 s2 的 子串 。



示例 1：

输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").
示例 2：

输入：s1= "ab" s2 = "eidboaoo"
输出：false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutation-in-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func checkInclusion(s1 string, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)

	char1 := s1[0]
	char2 := s1[l1-1]

	for i := 0; i < l2; i++ {
		if s2[i] == char1 {
			if i+l1 > l2 {
				return false
			}

			if str1(s2[i:i+l1], s1) {
				return true
			}
		}

		if s2[i] == char2 {
			if i+l1 > l2 {
				return false
			}

			if str2(s2[i:i+l1], s1) {
				return true
			}
		}
	}

	return false
}

// 字符串相等
func str1(s1, s2 string) bool {
	return s1 == s2
}

// 字符串反转相对
func str2(s1, s2 string) bool {
	size := len(s1)
	j := size - 1
	for i := 0; i < size; i++ {
		if s1[i] != s2[j] {
			return false
		}
		j--
	}

	return true
}

func main() {
	r := checkInclusion("adc",
		"dcda")

	fmt.Println(r)

	r = str2("cda", "adc")
	fmt.Println(r)
}
