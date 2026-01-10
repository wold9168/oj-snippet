package main

import "math"

func main() {
	println(checkPowersOfThree(27))
}

func checkPowersOfThree(n int) bool {
	// find the highest mi
	mi := 0
	for math.Pow(3, float64(mi)) <= float64(n) {
		mi += 1
	}
	mi -= 1
	// now the mi is the biggest
	for ; mi >= 0; mi-- {
		if float64(n) >= math.Pow(3, float64(mi)) {
			n -= int(math.Pow(3, float64(mi)))
		}
	}
	return n == 0
}
