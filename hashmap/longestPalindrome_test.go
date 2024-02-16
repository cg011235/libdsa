package hashmap

import "testing"

func TestLongestPalindromeLength(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"opinion", 7},
		{"GooooooOOOOODdddD", 15},
		{"AbcDeFGhAachDeFG", 15},
		{"REaccaR", 7},
	}

	for _, test := range tests {
		if result := LongestPalindromeLength(test.s); result != test.expected {
			t.Errorf("LongestPalindromeLength(%s) = %d; want %d", test.s, result, test.expected)
		}
	}
}
