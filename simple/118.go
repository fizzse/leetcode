package main

/*
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。



在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func generate(numRows int) [][]int {
	/*
	 边界都是1 其余位置为 sub[j] = res[i-1][j-1] + res[i-1][j]
	*/

	res := make([][]int, 0, numRows)
	for i := 0; i < numRows; i++ {
		sub := make([]int, i+1)
		sub[0] = 1
		sub[i] = 1

		for j := 1; j < i; j++ {
			sub[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, sub)
	}

	return res
}
