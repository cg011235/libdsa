package jobscheduling

import (
	"sort"
)

type Job struct {
	start, end, profit int
}

// Function to sort jobs based on end time
func (jobs Jobs) Len() int           { return len(jobs) }
func (jobs Jobs) Less(i, j int) bool { return jobs[i].end < jobs[j].end }
func (jobs Jobs) Swap(i, j int)      { jobs[i], jobs[j] = jobs[j], jobs[i] }

type Jobs []Job

// Binary Search to find the latest job (before the current job) that doesn't conflict
func binarySearch(jobs Jobs, index int) int {
	low, high := 0, index-1
	for low <= high {
		mid := (low + high) / 2
		if jobs[mid].end <= jobs[index].start {
			if jobs[mid+1].end <= jobs[index].start {
				low = mid + 1
			} else {
				return mid
			}
		} else {
			high = mid - 1
		}
	}
	return -1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Function to find the maximum profit from scheduling jobs
func findMaxProfit(jobs Jobs) int {
	sort.Sort(jobs) // Sort jobs based on finish time
	n := len(jobs)
	dp := make([]int, n)
	dp[0] = jobs[0].profit

	for i := 1; i < n; i++ {
		inclProf := jobs[i].profit
		l := binarySearch(jobs, i)
		if l != -1 {
			inclProf += dp[l]
		}
		dp[i] = max(inclProf, dp[i-1])
	}

	return dp[n-1]
}
