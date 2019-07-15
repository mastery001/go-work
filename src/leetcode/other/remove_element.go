/**
* Remove Element - https://leetcode.com/problems/remove-element/
*/

package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
 	n := len(nums)
 	i := 0
 	for j := 0 ; j < n ; j++ {
 		if nums[j] != val {
 			nums[i] = nums[j]
 			i++
 		}
 	}
 	return i  
}

func main() {
	fmt.Println(removeElement([]int{0,1,2,2,3,0,4,2} , 2))
}
