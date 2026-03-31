package stringutil

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple", "hello", "olleh"},
		{"empty", "", ""},
		{"single char", "a", "a"},
		{"palindrome", "racecar", "racecar"},
		{"unicode", "Hello, 世界", "界世 ,olleH"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"lowercase palindrome", "racecar", true},
		{"mixed case palindrome", "RaceCar", true},
		{"not palindrome", "hello", false},
		{"empty string", "", true},
		{"single char", "a", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestWordCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"simple", "hello world", 2},
		{"empty", "", 0},
		{"extra spaces", "  hello   world  ", 2},
		{"single word", "hello", 1},
		{"tabs and newlines", "hello\tworld\nfoo", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WordCount(tt.input)
			if result != tt.expected {
				t.Errorf("WordCount(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		n        int
		expected string
	}{
		{"truncate", "hello world", 5, "hello"},
		{"shorter than n", "hi", 10, "hi"},
		{"zero", "hello", 0, ""},
		{"negative", "hello", -1, ""},
		{"unicode", "Hello, 世界", 8, "Hello, 世"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Truncate(tt.input, tt.n)
			if result != tt.expected {
				t.Errorf("Truncate(%q, %d) = %q; want %q", tt.input, tt.n, result, tt.expected)
			}
		})
	}
}
