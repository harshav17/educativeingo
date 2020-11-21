package educativeingo

//practice dynamic programming with 0/1 knapsack
func knapsackNaive(profits []int, weight []int, capacity int) int {
	return 1
}

func knapsackNaiveRecursive(profits []int, weight []int, capacity int, currIdx int) int {
	if capacity <= 0 || currIdx <= len(profits) {

	}
	return 1
}

func zeroOneKnapsack(profits []int, weights []int, capacity int) int {
	memo := make([][]int, len(weights))
	for m := range memo {
		memo[m] = make([]int, capacity+1)
	}
	return knapsackRecursiveTopDown(profits, weights, capacity, 0, memo)
}

func knapsackRecursiveTopDown(profits []int, weights []int, capacity int, currIdx int, memo [][]int) int {
	//base checks
	if currIdx+1 > len(profits) || capacity <= 0 {
		return 0
	}

	if memo[currIdx][capacity] != 0 {
		return memo[currIdx][capacity]
	}

	//including
	var profit1 int
	if weights[currIdx] <= capacity {
		profit1 = profits[currIdx] + knapsackRecursiveTopDown(profits, weights, capacity-weights[currIdx], currIdx+1, memo)
	}

	//excluding
	profit2 := knapsackRecursiveTopDown(profits, weights, capacity, currIdx+1, memo)

	maxProfit := profit2
	if profit1 > profit2 {
		maxProfit = profit1
	}

	memo[currIdx][capacity] = maxProfit
	return maxProfit
}

//practive dynamic progressing with fibonacci
func fibNaiveRec(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibNaiveRec(n-1) + fibNaiveRec(n-2)
}

func findFibBottomUp(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	memo := make([]int, n+1)
	memo[1], memo[2] = 1, 1

	for i := 3; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}

func findFibTopDown(n int) int {
	memo := make([]int, n+1)
	return fibTopDown(n, memo)
}

func fibTopDown(n int, memo []int) int {
	if memo[n] != 0 {
		return memo[n]
	}

	var res int
	if n == 1 || n == 2 {
		res = 1
	} else {
		res = fibTopDown(n-1, memo) + fibTopDown(n-2, memo)
	}
	memo[n] = res
	return res
}
