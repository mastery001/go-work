/**
* Integer to Roman -- https://leetcode.com/problems/integer-to-roman/
*/

package main

import "fmt"

var (
	roman = [...]int{1000 , 900 , 500 , 400 , 100 , 90 , 50 , 40 , 10 , 9 , 5 , 4 , 1}
	romanStr = [...]string{"M" , "CM" , "D" , "CD" , "C" , "XC" , "L" , "XL" , "X" , "IX" , "V" , "IV" , "I"}
)

func intToRoman(num int) string {
    var res string 
    n := len(roman)
    i := 0
    for num > 0 && i < n {
    	if num - roman[i] >= 0 {
    		num -= roman[i]
    		res += romanStr[i]
    	}else {
    		i ++
    	}
    }
    return res
}

func romanToIntSequence(s string) int {
	var (
		n = len(s)
		start = 0
		m = len(romanStr)
		i = 0
		res = 0
	)
	for start < n && i < m {
		e := len(romanStr[i])
		if start + e <= n && romanStr[i] == s[start : start + e] {
			res += roman[i]
			start += e
			if e == 2 {
				i ++
			}
		}else {
			i ++
		}	
	}
	return res
}

func romanToIntReverse(s string) int {
	var (
		n = len(s)
		j = n
		m = len(romanStr)
		i = m - 1
		res = 0
	)
	for j >= 0 && i >= 0 {
		e := len(romanStr[i])
		if j - e >= 0 && romanStr[i] == s[j - e : j] {
			res += roman[i]
			j -= e
			if e == 2 {
				i --
			}
		}else {
			i --
		}
	}
	return res
}

func romanToInt(s string) int {
	return romanToIntReverse(s)
}

func main() {
	s := intToRoman(20846)
	fmt.Println(s)
	fmt.Println(romanToInt(s))
}