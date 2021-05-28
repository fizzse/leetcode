package main

/*
实现 strStr() 函数。

给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。



说明：

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。



示例 1：

输入：haystack = "hello", needle = "ll"
输出：2
示例 2：

输入：haystack = "aaaaa", needle = "bba"
输出：-1
示例 3：

输入：haystack = "", needle = ""
输出：0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-strstr
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func strStr(haystack string, needle string) int {
	index := -1

	if haystack == "" && needle == "" {
		return 0
	}

	size1 := len(haystack)
	size2 := len(needle)
	if size1 < size2 {
		return -1
	}

	firstByte := int32(0)
	for _, b := range needle {
		firstByte = b
		break
	}

	// needle == "" return 0
	if firstByte == 0 {
		return 0
	}

	for i, b := range haystack {
		if b == firstByte {
			if i+size2 > size1 {
				return -1
			}

			if compare(haystack[i:i+size2], needle) {
				return i
			}
		}
	}

	return index
}

// FIXME 比较逻辑就不写了
func compare(str1, str2 string) bool {
	return str1 == str2
}
