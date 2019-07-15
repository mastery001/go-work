/**
* Count and Say - https://leetcode.com/problems/count-and-say/
*/

package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
    var (
    	num = "1"
    	count = 1
    )
    if n < 1 {
    	n = 1
    }
    if n > 30 {
    	n = 30
    }
    for count < n {
    	var temp string = ""
    	var c int = 1
    	len := len(num)
    	for i := 1 ; i < len ; i++ {
    		if num[i - 1] == num[i] {
    			c ++ 
    		}else {
    			temp += strconv.Itoa(c) + string(num[i - 1])
    			c = 1
    		}
    	}
    	temp += strconv.Itoa(c) + string(num[len - 1])
    	num = temp
    	count ++
    }
    return num
}

func main() {
	fmt.Println(countAndSay(0))
}