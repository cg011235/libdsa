package pointers

import "testing"

func TestIsPalindrome(t *testing.T) {
	s := "bob"
	if isPalindrome(s) == false {
		t.Errorf("Expected %s to be palindrome", s)
	}
	s = "hello"
	if isPalindrome(s) == true {
		t.Errorf("Expected %s to be not a palindrome", s)
	}
	s = "ABCDCBA"
	if isPalindrome(s) == false {
		t.Errorf("Expected %s to be palindrome", s)
	}
	s = ""
	if isPalindrome(s) == false {
		t.Errorf("Expected %s to be palindrome", s)
	}
	s = "X"
	if isPalindrome(s) == false {
		t.Errorf("Expected %s to be palindrome", s)
	}
}

func TestValidPalindrome(t *testing.T) {
	s := "boba"
	if validPalindrome(s) == false {
		t.Errorf("Expected %s to be palindrome by removing at most one character from it", s)
	}
	s = "Holog"
	if validPalindrome(s) == true {
		t.Errorf("Expected %s to be not a palindrome", s)
	}
}
