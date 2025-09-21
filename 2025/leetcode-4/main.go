package main

import "log"

func nextele(nums []int, p int) int {
	if p+1 >= len(nums) {
		return nums[len(nums)-1]
	} else {
		return nums[p+1]
	}
}

func minright(nums1 []int, nums2 []int, p1 int) int {
	p2 := fetchp2(p1, len(nums1)+len(nums2))
	if (p1 == len(nums1)-1 && len(nums2) != 0) || p1 == -1 {
		return nextele(nums2, p2)
	} else if (p2 == len(nums2)-1 && len(nums1) != 0) || p2 == -1 {
		return nextele(nums1, p1)
	} else {
		return min(nextele(nums1, p1), nextele(nums2, p2))
	}
}

func maxleft(nums1 []int, nums2 []int, p1 int) int {
	p2 := fetchp2(p1, len(nums1)+len(nums2))
	if len(nums1) <= 0 || p1 == -1 {
		return nums2[p2]
	} else if len(nums2) <= 0 || p2 == -1 {
		return nums1[p1]
	} else if len(nums1) > 0 && len(nums2) > 0 {
		return max(nums1[p1], nums2[p2])
	}
	return -1
}

func judge(nums1 []int, nums2 []int, p1 int) bool {
	return !(maxleft(nums1, nums2, p1) <= minright(nums1, nums2, p1))
}

func fetchp2(p1, Nsum int) int {
	return (Nsum+1)/2 - p1 - 2 //grant nums1[0,p1] and nums2[0,p2] include leftmid
}

func getmid(l, r int) int {
	return (l + r) / 2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	N1 := len(nums1)
	l1 := 0
	r1 := N1 - 1
	p1 := getmid(l1, r1)
	if N1 == 0 {
		p1 = -1
	}
	for judge(nums1, nums2, p1) && l1 < r1 {
		if p1 != -1 && nums1[p1] <= minright(nums1, nums2, p1) {
			l1 = p1 + 1
		} else {
			r1 = p1
		}
		p1 = getmid(l1, r1)
	}

	if (len(nums1)+len(nums2))%2 == 0 {
		a := maxleft(nums1, nums2, p1)
		b := minright(nums1, nums2, p1)
		return float64(a+b) / 2.0
	} else {
		return float64(maxleft(nums1, nums2, p1))
	}
}

func main() {
	var n1 []int = []int{1, 2}
	var n2 []int = []int{3, 4}
	log.Printf("%v", findMedianSortedArrays(n1, n2))
}
