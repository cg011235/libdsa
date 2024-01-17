package eggdropping

import (
	"fmt"
	"testing"
)

// Test cases for the EggDrop function
func TestEggDrop(t *testing.T) {
	testCases := []struct {
		eggs, floors, expected int
	}{
		{1, 10, 10},
		{2, 10, 4},
		{3, 10, 4},
		{2, 6, 3},
		{3, 14, 4},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%dEggs%dFloors", tc.eggs, tc.floors), func(t *testing.T) {
			actual := eggDropBottomUp(tc.eggs, tc.floors)
			if actual != tc.expected {
				t.Errorf("EggDrop(%d, %d) = %d; expected %d", tc.eggs, tc.floors, actual, tc.expected)
			}
		})
	}
}
