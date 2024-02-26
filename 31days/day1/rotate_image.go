package day1

/*
 Rotate Image: You are given an n x n 2D matrix representing an image,
 rotate the image by 90 degrees (clockwise).
 You have to rotate the image in-place, which means you have to modify
 the input 2D matrix directly.
 Steps that happen during 90-degree clockwise rotation
 1. Transpose the Matrix
    1 2 3
    4 5 6
    7 8 9
  to
    1 4 7
    2 5 8
    3 6 9
 2. Flip Horizontally
    1 4 7
    2 5 8
    3 6 9
to
    7 4 1
    8 5 2
    9 6 3
*/

func rotateImage(image [][]int) {
	n := len(image)
	// Step 1: transpose
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			image[i][j], image[j][i] = image[j][i], image[i][j]
		}
	}
	// Step 2: Flip horizontally
	for i := 0; i < n; i++ {
		j := 0
		k := n - 1
		for j < k {
			image[i][j], image[i][k] = image[i][k], image[i][j]
			j = j + 1
			k = k - 1
		}
	}
	return
}
