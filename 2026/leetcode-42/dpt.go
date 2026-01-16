package main

func trap_dpt(height []int) int {
	ans := 0
	L, R := 0, len(height)-1
	Lh, Rh := height[L], height[R]
	for L < R {
		var cur,m int
		var update *int
		if Lh <= Rh {
			L++
			cur = L
			m=Lh
			update=&Lh
		} else {
			R--
			cur = R
			m=Rh
			update=&Rh
		}

		if m>height[cur]{
			ans+=m-height[cur]
		}else{
			*update=height[cur]
		}
	}
	return ans
}

// 双指针解法
