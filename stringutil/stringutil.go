package stringutil

import "strings"

// Reverse returns the reverse of the input string.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome checks whether a string reads the same forwards and backwards.
// Comparison is case-insensitive.
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	return s == Reverse(s)
}

// WordCount returns the number of words in a string.
// Words are delimited by whitespace.
func WordCount(s string) int {
	fields := strings.Fields(s)
	return len(fields)
}

// Truncate returns the first n characters of a string.
// If the string is shorter than n, it is returned as-is.
func Truncate(s string, n int) string {
	runes := []rune(s)
	if n < 0 {
		n = 0
	}
	if n >= len(runes) {
		return s
	}
	return string(runes[:n])
}
