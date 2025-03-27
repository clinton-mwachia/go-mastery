package main

import "fmt"

// longestCommonSubsequence calculates the length of the longest common subsequence between two strings.
func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	// Initialize a 2D slice to store lengths of LCS for substrings
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// Fill the dp table in a bottom-up manner
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	text1 := "abcde"
	text2 := "ace"
	fmt.Printf("The length of the longest common subsequence is: %d\n", longestCommonSubsequence(text1, text2))
}
