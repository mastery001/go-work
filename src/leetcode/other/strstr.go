/**
* Implement strStr() - https://leetcode.com/problems/implement-strstr/
*/

package main

import (
	"fmt"
)

func getNext(pattern string) []int {
	n := len(pattern)
	var next []int = make([]int , n)
	for i := 0 ; i < n ; i++ {
		next[i] = -1
	}
	var (
		i = 0
		j = -1
	)
	for i < n -1 {
		if j == -1 || pattern[i] == pattern[j] {
			i ++
			j ++
			next[i] = j
		}else {
			j = next[j]
		}
	}
	return next
}

func strStrKMP(haystack string, needle string) int { 
	var (
    	n = len(haystack)
    	m = len(needle)
    )	
    if m == 0 {
    	return 0
    }
	next := getNext(needle)
	var i , j = 0 , 0
    for i < n && j < m {
    	if j == -1 || haystack[i] == needle[j] {
    		i ++
    		j ++
    	}else {
    		j = next[j]
    	}
    }	
	if j == m {
        return i - j
    }
    return -1
}

/**
* O(m * n)
*/
func strStrBruteForce(haystack string, needle string) int { 
	var (
    	n = len(haystack)
    	m = len(needle)
    	i , j = 0 , 0
    )
    if m == 0 {
    	return 0
    }
    for i < n && j < m {
    	if haystack[i] == needle[j] {
    		i ++
    		j ++
    	}else {
    		i = i - j + 1
    		j = 0
    	}
    }
    if j == m {
        return i - j
    }
    return -1
}

func strStr(haystack string, needle string) int { 
	return strStrKMP(haystack , needle)
}

func main() {
	var (
		haystack = "hello"
		needle = "ll"
		// haystack = "mississippi"
		// needle = "pi"
	)
	fmt.Println(strStr(haystack, needle))
}