package main

/*
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。



示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
示例 3：

输入：digits = [0]
输出：[1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/plus-one
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func plusOne(digits []int) []int {
	size := len(digits)
	digits[size-1] += 1
	if digits[size-1] < 10 {
		return digits
	}

	ret := make([]int, size+1)
	carry := 0
	for i := size - 1; i >= 0; i-- {
		tmp := digits[i] + carry

		carry = 0
		if tmp >= 10 {
			tmp = tmp - 10
			carry = 1
		}
		ret[i+1] = tmp
	}

	if carry > 0 {
		ret[0] = carry
		return ret
	}

	return ret[1:]
}

func main() {
	plusOne([]int{8, 9, 9, 9})
}
