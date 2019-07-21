/*
* Jump Game II - https://leetcode.com/problems/jump-game-ii/
*/

package main

import (
	"math"
	"fmt"
)

func Max(a int , b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int , b int) int {
	if a > b {
		return b 
	}
	return a
}

func maxStep(nums []int) int {
	n := len(nums)
	var (
		end = 0 
		max = 0
		step = 0
	)
	// 每一次都选择最大的步数跳跃
	for i := 0 ; i < n - 1 ; i++ {
		max = Max(max , nums[i] + i)
		if i == end {
			end = max 
			step ++
		}
	}
	return step	
}

func jump(nums []int) int {
	// n := len(nums)
	// return recursive(nums, 0, n - 1) 
	return maxStep(nums) 
}

/**
* 速度太慢了。。。
*/
func recursive(nums []int , index int , target int) int {
	minStep := math.MaxInt32
	cur := nums[index]
	for j := cur ; j > 0 ; j --{
		jump := index + j
		// 到达目的地
		if jump == target {
			minStep = Min(1,minStep)
		}else if jump < target {
			minStep = Min(1 + recursive(nums , jump , target),minStep)
		}
		continue
	}
	return minStep
}

func main() {
	fmt.Println(jump([]int{5,2,1,1,4}))
}