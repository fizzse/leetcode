package main

/*
https://leetcode.cn/problems/asteroid-collision/
735. 行星碰撞
给定一个整数数组 asteroids，表示在同一行的行星。

对于数组中的每一个元素，其绝对值表示行星的大小，正负表示行星的移动方向（正表示向右移动，负表示向左移动）。每一颗行星以相同的速度移动。

找出碰撞后剩下的所有行星。碰撞规则：两个行星相互碰撞，较小的行星会爆炸。如果两颗行星大小相同，则两颗行星都会爆炸。两颗移动方向相同的行星，永远不会发生碰撞。

示例 1：

输入：asteroids = [5,10,-5]
输出：[5,10]
解释：10 和 -5 碰撞后只剩下 10 。 5 和 10 永远不会发生碰撞。
示例 2：

输入：asteroids = [8,-8]
输出：[]
解释：8 和 -8 碰撞后，两者都发生爆炸。
示例 3：

输入：asteroids = [10,2,-5]
输出：[10]
解释：2 和 -5 发生碰撞后剩下 -5 。10 和 -5 发生碰撞后剩下 10 。
*/

type Stack struct {
	store []int
	size  int
}

func NewStack(cap int) *Stack {
	return &Stack{
		store: make([]int, cap),
	}
}

func (s *Stack) Push(num int) {
	s.store[s.size] = num
	s.size++
}

func (s *Stack) Pop() int {
	s.size--
	return s.store[s.size]
}

func (s *Stack) Top() int {
	return s.store[s.size-1]
}

func (s *Stack) Store() []int {
	return s.store[:s.size]
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func absInt(num int) int {
	if num >= 0 {
		return num
	}

	return -num
}

func asteroidCollision(asteroids []int) []int {
	s := NewStack(len(asteroids))

	for _, num := range asteroids {
		if s.Empty() || num > 0 {
			s.Push(num)
			continue
		}

		remain := true
		for !s.Empty() {
			if s.Top() < 0 { // 方向相同 直接写入
				break
			}

			if s.Top() == absInt(num) { // 方向相反，且质量相同，两个都销毁
				s.Pop()
				remain = false
				break
			}

			if s.Top() > absInt(num) { // 新的被销毁
				remain = false
				break
			}

			s.Pop() // 小于后面的，本颗需要销毁，栈中其他相反方向的依次比较
		}

		if remain { // 这个星球是存活下来
			s.Push(num)
		}
	}

	return s.Store()
}
