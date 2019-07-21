/**
* Trapping Rain Water - https://leetcode.com/problems/trapping-rain-water/
*/

package main

import (
	"fmt"
)

func trap(height []int) int {
    var (
		l = len(height)
		maxLeft = 0
		maxRight = 0
		left = 0
		right = l - 1
		sum = 0
	)
	for left < right {
		if height[left] < height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			}else {
				sum += maxLeft - height[left]
			}
			left ++
		}else {
			if height[right] >= maxRight {
				maxRight = height[right]
			}else {
				sum += maxRight - height[right]
			}
			right --
		}
		
	}
	return sum 
}

func main() {
	fmt.Print(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}