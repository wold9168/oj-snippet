package main

import (
	"fmt"
)

type stack []int

func (st *stack) push(x int) {
	*st = append(*st, x)
}
func (st *stack) pop() {
	*st = (*st)[0 : len(*st)-1]
}
func (st stack) top() int {
	return st[len(st)-1]
}
func (st stack) top2() (int, error) {
	if len(st)-2 >= 0 {
		return st[len(st)-2], nil
	} else {
		return 0, fmt.Errorf("len(st)<2")
	}
}
func (st stack) empty() bool {
	return len(st) == 0
}

func trap_st(height []int) int {
	fmt.Println("TEST:", height)
	var st stack
	ans := 0
	for i := 0; i < len(height); i++ {
		cur := height[i]
		if len(st) == 0 {
			st.push(i)
		} else if cur <= height[st.top()] {
			st.push(i)
		} else {
			for !st.empty() && cur > height[st.top()] {
				if len(st) >= 2 { // 单调栈中元素小于2的话无法形成可以积水的结构
					left, _ := st.top2()
					width := i - left - 1
					h := min(cur, height[left]) - height[st.top()]
					ans += width * h
				}
				st.pop()
			}
			st.push(i)
		}
	}
	return ans
}

// 单调栈解法
