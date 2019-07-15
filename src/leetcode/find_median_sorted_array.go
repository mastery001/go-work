/**
* https://leetcode.com/problems/median-of-two-sorted-arrays/
* 两个数组的中位数 -- 分割法
* 中位数性质：中位数两侧的个数相同
*/
package main

import (
	"fmt"
	"math"
)

func lot(exp bool , a int , b func() int) float64 {
	if exp {
		return float64(a)
	}
	return float64(b())
}

func findMedianSortedArrays(nums1 []int , nums2 []int) float64 {
	var len1 = len(nums1)
	var len2 = len(nums2)
	if len1 > len2 {
		return findMedianSortedArrays(nums2 , nums1)
	}
	var twoArraysMedian = (len1 + len2 + 1) / 2
	var left = 0
	var right = len1
	for left < right {
		// 取左边数组的左半边
		var l1 = left + (right - left) / 2
		// 加上取右边数组的左半边
		var l2 = twoArraysMedian - l1
		if nums1[l1] < nums2[l2 - 1] {
			left = l1 + 1
		}else {
			right = l1
		}
	}
	// 此时得出的left就是左边应该分的个数
	var m1 = left 
	var m2 = twoArraysMedian - left
	var c1 = math.Max(lot(m1 <= 0 , math.MinInt32 , func() int {return nums1[m1 - 1]}), lot(m2 <= 0 , math.MinInt32 , func() int {return nums2[m2 - 1]}))
	if (len1 + len2) % 2 == 1 {
		return c1
	}
	var c2 = math.Min(lot(m1 >= len1 , math.MaxInt32 , func() int {return nums1[m1]}), lot(m2 >= len2 , math.MaxInt32 , func() int {return nums2[m2]}))
	return (c1 + c2) * 0.5
}

func main() {
	num1 := []int{1,2,5,8}
	num2 := []int{3,4,6,9}
	fmt.Println(findMedianSortedArray(num1 , num2))
}

