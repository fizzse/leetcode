package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

/*
https://leetcode-cn.com/problems/base-7/
504. 七进制数
给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。

示例 1:

输入: num = 100
输出: "202"
示例 2:

输入: num = -7
输出: "-10"
*/

func convertToBase7(num int) string {
	numAbs := int(math.Abs(float64(num)))

	x := 7
	var ret []int
	i := 0
	for numAbs != 0 {
		i++
		tmp := numAbs % x
		numAbs /= x
		ret = append(ret, tmp)
	}

	buf := bytes.Buffer{}
	for i := len(ret) - 1; i >= 0; i-- {
		buf.WriteString(strconv.Itoa(ret[i]))
	}

	res := buf.String()
	if res == "" {
		res = "0"
	}

	if num < 0 {
		res = "-" + res
	}
	return res
}

func main() {
	in := 0
	out := convertToBase7(in)

	fmt.Println(out)
}
