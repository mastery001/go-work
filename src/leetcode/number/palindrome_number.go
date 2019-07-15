/**
* Palindrome Number:https://leetcode.com/problems/palindrome-number/
*/
package main

import "fmt"

func isPalindrome(x int) bool {
    if x < 0 {
    	return false
    }
    if x < 10 {
    	return true
    }
    // 求整数的逆
    var (
    	reserve = 0
    	tempX = x
    )
    for tempX > 0 {
    	temp := reserve * 10 + tempX % 10
    	tempX /= 10
    	reserve = temp
    }
    fmt.Println(reserve)
    return x == reserve
}

func main() {
	isPalindrome(200)
}