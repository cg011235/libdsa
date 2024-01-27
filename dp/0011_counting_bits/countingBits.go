package countingbits

/*
 *For a given positive integer, n, your task is to return an array of length
�
+
1
n+1
 such that for each
�
x
 where
0
≤
�
≤
�
0≤x≤n
, result[x] is the count of
1
1
s in the binary representation of
�
x
.
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
