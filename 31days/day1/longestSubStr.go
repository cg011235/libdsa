package day1

/*
 Longest Substring Without Repeating Characters:
 Given a string s, find the length of the longest substring without repeating characters.
*/

func longestSubStrWithoutRepeatChar(s string) int {
	maxlen := 0
	start := 0
	m := make(map[rune]int)

	for i, c := range s {
		if lastIdx, ok := m[c]; ok && lastIdx >= start {
			start = lastIdx + 1
		}
		m[c] = i
		if curLen := i - start + 1; curLen > maxlen {
			maxlen = curLen
		}
	}
	return maxlen
}
