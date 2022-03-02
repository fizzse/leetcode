package main

import (
	"fmt"
	"strconv"
)

// 复杂度O(n)
func summaryRanges(nums []int) []string {
	var ret []string
	size := len(nums)

	for i := 0; i < size; {
		start := nums[i]
		j := 1
		for ; j < size-i; j++ {
			if nums[i+j] != start+j {
				break
			}
		}

		var r string
		if j == 1 {
			r = strconv.Itoa(start)
		} else {
			end := nums[i+j-1]
			r = strconv.Itoa(start) + "->" + strconv.Itoa(end)
		}

		ret = append(ret, r)
		i += j
	}

	return ret
}

func main() {
	in := []int{0, 1, 2, 4, 5, 7}
	out := summaryRanges(in)
	fmt.Println(out)
}
