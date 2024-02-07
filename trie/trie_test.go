package trie

import "testing"

func AssertTrue(t *testing.T, condition bool, message string) {
	if !condition {
		t.Errorf("Assertion failed: %s", message)
	}
}

func AssertFalse(t *testing.T, condition bool, message string) {
	if condition {
		t.Errorf("Assertion failed: %s", message)
	}
}

func TestTrieBasic(t *testing.T) {
	trie := new(Trie)
	trie.Init()

	result := trie.Insert("cat")
	AssertTrue(t, result, "Failed to insert string")
	result = trie.Insert("car")
	AssertTrue(t, result, "Failed to insert string")
	result = trie.Insert("call")
	AssertTrue(t, result, "Failed to insert string")
	result = trie.Insert("category")
	AssertTrue(t, result, "Failed to insert string")
	result = trie.Insert("calendar")
	AssertTrue(t, result, "Failed to insert string")
	result = trie.Insert("calendly")
	AssertTrue(t, result, "Failed to insert string")

	result = trie.Search("call")
	AssertTrue(t, result, "Failed to search string 'call'")
	result = trie.Search("calendar")
	AssertTrue(t, result, "Failed to search string 'calendar'")
	result = trie.Search("caw")
	AssertFalse(t, result, "String 'caw' should not be present")
	result = trie.Search("carr")
	AssertFalse(t, result, "String 'carr' should not be present")
	result = trie.Search("caller")
	AssertFalse(t, result, "String 'carr' should not be present")

	result = trie.SearchPrefix("ca")
	AssertTrue(t, result, "Failed to search prefix 'ca'")
	result = trie.SearchPrefix("calend")
	AssertTrue(t, result, "Failed to search prefix 'calend'")
}
