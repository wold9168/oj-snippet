package main

import "math"

func main() {
	println(separateSquares([][]int{{0, 0, 1}, {2, 2, 1}}))
	println(separateSquares([][]int{{0, 0, 2}, {1, 1, 1}}))
}

func c(total_area float64, squares [][]int, ky float64) bool {
	area := 0.0
	for _, sq := range squares {
		y, l := sq[1], sq[2]
		if float64(y) < ky {
			area += float64(l) * math.Min(ky-float64(y), float64(l))
		}
	}
	println(total_area/2.0, ";area:", area, ";ky:", ky, ";res:", area < (total_area/2.0))
	return area >= (total_area / 2.0)
}

func separateSquares(squares [][]int) float64 {
	total_area := 0.0
	ymin, ymax := 1.1e9, -0.1
	for _, sq := range squares {
		y, l := sq[1], sq[2]
		total_area += float64(l * l)
		ymin = math.Min(ymin, float64(y))
		ymax = math.Max(ymax, float64(y+l))
	}
	L := ymin
	R := ymax
	threshold := 1e-5
	for math.Abs(R-L) > threshold {
		m := (R-L)/2.0 + L
		if c(total_area, squares, m) {
			println("L", L, "R", R, "m", m, "\t", "High")
			R = m
		} else {
			println("L", L, "R", R, "m", m, "\t", "Low")
			L = m
		}
	}
	return R
}
