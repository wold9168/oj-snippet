package main

import (
	"errors"
	"log"
)

var p1chosen int
var N1 int
var N2 int

func fetchp2chosen() int {
	// [1 2] 3
	// [1 2] 3 4
	return (N1+N2+1)/2 - p1chosen
}

func fetchp1() int {
	return p1chosen - 1
}

func fetchp2() int {
	return fetchp2chosen() - 1
}

func weak_get(nums []int, pos int) (int, error) {
	if pos < 0 || pos >= len(nums) {
		return 0, errors.New("Exceed the range.")
	} else {
		return nums[pos], nil
	}

}

func maxleft(nums1 []int, nums2 []int) (int, error) {
	left1, err1 := weak_get(nums1, fetchp1())
	left2, err2 := weak_get(nums2, fetchp2())
	if err1 != nil {
		if err2 != nil {
			return 0, errors.New("No left")
		} else {
			return left2, nil
		}
	} else if err2 != nil {
		return left1, nil
	} else {
		return max(left1, left2), nil
	}
}

func minright(nums1 []int, nums2 []int) (int, error) {
	right1, err1 := weak_get(nums1, fetchp1()+1)
	right2, err2 := weak_get(nums2, fetchp2()+1)
	if err1 != nil {
		if err2 != nil {
			return 0, errors.New("No right")
		} else {
			return right2, nil
		}
	} else if err2 != nil {
		return right1, nil
	} else {
		return min(right1, right2), nil
	}
}

func found(nums1 []int, nums2 []int) bool {
	ml, errl := maxleft(nums1, nums2)
	mr, errr := minright(nums1, nums2)
	if errl != nil || errr != nil {
		return false
	} else {
		return ml <= mr
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) { //grant N1<=N2
		return findMedianSortedArrays(nums2, nums1)
	}
	N1 = len(nums1)
	N2 = len(nums2)
	if N1 == 0 {
		p1chosen = 0
	}
	L := 0
	R := N1
	p1chosen = (L + R) / 2
	for !found(nums1, nums2) && L < R {
		left1, err1 := weak_get(nums1, fetchp1())
		if err1 != nil {
			if fetchp1() < 0 {
				L = p1chosen + 1
			} else {
				R = p1chosen
			}
			p1chosen = (L + R) / 2
			continue
		}
		mr, errr := minright(nums1, nums2)
		if errr != nil {
			R = p1chosen
			p1chosen = (L + R) / 2
			continue
		}
		if left1 <= mr {
			L = p1chosen + 1 // Right
		} else {
			R = p1chosen // Left
		}
		p1chosen = (L + R) / 2
	}
	if (N1+N2)%2 == 0 {
		ml, _ := maxleft(nums1, nums2)
		mr, _ := minright(nums1, nums2)
		return float64(ml+mr) / 2.0
	} else {
		ml, _ := maxleft(nums1, nums2)
		return float64(ml)
	}
}

func main() {
	var n1 []int = []int{2, 3, 4}
	var n2 []int = []int{1}
	log.Printf("%v", findMedianSortedArrays(n1, n2))
}
