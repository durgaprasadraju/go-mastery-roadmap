// Package dp implements classic dynamic programming solutions.
package dp

import "math"

// Knapsack01 returns max value for 0/1 knapsack. O(n*W).
func Knapsack01(weights, values []int, capacity int) int {
	n := len(weights)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}
	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			dp[i][w] = dp[i-1][w]
			if weights[i-1] <= w {
				if v := dp[i-1][w-weights[i-1]] + values[i-1]; v > dp[i][w] {
					dp[i][w] = v
				}
			}
		}
	}
	return dp[n][capacity]
}

// LIS returns length of longest increasing subsequence. O(n log n).
func LIS(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
 tails := make([]int, 0, len(arr))
	for _, v := range arr {
		lo, hi := 0, len(tails)
		for lo < hi {
			mid := lo + (hi-lo)/2
			if tails[mid] < v {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		if lo == len(tails) {
			tails = append(tails, v)
		} else {
			tails[lo] = v
		}
	}
	return len(tails)
}

// LCS returns length of longest common subsequence. O(mn).
func LCS(a, b string) int {
	m, n := len(a), len(b)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else if dp[i-1][j] > dp[i][j-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

// CoinChange returns min coins needed. O(n*amount).
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if c <= i && dp[i-c] != math.MaxInt32 && dp[i-c]+1 < dp[i] {
				dp[i] = dp[i-c] + 1
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
