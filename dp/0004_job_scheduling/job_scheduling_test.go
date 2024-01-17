package jobscheduling

import (
	"sort"
	"testing"
)

// Test for binarySearch function
func TestBinarySearch(t *testing.T) {
	jobs := Jobs{
		{start: 1, end: 3, profit: 5},
		{start: 2, end: 5, profit: 6},
		{start: 4, end: 6, profit: 5},
		{start: 6, end: 7, profit: 4},
		{start: 5, end: 8, profit: 11},
		{start: 7, end: 9, profit: 2},
	}

	sort.Sort(jobs)

	lastNonConflicting := binarySearch(jobs, 3)
	if lastNonConflicting != 2 {
		t.Errorf("binarySearch failed, expected %v, got %v", 2, lastNonConflicting)
	}
}

// Test for findMaxProfit function
func TestFindMaxProfit(t *testing.T) {
	jobs := Jobs{
		{start: 1, end: 3, profit: 5},
		{start: 2, end: 5, profit: 6},
		{start: 4, end: 6, profit: 5},
		{start: 6, end: 7, profit: 4},
		{start: 5, end: 8, profit: 11},
		{start: 7, end: 9, profit: 2},
	}

	expectedProfit := 17
	calculatedProfit := findMaxProfit(jobs)

	if calculatedProfit != expectedProfit {
		t.Errorf("findMaxProfit failed, expected %v, got %v", expectedProfit, calculatedProfit)
	}
}
