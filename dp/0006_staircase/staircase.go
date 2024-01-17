package staircase

func staircase(steps int, memo []int) int {
	if steps == 0 || steps == 1 || steps == 2 {
		return steps
	}
	if memo[steps] != -1 {
		return memo[steps]
	}
	memo[steps] = staircase(steps-1, memo) + staircase(steps-2, memo)
	return memo[steps]
}

func staircaseTopDown(steps int) int {
	memo := make([]int, steps+1)
	return staircase(steps, memo)
}

func staircaseBottomUp(steps int) int {
	memo := make([]int, steps+1)

	memo[0] = 0
	memo[1] = 1
	memo[2] = 2

	for i := 3; i <= steps; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[steps]
}
