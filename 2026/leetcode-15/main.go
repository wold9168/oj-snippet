package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	fmt.Println(nums)
	for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1]{
            continue
        }
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
				println(i, j, k, " k--")
			} else if sum < 0 {
				j++
				println(i, j, k, " j++")
			} else if sum == 0 {
				x := []int{nums[i], nums[j], nums[k]}

				if len(ans) > 0 {
					if al := ans[len(ans)-1]; !(x[0] == al[0] && x[1] == al[1] && x[2] == al[2]) {
						ans = append(ans, x)
					}
				} else {
					ans = append(ans, x)
				}
				fmt.Println("ans: ", ans, ";i,j,k:", i, j, k)
				for j+1 < len(nums) && nums[j] == nums[j+1] && j < k {
					println("j++")
					j++
				}
				for k-1 > 0 && nums[k-1] == nums[k] && j < k {
					println("k--")
					k--
				}
				j++
			}
		}
	}
	return ans
}

// Output should be commented
