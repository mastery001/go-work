/**
* Letter Combinations of a Phone Number - https://leetcode.com/problems/letter-combinations-of-a-phone-number/
*/

package main

import (
	"fmt"
)

var phoneNumer = [...]string{
	"", 	// 1
	"abc",  // 2
	"def",  // 3
	"ghi",  // 4
	"jkl",  // 5
	"mno",  // 6
	"pqrs", // 7
	"tuv",  // 8
	"wxyz", // 9
}

func combine(digitsMap []string , index int , res []string , one string) []string {
	// 临界条件
	if index == len(digitsMap)  {
		return res
	}else {
		numbers := digitsMap[index]
		n := len(numbers)
		for i := 0 ; i < n ; i++ {
			temp := one + string(numbers[i])
			if len(temp) == len(digitsMap) {
				res = append(res , temp)
			}else {
				res = combine(digitsMap , index + 1 , res , temp)
			}		
		}		
	}
	return res
}

func letterCombinations(digits string) []string {
	var digitsMap []string
	for _,s := range digits {
		digitsMap = append(digitsMap , phoneNumer[int(s - '1')])
	}
	// fmt.Println(digitsMap)
    var res []string
    res = combine(digitsMap , 0 , res , "")
    return res
}


func main() {
	fmt.Println(letterCombinations("223"))
}