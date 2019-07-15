/**
* Divide Two Integers - https://leetcode.com/problems/divide-two-integers/
*/

package main

import (
	"fmt"
	"math"
)

func Abs(a int) int{
	if a < 0 {
		return -a
	}
	return a
}

func divideProgressive(dividend int, divisor int) int {
    var (
    	flag = 0
    	res = 0
    )
    if (dividend < 0 && divisor > 0) || (divisor < 0 && dividend > 0) {
    	flag = 1
    }
    dividend = Abs(dividend)
    divisor = Abs(divisor)
    for dividend >= divisor {
    	var temp ,i int = divisor , 1
    	for dividend >= temp {
    		dividend -= temp
    		res += i
    		i <<= 1
    		temp <<= 1
    	}
    }
    if flag == 1 {
    	return -res
    }
    if res <= math.MinInt32 {
    	return math.MinInt32
    }else if res >= math.MaxInt32 {
    	return math.MaxInt32
    }
    return res	
}

func divide(dividend int, divisor int) int {
    var (
    	flag = 0
    	res = 0
    )
    if (dividend < 0 && divisor > 0) || (divisor < 0 && dividend > 0) {
    	flag = 1
    }
    dividend = Abs(dividend)
    divisor = Abs(divisor)
    if divisor == 1 {
    	res =  dividend
    }else {
    	for dividend - divisor >= 0 {
	    	res ++
	    	dividend -= divisor
	    }
    }
    if flag == 1 {
    	return -res
    }
    if res <= math.MinInt32 {
    	return math.MinInt32
    }else if res >= math.MaxInt32 {
    	return math.MaxInt32
    }
    return res
}

func main() {
	fmt.Println(divideProgressive(7, -3))
}