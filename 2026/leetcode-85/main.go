package main

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	println(maximalRectangle(matrix))
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

func maximalRectangle(matrix [][]byte) int {
	columns := make([][]int, len(matrix))
	ans := 0
	for i, row := range matrix {
		columns[i] = make([]int, len(row))
		for j, v := range row {
			if v == '0' {
				columns[i][j] = 0
			} else {
				if i == 0 {
					columns[i][j] = 1
				} else {
					columns[i][j] = columns[i-1][j] + 1
				}
			}
		}
		ans = max(ans, largestRectangleArea(columns[i]))
	}
	return ans
}
