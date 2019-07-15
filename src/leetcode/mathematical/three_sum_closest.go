/**
* 3Sum Closest - https://leetcode.com/problems/3sum-closest/
*/

package main

import (
	"fmt"
	"sort"
)

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
    sort.Ints(nums)
    res := nums[0] + nums[1] + nums[2]
    for i := 0 ; i < n ; i++ {
    	start := i + 1
    	end := n - 1
    	for start < end {
    		temp := nums[i] + nums[start] + nums[end]
    		if Abs(temp - target) < Abs(res - target) {
    			res = temp
    		}
    		if temp > target {
    			end --
    		}else if temp < target {
    			start ++
    		}else {
    			return res
    		}
    	}
    }
    return res
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4} , 1))
}