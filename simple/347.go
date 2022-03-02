package main

/*
347. 前 K 个高频元素
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。



示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]


提示：

1 <= nums.length <= 105
k 的取值范围是 [1, 数组中不相同的元素的个数]
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
*/

type MinHeap struct {
	cap       int
	size      int
	container []int
}

func NewMinHeap(cap int) *MinHeap {
	return &MinHeap{
		cap:       cap,
		container: make([]int, cap+1),
	}
}

func (h *MinHeap) Push(num int) {
	if h.Full() {
		return
	}
	h.size++
	h.container[h.size] = num

	i := h.size
	for i/2 > 0 && h.container[i] < h.container[i/2] {
		h.container[i], h.container[i/2] = h.container[i/2], h.container[i]
		i = i / 2
	}
}

func (h *MinHeap) Pop() int {
	num := h.container[1]
	h.container[1] = h.container[h.size]
	h.size--
	h.heapify()
	return num
}

func (h *MinHeap) heapify() {
	i := 1
	n := h.size
	a := h.container
	for {
		maxPos := i
		if i*2 < n && a[i] > a[i*2] {
			maxPos = i * 2
		}

		if i*2+1 <= n && a[maxPos] > a[i*2+1] {
			maxPos = i*2 + 1
		}

		if maxPos == i {
			break
		}

		a[i], a[maxPos] = a[maxPos], a[i]
		i = maxPos
	}
}

func (h *MinHeap) Top() int {
	return h.container[1]
}

func (h *MinHeap) Full() bool {
	return h.size == h.cap
}

// 最小堆
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	ret := make([]int, 0, k)

	heap := NewMinHeap(k)
	for _, v := range m {
		if !heap.Full() {
			heap.Push(v)
			continue
		}

		if v > heap.Top() {
			heap.Pop()
			heap.Push(v)
		}
	}

	// 建一个容量为k的小顶堆
	top := heap.Top()
	// 取堆顶元素 map value >= 对顶的就是目标数据
	for k, v := range m {
		if v >= top {
			ret = append(ret, k)
		}
	}

	return ret
}

func topKFrequentBySort(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	ret := make([]int, 0, k)
	values := make([]int, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}

	topK := getTopk(values, k)
	//fmt.Println(topK, values)
	// 取堆顶元素 map value >= 对顶的就是目标数据
	for k, v := range m {
		if v >= topK {
			ret = append(ret, k)
		}
	}

	return ret
}

//----------------------------------------快排------------------------------------------//

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
