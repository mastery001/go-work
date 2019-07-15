/**
* Longest Common Prefix - https://leetcode.com/problems/longest-common-prefix/
*/

package main

import "fmt"

func longestCommonPrefix(strs []string) string {
    var (
    	n = len(strs)
    	i = 0
    	res = ""
    )
    if n == 0 {
    	return res
    }
    firstLen := len(strs[0])
    for i < firstLen {
	    for j := 1 ; j < n ; j ++ {
	    	e := len(strs[j])
	    	if i >= e  {
	    		return res
	    	}
	    	if strs[0][i] !=  strs[j][i] {
	    		return res
	    	}
	    }
	    res += string(strs[0][i])
	    i ++
    }
    return res
}

func main() {
	// fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	// fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
	// fmt.Println(longestCommonPrefix([]string{"",""}))
	fmt.Println(longestCommonPrefix([]string{"aa","a"}))
}