package algo

import (
	"math"
	"strings"
)

// Levenheistein distance: to get closest match in a list of strings.

func LevenheisteinRecursive(s1, s2 string, i, j int) int {
	if i < 0 {
		return j + 1
	}
	if j < 0 {
		return i + 1
	}
	if s1[i] == s2[j] {
		return LevenheisteinRecursive(s1, s2, i-1, j-1)
	}
	return min(LevenheisteinRecursive(s1, s2, i-1, j), LevenheisteinRecursive(s1, s2, i, j-1), LevenheisteinRecursive(s1, s2, i-1, j-1)) + 1
}

func LevenheisteinDP(s1, s2 string, i, j int, dp [][]int) int {
	if i < 0 {
		return j + 1
	}
	if j < 0 {
		return i + 1
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	if s1[i] == s2[j] {
		dp[i][j] = LevenheisteinDP(s1, s2, i-1, j-1, dp)
	}
	dp[i][j] = min(LevenheisteinDP(s1, s2, i-1, j-1, dp), LevenheisteinDP(s1, s2, i-1, j, dp), LevenheisteinDP(s1, s2, i, j-1, dp)) + 1
	return dp[i][j]
}

func Levenheistein(s1, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return LevenheisteinDP(s1, s2, len(s1)-1, len(s2)-1, dp)
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func ClosestDistance(s []string, input string, threshold int) string {
	min := math.MaxInt32;
	closest := "";

	for _, v := range s {
		if strings.HasPrefix(v, input) {
			return v
		}
	}
	for _, v := range s {
		distance := Levenheistein(v, input)
		if distance < min && distance <= threshold {
			min = distance
			closest = v
		}
	}
	return closest
}
