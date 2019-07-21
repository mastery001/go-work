/**
* Wildcard Matching - https://leetcode.com/problems/wildcard-matching/
*/
package main

import (
	"fmt"
)

/*
* 递归法，速度太慢
*/
func recursiveMatch(s string, p string) bool {
	sn := len(s)
	pn := len(p)
	if pn == 0 {
		return sn == 0
	}
	firstMatch := sn > 0 && (s[0:1] == p[0:1] || p[0:1] == "?" || p[0:1] == "*")
	if p[0:1] == "*" {
		return recursiveMatch(s , p[1:]) || (firstMatch && recursiveMatch(s[1:] , p))
	}else {
		return firstMatch && recursiveMatch(s[1:] , p[1:])
	}
}

func isMatch(s string, p string) bool {
    return dpMatch(s , p)
}

/*
* 动态规划
* dp[i][j]表示i到j匹配
* 1. 
*/
func dpMatch(s string, p string) bool {
	sl := len(s)
	pl := len(p)
	memory := make([][]int , sl + 1)
	for i := 0 ; i < sl + 1 ; i++ {
		memory[i] = make([]int , pl + 1)
		for j := range memory[i] {
			memory[i][j] = -1
		}
	}
	return dp(memory , 0 , 0  , s , p)
}

func dp(memory [][]int , si int , pi int , s string , p string) bool {
	if memory[si][pi] != -1 {
		return memory[si][pi] == 1
	}
	res := false
	// 已经匹配到末尾了
	if pi == len(p) {
		res = si == len(s)
	}else {
		firstMatch := si < len(s) && (s[si:si+1] == p[pi:pi + 1] || p[pi:pi + 1] == "?" || p[pi:pi + 1] == "*")
		if p[pi:pi + 1] == "*" {
			res = dp(memory , si , pi + 1 , s , p) || (firstMatch && dp(memory , si + 1 , pi, s , p) )
		}else {
			res = firstMatch && dp(memory , si + 1 , pi + 1 , s , p)
		}	
	}
	if res {
		memory[si][pi] = 1
	}else {
		memory[si][pi] = 0
	}
	return res
}

func main() {
	fmt.Println(isMatch("cb", "?a"))
	fmt.Println(isMatch("abceb", "*a*b"))
	fmt.Println(isMatch("aa", "*"))
	// recursive elapsed 9.269s
	// dp elapsed 
	fmt.Println(isMatch("bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab", "b*b*ab**ba*b**b***bba"))
	fmt.Println(isMatch("aa", "a"))
}