/**
* Multiply Strings - https://leetcode.com/problems/multiply-strings/
* 还可以使用Karatsuba乘法
*/

package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
 	l1 := len(num1)
 	l2 := len(num2)
 	n1 := make([]int , l1)
 	n2 := make([]int , l2)
 	for i := 0 ; i < l1 ; i ++ {
 		n1[i] = int(num1[i] - '0')
 	}   
 	 for i := 0 ; i < l2 ; i ++ {
 		n2[i] = int(num2[i] - '0')
 	}   
 	res := multiply0(n1 , l1 , n2 , l2)
 	i := 0
 	str := ""
 	for i < len(res) {
 		str += strconv.Itoa(res[i])
 		i++
 	}
 	return str
 	// fmt.Println(res)
 	// return hex.EncodeToString(res[:])
 	// return string(res[:])
}

func multiply0(num1 []int , l1 int , num2 []int , l2 int) []int {
	length := l1 + l2
	res := make([]int , length)
	for i := 0 ; i < l1 ; i ++ {
		for j := 0 ; j < l2 ; j++ {
			res[i + j + 1] += num1[i] * num2[j]
		}
	}
	// fmt.Println(res)
	// 处理进位
	for i := length - 1 ; i > 0 ; i-- {
		if res[i] >= 10 {
			res[i - 1] += res[i] / 10
			res[i] = res[i] % 10
		}
	}
	i := 0
	for i < length {
		if res[i] != 0 {
			return res[i:]
		}
		i++
	}
	if i == length {
		return res[length -1 :]
	}
	return res
}

func main() {
	fmt.Println(multiply("5", "12"))
}