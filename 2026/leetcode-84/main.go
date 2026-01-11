package main

func main() {
	println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

func buildMonostack(heights []int) ([]int, []int) {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	monostack := []int{}
	for i := 0; i < n; i++ {
		for len(monostack) > 0 && heights[monostack[len(monostack)-1]] >= heights[i] {
			monostack = monostack[:len(monostack)-1]
		}
		if len(monostack) == 0 {
			left[i] = -1
		} else {
			left[i] = monostack[len(monostack)-1]
		}
		monostack = append(monostack, i)
	}
	monostack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(monostack) > 0 && heights[monostack[len(monostack)-1]] >= heights[i] {
			monostack = monostack[:len(monostack)-1]
		}
		if len(monostack) == 0 {
			right[i] = n
		} else {
			right[i] = monostack[len(monostack)-1]
		}
		monostack = append(monostack, i)
	}
	return left, right
}

func largestRectangleArea(heights []int) int {
	l, r := buildMonostack(heights)
	ans := 0
	for i := 0; i < len(heights); i++ {
		ans = max(ans, (r[i]-l[i]-1)*heights[i])
	}
	return ans
}
