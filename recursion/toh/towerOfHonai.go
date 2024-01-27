package main

import "fmt"

func towerOfHanoi(n int, src, dst, tmp byte) {
	if n == 0 {
		return
	}
	towerOfHanoi(n-1, src, tmp, dst)
	fmt.Printf("Move %d disk from peg %c to peg %c \n", n, src, dst)
	towerOfHanoi(n-1, tmp, dst, src)
}

func main() {
	towerOfHanoi(3, 'A', 'C', 'B')
}
