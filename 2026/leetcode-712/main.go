package main

func main() {
	testData := [][]string{
		{"sea", "eat"},
		{"delete", "leet"},
	}
	for _, v := range testData {
		s1 := v[0]
		s2 := v[1]
		println(minimumDeleteSum(s1, s2))
	}
}

func calcSum(s string) (sum int) {
	for _, v := range s {
		sum += int(v)
	}
	return
}

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	for i := range m + 1 {
		dp[i][0] = calcSum(s1[0:i])
	}
	for j := range n + 1 {
		dp[0][j] = calcSum(s2[0:j])
	}
	for i, c1 := range s1 {
		for j, c2 := range s2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j+1]+int(s1[i]), dp[i+1][j]+int(s2[j]))
			}

		}
	}
	return dp[m][n]
}
