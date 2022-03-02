package main

import "fmt"

/*
https://leetcode-cn.com/problems/binary-watch/
401. 二进制手表
二进制手表顶部有 4 个 LED 代表 小时（0-11），底部的 6 个 LED 代表 分钟（0-59）。每个 LED 代表一个 0 或 1，最低位在右侧。
给你一个整数 turnedOn ，表示当前亮着的 LED 的数量，返回二进制手表可以表示的所有可能时间。你可以 按任意顺序 返回答案。

小时不会以零开头：

例如，"01:00" 是无效的时间，正确的写法应该是 "1:00" 。
分钟必须由两位数组成，可能会以零开头：

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-watch
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func readBinaryWatch1(turnedOn int) []string {
	var res []string
	for i := 0; i < 12; i++ {
		cntI := num1(i)
		for j := 0; j < 60; j++ {
			cntJ := num1(j)
			if cntI+cntJ != turnedOn { // 判断0：00——11：59 中所有1的个数 将1的数量满足条件的 拼接字符串
				continue
			}

			r := fmt.Sprintf("%d:%02d", i, j)
			res = append(res, r)
		}
	}

	return res
}

func num1(n int) int {
	ret := 0
	for n != 0 {
		ret++
		n &= n - 1
	}

	return ret
}

func main() {
	s := fmt.Sprintf("%02d", 10)
	fmt.Println(s)
}
