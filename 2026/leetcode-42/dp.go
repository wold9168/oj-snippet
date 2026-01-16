package main

func trap(height []int) int {
	ans := 0
	L, R := make([]int, len(height)), make([]int, len(height))
	L[0] = height[0]
	R[len(height)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		L[i] = max(L[i-1], height[i])
	}
	for i := len(height) - 2; i >= 0; i-- {
		R[i] = max(R[i+1], height[i])
	}
	for i, v := range height {
		ans += min(L[i], R[i]) - v
	}
	return ans
}

// 动态规划解法
