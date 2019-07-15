/**
* First Missing Positive - https://leetcode.com/problems/first-missing-positive/
Given an unsorted integer array, find the smallest missing positive integer.

Example 1:

Input: [1,2,0]
Output: 3
Example 2:

Input: [3,4,-1,1]
Output: 2
Example 3:

Input: [7,8,9,11,12]
Output: 1
Note:

Your algorithm should run in O(n) time and uses constant extra space.
*/

package main

import (
	"fmt"
	"math"
)

func firstMissingPositive(nums []int) int {
    n := len(nums)
    min := math.MaxInt32
    moreThanMin := math.MaxInt32
    for i := 0 ; i < n ; i ++ {
    	if nums[i] > 0 && nums[i] < min{
    		// moreThanMin = min
    		min = nums[i]
    	}else if nums[i] > min && moreThanMin > nums[i] {
			moreThanMin = nums[i]
    	}
    }
    fmt.Println(min , moreThanMin)
    return 0
}

func main() {
	fmt.Println(firstMissingPositive([]int{1,2,3}))
}