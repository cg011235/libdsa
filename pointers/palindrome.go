package pointers

func isPalindrome(inputString string) bool {
	n := len(inputString)
	i, j := 0, n-1
	for i < j {
		if inputString[i] != inputString[j] {
			return false
		}
		i++
		j--
	}
	return true
}

/*
 * check if given string is valid palindrome
 * by removing at most one character from it.
 */
func validPalindrome(s string) bool {
	isPalindrome := func(s string, start, end int) bool {
		for start < end {
			if s[start] != s[end] {
				return false
			}
			start++
			end--
		}
		return true
	}

	start, end := 0, len(s)-1
	for start < end {
		if s[start] != s[end] {
			return isPalindrome(s, start+1, end) || isPalindrome(s, start, end-1)
		}
		start++
		end--
	}
	return true
}
