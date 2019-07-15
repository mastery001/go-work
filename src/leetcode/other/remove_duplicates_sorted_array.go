/**
* Remove Duplicates from Sorted Array - https://leetcode.com/problems/remove-duplicates-from-sorted-array/
*/

package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums) 
	if n == 0 {
		return 0
	}
	i := 0
	for j := 1 ; j < n ; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}


func main() {
	fmt.Println(removeDuplicates([]int{1}))
}