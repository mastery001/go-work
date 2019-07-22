/*
* Permutations II - https://leetcode.com/problems/permutations-ii/
* Given a collection of numbers that might contain duplicates, return all possible unique permutations.
*/

package main

import (
	"fmt"
	"sort"
)

func backtracking(nums []int , n int , res [][]int , one []int , used []bool ) [][]int{
	if len(one) == n {
		t := make([]int, len(one))
        copy(t, one)
		return append(res , t)
	}
	for i := 0 ; i < n ; i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i - 1] && !used[i - 1] {
			continue
		}
		// one = append(one , nums[i])
		used[i] = true
		res = backtracking(nums, n, res, append(one , nums[i]), used)
		// one = one[0:len(one) - 1]
		used[i] = false
	}
	return res
}


func permuteUnique(nums []int) [][]int {
    n := len(nums)
    used := make([]bool , n)
    sort.Ints(nums)
    return backtracking(nums, n, [][]int{}, []int{}, used)
}

func main() {
	fmt.Println(permuteUnique([]int{1 , 1 , 3}))
	fmt.Println(permuteUnique([]int{2 , 2 , 1 , 1}))
}