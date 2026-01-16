package main

// import "fmt"

func findAnagrams(s string, p string) []int {
	// println("findAnagrams"," s:", s," p:", p)
	n := len(p)
	L := 0
	if len(s) < len(p) {
		return []int{}
	}
	R := L + n
	mp := make(map[byte]int)
	for _, v := range p {
		mp[byte(v)] += 1
	}
	win := make(map[byte]int)
	for i := L; i < R; i++ {
		win[s[i]] += 1
	}
	ans := []int{}
	if fulljudge(mp, win) {
		// println(0)
		ans = append(ans, 0)
	}
	for R < len(s) {
		win[s[L]] -= 1
		win[s[R]] += 1
		// println("-:",L,"\t+:",R)
		L++
		R++
		if fulljudge(mp, win) {
			// println(L)
			ans = append(ans, L)
		}
	}
	return ans

}
func fulljudge(a, b map[byte]int) bool {
	// println("fulljudge()")
	for k, v := range a {
		// fmt.Println(a,";",b)
		// fmt.Println("k,v: ", string(k), " ", v, "; d: ", b[k])
		if b[k] != v {
			return false
		}
	}
	for k, v := range b {
		// fmt.Println("k,v: ", string(k), " ", v, "; d: ", a[k])
		if a[k] != v {
			return false
		}
	}
	return true
}
