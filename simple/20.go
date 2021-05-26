package main

/*
 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。


示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false
示例 4：

输入：s = "([)]"
输出：false
示例 5：

输入：s = "{[]}"
输出：true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isValid(s string) bool {
	size := len(s)
	if size%2 != 0 {
		return false
	}

	push := map[byte]bool{
		'[': true,
		'(': true,
		'{': true,
	}

	pub := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}

	stack := make([]byte, size/2+1)
	stackSize := 0

	for i := 0; i < size; i++ {
		if push[s[i]] {
			stack[stackSize] = s[i]
			stackSize++
			if stackSize > size/2 {
				return false
			}
		} else {
			if stackSize <= 0 {
				return false
			}
			if pub[s[i]] != stack[stackSize-1] {
				return false
			}

			stackSize--
		}
	}

	return stackSize == 0
}
