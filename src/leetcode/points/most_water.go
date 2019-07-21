/**
* Container With Most Water-https://leetcode.com/problems/container-with-most-water/
*/
package main

import (
	"fmt"
)

func Max(a int , b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int , b int) int {
	if a > b {
		return b 
	}
	return a
}

// 暴力解法
func maxAreaForce(height []int) int {
    l := len(height)
    if l < 2 {
    	return 0
    }
    if l == 2 {
    	return Min(height[0] , height[1])
    }
    var max int = 0 
    for i := 0 ; i < l ; i++ {
    	for j := 1 ; j < l - 1 ; j++ {
    		if height[j] >= height[j + 1] {
    			// 记录一次
    			max = Max(max, (j - i) * Min(height[i] , height[j]))
    		}
    	}
    	max = Max(max, (l - 1 - i) * Min(height[i] , height[l - 1]))
    }
    return max
}

// 指针法
func maxArea(height []int) int {
	var (
		l = len(height)
		max = 0
		left = 0
		right = l - 1
	)
	for left < right {
		max = Max(max, (right - left) * Min(height[left] , height[right]))

		if height[left] < height[right] {
			left ++ 
		}else {
			right --
		}
	}
	return max
}

func main() {
	arrs := []int{1,1}
	fmt.Println(maxArea(arrs))
}