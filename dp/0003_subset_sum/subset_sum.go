package subsetsum

func subsetsum(set []int, n, sum int) bool {
	if sum == 0 {
		return true // empty set sums to zero
	}
	if n == 0 {
		return false
	}
	if set[n-1] > sum {
		return subsetsum(set, n-1, sum)
	}

	return subsetsum(set, n-1, sum) || subsetsum(set, n-1, sum-set[n-1])
}
