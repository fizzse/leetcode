package main

import (
	"bytes"
	"fmt"
)

/*
https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal/
405. 数字转换为十六进制数
给定一个整数，编写一个算法将这个数转换为十六进制数。对于负整数，我们通常使用 补码运算 方法。

注意:

十六进制中所有字母(a-f)都必须是小写。
十六进制字符串中不能包含多余的前导零。如果要转化的数为0，那么以单个字符'0'来表示；对于其他情况，十六进制字符串中的第一个字符将不会是0字符。
给定的数确保在32位有符号整数范围内。
不能使用任何由库提供的将数字直接转换或格式化为十六进制的方法。
示例 1：

输入:
26

输出:
"1a"
示例 2：

输入:
-1
*/

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	uNum := uint32(num)
	var buf []uint32

	for uNum > 0 {
		x := uNum % 16
		uNum /= 16
		buf = append(buf, x)
	}

	var buffer bytes.Buffer
	for j := len(buf) - 1; j >= 0; j-- {
		buffer.WriteString(fmt.Sprintf("%x", buf[j])) // 这里有库函数 省的写map了
	}

	return buffer.String()
}

func main() {
	in := -1
	out := toHex(in)
	fmt.Println(out)
}
