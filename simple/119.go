package main

/*
给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。



在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 3
输出: [1,3,3,1]
进阶：

你可以优化你的算法到 O(k) 空间复杂度吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func getRow(rowIndex int) []int {
	/*
		两行来回赋值 2*k的内存消耗 符合O(k)
	*/

	row1 := make([]int, rowIndex+1)
	row2 := make([]int, rowIndex+1)

	i := 0
	for ; i < rowIndex+1; i++ {
		row1[0] = 1
		row1[i] = 1

		row2[0] = 1
		row2[i] = 1

		if i%2 != 0 {
			for j := 1; j < i; j++ {
				row2[j] = row1[j-1] + row1[j]
			}
			continue
		}

		for j := 1; j < i; j++ {
			row1[j] = row2[j-1] + row2[j]
		}
	}

	if i%2 == 0 {
		return row2
	}

	return row1
}
