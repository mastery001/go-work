/**
* atoi算法 https://leetcode.com/problems/string-to-integer-atoi/
*/
package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	n := len(str)
	var res , sign int
	var i = 0
	for ; i < n ; i++  {
		if str[i] == ' ' || str[i] == '\t' || str[i] == '\n' {
			continue
		}
		break
	}
	if i < n {
		if str[i] == '-' {
			sign = 1
			i++
		}else if str[i] == '+' {
			i++
		}
	}
	for i < n {
		if str[i] < '0' || str[i] > '9' {
			return 0
		}else {
			break
		}
	}
	for ; i < n ; i++ {
		if str[i] >= '0' && str[i] <= '9' {
			res = res * 10 + int(str[i] - '0')
			if res > math.MaxInt32 {
				if sign == 1 {
					return math.MinInt32
				}else {
					return math.MaxInt32
				}
			}
		}else {
			break
		}
	}

	if sign == 1 {
		res = -res
	}
	return res
}

func main() {
	fmt.Println(myAtoi("9223372036854775808"))
}