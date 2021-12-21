package main

/*
https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/

反转字符串中的单词 III
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

func reverseWords(s string) string {
	buf := make([]byte, len(s))

	i := 0
	size := len(s)

	for j := 0; j < size; j++ {
		switch {
		case s[j] == ' ': // 单词结尾
			buf[j] = ' '

		case j == size-1: // 结尾 +1是为了执行同样的交换逻辑
			j = j + 1

		default:
			continue
		}

		for k := j - 1; i < j; k-- { // 进行赋值 (i等于起始单词位置，j是结束位置+1（空格或者总长度+1）)
			buf[k] = s[i]
			i++
		}

		i = j + 1 // 指向下一个单词
	}
	return string(buf)
}
