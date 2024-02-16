package hashmap

import (
	"testing"
)

func TestHashMap_InsertFind(t *testing.T) {
	hm := NewHashMap(10) // Create a new HashMap with size 10

	// Insert some key-value pairs
	hm.Insert(1, "one")
	hm.Insert(2, "two")
	hm.Insert(3, "three")

	// Test Find method
	if val := hm.Find(1); val != "one" {
		t.Errorf("Expected 'one', got '%v'", val)
	}
	if val := hm.Find(2); val != "two" {
		t.Errorf("Expected 'two', got '%v'", val)
	}
	if val := hm.Find(3); val != "three" {
		t.Errorf("Expected 'three', got '%v'", val)
	}
	if val := hm.Find(4); val != nil {
		t.Errorf("Expected 'nil', got '%v' for non-existent key", val)
	}
}

func TestHashMap_Remove(t *testing.T) {
	hm := NewHashMap(10) // Create a new HashMap with size 10

	// Insert some key-value pairs
	hm.Insert(1, "one")
	hm.Insert(2, "two")
	hm.Insert(3, "three")

	// Remove a key-value pair
	hm.Remove(2)

	// Test if the key-value pair was successfully removed
	if val := hm.Find(2); val != nil {
		t.Errorf("Expected 'nil', got '%v' after removal", val)
	}
}

func TestHashMap_Overwrite(t *testing.T) {
	hm := NewHashMap(10) // Create a new HashMap with size 10

	// Insert a key-value pair
	hm.Insert(1, "one")

	// Overwrite the value for the same key
	hm.Insert(1, "ONE")

	// Test if the value was overwritten
	if val := hm.Find(1); val != "ONE" {
		t.Errorf("Expected 'ONE', got '%v' after overwrite", val)
	}
}
