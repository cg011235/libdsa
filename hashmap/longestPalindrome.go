package hashmap

func LongestPalindromeLength(s string) int {
	charCounts := make(map[rune]int)
	for _, c := range s {
		charCounts[c]++
	}
	longestPalindromeLength := 0
	oddFound := false
	for _, count := range charCounts {
		if count%2 == 0 {
			longestPalindromeLength += count
		} else {
			longestPalindromeLength += count - 1
			oddFound = true
		}
	}
	if oddFound {
		longestPalindromeLength++
	}
	return longestPalindromeLength
}
