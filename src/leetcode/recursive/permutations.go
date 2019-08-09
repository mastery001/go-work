/*
* Permutations - https://leetcode.com/problems/permutations/
*/

package main

import (
	"fmt"
)

func backtracking(nums []int , n int , res [][]int , one []int) [][]int {
	if len(one) == n {
		return append(res , one)
	}
	for i := 0 ; i < n ; i++ {
		temp := nums[i]
		exist := false
		for _,value := range(one) {
			if value == temp {
				exist = true
				break;
			}
		}
		if exist {
			continue
		}
		res = backtracking(nums, n, res, append(one , temp))
	}
	return res
}

func permute(nums []int) [][]int {
    return backtracking(nums , len(nums) , [][]int{} , []int{})
}

func main() {
	fmt.Println(permute([]int{1 , 1 , 3}))
}