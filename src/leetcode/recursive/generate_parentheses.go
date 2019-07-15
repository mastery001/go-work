/*
* Generate Parentheses - https://leetcode.com/problems/generate-parentheses/
* Longest Valid Parentheses - https://leetcode.com/problems/longest-valid-parentheses/
* if s == 6 
*	...
* else 
* 	candidates = ['(' , ')']
*	for i in range(candidates) 
* 		a[k] = candidates[i]
*		dfs(k + 1)
*		a[k] = ''
*/ 
package main

import "fmt"

func backtracking(a []string , res string , open int , close int , n int) []string {
	if len(res) == n * 2{
		return append(a , res)
	}
	if open < n {
		a = backtracking(a , res + "(" , open + 1 , close , n)
	}
	if close < open {
		a = backtracking(a , res + ")" , open , close + 1 , n)
	}
	return a
}

func generateParenthesis(n int) []string {
	var a []string
    return backtracking(a , "" , 0 , 0 , n)
}

func main() {
	fmt.Println(generateParenthesis(3))
}