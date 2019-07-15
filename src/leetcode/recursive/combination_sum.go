/**
* Combination Sum - https://leetcode.com/problems/combination-sum/
*/

package main

import (
	"fmt"
	"sort"
)

/**
* allow duplicate number
*/
func dfsNotDuplicate(candidates []int, target int , res [][]int , one []int , index int) [][]int {
	n := len(candidates)
	for i := index ; i < n ; i++ {
		if i > index && candidates[i - 1] == candidates[i] {
			continue
		}
		temp := target - candidates[i]
		if temp < 0 {
			break
		}else if temp == 0 {
			t := make([]int, len(one))
			copy(t, one)
			t = append(t , candidates[i])
			// fmt.Println(t , res , i , index)
			return append(res , t)
		}
		next := i + 1
		if next < n {
			res = dfsNotDuplicate(candidates , temp , res , append(one , candidates[i]) , next)
		}
	}
	
	return res
}

/**
* allow duplicate number
*/
func dfs(candidates []int, target int , res [][]int , one []int , index int) [][]int {
	if target == 0 {
		t := make([]int, len(one))
        copy(t, one)
		// fmt.Println(one , res)
		res = append(res , t)
	}else {
		n := len(candidates)
		for i := index ; i < n ; i++ {
			temp := target - candidates[i]
			if temp < 0 {
				break
			}
			res = dfs(candidates , temp , res , append(one , candidates[i]) , i)
		}
	}

	return res
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var one []int
	sort.Ints(candidates)
	return dfsNotDuplicate(candidates , target , res , one , 0)
}

func main() {
	// fmt.Println(combinationSum([]int{3,2,4,7}, 7))
	// fmt.Println(combinationSum([]int{2 ,3 ,5}, 8))
	// fmt.Println(combinationSum([]int{7,3,2}, 18))
	fmt.Println(combinationSum([]int{10,1,2,7,6,1,5},8))
	// fmt.Println(combinationSum([]int{2,5,2,1,2},5))

}