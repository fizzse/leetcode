package main

import "fmt"

// topk 不存在重复数据(快排实现) FIXME(小顶堆实现)
func getTopk(arr []int, k int) int {
	topKMove(arr, k)
	return arr[k-1]
}

func topKMove(arr []int, k int) {
	size := len(arr)
	if size < 2 {
		return
	}

	i := partitionDesc(arr)
	if i+1 == k {
		return
	}

	if i+1 > k { // 位置靠前 从后面的切片中重新查找
		topKMove(arr[:i], k)
	}

	topKMove(arr[i+1:], k-i-1)
}

func partitionDesc(arr []int) int { // 倒叙partition
	end := len(arr) - 1
	pivot := arr[end]
	i := 0

	for j := 0; j < end; j++ {
		if arr[j] > pivot {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]
	return i
}

//-------------------------------------------------------------------------------//
// 先去重复
func thirdMax(nums []int) int {
	k := 3

	n := getFilter(nums)

	if k > len(n) {
		return getTopk(n, 1)
	}

	return getTopk(n, k)
}

func getFilter(nums []int) []int {
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	n := make([]int, 0, len(nums))
	for num := range m {
		n = append(n, num)
	}
	return n
}

func main() {
	a := []int{1, -2147483648, 2}
	//a := []int{5, 2, 4, 1, 3, 6, 0}
	//a := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println("ret", getTopk(getFilter(a), 3))
	//fmt.Println(thirdMax(a))
}
