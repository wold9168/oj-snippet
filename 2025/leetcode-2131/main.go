package leetcode2131

import "strings"

func reverse(rawword string) string {
	var res string
	for _, ch := range rawword {
		res = string([]rune{ch}) + res
	}
	return res
}
func longestPalindrome(words []string) int {
	m := make(map[string]int)
	res := 0
	for _, word := range words {
		m[word]++
	}
	flag := 0
	for word, cnt := range m {
		rever := reverse(word)
		if word == rever {
			flag = max(flag, cnt-cnt/2*2)
			res += 2 * (cnt / 2 * 2)
		} else if strings.Compare(word, rever) > 0 {
			res += 2 * 2 * min(m[word], m[rever])
		}
	}
	return res + flag*2
}
