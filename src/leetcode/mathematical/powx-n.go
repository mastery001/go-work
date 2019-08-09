/*
* Pow(x, n) - https://leetcode.com/problems/powx-n/
*/

package main

import (
	"fmt"
	// "math"
)

/*
* 暴力解法，乘法运算耗时很长
*/
func myPowForce(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
    var sign bool = n > 0
    if n < 0 {
    	n = -n
    }
    i := 1
    var result float64 = x
    for i < n {
    	result *= x
    	i++
    }
    if sign {
    	return result
    }else {
    	return 1 / result
    }
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	var (
		sign = n > 0
		result = 1.0
		base = x
	)
	if n < 0 {
    	n = -n
    }
	for n > 0 {
		if n & 1 != 0 {
			result *= base
		}
		n = n >> 1
		base *= base
	}
    if sign {
    	return result
    }else {
    	return 1 / result
    }
}

func main() {
	// fmt.Println(myPow(2.0, -2))
	fmt.Println(myPow(1.0, -2147483648))
}