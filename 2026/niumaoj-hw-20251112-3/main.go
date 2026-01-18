package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	var a, b []int
	for i := 0; i < N; i++ {
		var t int
		fmt.Scan(&t)
		a = append(a, t)
	}
	for i := 0; i < M; i++ {
		var t int
		fmt.Scan(&t)
		b = append(b, t)
	}
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, M+1)
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			dp[i][j] = imax(dp[i-1][j-1]+abs(a[i-1]-b[j-1]), imax(dp[i][j-1], dp[i-1][j]))
		}
	}
	fmt.Println(dp[N][M])
	return
}
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -1 * x
	}
}
func imax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
