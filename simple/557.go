package main

/*
https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/

给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。



示例：

输入："Let's take LeetCode contest"
输出："s'teL ekat edoCteeL tsetnoc"


提示：

在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reverseWords1(s string) string {
	buf := make([]byte, len(s))

	i := 0
	size := len(s)

	for j := 0; j < size; j++ {
		if s[j] != ' ' && j != size-1 {
			continue
		}

		if j == size-1 {
			j++
		}

		for k := j - 1; j > i; k-- {
			buf[k] = s[i]
			i++
		}

		buf[j] = s[j]
		i = j + 1
	}
	return string(buf)
}

func reverseWords(s string) string {
	buf := make([]byte, len(s))

	i := 0
	size := len(s)

	for j := 0; j < size; j++ {
		end := 0
		switch {
		case s[j] == ' ':
			end = j - 1
		case j == size-1:
			end = j
		default:
			continue
		}

		for k := end; end > i; k-- {
			buf[k] = s[i]
			i++
		}

		buf[j] = s[j]
		i = j + 1
	}
	return string(buf)
}
