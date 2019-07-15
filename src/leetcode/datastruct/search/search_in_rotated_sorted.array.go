/**
* Search in Rotated Sorted Array - https://leetcode.com/problems/search-in-rotated-sorted-array/
*/
package main

import (
	"fmt"
)

func findRotatedIndex(nums []int , left int , right int) int {
	if nums[left] < nums[right] {
		return 0
	}
    for left <= right {
    	mid := (right + left) / 2
    	if nums[mid] > nums[mid + 1] {
    		return mid + 1
    	}else {
    		if nums[mid] < nums[left] {
    			right = mid - 1 
    		}else {
    			left = mid + 1
    		}
    	}
    }
    return 0
}

func search(nums []int, target int) int {
    var (
    	n = len(nums)
    	left = 0
    	right = n - 1
    )
    if n == 0 {
        return -1
    }
    if n == 1 {
    	if nums[0] == target {
    		return 0
    	}else {
    		return -1
    	}
    }
    rotatedIndex := findRotatedIndex(nums, left, right)

    if nums[rotatedIndex] == target {
    	return rotatedIndex
    }
    if rotatedIndex == 0 {
    	return midSearch(nums, target, left, right)
    }
    if nums[0] > target {
    	return midSearch(nums , target , rotatedIndex , right)
    }
    return midSearch(nums , target , 0 , rotatedIndex)

}
func midSearch(nums []int, target int , left int , right int) int{
    for left <= right {
    	mid := (right + left) / 2
    	if nums[mid] == target {
    		return mid
    	}else if nums[mid] < target {
    		left = mid + 1
    	}else {
    		right = mid - 1
    	}
    }	
    return -1
}


func main() {
	fmt.Println(search([]int{1}, 4))
}