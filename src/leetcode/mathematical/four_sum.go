/**
* 4Sum - https://leetcode.com/problems/4sum/
*/

package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
    n := len(nums)
    var res [][]int
    if n < 4 {
    	return res
    }
    if n == 4 && target == nums[0] + nums[1] + nums[2] + nums[3] {
    	return append(res , nums)
    }
    sort.Ints(nums)
    // return if the sum of the largest four numbers also less than the value of targer
    if nums[n - 4] + nums[n - 3] + nums[n - 2] + nums[n - 1] < target {
    	return res
    }
    for i:= 0 ; i < n - 3 ; i++ {
    	// filter duplicate element
    	if i > 0 && nums[i] == nums[i - 1] {
    		continue
    	}
    	if nums[i] + nums[i + 1] + nums[i + 2] + nums[i + 3] > target {
    		break
    	}
    	if nums[i] + nums[n - 3] + nums[n - 2] + nums[n - 1] < target {
    		continue
    	}
    	for j := i + 1 ; j < n - 2 ; j ++ {
    		if j - i > 1 && nums[j] == nums[j - 1] {
    			continue
    		}
	     	if nums[i] + nums[j] + nums[j + 1] + nums[j + 2] > target {
	    		break
	    	}
	    	if nums[i] + nums[j] + nums[n - 2] + nums[n - 1] < target {
	    		continue
	    	}  
	    	// double point
	    	var (
	    		left = j + 1
	    		right = n - 1
	    	) 		
	    	for left < right {
	    		sum := nums[i] + nums[j] + nums[left] + nums[right]
	    		if sum == target {
	    			res = append(res , []int{nums[i] , nums[j] , nums[left] , nums[right]})
	    			for left < right && nums[left] == nums[left + 1] {
	    				left ++
	    			}
	    			for left < right && nums[right] == nums[right - 1] {
	    				right --	    			
	    			}
	    			left ++
	    			right --
	    		}else if sum > target {
	    			right --
	    		}else {
	    			left ++
	    		}
	    	}
    	}

    }
    return res
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}