package countingbits

/*
 * For a given positive integer, n, your task is to return an array of length
 * n+1 such that for each x where 0 ≤ x ≤ n, result[x] is the count of
 * 1s in the binary representation of x.
 * Example binary representation
 * 0 0000
 * 1 0001
 * 2 0010
 * 3 0011
 * 4 0100
 * 5 0101
 * 6 0110
 * 7 0111
 * So if you notice then odd number always had 1 additional 1bit than previous
 * number, thus using this we can build tabulation.
 */

func countingBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {
			result[i] = i
		}
		if i%2 == 0 {
			result[i] = result[i/2]
		} else {
			result[i] = 1 + result[(i-1)/2]
		}
	}
	return result
}
