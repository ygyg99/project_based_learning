package main

func min(a,b,c int)int{
    if a<b && a<c{
        return a
    }
    if b<a && b<c{
        return b
    }
    return c
}

func minDistance1(word1 string, word2 string) int {
	lg1, lg2 := len(word1), len(word2)
	// 初始化dp
	dp := make([][]int, lg1+1)
	for i := 0; i < lg1+1; i++ {
		dp[i] = make([]int, lg2+1)
	}
	// 赋初值
	for i := 1; i <= lg1; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}
	for j := 1; j <= lg2; j++ {
		dp[0][j] = dp[0][j-1] + 1
	}

	for i := 1; i <= lg1; i++ {
		for j := 1; j <= lg2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[lg1][lg2]
}

func minDistance2(word1 string, word2 string) int {
	n1 := len(word1)
	n2 := len(word2)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	// Initialize the first row
	for j := 1; j <= n2; j++ {
		dp[0][j] = dp[0][j-1] + 1
	}

	// Initialize the first column
	for i := 1; i <= n1; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}

	// Fill the DP table
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[n1][n2]
}