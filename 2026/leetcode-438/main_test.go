package main

import (
	"sort"
	"strconv"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want []int
	}{
		{
			s:    "cbaebabacd",
			p:    "abc",
			want: []int{0, 6},
		},
		{
			s:    "abab",
			p:    "ab",
			want: []int{0, 1, 2},
		},
	}
	for i, tt := range tests {
		t.Run("Test "+strconv.Itoa(i), func(t *testing.T) {
			if got := findAnagrams(tt.s, tt.p); !isSame(tt.want, got) {
				t.Errorf("[Result] ans = %v, want %v", got, tt.want)
			}
		})
	}
}

func isSame(a, b []int) bool {
	sort.Ints(a)
	sort.Ints(b)
	if len(a) == len(b) {
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
	} else {
		return false
	}
	return true
}
