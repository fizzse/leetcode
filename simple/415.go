package main

import (
	"bytes"
	"fmt"
	"strconv"
)

/*
https://leetcode-cn.com/problems/add-strings/
415. 字符串相加
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。

你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。



示例 1：

输入：num1 = "11", num2 = "123"
输出："134"
示例 2：

输入：num1 = "456", num2 = "77"
输出："533"
示例 3：

输入：num1 = "0", num2 = "0"
输出："0"
*/

func addStrings(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)

	i := l1 - 1
	j := l2 - 1

	var buf bytes.Buffer
	carry := 0
	for i >= 0 || j >= 0 || carry != 0 {
		n1 := 0
		n2 := 0
		if i >= 0 {
			n1 = binNum(num1[i])
		}

		if j >= 0 {
			n2 = binNum(num2[j])
		}

		sum := n1 + n2 + carry
		carry = sum / 10
		sum %= 10
		buf.WriteString(strconv.Itoa(sum))
		i--
		j--
	}

	r := buf.Bytes()
	l3 := len(r)
	for k := 0; k < l3/2; k++ {
		r[k], r[l3-1-k] = r[l3-1-k], r[k]
	}

	return string(r)
}

func binNum(b byte) int {
	return int(b - 48)
}

func main() {
	n1 := "11"
	n2 := "9"

	a := addStrings(n1, n2)
	fmt.Println(a)
}
