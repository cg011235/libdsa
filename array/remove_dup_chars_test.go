package array

import "testing"

/*
  Given a string S which may contain lowercase and uppercase characters.
	The task is to remove all duplicate characters from the string and
	find the resultant string.
*/

func RemoveDups(s string) string {
	runes := []rune(s)
	m := make(map[rune]bool)
	j := 0 // Index for the next unique rune's position

	for _, r := range runes {
		if !m[r] {
			m[r] = true
			runes[j] = r
			j++
		}
	}

	return string(runes[:j])
}

func TestRemoveDups(t *testing.T) {
	in := "geeksforgeeks"
	expected := "geksofr"

	out := RemoveDups(in)
	if out != expected {
		t.Errorf("Expected %s, got %s", expected, out)
	}
}
