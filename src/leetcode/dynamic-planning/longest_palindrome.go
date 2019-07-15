/**
* 最长回文子序列
* https://leetcode.com/problems/longest-palindromic-substring/
* 1. 中心扩展法
* 2. 动态规划
*/
package main

import (
	"fmt"
)

type Record struct {
	begin int 
	end int
}

func max(maxRecord Record , unknownRecord Record) Record {
	if maxRecord.end - maxRecord.begin < unknownRecord.end - unknownRecord.begin {
		return unknownRecord
	}
	return maxRecord
}

func findLeft(s string) string {
	    var len int = len(s)
    if len == 0 || len ==  1 {
    	return s
    }
	var r Record = Record{0,1}
	for i := 0 ; i < len - 1 ; i++ {
		var left = i
		var right = len - 1
		// 记录是否开始与end节点
		var begin = left
		var end = right
		var flag bool = false
		for left < right {
			if s[left] == s[right] {
				left ++
				right --
				flag = true
			}else {
				var stepLeft = left + 1
				for stepLeft < right {
					if s[stepLeft] == s[right] {
						break
					}
					stepLeft++
				}
				end = right
				if stepLeft < right {
					begin = stepLeft
					left = stepLeft + 1
				}else {
					left = i
					begin = left
				}
				right --
				flag = false
			}
		}
		if flag {
			r = max(r , Record{begin,end})	
		}
	}
	if r.end == 1 {
		r.end += 1
	}
	fmt.Println(r)
	return s[r.begin:r.end]
}

func centerFind(s string) string {
    var len int = len(s)
    var maxSize , res = 0 , ""
  	for i := range s {
  		// aba型
  		var j,k = i , i
  		for j >= 0 && k < len {
  			if s[j] == s[k] {
  				var size int = k - j + 1
  				if size > maxSize {
  					maxSize = size
  					res = s[j:k+1]
  				}
  				j--
  				k++
  			}else {
  				break
  			}
  		}	

  		// abba型
  		j,k = i , i+1
  		for j >= 0 && k < len {
  			if s[j] == s[k] {
  				var size int = k - j + 1
  				if size > maxSize {
  					maxSize = size
  					res = s[j:k+1]
  				}
  				j--
  				k++
  			}else {
  				break
  			}
  		}	
  	}
  	return res	
}

/**
动态规划
*/
func dpFind(s string) string {
	var n = len(s)
	if n == 0 {
		return s
	}
	var dp = make([][]int , n)
	var max , start = 1 , 0
	for i := range s {
		dp[i] = make([]int , n)
		dp[i][i] = 1
		if i < n - 1 && s[i] == s[i + 1] {
			dp[i][i+1] = 1
			start = i
			max = 2
		}
	}
	// 从长度为3子串开始
	for i:= 3 ; i <= n ; i++ {
		for j:= 0 ; i + j - 1 < n ; j++ {
			// 末尾的index
			k:= i + j - 1
			if s[j] == s[k] && dp[j + 1][k - 1] == 1 {
				dp[j][k] = 1
				max = i
				start = j
			}
		}
	}

	return s[start : start + max]
}

func longestPalindrome(s string) string {
	return dpFind(s)
}

func main() {
	fmt.Println(longestPalindrome("a"))
}