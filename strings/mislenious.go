package main

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	removeSpecialChar := func(char rune) rune {
		if char < 'a' || char > 'z' {
			return rune(' ')
		}
		return char
	}
	s = strings.Map(removeSpecialChar, s)
	s = strings.Join(strings.Fields(s), "")
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {

}
