package main

/*
521. 最长特殊序列 Ⅰ
给你两个字符串 a 和 b，请返回 这两个字符串中 最长的特殊序列  。如果不存在，则返回 -1 。

「最长特殊序列」 定义如下：该序列为 某字符串独有的最长子序列（即不能是其他字符串的子序列） 。

字符串 s 的子序列是在从 s 中删除任意数量的字符后可以获得的字符串。

例如，“abc” 是 “aebdc” 的子序列，因为您可以删除 “aebdc” 中的下划线字符来得到 “abc” 。 “aebdc” 的子序列还包括 “aebdc” 、 “aeb” 和 “” (空字符串)。
*/

/*
- 这题没理解
- 注意题目中的独有两个字，
- s1 = 'ab',s2 = 'a',因为ab是s1独有，所以最长子序列为ab，
- s1 = 'ab', s2 = 'ab', 因为ab是两个串都有，ab排除，a也是两个串都有，排除，b也是两个串都有，排除。所以最长特殊序列不存在，返回-1
- 通过以上分析，我们可以得出结论，如果：两个串相等（不仅长度相等，内容也相等），那么他们的最长特殊序列不存在。返回-1
- 如果两个串长度不一样，那么长的串   永远也不可能是   短串的子序列，即len(s1) > len(s2),则最长特殊序列为s1,返回长度大的数
*/

func findLUSlength(a string, b string) int {
	s1 := len(a)
	s2 := len(b)
	if s1 > s2 {
		return s1
	}

	if s1 < s2 {
		return s2
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return s1
		}
	}

	return 0
}

func main() {
	a := "abadd"
	b := "cdc"
	findLUSlength(a, b)
}
