/**
* Longest Valid Parentheses - https://leetcode.com/problems/longest-valid-parentheses/
*/
package main

import (
	"fmt"
)

func Max(a int , b int) int {
	if a > b {
		return a
	}
	return b
}

/**
* 第 ii 个元素表示以下标为 ii 的字符结尾的最长有效子字符串的长度
* if s[i] == '(' && s[i - 1] == ')'
* 	dp[i] = dp[i - 2] + 2
* else if s[i] == ')' && s[i - 1] == ')'
	// 这里i - dp[i - 1] - 1指的是跳过中间的已经匹配的长度
	dp[i] = d[i - 1] + dp[i - dp[i - 1] - 1] + 2
*/
func dp(s string) int {
	n := len(s)
	maxLen := 0
	dp := make([]int , n)
	for i := 1 ; i < n ; i ++ {
		if s[i] == ')' {
			if s[i - 1] == '(' {
				temp := 0
				if i >= 2 {
					temp = dp[i - 2]
				}
				dp[i] = temp + 2
			}else if i - dp[i - 1] > 0 && s[i - dp[i - 1] - 1] == '('{
				temp := 0
				if i - dp[i - 1] >= 2 {
					temp = dp[i - dp[i - 1] - 2]
				}
				dp[i] = dp[i - 1] + temp + 2
			}
			maxLen = Max(maxLen , dp[i])
		}
	}
	return maxLen
}


func longestValidParentheses(s string) int {
    return dp(s)
}

func main() {
	fmt.Println(longestValidParentheses("(()())"))
}