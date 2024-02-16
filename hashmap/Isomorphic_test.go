package hashmap

import "testing"

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"one", "two", true},
		{"ab", "aa", false},
		{"aba", "baa", false},
		{"", "", true},
	}

	for _, test := range tests {
		if result := IsIsomorphic(test.s1, test.s2); result != test.expected {
			t.Errorf("IsIsomorphic(%q, %q) = %t; want %t", test.s1, test.s2, result, test.expected)
		}
	}
}
