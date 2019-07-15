/**
* Find First and Last Position of Element in Sorted Array - https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
*/

package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
    var (
    	n = len(nums)
    	left = 0
    	right = n - 1
    )
    for left <= right {
    	mid := (left + right) / 2
    	if nums[mid] > target {
    		right = mid - 1
    	}else if nums[mid] < target {
    		left = mid + 1
    	}else {
    		start , end := mid , mid
    		for start - 1 >= left && nums[start] == target {
    			start --
    		}
    		for end + 1 <= right && nums[end + 1] == target {
    			end ++
    		}
    		return []int{start , end}
    	}
    }
    return []int{-1 , -1}
}

func main() {
	// fmt.Println(searchRange([]int{5,7,7,8,8,10}, 8))
	// fmt.Println(searchRange([]int{5,7,7,8,8,10}, 6))
	fmt.Println(searchRange([]int{1,2,3}, 2))
	
}