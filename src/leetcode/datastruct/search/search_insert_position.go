/**
* Search Insert Position - https://leetcode.com/problems/search-insert-position/
*/

package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
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
    		return mid
    	}
    }
    return left
}

func main() {
	fmt.Println(searchInsert([]int{1,3,5,6}, 2))
}