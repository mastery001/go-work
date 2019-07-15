/**
* Next Permutation - https://leetcode.com/problems/next-permutation/
*1. 先找出最大的索引 k 满足 nums[k] < nums[k+1]，如果不存在，就翻转整个数组；
*2. 再找出另一个最大索引 l 满足 nums[l] > nums[k]；
*3. 交换 nums[l] 和 nums[k]；
*4. 最后翻转 nums[k+1:]。
*/

package main

import (
	"fmt"
)

func swap(nums []int , i int , j int ) {
	nums[i] , nums[j] = nums[j] , nums[i]
}

func reverse(nums []int , start int , end int) {
	mid := (end - start) / 2 + start + 1
	for i := start ; i < mid ; i++ {
		swap(nums , i , end - i + start)
	}
}

func nextPermutation(nums []int)  {
	n := len(nums)
	firstIndex := -1
	for i := n -1 ; i > 0 ; i-- {
		if nums[i] > nums[i - 1] {
			firstIndex = i -1
			break 
		}
	}
	if firstIndex != -1 {
		for i := n - 1 ; i > firstIndex ; i-- {
			if nums[i] > nums[firstIndex] {
				swap(nums, firstIndex, i)
				break
			}
		}
	}
	reverse(nums, firstIndex + 1, n - 1)
}

func main() {
	// nums := []int{1 , 2 , 3}
	// nums := []int{1,2,7,4,3,1}
	nums := []int{1 , 3 , 2}
	nextPermutation(nums)
	fmt.Println(nums)
}