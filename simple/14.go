package main

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。


示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	minSize := 0
	// 求最小的size
	for i := 0; i < length; i++ {
		if i == 0 {
			minSize = len(strs[0])
			continue
		}

		size := len(strs[i])
		if size < minSize {
			minSize = size
		}
	}

	i := 0
	var bit byte
label:
	for ; i < minSize; i++ {
		bit = strs[0][i]
		for j := 0; j < length; j++ {
			if strs[j][i] != bit {
				break label
			}
		}
	}

	if i == 0 {
		return ""
	}

	return strs[0][:i]
}
