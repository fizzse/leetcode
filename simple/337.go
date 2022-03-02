package main

/*
给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。



示例 1：

输入：n = 2
输出：[0,1,1]
解释：
0 --> 0
1 --> 1
2 --> 10
示例 2：

输入：n = 5
输出：[0,1,1,2,1,2]
解释：
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/counting-bits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 求与法
// FIXME n & n-1 能去掉最后面的那个1,2的整数次幂只有1个1,如果n&(n-1)=0,n是2的整数次幂
func countBits(n int) []int {
	ret := make([]int, 0, n+1)
	ret = append(ret, 0)
	for i := 1; i <= n; i++ {
		cnt := 0
		tmp := i
		for tmp != 0 {
			cnt++
			tmp &= tmp - 1
		}

		ret = append(ret, cnt)
	}

	return ret
}
