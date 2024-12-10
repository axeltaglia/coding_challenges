package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"", true},
		{" ", true},
		{"12321", true},
		{"12345", false},
		{"0P", false},
		{"madam", true},
		{"No lemon, no melon", true},
		{"Was it a car or a cat I saw?", true},
		{"Able , was I saw eLba", true},
		{"palindrome", false},
		{"!!@@##$$", true}, // No valid alphanumeric characters, should return true.
		{"1a2", false},
		{"aa", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := isPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("isPalindrome(%q) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
